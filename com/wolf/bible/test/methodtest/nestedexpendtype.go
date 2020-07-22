package main

import (
	"fmt"
	"image/color"
	"sync"
)

type ColoredPoint struct {
	Point //内嵌，属性和方法都被引入
	Color color.RGBA
}

type ColoredPoint1 struct {
	*Point // 内嵌命名类型的指针
	Color  color.RGBA
}

//仅仅是组合has a，而没有继承概念，没有is a概念
func testNestType() {
	//使用内嵌变量
	var cp ColoredPoint
	cp.X = 1
	fmt.Println(cp.Point.X)
	cp.Point.Y = 2
	fmt.Println(cp.Y)

	//使用内嵌方法
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = ColoredPoint{Point{1, 1}, red}
	var q = ColoredPoint{Point{5, 4}, blue}
	fmt.Println(p.Distance(q.Point))
	p.ScaleBy(2)
	q.ScaleBy(2)
	fmt.Println(p.Distance(q.Point))
}

func testNestPointType() {
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	p := ColoredPoint1{&Point{1, 1}, red}
	q := ColoredPoint1{&Point{5, 4}, blue}
	fmt.Println(p.Distance(*q.Point))
	q.Point = p.Point
	p.ScaleBy(2)
	fmt.Println(*p.Point, *q.Point)
}

//但是多亏了内嵌，有些时候我们给匿名struct类型来定义方法也有了手段?
//将两个包级别的变量放在了cache这个struct一组内
var cache = struct { //匿名结构体
	sync.Mutex
	mapping map[string]string
}{ //初始化?
	mapping: make(map[string]string),
}

func Lookup(key string) string {
	cache.Lock()
	v := cache.mapping[key]
	cache.Unlock()
	return v
}

//go run nestedexpendtype.go commontype.go
func main() {
	//testNestType()
	//testNestPointType()

}
