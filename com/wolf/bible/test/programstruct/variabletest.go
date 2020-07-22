package programstruct

import (
	"fmt"
	"os"
)

//var声明语句可以创建一个特定类型的变量，然后给变量附加一个名字，并且设置变量的初始值
//var 变量名字 [类型] [= 表达式]

//包级别声明的变量会在main入口函数执行前完成初始化
var a int

func declareVar() {
	//局部变量将在声明语句被执行到的时候完成初始化
	var s string
	//声明则自动赋值
	fmt.Println(s)
	//""
	//声明多个变量
	var i, j, k int
	var b, f, s1 = true, 2.3, "four"
	fmt.Println(s, i, j, k, b, f, s1)
}

func shortDeclareVar() {
	//函数内部，简短变量声明,声明和初始化局部变量,变量的类型根据表达式来自动推导。简洁和灵活
	//var形式的声明语句往往是用于需要显式指定变量类型地方，或者因为变量稍后会被重新赋值而初始值无关紧要的地方
	a := 1
	var boling float64 = 100
	var names []string
	var err error

	//多变量声明，这种方式应该限制只在可以提高代码可读性的地方使用，比如for语句的循环的初始化语句部分
	i, j := 0, 1
	fmt.Println("a:", a)
	//交换i和j的值
	i, j = j, i

	//声明in,err
	in, err := os.Open("a.txt")
	//声明out，赋值err，简短变量声明语句中必须至少要声明一个新的变量
	out, err := os.Create("b.txt")

	fmt.Println(boling, names, err, in, out)
}

func newVar() {
	p := new(int) //创建int类型的匿名变量，初始化为零值，返回变量地址，类型是*int，赋值给p
	fmt.Println(*p)
	*p = 2
	fmt.Println(*p)

	//与new相同，只不过引入了一个指向变量的临时变量p
	var p1 int
	fmt.Println(&p1)

	fmt.Println(new(int), new(int)) //false，每次都是新的变量的地址

	//对于结构体来说，直接用字面量语法创建新变量的方法会更灵活
}

//由于new只是一个预定义的函数，它并不是一个关键字，因此我们可以将new名字重新定义为别的类型
//new被定义为int类型的变量名，因此在delta函数内部是无法使用内置的new函数的
func delta(old, new int) int {
	return new - old
}

//包级声明的变量的生命周期和整个程序的运行周期是一致的
//局部变量的声明周期则是动态的：每次从创建一个新变量的声明语句开始，直到该变量不再被引用为止，然后变量的存储空间可能被回收。
var a1 = 1

// 函数的参数变量(a,b)和返回值变量(c)都是局部变量。它们在函数每次被调用的时候创建
func testVarLife(a, b int) (c int) {
	return a + b
}

func newLine() {
	testVarLife(1, //遇到逗号不会合并行
		2, //最后插入逗号不会报错，为了下面小括号能另起一行，和大括号保持一致
	)
}

//Go语言的自动垃圾收集器是如何知道一个变量是何时可以被回收的呢？
// 基本的实现思路是，从每个包级的变量和每个当前运行函数的每一个局部变量开始，通过指针或引用的访问路径遍历，是否可以找到该变量。
// 如果不存在这样的访问路径，那么说明该变量是不可达的。一个变量的有效周期只取决于是否可达

//编译器会自动选择在栈上还是在堆上分配局部变量的存储空间
var global *int

func f1() {
	var x int
	x = 1 //x对应的变量1必须在堆上分配，因为它在函数退出后依然可以通过包级global变量找到，这个局部变量从函数中逃逸了
	global = &x
}

func g() { //函数返回时，变量*y不可达，可以被回收。编译器可以选择在栈上分配*y的存储空间
	y := new(int)
	*y = 1
}

func main() {
	//declareVar()
	//shortDeclareVar()
	//newVar()
	delta(1, 2)

}
