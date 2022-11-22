package main

import (
	"fmt"

	"simple-tools/simhash"
)

// this is a test for 3 texts and some words, and for shingled/weighted methods;
func main() {
	simhash.Alg = simhash.Weighted
	a := []byte{}
	b := []byte{}
	b = []byte(`用正则表达式限制只能输入数字和英文：onkeyup="value=value.replace(/[W]/g,') "onbeforepaste="clipboardData.setData('text',clipboardData.getData('text').replace(/[^d]/g,'`)
	a = []byte(`用正则表达式限制只能输入数字：onkeyup="value=value.replace(/[^d]/g,') "onbeforepaste= "clipboardData.setData('text',clipboardData.getData('text').replace(/[^d]/g,'))"`)
	b = []byte(`用正则表达式限制只能输入数字和英文：onkeyup="=value.replace(/[d]/g,') "onbeforepaste="clipboardData.setData('text',clipboardData.getData('text').replace(/[^d]/g,'`)

	simhash.Alg = simhash.Weighted
	fmt.Println(simhashCompare(a, b))
	simhash.Alg = simhash.Shingled
	fmt.Println(simhashCompare(a, b))
}

func simhashCompare(a, b []byte) uint8 {
	return simhash.Compare(simhash.SimhashByteArray(a), simhash.SimhashByteArray(b))
}
