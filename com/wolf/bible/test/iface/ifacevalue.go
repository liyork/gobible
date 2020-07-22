package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"time"
)

//接口值，由两个部分组成，一个具体的类型和那个类型的值。被称为接口的动态类型和动态值
//Go语言这种静态类型的语言，类型是编译期的概念；因此一个类型不是一个值
//在我们的概念模型中，一些提供每个类型信息的值被称为类型描述符，比如类型的名称和方法
//在一个接口值中，类型部分代表与之相关类型的描述符

func testIfaceValue() {
	//定义变量w，go中变量总是被一个定义明确的值初始化。接口的零值即类型和值的部分都是nil
	//一个接口值基于它的动态类型被描述为空或非空，所以第一行的w是一个空的接口值。
	var w io.Writer
	//w.Write([]byte("hello"))//调用一个空接口值上的任意方法都会产生panic
	fmt.Printf("%T\n", w)
	// 接口值的动态类型，"<nil>"
	//将一个*File类型的值赋给变量w，赋值过程调用了一个具体类型到接口类型的隐式转换
	//这个接口值的动态类型被设为*os.File指针的类型描述符，它的动态值持有os.Stdout的拷贝
	w = os.Stdout
	w.Write([]byte("hello"))
	//使得(*os.File).Write方法被调用
	fmt.Printf("%T\n", w)
	// "*os.File"
	//通常在编译期，我们不知道接口值的动态类型是什么，所以一个接口上的调用必须使用动态分配。
	// 因为不是直接进行调用，所以编译器必须把代码生成在类型描述符的方法Write上，然后间接调用那个地址。
	// 这个调用的接收者是一个接口动态值的拷贝，os.Stdout
	os.Stdout.Write([]byte("hello"))
	//等同于上面
	//给接口值赋了一个*bytes.Buffer类型的值
	//动态类型是*bytes.Buffer并且动态值是一个指向新分配的缓冲区的指针
	w = new(bytes.Buffer)
	w.Write([]byte("hello"))
	//类型描述符是*bytes.Buffer，调用了(*bytes.Buffer).Write方法，接收者是该缓冲区的地址
	fmt.Printf("%T\n", w)
	// "*bytes.Buffer"
	w = nil
	//一个接口值可以持有任意大的动态值，从概念上讲，不论接口值多大，动态值总是可以容下它
	var x interface{} = time.Now()
	//接口值用==比较，相等则都是nil或者他们的动态类型相同并且动态值也根据这个动态类型的==操作相等，
	// 因为接口值可比较，所以用在map的key或者switch语句
	//若接口值的动态类型相同，但这个动态类型是不可比较的(如切片)，比较就失败并且panic
	var y interface{} = []int{1, 2, 3}
	fmt.Println(y == y)
	// panic: comparing uncomparable type []int
	//其它类型要么是安全的可比较类型（如基本类型和指针）要么是完全不可比较的类型（如切片，映射类型，和函数），
	// 但是在比较接口值或者包含了接口值的聚合类型时，我们必须要意识到潜在的panic，以及map、switch。
	fmt.Println(x)
}

//一个不包含任何值的nil接口值和一个刚好包含nil指针的接口值是不同的
func testIfaceNil() {
	var buf *bytes.Buffer
	//调用时给f函数的out参数赋了一个*bytes.Buffer的空指针，所以out的动态值是nil，动态类型是*bytes.Buffer
	f(buf)
}

func f(out io.Writer) {
	if out != nil { //失败,out变量是一个包含空指针值的非空接口
		out.Write([]byte("done!\n")) // panic: nil pointer dereference
		//解决方案就是参数改成*bytes.Buffer或者buf声明改成io.Writer，这样就不会调用f时动态类型有值
	}
}

func main() {
	//testIfaceValue()
	testIfaceNil()
}
