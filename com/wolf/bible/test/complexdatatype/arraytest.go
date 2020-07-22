package main

import (
	"crypto/sha256"
	"fmt"
)

//数组是一个由固定长度的特定类型元素组成的序列，一个数组可以由零个或多个元素组成。

func main() {
	//testArrayBase()
	//testArrayLiteral()
	//testArrayInit()
	//testArrayEqual()
	//testArrayPassCopy()
	testArrayPassPoint()

}

func testArrayLiteral() {
	//默认情况下，数组的每个元素都被初始化为元素类型对应的零值，对于数字类型来说就是0
	var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2}
	fmt.Println(r[2])
	//...表示数组的长度根据初始化值个数确定
	q = [...]int{1, 2, 3}
	fmt.Printf("%T\n", q)
	//数组的长度是数组类型的一个组成部分，因此[3]int和[4]int是两种不同的数组类型。
	// 数组的长度必须是常量表达式，因为数组的长度需要在编译阶段确定。
	q = [3]int{1, 2, 3}
	//q = [4]int{1, 2, 3, 4} // compile error: cannot assign [4]int to [3]int
}

func testArrayBase() {
	var a [3]int
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])
	// Print the indices and elements
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}
	// Print the elements only.
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}
}

type Currency int

const (
	USD Currency = iota // 美元
	EUR                 // 欧元
	GBP                 // 英镑
	RMB                 // 人民币
)

func testArrayInit() {
	//通过索引初始化，顺序不重要了
	test := [3]int{1: 2, 0: 1, 2: 3}
	fmt.Println(test)

	symbol := [...]string{EUR: "€", USD: "$", GBP: "￡", RMB: "￥"}
	fmt.Println(RMB, symbol[RMB]) // "3 ￥"

	//定义了一个含有100个元素的数组r，最后一个元素被初始化为-1，其它元素都是用0初始化。
	r := [...]int{99: -1}
	fmt.Println(r)
}

//只有当两个数组的所有元素都是相等的时候数组才是相等的
func testArrayEqual() {
	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(a == b, a == c, b == c) // "true false false"
	d := [3]int{1, 2}
	//fmt.Println(a == d) // compile error: cannot compare [2]int == [3]int
	fmt.Println(d)

	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	//%x以16进制格式打印，%t打印boo类型数据，%T打印值对应的数据类型
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
}

func testArrayPassCopy() {
	arr := [3]int{1, 2, 3}
	testInterArrayPass(arr)
	fmt.Println("testArrayPass:", arr)
}

//当调用一个函数的时候，函数的每个调用参数将会被赋值给函数内部的参数变量，所以函数参数变量接收的是一个复制的副本。
// 因为函数参数传递的机制导致传递大的数组类型将是低效的，并且对数组参数的任何的修改都是发生在复制的数组上，并不能直接修改调用时原始的数组变量。
func testInterArrayPass(arr [3]int) {
	arr[1] = 4
	fmt.Println("testInterArrayPass:", arr)
}

func testArrayPassPoint() {
	arr := [3]int{1, 2, 3}
	testInterArrayPassPoint(&arr)
	fmt.Println("testArrayPassPoint:", arr)
}

//传入一个数组指针，调用拷贝的仅仅是指针，通过指针对数组的任何修改都可以直接反馈到调用者
func testInterArrayPassPoint(arr *[3]int) {
	arr[1] = 4
	fmt.Println("testInterArrayPassPoint:", *arr)
}
