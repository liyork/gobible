package main

import (
	"io"
	"log"
	"net"
	"time"
)

//紧急演示用go并发接收client连接(每个client开启一个goroutine)

//telnet localhost 8000
func main() {
	clock()
}

//顺序执行
func clock() {
	listener, err := net.Listen("tcp", "localhost:8000") //监听
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept() //阻塞直到有新连接被创建
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		//handleConn(conn)
		go handleConn(conn) // handle connections concurrently
	}
}

//处理一个完整的客户端连接，一直写入客户端当前时间，直到写入失败(可能客户端主动断开连接)
func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
