package main

import (
	"fmt"
	"math"
)

//func name(parameter-list) (result-list) {
//    body
//}

//如果一个函数在声明时，包含返回值列表，该函数必须以 return语句结尾
//x,y形参
//函数的形参和有名返回值作为函数最外层的局部变量，被存储在相同的词法块中。
//实参通过值的方式传递，因此函数的形参是实参的拷贝。对形参进行修改不会影响实参。但是，如果实参包括引用类型，
// 如指针，slice(切片)、map、function、channel等类型，实参可能会由于函数的间接引用被修改。
func hypot(x, y float64) float64 { //x，y相同类型则最后声明类型，返回值若被命名则被初始化为0，
	return math.Sqrt(x*x + y*y)
}

//没有函数体的函数声明，表示该函数不是以Go实现的。这样的声明定义了函数标识符
//func Sin(x float64) float

//函数的类型被称为函数的标识符。形参+返回值中变量类型--对应则两个函数有相同的类型和标识符
func add(x int, y int) int   { return x + y }
func sub(x, y int) (z int)   { z = x - y; return }
func first(x int, _ int) int { return x }
func zero(int, int) int      { return 0 }

func main() {
	//3,4调用时传入的实际参数
	fmt.Println(hypot(3, 4))

	fmt.Printf("%T\n", add)
	fmt.Printf("%T\n", sub)
	fmt.Printf("%T\n", first)
	fmt.Printf("%T\n", zero)
}
