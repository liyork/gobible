package main

import (
	"fmt"
	"golang.org/x/net/html" //下载
	"os"
)

//函数可以是递归的，这意味着函数可以直接或间接的调用自身
//Go语言使用可变栈，栈的大小按需增加(初始时很小)。使用递归时不必考虑溢出和安全问题

//HTML拥有很多类型的结点，参见html/node.go

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	//go build ../../demo/fetch.go
	//go build recusivetest.go
	//./fetch https://golang.org | ./recusivetest
	printVisit(doc)

	//go build recusivetest.go
	//./fetch https://golang.org | ./recusivetest
	//outline(nil, doc)
}

func printVisit(doc *html.Node) {
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

//深度优先遍历，寻找a标签(links链接)
func visit(links [] string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

//递归遍历整个HTML结点熟，输出树的结构
func outline(stack []string, n *html.Node) { //调用outline会触发stack的拷贝
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) //push tag
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}
