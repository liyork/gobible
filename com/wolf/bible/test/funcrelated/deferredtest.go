package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"sync"
	"time"
)

//defer语句经常被用于处理成对的操作，如打开、关闭、连接、断开连接、加锁、释放锁。

func title(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close() //函数被延迟执行。保证title执行完毕后一定执行.

	// Check Content-Type is HTML (e.g., "text/html;charset=utf-8).
	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
		return fmt.Errorf("%s has type %s, not text/html", url, ct)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	fmt.Println("doc:", doc)
	return nil
}

//互斥锁使用defer
var mu sync.Mutex
var m = make(map[string]int)

func lookup(key string) int {
	mu.Lock()
	defer mu.Unlock()

	return m[key]
}

//用defer记录耗时
func bigSlowOperation() {
	defer trace("bigSlowOperation")() //trace("bigSlowOperation")返回函数值，然后defer 函数值()会被延迟执行
	fmt.Println("business logic start")
	time.Sleep(10 * time.Second)
	fmt.Println("business logic end")
}

//返回函数值，有start的引用
func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

//defer查看函数返回值，多return的函数有用
func double(x int) (result int) {
	defer func() { fmt.Printf("double(%d) = %d\n", x, result) }()
	return x + x
}

//用defer修改返回值
func triple(x int) (result int) {
	defer func() { result += x }()
	return double(x)
}

//注意for中的defer
func deferInFor() error {
	filenames := []string{}
	for _, filename := range filenames {
		f, err := os.Open(filename)
		if err != nil {
			return err
		}
		//不能执行，因为deferInFor方法没有退出
		//解决方案：将for中的语句单独提取到一个方法中
		defer f.Close()
		fmt.Println(f)
	}
	return nil
}

//使用defer进行关闭资源
func fetch(url string) (file string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(f, resp.Body)
	//在关闭文件时，没有对f.close采用defer机制，因为这会产生一些微妙的错误。
	// 许多文件系统，尤其是NFS，写入文件时发生的错误会被延迟到文件关闭时反馈。
	// 如果没有检查文件关闭时的反馈信息，可能会导致数据丢失，而我们还误以为写入操作成功。
	// 如果io.Copy和f.close都失败了，我们倾向于将io.Copy的错误信息反馈给调用者，因为它先于f.close发生，更有可能接近问题的本质。
	if closeErr := f.Close(); err == nil { //prefer error from Copy, if any.
		err = closeErr
	}
	return local, n, err

}

func main() {
	//bigSlowOperation()
	//double(4)
	testInvokeMethodDefer()
}

//只有真正退出方法时才会调用defer，内部有函数调用没表示方法退出，不能执行
func testInvokeMethodDefer() {
	defer func() { fmt.Println("defer in testInvokeMethodDefer") }()
	fmt.Println("invoke testInvokeMethodDefer1 start")
	testInvokeMethodDefer1()
	fmt.Println("invoke testInvokeMethodDefer1 end")
}

func testInvokeMethodDefer1() {
	defer func() { fmt.Println("defer in testInvokeMethodDefer1") }()
	fmt.Println("testInvokeMethodDefer1")
}
