package main

import (
	"fmt"
	"golang.org/x/net/html"
	"strings"
)

//在Go中，函数被看作第一类值（first-class values）：函数像其他值一样，拥有类型，可以被赋值给其他变量，传递给函数，从函数返回。
//对函数值（function value）的调用类似函数调用

func square(n int) int     { return n * n }
func negative(n int) int   { return -n }
func product(m, n int) int { return m * n }

func testFunCompare() {
	var f1 func(int) int
	if f1 != nil { //函数值可以与nil比较
		f1(3)
	}
	//函数值之间是不可比较的，也不能用函数值作为map的key。
}

func testFuncNil() {
	var f1 func(int) int //函数类型的零值是nil
	f1(3)                //调用值为nil的函数会引起panic错误
}

func testFuncValue() {
	f := square
	fmt.Println(f(3))
	// "9"
	f = negative
	fmt.Println(f(3))
	fmt.Printf("%T\n", f)
	//f = product// compile error: can't assign func(int, int) int to func(int) int
}

func add1(r rune) rune { return r + 1 }

//函数值使得我们不仅仅可以通过数据来参数化函数，亦可通过行为
func testParamFunc() {
	fmt.Println(strings.Map(add1, "HAL-900"))
	fmt.Println(strings.Map(add1, "VMS"))
	fmt.Println(strings.Map(add1, "Admix"))
}

//参数化行为，复用逻辑，分离逻辑
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

//func main() {
//	//testFuncValue()
//	//testFuncNil()
//	//testFunCompare()
//	testParamFunc()
//}
