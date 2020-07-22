package main

import (
	"fmt"
	"time"
)

//结构体是一种聚合的数据类型，是由零个或多个任意类型的值聚合成的实体。每个值称为结构体的成员。

//成员顺序不同则表示不同结构体类型
//结构体类型的零值是每个成员都是零值。通常会将零值作为最合理的默认值
type Employee struct {
	ID int //每行对应一个成员
	//Name          string
	//Address       string
	Name, Address string //可以合并
	DoB           time.Time
	Position      string
	Salary        int
	ManagerID     int
}

type tree struct {
	value int
	//t           tree //聚合的值不能包含自身
	left, right *tree //可以包含tree类型指针成员，用于创建递归的数据结构
}

// Sort sorts values in place.
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues appends the elements of t to value in order and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func testStructBase() {
	//变量
	var dilbert Employee
	//成员也是变量
	dilbert.Salary -= 5000
	position := &dilbert.Position
	//取变量地址
	*position = "Senior" + *position
	//解指针，操作变量
	var employeeOfTheMonth *Employee = &dilbert
	//点操作符和指向结构体的指针工作
	employeeOfTheMonth.Position += "(proactive team player)"
}

type Point struct {
	X, Y int
}

func testStructLiteral() {
	//以结构体成员定义的顺序指定字面值，必须全部指定，若结构体成员有变动则不能编译，一般只在定义结构体的包内部使用，或者是在较小的结构体中使用。
	p := Point{1, 2}
	//以成员名字和相应的值来初始化，可以包含部分字段，剩余字段用零值，更常用。
	p1 := Point{X: 1, Y: 2}

	fmt.Println(p, p1)

	//因为结构体通常通过指针处理，所以如下创建并初始化一个结构体变量，并返回结构体的地址
	pp := &Point{1, 2}
	//等同上面
	pp1 := new(Point)
	*pp1 = Point{1, 2}
	fmt.Println(pp, pp1)
}

//结构体作为函数的参数和返回值
func Scale(p Point, factor int) Point {
	return Point{p.X * factor, p.Y * factor}
}

//考虑效率的话，较大的结构体通常会用指针的方式传入和返回
func Bonus(e *Employee, percent int) int {
	return e.Salary * percent / 100
}

//要在函数内部修改结构体成员的话，用指针传入是必须的；
// Go语言中，所有的函数参数都是值拷贝传入的，函数参数将不再是函数调用时的原始变量。
func AwardAnnualRaise(e *Employee) {
	e.Salary = e.Salary * 105 / 100
}

//如果结构体的全部成员都是可以比较的，那么结构体也是可以比较的，那样的话两个结构体将可以使用==或!=运算符进行比较
func testStructCompare() {
	p := Point{1, 2}
	q := Point{2, 1}
	z := Point{1, 2}
	fmt.Println(p.X == q.X && p.Y == q.Y)
	fmt.Println(p == q)

	fmt.Println(p == z) //对成员比较

	//对象的指针不同
	p1 := &Point{1, 2}
	q1 := &Point{1, 2}
	fmt.Println(p1 == q1)
}

type address struct {
	hostname string
	port     int
}

//可比较的结构体类型可以用于map的key类型
func testMapKeyStruct() {
	hits := make(map[address]int)
	hits[address{"golang.org", 443}]++
}

//代表的圆形类型包含了标准圆心的X和Y坐标信息，和一个Radius表示的半径信息
type Circle struct {
	//Center Point //复用
	//只声明一个成员对应的数据类型而不指名成员的名字，结构体类型的匿名成员
	//匿名成员必须是命名的类型或指向一个命名的类型的指针
	Point //匿名成员Point有自己的名字——就是命名的类型名字——但是名字在点操作符中是可选的。
	//Point //因为匿名成员也有一个隐式的名字，因此不能同时包含两个类型相同的匿名成员，这会导致名字冲突。
	//因为匿名成员的名字是由其类型隐式地决定的，所有匿名成员也有可见性的规则约束，大写暴露
	Radius int
}

//轮形除了包含Circle类型所有的全部成员外，还增加了Spokes表示径向辐条的数量
type Wheel struct {
	//Circle Circle //复用
	Circle
	Spokes int
}

func testAnonyStruct() {
	var w Wheel
	//复用，但是访问繁琐
	//w.Circle.Center.X = 8
	//w.Circle.Center.Y = 8
	//w.Circle.Radius = 5
	//w.Spokes = 20

	//匿名成员使得访问简单
	w.X = 8      // equivalent to w.Circle.Point.X = 8
	w.Y = 8      // equivalent to w.Circle.Point.Y = 8
	w.Radius = 5 // equivalent to w.Circle.Radius = 5
	w.Spokes = 20
	//简短的点运算符语法可以用于选择匿名成员嵌套的成员，也可以用于访问它们的方法

	//结构体字面值并没有简短表示匿名成员的语法
	//w = Wheel{8, 8, 5, 20}                       // compile error: unknown fields
	//w = Wheel{X: 8, Y: 8, Radius: 5, Spokes: 20} // compile error: unknown fields}

	//结构体字面值
	w = Wheel{Circle{Point{8, 8}, 5}, 20}
	w = Wheel{
		Circle: Circle{
			Point:  Point{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20, // NOTE: trailing comma necessary here (and at Radius)
	}
	//#表示用和go语言类似的语法打印值，对于结构体类型，打印包含每个成员的名字
	fmt.Printf("%#v\n", w)
	w.X = 42
	fmt.Printf("%v\n", w)
}

func main() {
	//testStructBase()
	//testStructLiteral()
	//fmt.Println(Scale(Point{1, 2}, 5))
	testStructCompare()
	//testMapKeyStruct()
	//testAnonyStruct()
}
