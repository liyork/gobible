package main

import (
	"log"
	"net"
	"os"
)

//演示main通过channel等待子goroutine执行完  --没成功，os.Stdin好像不执行。。

//go run echoClient1.go common.go
//a
//b
// 当用户关闭了标准输入，mustCopy函数返回，调用close关闭读和写方向的网络连接。
// 关闭网络连接中的写方向的连接将导致server程序收到一个文件（end-of-ﬁle）结束的信号。
// 关闭网络连接中读方向的连接将导致后台goroutine的io.Copy函数调用返回一个
// “read from closed connection”（“从关闭的连接读”）类似的错误
func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		mustCopy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{} //先发送，然后main中才能接收到
	}()
	mustCopy(conn, os.Stdin)
	conn.Close()
	<-done // wait for background goroutine to finish
}
