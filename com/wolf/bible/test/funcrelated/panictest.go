package main

import (
	"fmt"
	"golang.org/x/text/collate"
	"os"
	"runtime"
)

//运行时错误会引起painc异常，如数组访问越界、空指针引用等
//一般而言，当panic异常发生时，程序会中断运行，并立即执行在该goroutine中被延迟的函数（defer 机制）。

//由于panic会引起程序的崩溃，因此panic一般用于严重错误，如程序内部的逻辑不一致。
//在健壮的程序中，任何可以预料到的错误，如不正确的输入、错误的配置或是失败的I/O操作都应该被优雅的处理，最好的处理方式，就是使用Go的错误机制。
//当调用者明确的知道正确的输入不会引起函数错误时，要求调用者检查这个错误是不必要和累赘的。
//MustCompile不能接收不合法的输入。函数名中的Must前缀是一种针对此类函数的命名约定

//当某些不应该发生的场景发生时，程序到达了某条逻辑上不可能到达的路径，用panic
func switchUsePanic() {
	var suit string
	switch s := suit; s {
	case "Spades":
	case "Hearts":
	case "Diamonds":
	case "Clubs":
	default:
		panic(fmt.Sprintf("invalid suit %q", s)) // Joker?
	}
}

//断言函数用于必须满足的前置条件，很容易被滥用。除非你能提供更多的错误信息，或更快速的发现错误，否则不需要用断言，编译器在运行时会帮你检查代码
func Reset(x *collate.Buffer) {
	if x == nil {
		panic("x is nil") // unnecessary!
	}
}

func f(x int) {
	if x == 0 {
		return
	}
	fmt.Printf("f(%d)\n", x+0/x) // panics if x == 0
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}

//在Go的panic机制中，延迟函数的调用在释放堆栈信息之前
//为了方便诊断问题，runtime包允许输出堆栈信息
func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}

func main() {
	defer printStack()
	f(3)
}
