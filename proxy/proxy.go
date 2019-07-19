package main

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

func handler(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	//path := r.URL.Path
	//path = "http://10.39.40.94:9220/logstash-cms-esdoc*/_search"
	//path = "http://10.83.0.44:9201/zipkin/_search"
	//querystr := r.URL.RawQuery
	//if strings.Contains(querystr, "q=") {
	//	querystr = strings.Replace(querystr, "q=", "q=timestamp_millis:* AND ", 1)
	//} else {
	//	querystr += "q=timestamp_millis:*"
	//}
	path := r.URL.Path
	//url := "http://10.39.40.94:9220" + path + "?" + r.URL.RawQuery
	url := "http://10.83.0.44:9201" + path + "?" + r.URL.RawQuery
	req, err := http.NewRequest("GET", url, r.Body)

	logrus.Info(url, "     ", r.Body)
	if err != nil {
		w.Write([]byte(err.Error()))
		logrus.Error(err.Error())
		w.WriteHeader(500)
		return
	}
	//for k, v := range r.Header {
	//	vs := strings.JoinStrings(",", v...)
	//	req.Header.Set(k, vs)
	//}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		w.Write([]byte(err.Error()))
		logrus.Error(err.Error())
		w.WriteHeader(500)
		return
	}
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	if bodystr, err := ioutil.ReadAll(resp.Body); err != nil {
		w.Write([]byte(err.Error()))
		logrus.Error(err.Error())
		w.WriteHeader(500)
		return
	} else {
		w.Write([]byte(bodystr))
		logrus.Info(string(bodystr))
		return
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":19999", nil)
}
