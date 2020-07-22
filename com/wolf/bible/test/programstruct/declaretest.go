package programstruct

import "fmt"

//包级范围声明语句
const bolingF = 212.0

func main() {
	var f = bolingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boing point = %g or %g\n", f, c)

	const frezingF, bolingF = 32.0, 212.0
	fmt.Printf("%g = %g\n", frezingF, fToC(frezingF))
	fmt.Printf("%g = %g\n", bolingF, fToC(bolingF))
}

//函数
func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
