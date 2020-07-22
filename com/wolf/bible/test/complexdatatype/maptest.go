package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"unicode"
	"unicode/utf8"
)

//无序的key/value对的集合，其中所有的key都是不同的。map[K]V。map中所有的key都有相同的类型，所有的value也有着相同的类型
//K对应的key必须是支持==比较运算符的数据类型

func main() {
	//testCreate()
	//testMapOpt()
	//testMapTravel()
	//testMapSort()
	//testMapNil()
	//testKeySlice()
	testGraph()
}

//排序
func testMapSort() {
	args := map[string]int{}
	names := make([]string, 0, len(args))
	for name := range args {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, args[name])
	}
}

func testMapTravel() {
	args := map[string]int{"1": 1, "2": 2}
	//遍历顺序随机
	for name, age := range args {
		fmt.Printf("%s\t%d\n", name, age)
	}
}

func testMapOpt() {
	args := make(map[string]int)
	args["alice"] = 32
	fmt.Println(args["alice"])
	//允许不存在
	delete(args, "alice")
	fmt.Println(args)
	//不存在则返回零值
	args["bob"] = args["bob"] + 1
	//简短赋值
	args["bob"] += 1
	args["bob"] ++
	//map中的元素并不是一个变量，不能对map的元素进行取址操作：
	//_ = &ages["bob"] // compile error: cannot take address of map element
	//禁止对map元素取址的原因是map可能随着元素数量的增长而重新分配更大的内存空间，从而可能导致之前的地址无效。
	fmt.Println(args)
}

func testCreate() {
	args := make(map[string]int)
	args = map[string]int{}
	args = map[string]int{
		"alice":   31,
		"chalsld": 34,
	}
	fmt.Println(args)
}

//查找、删除、len和range循环都可以安全工作在nil值的map上。在向map存数据前必须先创建map
func testMapNil() {
	var ages map[string]int
	fmt.Println(ages == nil)
	fmt.Println(len(ages) == 0)

	//ages["carol"] = 21 // panic: assignment to entry in nil map
	fmt.Println(ages["carol"]) //安全

	//区分不存在与0的场景
	if age, ok := ages["bob"]; !ok {
		fmt.Println("not exists bob")
	} else {
		fmt.Println("exists bob:", age)
	}
}

//map之间也不能进行相等比较
func equalMap(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}

func demoSet() {
	seen := make(map[string]bool) // a set of strings
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}
}

var m = make(map[string]int)

func k(list []string) string { return fmt.Sprintf("%q", list) }

func Add(list []string)       { m[k(list)]++ }
func Count(list []string) int { return m[k(list)] }

func testKeySlice() {
	s1 := []string{"1", "2"}
	s2 := []string{"2", "3"}
	s3 := []string{"1", "2"}

	Add(s1)
	Add(s2)
	Add(s3)

	fmt.Println(Count(s1), Count(s2), Count(s3))
}

// computes counts of Unicode characters.
func charCount() {
	counts := make(map[rune]int)    // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes-字符UTF8编码后的长度, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 { //无效的UTF-8编码的字符
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

//graph将一个字符串类型的key映射到一组相关的字符串集合
func testGraph() {
	addEdge("1.1", "1.2")
	addEdge("1.2", "1.3")
	fmt.Println(hasEdge("1.1", "1.2"))
}

//key:string,value:map[string]bool
var graph = make(map[string]map[string]bool)

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges //map插入前必须先初始化
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to] //允许不存在
}
