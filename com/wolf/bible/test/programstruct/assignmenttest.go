package programstruct

import "fmt"

func testAssignment() {
	//x = 1 //命名变量的赋值
	//*p = true //通过指针间接赋值
	//person.name = "bob" //结构体字段赋值
	//count[x] = count[x] * scale //数组、slice或map的元素赋值
	//count[x] *=scale//符合操作：二元算数运算符+赋值语句，省去对变量表达式的重复计算

	//自增和自减是语句，不是表达式
	v := 1
	v++
	v--

	var x, y, k int
	//元组赋值，多个变量赋值，右边下先进行求值，然后统一更新左边变量值
	x, y = y, x
	x, y, k = 2, 3, 4
	fmt.Println(k)

	//nil可以赋值给任何指针或引用类型的变量
}
