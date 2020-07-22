package main

import (
	"fmt"
	"math"
)

//在函数声明时，在其名字之前放上一个变量，即是一个方法。这个附加的参数会将该函数附加到这种类型上，即相当于为这种类型定义了一个独占的方法。

// traditional function
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			//编译器会根据方法的名字以及接收器来决定具体调用的是哪一个函数
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

//方法比之函数的一些好处：方法名可以简短，不用携带报名，仅仅用短的变量名
func main() {
	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println(Distance(p, q)) // function call
	fmt.Println(p.Distance(q))  // method call，p.Distance表达式叫做选择器，会选择合适的对应p这个对象的Distance方法执行
	fmt.Println(q.Distance(p))

	perim := Path{
		{1, 1},
		{5, 1},
		{2, 2},
	}
	fmt.Println(perim.Distance())
}
