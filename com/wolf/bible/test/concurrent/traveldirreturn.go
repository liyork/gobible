package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

//使用chan的关闭方式来广播相关goroutine进行停止
// Cancel traversal when input is detected

func walkDir1(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	if cancelled() { //goroutine中判断退出标识
		return
	}

	for _, entry := range dirents1(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			walkDir1(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func sleep1Second1() {
	time.Sleep(1 * time.Second)
}

var sema1 = make(chan struct{}, 20)

func dirents1(dir string) []os.FileInfo {
	//阻塞，直到有一个case满足
	select {
	case sema1 <- struct{}{}: // acquire token
	case <-done:
		return nil // cancelled
	}
	defer func() { <-sema1 }() // release token
	entries, err := ioutil.ReadDir(dir)
	sleep1Second1()
	if err != nil {
		fmt.Fprint(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func printDiskUsage1(prefix string, nfiles, nbytes int64) {
	fmt.Printf(prefix+"%d files  %.1f KB\n", nfiles, float64(nbytes)/1024)
}

var verbose1 = flag.Bool("v", false, "show verbose progress messges")

// go run traveldirreturn.go -v $GOPATH/src/test
func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	//响应用户输入进行终止
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		close(done)
	}()

	fileSizes := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir1(root, &n, fileSizes)
	}

	//正常退出
	go func() {
		n.Wait()
		close(fileSizes)
	}()

	var tick <-chan time.Time
	if *verbose1 {
		tick = time.Tick(1000 * time.Microsecond)
	}
	var nfiles, nbytes int64
loop:
	for {
		select {
		case <-done: //有数据可读或者关闭
			// Drain fileSizes to allow existing goroutines to finish.
			// fileSizes会在goroutine中放入，而由下面case进行消费，若是进了此case那么需要把之前fileSizes中数据读取，
			// 然后不会卡住goroutine的退出，进而就可以关闭chan了
			for range fileSizes {
				// Do nothing.
			}
			return
		case size, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes was closed
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage1("tick:", nfiles, nbytes)
		}
	}
	printDiskUsage1("total:", nfiles, nbytes)
}

//现在当取消发生时，所有后台的goroutine都会迅速停止并且主函数会返回。当主函数返回时，而我们又无法在主函数退出的时候确认其已经释放了所有的资源。
// 小窍门可以一用：取代掉直接从主函数返回，我们调用一个panic，然后runtime会把每一个goroutine的栈dump下来。
// 如果main goroutine是唯一一个剩下的goroutine的话，他会清理掉自己的一切资源。
// 但是如果还有其它的goroutine没有退出，他们可能没办法被正确地取消掉，也有可能被取消但是取消操作会很花时间；所以这里的一个调研还是很有必要的。
// 感觉有共享内存存放状态，main退出时可以进行销毁，但是每个子goroutine中的资源处理，应该交给他们自己的defer处理，main中似乎不知道他们的资源如何释放
