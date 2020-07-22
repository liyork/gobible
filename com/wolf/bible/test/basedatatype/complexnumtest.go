package main

import (
	"fmt"
	"math/cmplx"
)

//complex64和complex128，分别对应float32和float64两种浮点数精度

func main() {
	//构建复数
	var x complex128 = complex(1, 2) // 1+2i
	var y complex128 = complex(3, 4) // 3+4i
	fmt.Println(x * y)               // "(-5+10i)"
	//实部
	fmt.Println(real(x * y)) // "-5"
	//虚部
	fmt.Println(imag(x * y)) // "10"

	//一个浮点数面值或一个十进制整数面值后面跟着一个i，构成一个复数的虚部，实部是0
	fmt.Println(1i * 1i) // "(-1+0i)", i^2 = -1

	//算数运算
	x := 1 + 2i
	y := 3 + 4i
	fmt.Println(x, y)

	//运算
	fmt.Println(cmplx.Sqrt(-1)) // "(0+1i)"
}
