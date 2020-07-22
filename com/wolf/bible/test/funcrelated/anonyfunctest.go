package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
	"sort"
)

//拥有函数名的函数只能在包级语法块中被声明
//通过函数字面量（function literal），我们可绕过这一限制，在任何表达式中表示一个函数值
//函数值字面量是一种表达式，它的值被称为匿名函数（anonymous function）

//通过这种方式定义的函数可以访问完整的词法环境（lexical environment），这意味着在函数中定义的内部函数可以引用该函数的变量
//squares返回一个匿名函数
func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

//squares的例子证明，函数值不仅仅是一串代码，还记录了状态
//在squares中定义的匿名内部函数可以访问和更新squares中的局部变量，这意味着匿名函数和squares中，存在变量引用
//这就是函数值属于引用类型和函数值不可比较的原因。
// Go使用闭包（closures）技术实现函数值，Go程序员也把函数值叫做闭包。
func testSquares() {
	f := squares() //对squares的一次调用会生成一个局部变量x并返回一个匿名函数。
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

//这类问题被称作拓扑排序
// prereqs记录了每个课程的前置课程
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

//给定一些计算机课程，每个课程都有前置课程，只有完成了前置课程才可以开始当前课程的学习；
// 我们的目标是选择出一组课程，这组课程必须确保按顺序学习时，能全部被完成。
//这类问题被称作拓扑排序。从概念上说，前置条件可以构成有向图。图中的顶点表示课程，边表示课程间的依赖关系。
//topoSort用深度优先搜索了整张图，获得了符合要求的课程序列
func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)

	//当匿名函数需要被递归调用时，必须首先声明一个变量,再将匿名函数赋值给这个变量进行绑定
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	visitAll(keys)
	return order
}

//网页抓取的核心问题就是如何遍历图
func Extract(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	//匿名函数负责将新连接(a标签，links)添加到切片中
	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				links = append(links, link.String())
			}
		}
	}

	forEachNode(doc, visitNode, nil)
	return links, nil
}

//使用广度优先算法
// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
			}
			worklist = append(worklist, f(item)...)
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

//捕获迭代变量
//例如：创建一些目录然后删除，使用函数值
func testCatechVar() {
	tempDirs := [2]string{"/a", "b"}

	var rmdirs []func()
	//for循环引入新的词法块。在该循环中生成的所有函数值都共享相同的循环变量。
	for _, d := range tempDirs {
		dir := d               //Necessary
		os.MkdirAll(dir, 0755) // creates parent directories too
		rmdirs = append(rmdirs, func() {
			//os.RemoveAll(d) //错误，函数值中记录的是循环变量的内存地址，而不是循环变量某一时刻的值，当for执行完后就是最后一个元素地址了
			os.RemoveAll(dir) //基于上述问题，引入局部变量dir作为副本。
		})
	}
	// ...to some work...
	//等待上面循环结束后再统一执行函数
	for _, rmdir := range rmdirs {
		rmdir()
	}
}

func main() {
	//函数字面量允许我们在使用函数时，再定义它
	//strings.Map(func(r rune) rune { return r + 1 }, "HAL-9000")
	//testSquares()

	//for i, course := range topoSort(prereqs) {
	//	fmt.Printf("%d:\t%s\n", i+1, course)
	//}

	//go build anonyfunctest.go funcvaluetest.go
	//./anonyfunctest https://golang.org
	breadthFirst(crawl, os.Args[1:])
}
