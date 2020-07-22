package main

import (
	"io"
	"log"
)

//将读到的内容写入writer中，直到遇到end of file的条件或发生错误。
func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}