package main

import (
	"fmt"
)

//Slice（切片）代表变长的序列，序列中每个元素都有相同的类型。一个slice类型一般写作[]T
//一个slice是一个轻量级的数据结构，提供了访问数组子序列（或者全部）元素的功能，而且slice的底层确实引用一个数组对象
//一个slice由三个部分构成：指针、长度和容量。
// 指针指向第一个slice元素对应的底层数组元素的地址(并不一定是数组第一个)
// len:长度对应slice中元素的数目,长度不能超过容量
// cap:容量一般是从slice的开始位置到底层数据的结尾位置

//多个slice之间可以共享底层的数据，并且引用的数组部分区间可能重叠

func main() {
	//testSliceBase()
	//reverse([]int{0, 1, 2, 3, 4, 5})
	//testSliceNil()
	//testMakeSlice()
	//testAppendSlice()
	testAppendInt()
}

func testSliceBase() {
	//months[0]=""
	months := [...]string{1: "January", 2: "February", 3: "March", 4: "April", 5: "May",
		6: "June", 7: "July", 8: "August", 9: "September", 10: "October", 11: "November", 12: "December"}
	fmt.Println(len(months))
	//创建一个新的slice，s[i:j]，其中0 ≤ i≤ j≤ cap(s)，引用s的从第i个元素开始到第j-1个元素的子序列。新的slice将只有j-i个元素
	//months[1:13] 等价 months[1:]
	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(Q2)
	// ["April" "May" "June"]
	fmt.Println(summer)
	// ["June" "July" "August"]
	//fmt.Println(summer[:20]) // panic: out of range
	endlessSummer := summer[:5]
	// extend a slice (within capacity)
	fmt.Println(endlessSummer)
	// "[June July August September October]"
}

//复制一个slice只是对底层的数组创建了一个新的slice别名
// reverse reverses a slice of ints in place.
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	fmt.Println(s) // "[5 4 3 2 1 0]"
}

//将slice元素循环向左旋转n个元素的
// 方法是三次调用reverse反转函数，第一次是反转开头的n个元素，然后是反转剩下的元素，最后是反转整个slice的元素。
//（如果是向右循环旋转，则将第三个函数调用移到第一个调用位置就可以了。）
func move2Left() {
	s := []int{0, 1, 2, 3, 4, 5}
	// Rotate s left by two positions.
	reverse(s[:2])
	reverse(s[2:])
	reverse(s)
	fmt.Println(s) // "[2 3 4 5 0 1]"
}

//slice之间不能比较，对于byte使用bytes.Equal，其他类型则自己定义
//slice不能比较原因：第一个原因，一个slice的元素是间接引用的，一个slice甚至可以包含自身
//第二个原因，因为slice的元素是间接引用的，一个固定的slice值(指slice本身的值)在不同的时刻可能包含不同的元素，
// 因为底层数组的元素可能会被修改引发的扩容导致slice指向新地址
func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

//一个nil值的slice并没有底层数组
//以相同的方式对待nil值的slice和0长度的slice
func testSliceNil() {
	var s []int              // len(s) == 0, s == nil
	s = nil                  // len(s) == 0, s == nil
	s = []int(nil)           // len(s) == 0, s == nil
	s = []int{}              // len(s) == 0, s != nil
	fmt.Println(len(s) == 0) //判断是否空的
}

//底层make创建了一个匿名的数组变量，然后返回一个slice
func testMakeSlice() {
	a := make([]int, 2)
	b := make([]int, 2, 2) //与上面相同
	c := make([]int, 2, 4)
	d := make([]int, 2)[:4] //与上面相等
	fmt.Println(a, b, c, d)
}

//内置的append函数可能使用比appendInt更复杂的内存扩展策略。因此，通常我们并不知道append调用是否导致了内存的重新分配，
// 因此我们也不能确认新的slice和原始的slice是否引用的是相同的底层数组空间。同样，我们不能确认在原先的slice上的操作是否会影响到新的slice。
// 因此，通常是将append返回的结果直接赋值给输入的slice变量：
func testAppendSlice() {
	var runes []rune
	for _, r := range "Hello,世界" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)

	var x []int
	x = append(x, 1)
	x = append(x, 2, 3)
	x = append(x, 4, 5, 6)
	x = append(x, x...) // append the slice x
	fmt.Println(x)      // "[1 2 3 4 5 6 1 2 3 4 5 6]"
}

//演示扩展slice
func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) { //还有一个元素
		// There is room to grow. Extend the slice.
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

func appendInt1(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	copy(z[len(x):], y)
	return z
}

func testAppendInt() {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d oldCap=%d\tnewSlice:%v\n", i, cap(x), y)
		x = y
	}
}
