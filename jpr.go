package main

import (
	"fmt"
	"bufio"
	"os"
	"encoding/json"
	. "github.com/logrusorgru/aurora"
)
/*
https://stackoverflow.com/questions/43843477/scanln-in-golang-doesnt-accept-whitespace
fmt.scanln does not accept white space
instead, use bufio.NewScanner

 */

//str,_ := jsonpath.JsonPathLookup(jsonObject, "$.msg")
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		line := scanner.Text()
		var jsonObject interface{}
		json.Unmarshal([]byte(line), &jsonObject)
		recursivePrint2(jsonObject, 0, "\n")
	}
}

func recursivePrint2(data interface{}, n int, tail string) {
	switch data.(type) {
		default:
			jsonStr, _ :=json.MarshalIndent(data, "", "  ")
			fmt.Printf(`"%v"`, Red(string(jsonStr)))
		case bool:
			fmt.Printf(`"%v"`, Brown(data))
		case int:
			fmt.Printf(`"%v"`, Gray(data))
		case string:
			fmt.Printf(`"%v"`, Magenta(data.(string)))
		case map[string] interface{}:
			fmt.Println("{")
			dataArr := data.(map[string] interface{})
			for k2, v2 := range dataArr {
				recursivePrintWhiteSpace(n + 1)
				fmt.Printf(`"%v":`, Green(k2))
				recursivePrint2(v2, n + 1, ",\n")
			}
			fmt.Print(w(n) + "}")
		case [] interface{}:
			fmt.Println("[")
			dataArr := data.([]interface{})
			for _, v2 := range dataArr {
				recursivePrintWhiteSpace(n + 1)
				recursivePrint2(v2, n + 1, ",\n")
			}
			fmt.Print(w(n) + "]")
	}
	fmt.Print(tail)
}
func recursivePrintWhiteSpace(n int) {
	for i := 0; i<n; i++ {
		fmt.Print("  ")
	}
}

func w(n int) string{
	var a string = ""
	for i := 0; i<n; i++ {
		a += "  "
	}
	return a
}



