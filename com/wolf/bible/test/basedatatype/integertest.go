package main

import "fmt"

//int8、int16、int32和int64，8、16、32、64bit大小的有符号整数
//uint8、uint16、uint32和uint64四种无符号整数类型
//Unicode字符rune类型是和int32等价的类型，byte和uint8类型是等价的,byte类型一般用于强调数值是一个原始的数据而不是一个小的整数。
//无符号的整数类型uintptr，没有指定具体的bit大小但是足以容纳指针。在底层编程时才需要，函数库或操作系统接口相交互的地方
//int、uint和uintptr、int32是不同类型

//有符号整数采用2的补码形式表示，最高bit位用来表示符号位，一个n-bit的有符号数的值域是从-2^{n-1}到2^{n-1}-1。
// 无符号整数的所有bit位都用于表示非负数，值域是0到2^n-1
//如：int8类型整数的值域是从-128到127，而uint8类型整数的值域是从0到255。

func main() {
	//test1()
	//testOverflow()
	//testBit()
	//bitMove()
	//unSigned()
	//testSameType()
	//testFmtDecimal()
	testFmtString()
}

func test1() {
	//被取模数的符号总是一致的
	fmt.Println(-5%3 == -5%-3)
	//除法运算符/的行为依赖于操作数是否为全为整数
	fmt.Println(5.0 / 4.0)
	//1.25
	fmt.Println(5 / 4)
	//1,整数除法会向着0方向截断余数
}

//溢出
func testOverflow() {
	var u uint8 = 255
	fmt.Println(u, u+1, u*u)
	// "255 0 1"
	var i int8 = 127
	fmt.Println(i, i+1, i*i)
	// "127 -128 1"
}

func testBit() {
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2

	fmt.Printf("%08b\n", x) // "00100010", the set {1, 5}
	fmt.Printf("%08b\n", y) // "00000110", the set {1, 2}

	fmt.Printf("%08b\n", x&y)  // "00000010", the intersection {1}
	fmt.Printf("%08b\n", x|y)  // "00100110", the union {1, 2, 5}
	fmt.Printf("%08b\n", x^y)  // "00100100", the symmetric difference {2, 5}
	fmt.Printf("%08b\n", x&^y) // "00100000", the difference {5}

	//每次移动i位然后&x得到x对应的位置为1的位置
	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 { // membership test
			fmt.Println(i) // "1", "5"
		}
	}

	fmt.Printf("%08b\n", x<<1) // "01000100", the set {2, 6}
	fmt.Printf("%08b\n", x>>1) // "00010001", the set {0, 4}
}

//x<<n和x>>n，其中n必须是无符号数
//无符号数的右移运算也是用0填充左边空缺的bit位，但是有符号数的右移运算会用符号位的值填充左边空缺的bit位，最好用无符号运算
func bitMove() {

}

//无符号数往往只有在位运算或其它特殊的运算场景才会使用，就像bit集合、分析二进制文件格式或者是哈希和加密操作等
func unSigned() {
	medals := []string{"gold", "silver", "bronze"}
	//若len返回无符号数，那么i也是unit类型，3次迭代后i==0时，i--语句不产生-1而是unit类型的最大值
	for i := len(medals) - 1; i >= 0; i-- {
		fmt.Println(medals[i]) // "bronze", "silver", "gold"
	}
}

//算术和逻辑运算的二元操作中必须是相同的类型
func testSameType() {
	var apples int32 = 1
	var oranges int16 = 2
	//var compote int = apples + oranges // compile error
	var compote = int(apples) + int(oranges)
	fmt.Println("compote:", compote)

	//注意转换是否丢失数据。浮点数到整数，丢失小数部分
	f := 3.141 // a float64
	i := int(f)
	fmt.Println(f, i) // "3.141 3"
	f = 1.99
	fmt.Println(int(f)) // "1"
}

func testFmtDecimal() {
	o := 0666
	//%[1]再次使用第一个操作数
	fmt.Printf("%d %[1]o %#[1]o\n", o) // "438 666 0666"
	x := int64(0xdeadbeef)
	//%#表明，用%o、%x或%X输出时生成0、0x或0X前缀
	fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x) // 3735928559 deadbeef 0xdeadbeef 0XDEADBEEF
}

func testFmtString() {
	ascii := 'a'
	unicode := '国'
	newline := '\n'
	//%c打印字符,%q打印带单引号的字符
	fmt.Printf("%d %[1]c %[1]q\n", ascii)   // "97 a 'a'"
	fmt.Printf("%d %[1]c %[1]q\n", unicode) // "22269 国 '国'"
	fmt.Printf("%d %[1]q\n", newline)       // "10 '\n'"
}
