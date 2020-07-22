package main

import (
	"log"
	"net"
	"os"
)

//演示两个goroutine，一个写一个读

//go run echoClient.go common.go
//a
//b 快速输入可以实验一下同一个连接对于多次输出的处理
func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	//开启goroutine拷贝conn信息到标准输出
	go mustCopy(os.Stdout, conn)
	//接收标准输入拷贝到conn
	mustCopy(conn, os.Stdin)
}
