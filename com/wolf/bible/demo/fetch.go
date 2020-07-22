package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func fetch() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1) //终止进程，返回错误码
		}
		b, err := ioutil.ReadAll(resp.Body) //从body中读取到全部内容
		resp.Body.Close()                   //关闭resp的Body流
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch:reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}

//使用io.Copy直接写入stdout中，避免申请缓冲区b
func practiceFetch1() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "practiceFetch1:io.Copy %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}

//url判断并添加http://
func practiceFetch2() {
	for _, url := range os.Args[1:] {
		fmt.Println("pre url:", url)
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		fmt.Println("converted url:", url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "practiceFetch1:io.Copy %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}

//打印resp.Status
func practiceFetch3() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "practiceFetch1:io.Copy %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Println("resp.Status:", resp.Status)
	}
}

//go run fetch.go http://www.baidu.com 正确
//go run fetch.go http://bad.gopl.io 错误
func main() {
	//fetch()
	//practiceFetch1()
	//practiceFetch2()
	practiceFetch3()
}
