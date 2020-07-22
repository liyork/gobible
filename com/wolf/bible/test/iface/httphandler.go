package main

import (
	"fmt"
	"log"
	"net/http"
)

//http.Handler

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

//http://localhost:8000/list
//http://localhost:8000/price?item=socks
//http://localhost:8000/price?item=shoes
func main() {
	db := database{"shoes": 50, "socks": 5}
	//HandlerFunc(handler)，HandlerFunc和db.list具有相同方法，进行转换。
	// HandlerFunc是一个实现了接口http.Handler方法的函数类型，ServeHTTP方法调用了它本身的函数，
	// HandlerFunc是一个让函数值满足一个接口的适配器，这个函数和接口仅是方法有相同的函数签名
	//是一个实现了handler类似行为的函数，但是没有方法，不满足http.Handler接口并且不能直接传给mux.Handle
	list := db.list //方法值，类型func(w http.ResponseWriter, req *http.Request)
	http.HandleFunc("/list", list)
	http.HandleFunc("/price", db.price)
	//web服务器在一个新的协程中调用每一个handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil)) //DefaultServerMux
}
