package sharevar

import "fmt"

//为什么Balance方法需要用到互斥条件, 只由一个简单的操作组成, 没有竞争条件才对?
//使用mutex有两方面考虑:
//第一:Balance不会在其它操作比如Withdraw“中间”执行。
//第二:"同步"不仅仅是一堆goroutine执行顺序的问题；同样也会涉及到内存的问题。
//
//在现代计算机中可能会有一堆处理器，每一个都会有其本地缓存(local cache)。为了效率，对内存的写入一般会在每一个处理器中缓冲，并在必要时一起flush到主存。
//这种情况下这些数据可能会以与当初goroutine写入顺序不同的顺序被提交到主存。而像channel通信或者互斥量操作这样的原语会使处理器将其聚集的写入flush并commit，
//这样goroutine在某个时间点上的执行结果才能被其它处理器上运行的goroutine得到。

func main() {
	var x, y int
	go func() {
		x = 1                   // A1
		fmt.Print("y:", y, " ") // A2
	}()
	go func() {
		y = 1                   // B1
		fmt.Print("x:", x, " ") // B2
	}()
}

//以上代码两个goroutine看到的结果各种各样，可能被编译器优化，因为从自己goroutine看没有依赖条件
//两个goroutine是并发执行，并且访问共享变量时也没有互斥，会有数据竞争
//在一个独立的goroutine中，每一个语句的执行顺序是可以被保证的；也就是说goroutine是顺序连贯的。但是在不使用channel且不使用mutex这样的显式同步操作时，
//我们就没法保证事件在不同的goroutine中看到的执行顺序是一致的了。
//并发的问题都可以用一致的、简单的既定的模式来规避。所以可能的话，将变量限定在goroutine内部；如果是多个goroutine都需要访问的变量，使用互斥条件来访问
