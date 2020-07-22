package main

import (
	"fmt"
	"os"
	"time"
)

//无缓存channel更强地保证了每个发送操作与相应的同步接收操作；但是对于带缓存channel，这些操作是解耦的
// 即使知道将要发送到一个channel的信息的数量上限，创建一个对应容量大小的带缓存channel也是不现实的，因为这要求在执行任何接收
// 操作之前缓存所有已经发送的值。如果未能分配足够的缓冲将导致程序死锁。
// 衡量好生产和消费之间的速率，不然生产快则缓冲队列是满的，要是消费快则队列是空的。基本保持两者平衡。若是有环节需要慢做可以并发
// 那么可以添加goroutines操作相同channel

// channel和net.Conn实例是并发安全的

func main() {
	testBaseChannel()
	//nobuffChan()
	//pipeline()
	//bufChan()
	//fmt.Println(mirroredQuery())
	//testchannel1()
	//testchannel2_1()
	//testchannel2_2()
	//testchannel3()
	//testchannel4()
	//testchannel5()
	//testchanne6()
}

func testBaseChannel() {
	//channel对应一个make创建的底层数据结构的应用，赋值channel或用于函数参数传递时，只是拷贝了一个channel引用。零值也是nil
	ch := make(chan int)
	//unbuffered channel
	ch = make(chan int, 3)
	// buffered channel with capacity 3
	fmt.Println(ch)
	//发送和接收两个操作都使用<-运算符
	ch <- 1
	//send
	x := <-ch
	//receive
	<-ch
	// receive, result is discarded
	//close后，若再写入则panic。可以已经发送成功的数据，没有的话则是零值。
	close(ch)
	fmt.Println(x)
}

//无缓存的channel发送时阻塞，直到有人读取。读取也会阻塞直到有人发送。
// 当通过一个无缓存Channels发送数据时，接收者收到数据发生先发生(相关事情完成)，唤醒发送者goroutine后发生(可以使用事情)
func nobuffChan() {

}

//不管一个channel是否被关闭，当它没有被引用时将会被Go语言的垃圾自动回收器回收。
// 对于每个打开的文件，都需要在不使用的使用调用对应的Close方法来关闭文件。
// 重复关闭一个channel将导致panic异常，试图关闭一个nil值的channel也将导致panic异常

// chan<- int 表示只发送int的channel，不能接收
//  <-chan int 表示只接收int的channel，不能发送
func pipeline() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals) // 调用时，隐式从chan int转换成chan<- int
	go squarer(squares, naturals)
	printer(squares)
}

func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for v := range in { //当channel被关闭并且没有值时跳出
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

// 带缓存的Channel内部持有一个元素队列。队列最大容量在make时设定
// 发送则到队尾，接收则从队头，满了则放入会阻塞，空的则读取会阻塞
// channel的缓冲队列解耦了接受和发送的goroutine
//注意：
//要使用队列则用slice，对于同一个goroutine不要用带缓存的channel作为队列使用
// Channel和goroutine的调度器机制是紧密相连的，一个发送操作或许使整个程序可能会永远阻塞
func bufChan() {
	ch := make(chan string, 3) // 持有三个字符元素的带缓存channel

	ch <- "A"
	ch <- "B"
	ch <- "C"

	//len-内部缓冲队列中有效元素个数，cap-队列容量
	fmt.Println(<-ch, len(ch), cap(ch))
}

//演示通过三个并发goroutine获取数据，最后只要最快返回的数据
// 若用无缓冲的channel，则两个慢的goroutines因为没有人接收而被永远卡住，称为goroutines泄漏bug，不会被自动回收，
// 因此确保每个不再需要的goroutine能正常退出是重要的
func mirroredQuery() string {
	responses := make(chan string, 3)
	go func() { responses <- request("a.com") }()
	go func() { responses <- request("b.com") }()
	go func() { responses <- request("c.com") }()
	return <-responses // return the quickest response
}

func request(hostname string) (response string) { return }

//正常
func testchannel1() {
	ch := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		ch <- "1"
	}()
	x := <-ch

	fmt.Println("11111", x)
}

func testchannel2_1() {
	ch := make(chan string)
	ch <- "1" //报错，应该是检测到没有其他地方放入?。本身构造完ch不开启其他goroutine那么自己写入就会阻塞，然后就卡住了，需要有其他地方先进行读取
}

//错误示范，一个goroutine中读取无缓冲channel必然一直阻塞
func testchannel2_2() {
	ch := make(chan string)
	fmt.Println(<-ch)
}

//看来监测的是有其他goroutine进行反向操作，即使不是当时执行
func testchannel3() {
	ch := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		println(<-ch)
	}()
	ch <- "1"
	fmt.Println("main;.....")
}

//测试range chan
func testchannel4() {
	ch := make(chan string, 2)
	go func() {
		for x := range ch { //阻塞直到有数据可读，或者chan关闭了
			fmt.Println("x===>", x)
		}
	}()
	ch <- "1"
	ch <- "1"
	time.Sleep(1 * time.Second)
	ch <- "2"

	time.Sleep(5 * time.Second)
	fmt.Println("main;.....")
}

func testchannel5() {
	ch := make(chan string)
	ch <- "1"

	for x := range ch {
		fmt.Println("x===>", x)
	}

	fmt.Println("main;.....")
}

//因为os.Args[1:]这个是启动命令时传入的参数，只是一次启动就传完了，所以main中for range chan会报错，显示main chann不会再有其他传值
func testchanne6() {
	worklist := make(chan string)
	fmt.Println("os.Args[1:]:", os.Args[1:])

	go func() {
		for _, v := range os.Args[1:] {
			fmt.Println("v:", v)
			worklist <- v
		}
	}()

	for list := range worklist {
		time.Sleep(1 * time.Second)
		fmt.Println("list:", list)
	}
}
