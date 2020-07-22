package main

import (
	"fmt"
	"log"
	"os"
)

//并发web爬虫

// tokens is a couting semaphore used to
// enforce a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 20)

//比如CPU核心数会限制你的计算负载，比如你的硬盘转轴和磁头数限制了你的本地磁盘IO操作频率，比如你的网络带宽限制了你的下载速度上限，
// 或者是你的一个web服务的服务容量上限等等
//增加限速，同一时间调用Extract最多不会超过n次
func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // acquire a token
	list, err := Extract(url)
	<-tokens // release the token
	if err != nil {
		log.Print(err)
	}
	return list
}

func crawl1(url string) []string {
	fmt.Println(url)
	list, err := Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func Extract(s string) ([]string, error) {
	return []string{"url1", "url2"}, nil
}

func main() {
	//可以让程序完成：worklist为空或者没有crawl的goroutine在运行时退出
	//finish1()
	finish2()
}

func finish1() {
	worklist := make(chan []string)
	var n = 1
	// number of pending sends to worklist
	fmt.Println("os.Args[1:]:", os.Args[1:])
	// Start with the command-line arguments.
	go func() { worklist <- os.Args[1:] }()
	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		fmt.Println("range worklist:", list)
		for _, link := range list {
			fmt.Println("link,seen[link]:", link, seen[link])
			if !seen[link] {
				seen[link] = true
				n++ // 准备启动goroutine前+1表示还有数据准备发送到worklist，即有goroutine在运行
				go func(link string) {
					worklist <- crawl(link)
				}(link) // 显示传入，避免循环变量相同引用问题
			}
		}
	}
}

//启用20个goroutine执行crawl，for link := range unseenLinks会引发报错，由于没有人进行放入了，goroutine会卡在这里
func finish2() {
	worklist := make(chan []string)  // list of URLs, may have duplicates
	unseenLinks := make(chan string) // de-duplicated URLs

	go func() { worklist <- os.Args[1:] }()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl1(link)
				go func() { worklist <- foundLinks }() //专用goroutine发送，避免死锁
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to te crawlers.与其他goroutine只通过chan交互，不暴露seen
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}

}
