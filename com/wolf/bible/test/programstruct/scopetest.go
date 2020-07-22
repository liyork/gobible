package programstruct

import (
	"fmt"
	"log"
	"os"
)

//一个声明语句将程序中的实体和一个名字关联，比如一个函数或一个变量

func f() {}

var g = "g"

func f1() int      { return 1 }
func g1(x int) int { return 2 }

func main() {
	//test1()
	//test2()
	test3()
	log.Printf("main - Working directory = %s", cwd)
}

func test1() {
	//覆盖
	f := "f"
	fmt.Println(f)
	// "f"; local var f shadows package-level func f
	fmt.Println(g)
	// "g"; package-level var
	//fmt.Println(h) // compile error: undefined: h
}

func test2() {
	//作用域在if内
	if x := f1(); x == 0 {
		fmt.Println(x)
	} else if y := g1(x); x == y {
		fmt.Println(x, y)
	} else {
		fmt.Println(x, y)
	}
	//fmt.Println(x, y) // compile error: x and y are not visible here
}

var cwd string

func test3() {
	var err error
	cwd, err = os.Getwd()
	//cwd, err := os.Getwd() //短变量:=作用范围是局部的，将cwd和err重新声明为新的局部变量，不会使用包变量
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
	log.Printf("Working directory = %s", cwd)
}
