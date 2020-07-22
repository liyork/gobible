package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

//演示连接来不仅新开goroutine，内部处理时也再开goroutine进行并发处理

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		fmt.Println("handleConn1 main():") //o, 是一个连接。。。
		go handleConn1(conn)
	}
}

//在使用go关键词的同时，需要慎重地考虑net.Conn中的方法在并发地调用时是否安全，事实上对于大多数类型来说也确实不安全
func handleConn1(c net.Conn) {
	fmt.Println("handleConn1():")
	input := bufio.NewScanner(c)
	for input.Scan() {
		fmt.Println("input.Text():", input.Text())
		//echo(c, input.Text(), 1*time.Second) // 对于同一个连接内，多次的input，服务器只能等待之前的处理完再处理之后的
		go echo(c, input.Text(), 1*time.Second) // 一个连接，并发处理多个input
	}
	c.Close()
}

func echo(c net.Conn, shuout string, delay time.Duration) {
	fmt.Println("echo() upper:", strings.ToUpper(shuout))
	fmt.Fprintln(c, "serverreplay:\t", strings.ToUpper(shuout))
	time.Sleep(delay)
	fmt.Println("echo():", shuout)
	fmt.Fprintln(c, "serverreplay:\t", shuout)
	time.Sleep(delay)
	fmt.Println("echo() lower:", strings.ToLower(shuout))
	fmt.Fprintln(c, "serverreplay:\t", strings.ToLower(shuout))
}
