package main

import "fmt"

//接口类型是一种抽象的类型。它不会暴露出它所代表的对象的内部值的结构和这个对象支持的基础操作的集合；它们只会展示出它们自己的方法。
// 当看到一个接口类型的值时，你不知道它是什么，唯一知道的就是可以通过它的方法来做什么

//Printf和Sprintf都使用Fprintf，fmt.Fprintf函数没有对具体操作的值做任何假设而是仅仅通过io.Writer接口的约定来保证行为

//fmt.Stringer用于print时调用

func main() {
	var c ByteCounter
	c.Write([]byte("hello")) // 隐式转换为&c
	fmt.Println(c)

	fmt.Println("Get:", c.Get())
	c = 0 //上面操作的是指针，而这里直接赋值，应该用的是int本身特性
	fmt.Println("Get:", c.Get())
	fmt.Printf("%T\n", c) // main.ByteCounter
	var name = "Dolly"
	//*ByteCounter满足io.Writer的约定，我们可以把它传入Fprintf函数中
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)
}
