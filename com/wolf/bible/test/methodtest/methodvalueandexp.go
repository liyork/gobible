package main

import (
	"fmt"
	"time"
)

//经常选择一个方法，并且在多个表达式里执行，比如常见的p.Distance()形式，实际上将其分成两步来执行也是可能的。
//p.Distance叫作“选择器”，选择器会返回一个方法"值"————一个将方法绑定到特定接收器变量的函数。之后进行调用

func testMethodValue() {
	p := Point{1, 2}
	q := Point{4, 6}

	distanceFromP := p.Distance   // method value，绑定到了p上
	fmt.Println(distanceFromP(q)) //调用，之前已经绑定所以调用的就是p的方法
	var origin Point
	fmt.Println(distanceFromP(origin))

	scaleP := p.ScaleBy // method value，编译器自动转换成了&p
	fmt.Printf("%T\n", scaleP)
	scaleP(2)  // p becomes (2, 4)
	scaleP(3)  //      then (6, 12)
	scaleP(10) //      then (60, 120)

	//使用方法值作为参数
	time.AfterFunc(10*time.Second, p.test)
}

//当调用一个方法时，与调用一个普通的函数相比，我们必须要用选择器(p.Distance)语法来指定方法的接收器
func testMethodExpress() {
	p := Point{1, 2}
	q := Point{4, 6}

	distance := Point.Distance   // method expression
	fmt.Println(distance(p, q))  //将第一个参数用作接收器，用通常的函数调用方式调用
	fmt.Printf("%T\n", distance) // "func(main.Point, main.Point) float64"
	fmt.Println(p)

	scale := (*Point).ScaleBy
	scale(&p, 2)
	fmt.Println(p)
	fmt.Printf("%T\n", scale) // "func(*main.Point, float64)"
}

func (p Point) Add(q Point) Point { return Point{p.X + q.X, p.Y + q.Y} }
func (p Point) Sub(q Point) Point { return Point{p.X - q.X, p.Y - q.Y} }

//根据一个变量来决定调用同一个类型的哪个函数，不用写死调用
func (path Path) TranslatedBy(offset Point, add bool) {
	var op func(p, q Point) Point
	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}
	for i := range path {
		// Call either path[i].Add(offset) or path[i].Sub(offset).
		path[i] = op(path[i], offset)
	}
}

func main() {
	//testMethodValue()
	//testMethodExpress()

	perim := Path{
		{1, 1},
		{5, 1},
		{2, 2},
	}
	perim.TranslatedBy(Point{1, 2}, true)
	fmt.Println(perim)
}
