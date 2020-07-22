package sharevar

import (
	"fmt"
)

var balance int

func Deposit(amount int) { balance = balance + amount }
func Balance() int       { return balance }

func main() {
	//dataRace1()
	//dataRace2()
	//dataRace2Right()
	//dataRace3Right()
	dataRace4Right()
}

// 数据竞争：A的操作涉及两个(读、写)，这时中间插入B，那么最后B可能会丢了
// 数据竞争：数据竞争会在两个以上的goroutine并发访问相同的变量且至少其中一个为写操作时发生
// 无论任何时候，只要有两个goroutine并发访问同一变量，都有写操作，且一个为读后写，会发生数据竞争
func dataRace1() {
	// Alice:
	go func() {
		Deposit(200)                // A1
		fmt.Println("=", Balance()) // A2
	}()
	// Bob:
	go Deposit(100)
	// B
}

var icons = make(map[string]string)

func loadIcon(name string) string {
	return "load"
}

func Icon(name string) string {
	icon, ok := icons[name]
	if !ok {
		icon = loadIcon(name)
		icons[name] = icon
	}
	return icon
}

// Icon不是线程安全操作
func dataRace2() {
	go func() { loadIcon("xxx") }()
	go func() { Icon("xxx") }()
}

// 包初始化阶段已经赋值，在main函数执行之前完成。
var icons1 = map[string]string{
	"spades.png":   loadIcon("spades.png"),
	"hearts.png":   loadIcon("hearts.png"),
	"diamonds.png": loadIcon("diamonds.png"),
	"clubs.png":    loadIcon("clubs.png"),
}

// 直接初始化所有map，这样没有并发写入，就不会产生问题
func dataRace2Right() {
	fmt.Println(icons1)
	go func() { Icon("spades.png") }()
}

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance

func Deposit1(amount int) { deposits <- amount }
func Balance1() int       { return <-balances }

// 避免
func teller() {
	var balance1 int // balance1 is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance1 += amount
		case balances <- balance:
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}

// 避免从多个goroutine访问变量，使用通道来进行多goroutine共享变量
// 不要使用共享数据来通信；使用通信来共享数据
func dataRace3Right() {
	teller()
}

type Cake struct {
	state string
}

func baker(cooked chan<- *Cake) {
	for {
		cake := new(Cake)
		cake.state = "cooked"
		cooked <- cake // baker never touches this cake again
	}
}

func icer(iced chan<- *Cake, cooked <-chan *Cake) {
	for cake := range cooked {
		cake.state = "iced"
		iced <- cake // icer never touch this cake again
	}
}

// 流水线操作，每条流水线保证传递到chan之后不再访问
func dataRace4Right() {
	cooked := make(chan *Cake)
	go func() { baker(cooked) }()
	iced := make(chan *Cake)
	go func() { icer(iced, cooked) }()
}
