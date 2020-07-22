package main

import (
	"flag"
	"fmt"
	"time"
)

//标准接口类型flag.Value帮助命令行标记定义新的符号

//创建time.Duration类型的标记变量并且允许用户通过多种用户友好的方式来设置这个变量的大小，这种方式还包括和String方法相同的符号排版形式。
// 这种对称设计使得用户交互良好。
var period = flag.Duration("period", 1*time.Second, "sleep period")

func testFlag() {
	flag.Parse()
	fmt.Printf("Sleeping for %v...", *period)
	time.Sleep(*period)
	fmt.Println()
}

// 为数据celsiusFlag类型定义新的标记符号，实现flag.Value接口的类型

type Celsius float64

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

// *celsiusFlag satisfies the flag.Value interface.
type celsiusFlag struct{ Celsius }

func (f *celsiusFlag) Set(s string) error {
	var uint string
	var value float64
	//从输入s中解析一个浮点数（value）和一个字符串（unit）
	fmt.Sscanf(s, "%f%s", &value, &uint)
	switch uint {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = Celsius(value + 1)
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

// CelsiusFlag defines a Celsius flag with the specified name,
// default value, and usage, and returns the address of the flag variable.
// The flag argument must have a quantity and the unit, e.g., "100C"
func CelsusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	//将标记加入应用的命令行标记集合中
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

func testCelsiusFlag() {
	var temp = CelsusFlag("temp", 20.0, "the temperature")
	flag.Parse()
	fmt.Println("temp:", *temp)
}

func main() {
	//go run flagvaluetest.go
	//go run flagvaluetest.go -period 2000ms
	//testFlag()

	//使用新标记
	//go run flagvaluetest.go
	//go run flagvaluetest.go -temp -18C
	//go run flagvaluetest.go -temp -18k
	testCelsiusFlag()
}
