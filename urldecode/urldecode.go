package main

import (
	"bufio"
	"github.com/neoxue/goutils/rerrors"
	"github.com/sirupsen/logrus"
	"go.etcd.io/etcd/client"
	"io/ioutil"
	"net/http"

	"os"
	"fmt"
	"net/url"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		decoded, _ := url.PathUnescape(scanner.Text())
		fmt.Println(string(decoded))
	}


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