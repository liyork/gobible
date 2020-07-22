package main

import "math"

type Point struct {
	X, Y float64
}

// same thing, but as a method of the Point type
//附加参数p，叫做方法的接收器(receiver)，go中不用this或self作为接收器，可以任意选择接收器的名字
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Point) test() {
}

//要更新对象p必须用指针，否则就拷贝对象修改不了原值
func (p *Point) ScaleBy(factor float64) { //方法的名字是(*Point).ScaleBy
	p.X *= factor
	p.Y *= factor
}

//每种类型都有其各自的命名空间
//Path是一个命名的slice类型，go能够给任意类型定义方法
//go为一些简单的数值、字符串、slice、map来定义一些附加行为很方便
//可以给同一个包内的任意命名类型定义方法，只要这个命名类型的底层类型不是指针或者interface。
type Path []Point //底层类型是[]Point这个slice，Path就是命名类型
