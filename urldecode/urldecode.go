package main

import (
	"bufio"
	"os"
	"fmt"
	"net/url"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		decoded, _ := url.PathUnescape(scanner.Text())
		fmt.Println(string(decoded))
	}
}