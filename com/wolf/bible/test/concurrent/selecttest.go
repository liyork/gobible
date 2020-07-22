package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	//testTick()
	//testNewTick()
	//testTimeAfter()
	//testSelect()
	//testSelectDefault()
	testNonMeetSelect()
	//testLaunch()
	//testSelectChannel()

}

func testTimeAfter() {

	abortchan := make(chan struct{})
	abort(abortchan)

	select {
	//time.After立即返回一个channel，并启动一个goroutine在经过特定的时间后向channel发送一个独立的值
	case <-time.After(10 * time.Second):
	case <-abortchan:
		fmt.Println("Launch aborted!")
		return
	}
}

//多路复用(multiplex)，多个chan选择一个，不会被第一个而阻塞。每个case代表通信操作
//select会等待case中有能够执行的case时才去执行。
//只有一个case执行，交替执行，从0开始放入，i为基数时取，i为偶数时放入
func testSelect() {
	//ch := make(chan int, 1)
	ch := make(chan int, 3) //如果多个case同时就绪时，select会随机地选择一个执行
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println("x := <-ch,,", x, ",,", i)
		case ch <- i:
			//fmt.Println("ch <- i,,", i)
		}
	}
}

//有时候我们希望能够从channel中发送或者接收值，并避免因为发送或者接收导致的阻塞，尤其是当channel没有准备好写或者读时。
// select语句就可以实现这样的功能。select会有一个default来设置当其它的操作都不能够马上被处理时程序需要执行哪些逻辑。
// channel的零值是nil，对一个nil的channel发送和接收操作会永远阻塞
// 在select语句中操作nil的channel永远都不会被select到
func testSelectDefault() {
	abortchan := make(chan struct{})
	abort(abortchan)

	//轮询channel
	for {
		select {
		case <-abortchan:
			fmt.Println("Launch aborted!")
			return
		default:
			fmt.Println("Do nothing...")
		}
		time.Sleep(1 * time.Second)
	}
}

func testNonMeetSelect() {
	abortchan1 := make(chan struct{})
	abort(abortchan1) // 要是没有关闭操作，运行时直接报错，因为select会卡住无法退出
	abortchan2 := make(chan struct{})

	for {
		select {
		case <-abortchan1:
			fmt.Println("Launch aborted1!")
			return
		case <-abortchan2:
			fmt.Println("Launch aborted2!")
			return
		}
	}
}

func testLaunch() {
	fmt.Println("Commencing countdown. Press return to abort.")
	tick := time.Tick(1 * time.Second)

	abortchan := make(chan struct{})
	abort(abortchan)

	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick:
			fmt.Println("Do nothing")
		case <-abortchan:
			fmt.Println("Launch aborted!")
			return
		}
		<-tick //读取
	}
	//launch()
}

func abort(abortchan chan struct{}) {
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		abortchan <- struct{}{}
	}()
}

//周期性地发送事件(时间戳)
//若没有人接收，但是ticker这个goroutine还会存活，依然尝试向channel中发值--goroutine泄漏
//只有当程序整个生命周期都需要这个时间时我们使用它才比较合适
func testTick() {
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		<-tick
	}
}

func testNewTick() {
	ticker := time.NewTicker(1 * time.Second)
	<-ticker.C    // receive from the ticker's channel
	ticker.Stop() // cause the ticker's goroutine to terminate
}

func testSelectChannel() {
	fmt.Println(cancelled1()) // false

	go func() {
		done1 <- struct{}{}
	}()
	time.Sleep(1 * time.Second)
	fmt.Println(cancelled1()) // true

	close(done1)
	time.Sleep(1 * time.Second)
	fmt.Println(cancelled1()) // true
}

var done1 = make(chan struct{})

func cancelled1() bool {
	select {
	case <-done1: //有值或者关闭情况
		return true
	default:
		return false
	}
}
