package main

import (
	"bufio"
	"os"
	"encoding/base64"
	"fmt"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var line = ""
	for scanner.Scan() {
		line += scanner.Text()
	}
	decoded, _ := base64.StdEncoding.DecodeString(line)
	fmt.Println(string(decoded))
}