//命令行参数测试
package main

import (
	"fmt"
	"os"
	"strings"
)

//os.Args变量是一个字符串（string）的切片（slice）
//用s[i]访问单个元素，用s[m:n]获取子序列。序列的元素数目为len(s)
//区间索引时，左闭右开。如s[m:n]这个切片，0 ≤ m ≤ n ≤ len(s)，包含n-m个元素
//os.Args[0], 是命令本身的名字。os.Args为os.Arg[0,len(os.Arg)]

func useArgs1() {
	//var声明定义变量，变量在声明时直接初始化为指定值或默认值(0或"")
	var s, sep string
	//":="短变量声明，定义并根据初始值为这些变量赋予适当类型
	//只有for循环一种形式：for [initialization];[condition];[post]{statements}
	for i := 1; i < len(os.Args); i++ { //i++是语句，不是表达式，所以j = i++非法，++只能放在变量名后面
		s += sep + os.Args[i] //"+"连接，+=是赋值运算符
		sep = " "
	}

	fmt.Println(s)
}

func useArgs2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] { //每次循环，range产生index,value，_是空标识
		s += sep + arg //每次都构造新数据，旧数据等待垃圾回收，代价高昂。使用useArgs3
		sep = " "
	}

	fmt.Println(s)
}

func useArgs3() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

//有值需要显示赋值初始化则用第一种，否则用第二种
func varDefine() {
	s1 := ""           //短变量声明，简洁，只能用在函数内部，不能用于包变量
	var s2 string      //依赖默认初始化值
	var s3 = ""        //用得少，除非同时声明多个变量
	var s4 string = "" //当变量类型与初始值类型相同则类型冗余

	fmt.Println(s1, s2, s3, s4)
}

func practiceArgs() {
	fmt.Println("args[0]:", os.Args[0])

	for i, v := range os.Args {
		fmt.Println(i, v)
	}
}

//go run commandargstest.go a b c
func main() {
	//useArgs1()
	//useArgs2()
	//useArgs3()

	practiceArgs()
}
