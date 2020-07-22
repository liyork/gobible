package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

//类型断言是一个使用在接口值上的操作，x.(T)
//一个类型断言检查它操作对象的动态类型是否和断言的类型匹配

func main() {
	//tIsConcrete()
	//tIsInterface()
	//testAssertIfaceNil()
	testAssertTwoResult()
}

func tIsInterface() {
	//如果断言的类型T是一个接口类型，类型断言检查x的动态类型满足T，成功则只能暴露目标接口方法
	var w io.Writer
	w = os.Stdout
	rw := w.(io.ReadWriter) // success: *os.File has both Read and Write
	w = new(ByteCounter)
	//rw = w.(io.ReadWriter) // panic: *ByteCounter has no Read method
	fmt.Println(rw)
}

func tIsConcrete() {
	//断言的类型T是一个具体类型，类型断言检查x的动态类型是否和T相同，结果是x的动态值，类型是T
	var w io.Writer
	w = os.Stdout
	f := w.(*os.File)
	// success: f == os.Stdout
	//c := w.(*bytes.Buffer) // panic: interface conversion: io.Writer is *os.File, not *bytes.Buffer
	fmt.Println(f)
}

//断言操作nil则报错
//go run typeassert.go common.go
func testAssertIfaceNil() {
	var w io.ReadWriter
	q := w.(io.Writer) //panic: interface conversion: interface is nil, not io.Writer
	fmt.Println(q)
}

//有标识返回是否成功，不会panic
func testAssertTwoResult() {
	var w io.Writer = os.Stdout
	f, ok := w.(*os.File) // success:  ok, f == os.Stdout
	fmt.Println(f, ok)
	b, ok := w.(*bytes.Buffer) // failure: !ok, b == nil
	fmt.Println(b, ok)

	if q, ok := w.(*os.File); ok {
		fmt.Println(q)
	}
}
