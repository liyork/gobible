package main

import (
	"bytes"
	"fmt"
	"strconv"
	"unicode/utf8"
)

//一个字符串是一个不可改变的字节序列
//不变性意味如果两个字符串共享相同的底层数据的话也是安全的，这使得复制任何长度的字符串代价是低廉的。
// 同样，一个字符串s和对应的子字符串切片s[7:]的操作也可以安全地共享相同的内存，因此字符串切片操作代价也是低廉的。在这两种情况下都没有必要分配新的内存。

func main() {
	//testString1()
	//testSubString()
	//noChange()
	//stringliteral()
	//testUnicodeShow()
	//strings中，HasPrefix,HasSuffix,Contains
	//optUnicode()
	//testRune()
	//fmt.Println(comma("12345"))
	//testStringSlice()
	//fmt.Println(intsToString([]int{1, 2, 3}))
	//testConvert()
	testCompare()
}

func testUnicodeShow() {
	//"世界"
	//"\xe4\xb8\x96\xe7\x95\x8c"
	//通过Unicode码点输入特殊的字符
	//"\u4e16\u754c"//\uhhhh对应16bit的码点值
	//"\U00004e16\U0000754c"//\Uhhhhhhhh对应32bit的码点值
	//Unicode转义也可以使用在rune字符中
	//'世' '\u4e16' '\U00004e16'
}

func stringliteral() {
	//字符串面值，Go语言源文件总是用UTF8编码，并且Go语言的文本字符串也以UTF8编码的方式处理，因此我们可以将Unicode码点也写到字符串面值中
	var s = "Hello, 世界"
	//一个十六进制的转义形式是\xhh(h表示十六进制数字)，一个八进制转义形式是\ooo(三个八进制的o数字，0-7)
	//一个原生的字符串面值形式是`...`，使用反引号代替双引号。在原生的字符串面值中，没有转义操作；全部的内容都是字面的意思，包含退格和换行，
	// 因此一个程序中的原生字符串面值可能跨越多行
	//原生字符串面值用于编写正则表达式会很方便，因为正则表达式往往会包含很多反斜杠。原生字符串面值同时被广泛应用于HTML模板、
	// JSON面值、命令行提示信息以及那些需要扩展到多行的场景
	fmt.Println(s)
}

func noChange() {
	//不可变
	s := "left foot"
	t := s
	s += ", right foot"
	fmt.Println(s)
	// "left foot, right foot"
	fmt.Println(t)
	// "left foot"
	//s[0] = 'L' // compile error: cannot assign to s[0]
}

func testSubString() {
	s := "hello, world"
	//包含0位置,不包含5位置
	fmt.Println(s[0:5])
	// "hello"
	//忽略start则为0，end则为len
	fmt.Println(s[:5])
	// "hello"
	fmt.Println(s[7:])
	// "world"
	fmt.Println(s[:])
	// "hello, world"
	fmt.Println("goodbye" + s[5:])
	// "goodbye, world"
}

func testString1() {
	s := "hello, world"
	//返回字节数目,不是rune字符数目
	fmt.Println(len(s))
	// "12"
	fmt.Println(s[0], s[7])
	// "104 119" ('h' and 'w')
	fmt.Printf("%c,%c\n", s[0], s[7])
	// "104 119" ('h' and 'w')
	//c := s[len(s)] // panic: index out of range
	//第i个字节并不一定是字符串的第i个字符，因为对于非ASCII字符的UTF8编码会要两个或多个字节
}

func optUnicode() {
	s := "Hello, 世界"
	fmt.Println(len(s))                    // 字符串包含13个字节
	fmt.Println(utf8.RuneCountInString(s)) // 对应9个Unicode字符
	fmt.Println()
	//笨拙
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c,size:%d\n", i, r, size)
		i += size
	}
	fmt.Println()
	//简化,range循环在处理字符串时，会自动隐式解码UTF8字符串,索引更新的步长超过1个字节
	for i, r := range s {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

	fmt.Println("字符数目:", utf8.RuneCountInString(s))
}

//UTF8字符串作为交换格式是非常方便的，但是在程序内部采用rune序列可能更方便，因为rune大小一致，支持数组索引和方便切割。
func testRune() {
	s := "中国你好"
	fmt.Printf("% x\n", s) // "e4 b8 ad e5 9b bd e4 bd a0 e5 a5 bd"
	r := []rune(s)         //字符串编码的Unicode码点序列
	fmt.Printf("%c\n", r)
	fmt.Printf("%x\n", r)  // "[4e2d 56fd 4f60 597d]"
	fmt.Println(string(r)) //进行UTF8编码

	//将一个整数转型为字符串意思是生成以只包含对应Unicode码点字符的UTF8字符串：
	fmt.Println(string(65))     // "A", not "65"
	fmt.Println(string(0x4eac)) // "京"
	//无效
	fmt.Println(string(1234567)) // "�"
}

// comma inserts commas in a non-negative decimal integer string
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	//倒数第三个字符为分隔
	return comma(s[:n-3]) + "," + s[n-3:]
}

//一个字符串是包含的只读字节数组，一旦创建，是不可变的。一个字节slice的元素则可以自由地修改。
func testStringSlice() {
	s := "abc"
	b := []byte(s) //不会对原有string影响
	s2 := string(b)

	b[2] = 'd'
	fmt.Println(s, s2)
	fmt.Printf("%c\n", b)
}

//Buffer类型用于字节slice的缓存
// intsToString is like fmt.Sprint(values) but adds commas.
func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
		//buf.WriteString(strconv.Itoa(v))
	}
	buf.WriteByte(']')
	return buf.String()
}

func testConvert() {
	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x))

	//用不同进制格式化
	fmt.Println(strconv.FormatInt(int64(x), 2)) // "1111011"
	//fmt.Printf函数的%b、%d、%o和%x等参数提供功能往往比strconv包的Format函数方便很多
	s := fmt.Sprintf("x=%b", x) // "x=1111011"
	fmt.Println(s)

	x, err := strconv.Atoi("123") // x is an int
	//16表示int16，0则表示int，返回的结果z总是int64类型，可以通过强制类型转换将它转为更小的整数类型。
	z, err := strconv.ParseInt("123", 10, 64) // base 10, up to 64 bits
	fmt.Println(x, z, err)
}

//看来用字符值比较
func testCompare() {
	fmt.Println("a" == "a")

	var a1 = "a"
	var b1 = "a"
	fmt.Println(a1 == b1)

	var a2 = getA()
	var b2 = getB()
	fmt.Println(a2 == b2)
}

func getB() string {
	return "a"
}

func getA() string {
	return "a"
}
