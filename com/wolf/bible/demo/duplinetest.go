//寻找相同行并打印重复数
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//go run duplinetest.go 分别每行输入每行，之后回车，最后退出使用ctrl+d
func dup1() {
	counts := make(map[string]int)      //make创建，key=string/value=int,key需要能用==比较
	input := bufio.NewScanner(os.Stdin) //声明并创建Scanner
	for input.Scan() {                  //读入一行
		line := input.Text()
		//if line == "q" {
		//	break
		//}
		counts[line]++ //map中读取不存在的key时为value类型的默认值
	}

	for line, n := range counts {
		if n > 1 {
			//制表符\t和\n，代表不可见字符的转义字符
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

//go run duplinetest.go
//go run duplinetest.go ./testfile/a.txt
func dup2() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg) //打开文件
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

//函数和包级别的变量可以任意顺序声明，并不影响其被调用。
//map传递的是引用，指针的拷贝
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() { //流式读取
		counts[input.Text()]++
	}
}

//把全部输入数据读到内存中，一次分割为多行，然后处理它们
func dup3() {
	counts := computeCounts()

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func computeCounts() map[string]int {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	return counts
}

func practiceDup() {
	counts := computeCounts()

	isDup := false
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
			isDup = true
		}
	}

	if isDup {
		fmt.Println("file:", os.Args[0], "has dup")
	}
}

//bufio.Scanner、ioutil.ReadFile和ioutil.WriteFile都使用*os.File的Read和Write方法
func main() {
	//dup1()
	//dup2()
	//dup3()
	practiceDup()
}
