package programstruct

import (
	"fmt"
	"github.com/liyork/bible/com/wolf/bible/test/programstruct/package1"
)

//Go语言中的包和其他语言的库或模块的概念类似，目的都是为了支持模块化、封装、单独编译和代码重用。
//一个包的源代码保存在一个或多个以.go为文件后缀名的源文件中
//每个包都对应一个独立的名字空间

//每个包在解决依赖的前提下，以导入声明的顺序初始化，每个包只会被初始化一次。
//初始化工作是自下而上进行的，main包最后被初始化

//访问外部代码
func accessOutPackage() {
	fmt.Printf("Brrrr! %v\n", package1.AbsoluteZeroC)
	fmt.Println(package1.CToF(package1.BoilingC))
}

// pc[i] is the population count of i.
var pc [256]byte

//在每个文件中的init初始化函数，在程序开始执行时按照它们声明的顺序被自动调用。
func init() {
	//for i := range pc {
	//	pc[i] = pc[i/2] + byte(i&1)
	//}

	pc = func() (pc [256]byte) {
		for i := range pc {
			pc[i] = pc[i/2] + byte(i&1)
		}
		return
	}()
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func main() {
	accessOutPackage()

	fmt.Println("a2:", a2)
}

//包的初始化首先是解决包级变量的依赖顺序，然后按照包级变量声明出现的顺序依次初始化
//如果包中含有多个.go源文件，它们将按照发给编译器的顺序进行初始化，Go语言的构建工具首先会将.go文件根据文件名排序，然后依次调用编译器编译。
var a2 = b + c // a 第三个初始化, 为 3(依赖b、c)
var b = f2()   // b 第二个初始化, 为 2, 通过调用 f2 (依赖c)
var c = 1      // c 第一个初始化, 为 1
func f2() int  { return c + 1 }
