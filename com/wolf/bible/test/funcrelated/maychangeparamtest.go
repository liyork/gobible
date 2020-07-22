package main

import (
	"fmt"
	"os"
)

//参数数量可变的函数称为为可变参数函数

//在参数列表的最后一个参数类型之前加上省略符号“...”
func sum(vals ...int) int { //函数体中，vals被看做是类型为[]int的切片。仅仅是看做
	fmt.Printf("%T\n", vals) // "func(...int)"
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

//函数名的后缀是f是一种通用的命名规范，代表该可变参数函数可以接收Printf风格的格式化字符串
func errorf(linenum int, format string, args ...interface{}) { //interfac{}表示任意类型
	fmt.Fprintf(os.Stderr, "Line %d: ", linenum)
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintln(os.Stderr)
}

func main() {
	fmt.Println(sum())
	//调用者隐式创建一个数组并将原始参数复制到数组中，再把数组的一个切片作为参数传给被调函数
	fmt.Println(sum(3))
	fmt.Println(sum(1, 2, 3))

	//原始参数已经是切片类型则在最后加上...
	values := []int{1, 2, 3, 4}
	fmt.Println(sum(values...))

	linenum, name := 12, "count"
	errorf(linenum, "undefined: %s", name)
}
