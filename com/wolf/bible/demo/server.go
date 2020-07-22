//服务器
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

//go run server1.go
func main() {
	//server1()
	//server2()
	//server3()
	server4()
}

//返回用户访问的url
//浏览器访问：http://localhost:8000/hello
func server1() {
	http.HandleFunc("/", handler)
	//each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the request URL r
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

//统计每个url数量

var mu sync.Mutex
var count int

//浏览器访问：http://localhost:8000/hello
//浏览器访问：http://localhost:8000/ddsf
//浏览器访问：http://localhost:8000/count
//在这些代码的背后，服务器每一次接收请求处理时都会另起一个goroutine，这样服务器就可以同一时间处理多个请求
func server2() {
	http.HandleFunc("/", handler1)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the requested URL.
func handler1(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// counter echoes the number of calls so far.
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

//http://localhost:8000/?q=query
func server3() {
	http.HandleFunc("/", handler2)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//handle echoes the HTTP request.
func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)

	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}

	fmt.Fprintf(w, "Host = %q\n", r.Host)

	fmt.Fprintf(w, "RemoveAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil { //两句合并,err := r.ParseForm();,if err!=nil
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

// go run server.go lissajous.go
//http://localhost:8000/
func server4() {
	http.HandleFunc("/", handler4)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler4(w http.ResponseWriter, r *http.Request) {
	lissajous(w)
}
