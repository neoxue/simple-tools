package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println(1 << 0)

	r := regexp.MustCompile("[\u4e00-\u9fa5]|[a-zA-Z0-9]*")
	ret := r.FindAll([]byte(`here is aaa 999 中文... 人啊。。。sss。\。sss{}.

onkeyup="value=value.replace(/[^d]/g,') "onbeforepaste= "clipboardData.setData('text',clipboardData.getData('text').replace(/[^d]/g,'))"

..`), -1)

	for _, k := range ret {
		fmt.Println(string(k))

	}

}
