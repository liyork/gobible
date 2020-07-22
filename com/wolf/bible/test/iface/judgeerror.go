package main

import (
	"fmt"
	"os"
)

//I/O可以因为任何数量的原因失败，但是有三种经常的错误必须进行不同的处理：文件已经存在（对于创建操作），找不到文件（对于读取操作），和权限拒绝
//通过检查错误消息的子字符串来保证特定的函数以期望的方式失败对于线上的代码是不够的
//一个更可靠的方式是使用一个专门的类型来描述结构化的错误,PathError

func main() {
	_, err := os.Open("/xx/yy")
	fmt.Println(err)         //String
	fmt.Printf("%#v\n", err) //类型

	//区别错误通常必须在失败操作后，错误传回调用者前进行
	fmt.Println(os.IsNotExist(err))
}
