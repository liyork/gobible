package programstruct

import "fmt"

//变量或表达式的类型定义了对应存储值的属性特征，例如数值在内存的存储大小（或者是元素的bit个数），它们在内部是如何表达的，
// 是否支持一些操作符，以及它们自己关联的方法集等。

//一个类型声明语句创建了一个新的类型名称，和现有类型具有相同的底层结构。与原有类型不兼容。类型声明语句一般出现在包一级
//type 类型名字 底层类型

//展示：将不同温度单位分别定义为不同的类型：

//命名类型
//声明两种类型，不同数据类型，不可以被相互比较或混在一个表达式运算
//刻意区分类型，可以避免一些像无意中使用不同单位的温度混合计算导致的错误
type Celsius float64    // 摄氏温度
type Fahrenheit float64 // 华氏温度

const (
	AbsoluteZeroC Celsius = -273.15 // 绝对零度
	FreezingC     Celsius = 0       // 结冰点温度
	BoilingC      Celsius = 100     // 沸水温度
)

//显示类型转换
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

//命名类型还可以为该类型的值定义新的行为。这些行为表示为一组关联到该类型的函数集合，我们称为类型的方法集。
//Celsius类型的参数c出现在了函数名的前面，表示声明的是Celsius类型的一个名叫String的方法
func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func main() {
	//底层数据类型决定了内部结构和表达方式，也决定是否可以像底层类型一样对内置运算符的支持。
	// Celsius和Fahrenheit类型的算术运算行为和底层的float64类型是一样的
	fmt.Printf("%g\n", BoilingC-FreezingC)
	bolingF := CToF(BoilingC)
	fmt.Printf("%g\n", bolingF-CToF(FreezingC))
	//fmt.Printf("%g\n", bolingF-FreezingC)// compile error: type mismatch

	var c Celsius
	var f Fahrenheit
	fmt.Println(c == 0)
	fmt.Println(f >= 0)
	//fmt.Println(c == f) // compile error: type mismatch
	fmt.Println(c == Celsius(f))

	//许多类型都会定义一个String方法，因为当使用fmt包的打印方法时，将会优先使用该类型对应的String方法返回的结果打印
	c = FToC(212.0)
	fmt.Println(c.String())
	fmt.Printf("%v\n", c)
	fmt.Printf("%s\n", c)
	fmt.Println(c)
	fmt.Printf("%g\n", c)
	fmt.Println(float64(c))
}
