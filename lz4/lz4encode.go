//compress project main.go
package main

import "fmt"
import (
	"github.com/pierrec/lz4"
	"bytes"
)

var fileContent = `CompressBlock compresses the source buffer starting at soffet into the destination one.
This is the fast version of LZ4 compression and also the default one.
The size of the compressed data is returned. If it is 0 and no error, then the data is incompressible.
An error is returned if the destination buffer is too small.`

func main() {
	toCompress := []byte(fileContent)
	compressed := make([]byte, len(toCompress))

	//compress
	l, err := lz4.CompressBlock(toCompress, compressed, 0)
	fmt.Println(bytes.Count(compressed))
	if err != nil {
		panic(err)
	}
	fmt.Println("compressed Data:", string(compressed[:l]))

	//decompress
	decompressed := make([]byte, len(toCompress))
	l, err = lz4.UncompressBlock(compressed[:l], decompressed, 0)
	if err != nil {
		panic(err)
	}
	fmt.Println("\ndecompressed Data:", string(decompressed[:l]))
}
