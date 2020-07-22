package main

import (
	"fmt"
	"math"
)

//float32和float64

func main() {
	//testFloatOverflow()
	//testLiteral()
	//testPrintFloat()
	testInfinity()
}

//通常应该优先使用float64类型，因为float32类型的累计计算误差很容易扩散，并且float32能精确表示的正整数并不是很大(23位)
func testFloatOverflow() {
	var f float32 = 16777216
	// 1 << 24
	fmt.Println(f == f+1)
	// "true"!
}

func testLiteral() {
	const e = 2.71828 // (approximately)
	//很小或很大的数最好用科学计数法书写，通过e或E来指定指数部分：
	const Avogadro = 6.02214129e23 // 阿伏伽德罗常数
	const Planck = 6.62606957e-34  // 普朗克常数
}

//%g参数打印浮点数，对于表格使用%e或%f
func testPrintFloat() {
	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d e^x = %8.3f\n", x, math.Exp(float64(x)))
	}
}

//特殊值
func testInfinity() {
	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z) // "0 -0 +Inf -Inf NaN"
}

//如果一个函数返回的浮点数结果可能失败，最好的做法是用单独的标志报告失败，而不是返回math.NaN，因为math.NaN和任何数都不相等
func compute() (value float64, ok bool) {
	// ...
	var failed bool
	if failed {
		return 0, false
	}
	return 1, true
}
