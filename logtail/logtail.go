package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/Shopify/sarama"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
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
	logrus.Info("(re)start taillog master")
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
	logrus.Info("(re)start dotaillog:" + path)
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
		return
	}
	defer f.Close()
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

//var producer *kafka.Producer
var producer sarama.AsyncProducer

func newproducer(bootstrapservers string) {
	var err error
	logrus.Info("kafka bootstrapservers:" + bootstrapservers)

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	bootstrapserverarr := strings.Split(bootstrapservers, ",")
	producer, err = sarama.NewAsyncProducer(bootstrapserverarr, config)
	if err != nil {
		logrus.Fatal(err)
	}
	// Trap SIGINT to trigger a graceful shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	var (
		wg sync.WaitGroup
	)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for range producer.Successes() {
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for err := range producer.Errors() {
			logrus.Error(err)
		}
	}()
}

func toKafka(str string, topic string) {
	//topic = "logtail_test"
	message := &sarama.ProducerMessage{Topic: topic, Value: sarama.StringEncoder(str)}
	//fmt.Println("int logging"+time.Now().Format(time.RFC3339), str)
	select {
	case producer.Input() <- message:
	}
}
