package main

import (
	"fmt"
	"math"
	"time"
)

//常量表达式的值在编译期计算，而不是在运行期。每种常量的潜在类型都是基础类型：boolean、string或数字。

func main() {
	//testConst()
	//testIota()
	testUnType()
}

func testConst() {
	//一个常量的声明语句定义了常量的名字，常量的值不可修改
	const pi = 3.14159
	// approximately; math.Pi is a better approximation
	//批量
	const (
		e   = 2.718281
		pi2 = 3.141592
	)
	//声明包含一个类型和一个值
	const noDelay time.Duration = 0
	const timeout = 5 * time.Minute
	//%T打印参数类型
	fmt.Printf("%T %[1]v\n", noDelay)
	// "time.Duration 0"
	fmt.Printf("%T %[1]v\n", timeout)
	// "time.Duration 5m0s"
	fmt.Printf("%T %[1]v\n", time.Minute)
	// "time.Duration 1m0s"
	const (
		a = 1 //不可省略初始值
		b     //使用前面值1
		c = 2
		d
	)
	fmt.Println(a, b, c, d)
	// "1 1 2 2"
}

//使用iota常量生成器
func testIota() {
	const (
		Sunday int = iota //0
		Monday            //1
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	fmt.Println(Monday, Saturday)

	fmt.Printf("%b,%b,%b\n", FlagUp, FlagBroadcast, FlagLoopback)

	//每个常量对应表达式1 << iota，是连续的2的幂，分别对应一个bit位置。使用这些常量可以用于测试、设置或清除对应的bit位的值：
	var v Flags = FlagMulticast | FlagUp
	fmt.Printf("%b %t\n", v, IsUp(v)) // "10001 true"
	TurnDown(&v)
	fmt.Printf("%b %t\n", v, IsUp(v)) // "10000 false"
	SetBroadcast(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))   // "10010 false"
	fmt.Printf("%b %t\n", v, IsCast(v)) // "10010 true"
}

//别名
type Flags uint

//给一个无符号整数的最低5bit的每个bit指定一个名字：
const (
	FlagUp        Flags = 1 << iota // is up
	FlagBroadcast                   // 1 << iota+1，移动+1位
	FlagLoopback
	FlagPointToPoint
	FlagMulticast
)

func IsUp(v Flags) bool     { return v&FlagUp == FlagUp }
func TurnDown(v *Flags)     { *v &^= FlagUp }
func SetBroadcast(v *Flags) { *v |= FlagBroadcast }
func IsCast(v Flags) bool   { return v&(FlagBroadcast|FlagMulticast) != 0 }

const (
	_   = 1 << (10 * iota)
	KiB // 1024
	MiB // 1048576
	GiB // 1073741824
	TiB // 1099511627776             (exceeds 1 << 32)
	PiB // 1125899906842624
	EiB // 1152921504606846976
	ZiB // 1180591620717411303424    (exceeds 1 << 64)
	YiB // 1208925819614629174706176
)

//六种未明确类型的常量类型，无类型的布尔型、无类型的整数、无类型的字符、无类型的浮点数、无类型的复数、无类型的字符串
//通过延迟明确常量的具体类型，无类型的常量不仅可以提供更高的运算精度，而且可以直接用于更多的表达式而不需要显式的类型转换。
func testUnType() {
	fmt.Println(YiB / ZiB) // "1024"

	//math.Pi无类型的浮点数常量
	var x float32 = math.Pi
	var y float64 = math.Pi
	var z complex128 = math.Pi

	//如果math.Pi被确定为特定类型，比如float64，那么结果精度可能会不一样，同时对于需要float32或complex128类型值的地方则会强制需要一个明确的类型转换：
	const Pi64 float64 = math.Pi
	var x float32 = float32(Pi64)
	var y float64 = Pi64
	var z complex128 = complex128(Pi64)

	//对于常量面值，不同的写法可能会对应不同的类型。
	// 例如0、0.0、0i和\u0000虽然有着相同的常量值，但是它们分别对应无类型的整数、无类型的浮点数、无类型的复数和无类型的字符等不同的常量类型。
	// 同样，true和false也是无类型的布尔类型，字符串面值常量是无类型的字符串类型。

	//除法运算符/会根据操作数的类型生成对应类型的结果
	var f float64 = 212
	fmt.Println((f - 32) * 5 / 9)     // "100"; (f - 32) * 5 is a float64
	fmt.Println(5 / 9 * (f - 32))     // "0";   5/9 is an untyped integer, 0
	fmt.Println(5.0 / 9.0 * (f - 32)) // "100"; 5.0/9.0 is an untyped float

	//只有常量可以是无类型的。
	//无类型常量被赋值给变量时，自动转换
	var f float64 = 3 + 0i // untyped complex -> float64  //float64(3 + 0i)
	f = 2                  // untyped integer -> float64  //float64(2)
	f = 1e123              // untyped floating-point -> float64  //float64(1e123)
	f = 'a'                // untyped rune -> float64  //float64('a')

	//没有显示类型的变量声明，常量的形式将隐式决定变量的默认类型
	i := 0      // untyped integer;        implicit int(0)
	r := '\000' // untyped rune;           implicit rune('\000')
	f := 0.0    // untyped floating-point; implicit float64(0.0)
	c := 0i     // untyped complex;        implicit complex128(0i)

	//默认类型
	fmt.Printf("%T\n", 0)      // "int"
	fmt.Printf("%T\n", 0.0)    // "float64"
	fmt.Printf("%T\n", 0i)     // "complex128"
	fmt.Printf("%T\n", '\000') // "int32" (rune)
}
