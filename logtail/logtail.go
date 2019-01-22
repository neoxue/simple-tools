package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sirupsen/logrus"
)

var (
	logtaildir              = "/data1/ms/log/"
	logtailinfofile         = "/tmp/logtailinfofile"
	logtailtoremoteinfofile = "/data1/ms/sys/logtailtoremoteinfofile"
	// key均为绝对路径
	logtailinfos        = map[string]*logtailinfo_t{}
	logtailtoremoteinfo = logtailtoremoteinfo_t{}
	mu                  sync.RWMutex
)

type logtailtoremoteinfo_t struct {
	Typ              string
	Bootstrapservers string
}

// 统计地址
func main() {
	// 设置日志
	rl, _ := rotatelogs.New(logtaildir + "logrus.%Y%m%d")
	logrus.SetOutput(rl)
	logrus.SetLevel(logrus.InfoLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	initConfigs()
	newproducer(logtailtoremoteinfo.Bootstrapservers)
	monitorlogs()
}

func monitorlogs() {
	if logtaildir == "" {
		fmt.Println("no logtaildir found")
		logrus.Fatal("no logtaildir found")
	}
	if fileinfo, err := os.Stat(logtaildir); err != nil {
		fmt.Println(err)
		logrus.Fatal(err)
	} else {
		if !fileinfo.IsDir() {
			fmt.Println("logtaildir:" + logtaildir + "is not a dir")
			logrus.Fatal("logtaildir:" + logtaildir + " is not a dir")
		}
	}
	for true {
		go filepath.Walk(logtaildir, taillog)
		time.Sleep(10 * time.Second)
		//刷logtailinfos 进入文件
		loglogtailinfos()
	}
}

type logtailinfo_t struct {
	Topic   string
	File    string
	Cursor  int64
	Logtime time.Time
}

func computetopic(path string) string {
	if strings.Contains(path, logtaildir) {
		strarr := strings.Split(path, "/")
		if len(strarr) > 5 {
			return "logtail_" + strarr[4]
		}
	}
	return "logtail_default"
}
func taillog(path string, info os.FileInfo, err error) error {
	mu.Lock()
	defer mu.Unlock()
	topic := computetopic(path)
	if err != nil {
		logrus.Error(err)
	}
	// dir continue
	if info.IsDir() {
		return nil
	}
	// 一天以前的日志不做处理
	if time.Now().Sub(info.ModTime()).Seconds() > 86400 {
		return nil
	}
	// 已经读完不做处理
	if logtailinfos[path] != nil && logtailinfos[path].Cursor == info.Size() {
		return nil
	}
	// 文件曾被重置
	if logtailinfos[path] != nil && logtailinfos[path].Cursor > info.Size() {
		logtailinfos[path].Cursor = 0
	}
	// cursor 现在肯定小于 文件size, 判断3分钟内读取过，则返回
	if logtailinfos[path] != nil && time.Now().Sub(logtailinfos[path].Logtime).Seconds() < 3*60 {
		return nil
	}
	// 若信息不存在，记录信息
	if logtailinfos[path] == nil {
		logtailinfos[path] = &logtailinfo_t{File: path, Logtime: time.Now(), Cursor: 0, Topic: topic}
	}
	go dotaillog(path, info)
	logrus.Info("(restart) dotaillog:" + path)
	return nil
}
func dotaillog(path string, info os.FileInfo) {
	mu.Lock()
	ptr := logtailinfos[path].Cursor
	topic := logtailinfos[path].Topic
	mu.Unlock()
	// truely tail file
	f, err := os.Open(path)
	if err != nil {
		logrus.Error(err)
	}
	_, err = f.Seek(ptr, io.SeekStart)
	if err != nil {
		logrus.Error(err)
	}
	// Start reading from the file with a reader.
	lastTime := time.Now()
	reader := bufio.NewReader(f)
	for true {
		var line string
		for {
			line, err = reader.ReadString('\n')
			if err != nil {
				break
			}
			toremote(line, topic)
			cursor, _ := f.Seek(0, io.SeekCurrent)
			mu.Lock()
			lastTime = time.Now()
			logtailinfos[path].Logtime = lastTime
			logtailinfos[path].Cursor = cursor
			mu.Unlock()
		}
		if err != io.EOF {
			logrus.Error(err)
		}
		time.Sleep(3 * time.Second)
		// 距离上一次读取30s以上, 则不再读取
		if err == nil && time.Now().Sub(lastTime).Seconds() > 60 {
			return
		}
	}
}
func toremote(str string, tpc string) {
	toKafka(str, tpc)
}

func initConfigs() {
	switch len(os.Args) {
	case 1:
	case 2:
		logtaildir = os.Args[1]
	case 3:
		logtaildir = os.Args[1]
		logtailinfofile = os.Args[2]
	default:
		logtaildir = os.Args[1]
		logtailinfofile = os.Args[2]
		logtailtoremoteinfofile = os.Args[3]
	}
	if err := fetchlogtailinfo(); err != nil {
		logrus.Error(err)
	}
	if err := fetchremoteinfo(); err != nil {
		fmt.Println(err)
		logrus.Fatal(err)
	}
}

/*
if file empty,     initialize empty logtailinfos
if file not empty, initialize logtailinfos
*/
func fetchlogtailinfo() (err error) {
	info, err := ioutil.ReadFile(logtailinfofile)
	if err != nil {
		return err
	}
	err = json.Unmarshal(info, &logtailinfos)
	if err != nil {
		return err
	}
	return nil
}

func fetchremoteinfo() (err error) {
	info, err := ioutil.ReadFile(logtailtoremoteinfofile)
	if err != nil {
		return err
	}
	err = json.Unmarshal(info, &logtailtoremoteinfo)
	if err != nil {
		return err
	}
	return nil
}

func loglogtailinfos() {
	mu.Lock()
	defer mu.Unlock()
	// 删除过期文件信息
	for k, v := range logtailinfos {
		if time.Now().Sub(v.Logtime).Seconds() > 2*86400 {
			delete(logtailinfos, k)
		}
	}
	logtailinfos_bts, err := json.Marshal(logtailinfos)
	if err != nil {
		fmt.Println(err)
		logrus.Fatal(err)
	}
	err = ioutil.WriteFile(logtailinfofile, logtailinfos_bts, 0777)
	if err != nil {
		fmt.Println(err)
		logrus.Fatal(err)
	}
}

var producer *kafka.Producer

func newproducer(bootstrapservers string) {
	logrus.Info("kafka bootstrapservers:" + bootstrapservers)
	var err error
	producer, err = kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": bootstrapservers})
	if err != nil {
		panic(err)
	}
	defaultTopic := "logtail_default"
	_, err = producer.GetMetadata(&defaultTopic, true, 1000)
	if err != nil {
		fmt.Println(err)
		logrus.Fatal(err)
	}
	// Delivery report handler for produced messages
	go func() {
		for e := range producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					logrus.Error("Delivery failed: %v\n", ev.TopicPartition)
				}
				//fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
			}
		}
	}()
}

func toKafka(str string, topic string) {
	producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(str),
	}, nil)
	// Wait for message deliveries before shutting down
	//p.Flush(15 * 1000)
}
