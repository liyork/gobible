package main

import "fmt"

//true和false

func main() {
	var s string
	//布尔值可以和&&（AND）和||（OR）操作符结合，并且有短路行为，左边确定就不计算右边
	fmt.Println(s != "" && s[0] == 'x') //安全

	//&&的优先级比||高
}

//布尔值和数字值0或1，不会隐式转换
// btoi returns 1 if b is true and 0 if false.
func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

// itob reports whether i is non-zero.
func itob(i int) bool { return i != 0 }
