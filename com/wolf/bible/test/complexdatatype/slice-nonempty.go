// Nonempty is an example of an in-place slice algorithm
package main

import "fmt"

// nonempty returns a slice holding only the non-empty strings.
// The underlying array is modified during the call.
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

//使用append实现
func nonempty2(strings []string) []string {
	out := strings[:0] // zero-length slice of original
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func main() {
	//testNonempty()
	testStack()
}

//输入的slice和输出的slice共享一个底层数组
func testNonempty() {
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty(data)) // `["one" "three"]`
	fmt.Printf("%q\n", data)           // `["one" "three" "three"]`
	data = nonempty(data)
	fmt.Printf("%q\n", data)
}

func testStack() {
	var stack []int = nil
	var a = 1
	//入栈
	stack = append(stack, a)
	//top
	top := stack[len(stack)-1]
	fmt.Println(top)
	//pop
	stack = stack[:len(stack)-1]
}

//删除元素并保证原来顺序
func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

//删除元素不保证原来顺序
func remove2(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}
