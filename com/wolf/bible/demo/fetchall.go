//并发获取
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

//goroutine是一种函数的并发执行方式，而channel是用来在goroutine之间进行参数传递。main函数本身也运行在一个goroutine中

//practiceFetchall1，调研网站的缓存策略：go run fetchall.go http://www.baidu.com http://www.baidu.com
//practiceFetchall2，url中有错误如何处理：go run fetchall.go http://www.baidu.com asdfds http://www.baidu.com
func main() {
	start := time.Now()
	ch := make(chan string) //创建传递string类型参数的channel。
	// 当一个goroutine尝试往channel上做send或receive时，这个goroutine会阻塞在调用处，直到另一个goroutine做了写入、接收值。
	for _, url := range os.Args[1:] {
		go fetch1(url, ch) //start a goroutine,创建一个新的goroutine，并在这个新的goroutine中执行这个函数
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) //receive from channel ch
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch1(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) //send to channel ch
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body) //将数据拷贝到ioutil.Discard，丢弃，只要nbytes返回的字节数
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprint("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
