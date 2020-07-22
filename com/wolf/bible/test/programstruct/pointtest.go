package programstruct

import (
	"flag"
	"fmt"
	"strings"
)

//一个变量对应一个保存了变量对应类型值的内存空间。
func testPoint() {
	//普通变量(33)在声明语句创建时被绑定到一个变量名(y)
	var y = 33
	fmt.Println(y)

	x := 1
	//一个指针的值是另一个变量的地址。一个指针对应变量在内存中的存储位置。
	// 通过指针，我们可以直接读或更新对应变量的值，而不需要知道该变量的名字（如果变量有名字的话）。
	var p *int = &x //指向整数变量1的指针，p指针保存了x变量的内存地址
	fmt.Println(*p) //p指针指向的变量的值
	*p = 2          //*p对应一个变量
	fmt.Println(x)

	//任何类型的指针的零值都是nil,指针比较只有当它们指向同一个变量或全部是nil时才相等
	var x1, y1 int
	fmt.Println(&x1 == &x1, &x1 == &y1, &x1 == nil)
}

//false
func returnPoint() {
	fmt.Println(f() == f())
}

func f() *int {
	v := 1
	return &v
}

//指针包含了一个变量的地址
func incr(p *int) int {
	*p++ //增加p指向的变量的值
	return *p
}

func updatePoint() {
	v := 1
	incr(&v)
	fmt.Println(incr(&v))
}

//go run pointtest.go a b
//go run pointtest.go -s / a b
//go run pointtest.go n=1 a b
//go run pointtest.go -help
func testFlag() {
	//返回变量的指针
	var n = flag.Bool("n", false, "omit trailing newline")
	var sep = flag.String("s", " ", "separator")

	flag.Parse()
	fmt.Println(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println("*n is not nil")
	}
}

func main() {
	//testPoint()
	//returnPoint()
	//updatePoint()
	testFlag()
}
