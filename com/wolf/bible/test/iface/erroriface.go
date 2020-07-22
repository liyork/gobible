package main

import (
	"errors"
	"fmt"
	"syscall"
)

//type error interface {
//	Error() string
//}

type Point struct {
	a string
}

func main() {
	//返回变量的指针，==比较不同
	err := errors.New("xxxx")
	//封装函数
	err = fmt.Errorf("errormsg %v..", err)
	fmt.Println("err:", err)

	//syscall包提供了Go语言底层系统调用API。在多个平台上，它定义一个实现error接口的数字类型Errno
	//Errno是一个系统调用错误的高效表示方式，它通过一个有限的集合进行描述，并且它满足标准的错误接口。
	err = syscall.Errno(2) //接口值，动态类型为syscall.Errno，动态值为2
	fmt.Println(err.Error())
	fmt.Println(err)

	fmt.Println(Point{"a"} == Point{"a"})
	fmt.Println(&Point{"a"} == &Point{"a"})
}
