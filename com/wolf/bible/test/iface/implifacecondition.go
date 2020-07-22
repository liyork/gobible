package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

//一个类型如果拥有一个接口需要的所有方法，那么这个类型就实现了这个接口

func assignment() {
	//变量属于某个接口
	var w io.Writer
	w = os.Stdout
	w = new(bytes.Buffer)
	//w = time.Second //错误，time.Duration没有实现Write方法
	var rwc io.ReadWriteCloser
	rwc = os.Stdout
	//rwc = new(bytes.Buffer)//缺少Close方法
	//接口属于某个接口
	w = rwc
	//rwc = w//缺少Read/Close方法
	fmt.Println(w)
}

//类型持有某些方法，指针类型持有另一些方法

type IntSet struct{ /* ... */ }

func (*IntSet) String() string { return "" }

func typeAndPointHasDiffMethod() {
	//临时变量没有变量指向故没有指针指向地址
	var _ = IntSet{}.String()
	// cannot call pointer method on IntSet literal, cannot take the address of IntSet literal
	var s IntSet
	var _ = s.String()
	//编译器转换成&s
	var _ fmt.Stringer = &s
	//var _ fmt.Stringer = s//类型IntSet并没有实现String，是*IntSet类型实现了String方法
}

//接口类型封装和隐藏具体类型和它的值。即使具体类型有其它的方法也只有接口类型暴露出来的方法会被调用到
func exportMethod() {
	os.Stdout.Write([]byte("hello")) // *os.File has Write method
	os.Stdout.Close()                // *os.File has Close method

	var w io.Writer
	w = os.Stdout
	w.Write([]byte("hello")) // io.Writer has Write method
	//w.Close()// compile error: io.Writer lacks Close method
}

//interface{}类型，它没有任何方法，被称为空接口类型。空接口类型对实现它的类型没有要求，所以我们可以将任意一个值赋给空接口类型。
func testInterface() {
	var any interface{}
	any = true
	any = 12.34
	any = "hello"
	any = map[string]int{"one": 1}
	any = new(bytes.Buffer)

	fmt.Println(any)

	// 编译器断言：*bytes.Buffer must satisfy io.Writer
	var _ io.Writer = (*bytes.Buffer)(nil)
}

//因为接口实现只依赖于判断的两个类型的方法，所以没有必要定义一个具体类型和它实现的接口之间的关系
//非空的接口类型比如io.Writer经常被指针类型实现，多个接口隐式的给接收者带来变化。
//一个结构体的指针是非常常见的承载方法的类型。一个具体的类型可能实现了很多不相关的接口
//把每个抽象的特点用接口来表示。一些特性对于所有的产品都是共通的，其它的一些特性只对特定类型的产品才有。
//这些接口不止是一种有用的方式来分组相关的具体类型和表示他们之间的共同特定。后续发现其他共性则可以定义一个新接口来而不必对已经存在的类型做改变
//每一个具体类型的组基于它们相同的行为可以表示成一个接口类型

func main() {
	//assignment()
	//typeAndPointHasDiffMethod()
	//exportMethod()
	testInterface()
}
