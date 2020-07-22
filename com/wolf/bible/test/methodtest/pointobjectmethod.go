package main

import (
	"fmt"
	"net/url"
)

//当调用一个函数时，会对其每一个参数值进行拷贝，如果一个函数需要更新一个变量，
// 或者函数的其中一个参数实在太大我们希望能够避免进行这种默认的拷贝，这种情况下我们就需要用到指针了。

//一般会约定如果Point这个类有一个指针作为接收器的方法，那么所有Point的方法都必须有一个指针接收器，即使是那些并不需要这个指针接收器的函数。
//只有类型(Point)和指向他们的指针(*Point)，才是可能会出现在接收器声明里的两种接收器

//type P *int
//func (p P)f(){}// compile error: invalid receiver type，类型本身是指针不能作为接收器

func testPoint() {
	//point类型变量的指针
	r := &Point{1, 2}
	//调用指针类型方法(*Point)ScaleBy
	r.ScaleBy(2) //实参和形参一样都是*Point
	fmt.Println(*r)

	p := Point{1, 2}
	//实参是Point而形参是*Point，所以编译器隐式用&p调用ScaleBy,这种简写方法只适用于“变量”,包括struct里的字段比如p.X，以及array和slice内的元素比如perim[0]
	p.ScaleBy(3)

	//不能通过一个无法取到地址的接收器来调用指针方法，比如临时变量的内存地址就无法获取得到：临时变量没有被变量引用就没有变量地址
	//Point{1, 2}.ScaleBy(3) // cannot call pointer method on composite literal, cannot take the address of composite literal

	pptr := &p
	//实参*Point而形参是Point,则编译器隐式地插入*，解引用得到变量
	pptr.Distance(Point{2, 3}) //等价于：(*pptr).Distance

}

//如果命名类型T的所有方法都是用T类型自己来做接收器(而不是*T)，那么拷贝这种类型的实例就是安全的；调用他的任何一个方法也就会产生一个值的拷贝。
// 比如time.Duration的这个类型，在调用其方法时就会被全部拷贝一份，包括在作为参数传入函数的时候。
// 但是如果一个方法使用指针作为接收器，内部会被改变。
// 比如你对bytes.Buffer对象进行了拷贝，那么可能会引起原始对象和拷贝对象只是别名而已，但实际上其指向的对象是一致的
//声明method的receiver是用指针还是非指针类型，取决于是否要进行拷贝还是要进行修改

//当定义一个允许nil作为接收器值的方法的类型时，在类型前面的注释中指出nil变量代表的意义是很有必要的
// An IntList is a linked list of integers.
// A nil *IntList represents the empty list.
type IntList struct {
	Value int
	Tail  *IntList
}

// sum returns the sum of the list elements.
func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}
	return list.Value + list.Tail.Sum()
}

// 使用url.Values查看是如何应对nil
//方法是属于Value类型而非指针，但map本身就是引用，所以拷贝引用，底层用的还是一个map结构
func testValues() {
	m := url.Values{"lang": {"en"}}
	q := m
	m.Add("item", "1")
	m.Add("item", "2")

	fmt.Println(m.Get("lang"))
	fmt.Println(m.Get("q"))
	fmt.Println(m.Get("item"))
	fmt.Println(m["item"])

	fmt.Println("q:", q)

	m = nil
	fmt.Println(m.Get("item"))
	m.Add("item", "3") // panic: assignment to entry in nil map
}

//go run pointobjectmethod.go commontype.go
func main() {
	//testPoint()
	testValues()
}
