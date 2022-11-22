package main

import (
	"bufio"
	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/neoxue/goutils/rerrors"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

func ReadLine(fileName string, handler func(string)) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		ch <- line
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
}

// the reason add channel and sleep is
// if your os could only maintain 65536 (out) connections the same time (65535 ports..)
// it's probabely neccessary to limit your requests;-(
var ch = make(chan string, 1000)

func readChan() {
	for true {
		time.Sleep(1 * time.Millisecond)
		line := <-ch
		go fuckkv(line)
	}
}

var client = &http.Client{}

func fuckkv(line string) {
	urlpath := "http://10.41.11.119:20001/comos_urls/" + url.QueryEscape(line)
	if req, err := http.NewRequest("GET", urlpath, strings.NewReader("")); err != nil {
		err = rerrors.WrapErrors(err, "new request failed", "x302")
		logrus.WithFields(logrus.Fields{"actionstr": "empty"}).Error(err)
	} else {
		resp, err := client.Do(req)
		if err == nil {
			defer resp.Body.Close()
			bts, _ := ioutil.ReadAll(resp.Body)
			logrus.Info(urlpath + "    " + string(bts))
		} else {
			logrus.Error(err)
		}
	}
}

func main() {
	go readChan()
	logrus.SetFormatter(&logrus.TextFormatter{QuoteEmptyFields: false, ForceColors: true, FullTimestamp: true, DisableColors: false})
	rl, _ := rotatelogs.New("/data1/ms/log/urltest.%Y%m%d")
	logrus.SetOutput(rl)
	logrus.SetLevel(logrus.DebugLevel)
	ReadLine("pure_urls.txt", fuckkv)
}
