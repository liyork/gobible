package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

//聊天服务器，广播消息。

// go run chardemo.go
// go run echoClient.go common.go
// go run echoClient.go common.go
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		//每个新建连接使用一个新goroutine
		go handleConn2(conn)
	}
}

type client chan<- string // an outgoing message channel

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

func broadcaster() {
	clients := make(map[client]bool) // all connected clients

	for {
		//监听消息、客户端来、客户端离开
		select {
		case msg := <-messages:
			// Broadcast incoming message to call
			// client's outgoing message channels.
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

func handleConn2(conn net.Conn) {
	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	//欢迎信息
	ch <- "Yor are " + who
	//广播消息
	messages <- who + " has arrived"
	//保存客户端
	entering <- ch

	input := bufio.NewScanner(conn)
	//读取客户端发来信息,一直读取
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}

	leaving <- ch
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
