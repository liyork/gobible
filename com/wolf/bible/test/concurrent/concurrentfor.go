package main

import (
	"log"
	"math/rand"
	"strconv"
	"sync"
)

func ImageFile(infile string) (string, error) {
	return "", nil
}

//彼此独立，顺序无关，并执行
func makeThumbnails(filenames []string) {
	ch := make(chan struct{})

	for _, f := range filenames {
		go func(q string) {
			ImageFile(q) //使用传递给匿名函数f，不能直接使用f，因为匿名函数共享f则随着for执行都是最后一个了
			ch <- struct{}{}
		}(f)
	}

	// Wait for goroutines to complete.
	for range filenames {
		<-ch
	}
}

//带有错误返回。--错误方式
func makeThumbnails2(filenames []string) error {
	errors := make(chan error)

	for _, f := range filenames {
		go func(q string) {
			_, err := ImageFile(q)
			errors <- err
		}(f)
	}

	for range filenames {
		if err := <-errors; err != nil {
			return err // incorrect:goroutine leak!，遇到第一个就返回，那么若其他goroutine再想放入chann就阻塞，就会卡住，不退出，泄漏
		}
	}
	return nil
}

//goroutine返回结构体。--正确方式，使用缓存chan
func makeThumbnails3(filenames []string) (thumbfiles []string, err error) {
	type item struct {
		thumbfile string
		err       error
	}

	ch := make(chan item, len(filenames))
	for _, f := range filenames {
		go func(q string) {
			var it item
			it.thumbfile, err = ImageFile(q)
			ch <- it
		}(f)
	}

	for range filenames {
		it := <-ch
		if it.err != nil {
			return nil, it.err
		}
		thumbfiles = append(thumbfiles, it.thumbfile)
	}

	return thumbfiles, nil
}

//返回文件总大小，参数是chan，不能直接得到个数。使用WaitGroup
func makeThumbnails4(filenames <-chan string) int {
	sizes := make(chan int, 10000)
	var wg sync.WaitGroup

	//fmt.Println("makeThumbnails4")

	for f := range filenames {
		//fmt.Println("for range")
		wg.Add(1) //确保在wg.Wait之前调用
		// worker
		go func(f string) {
			defer wg.Done()
			_, err := ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			//info, _ := os.Stat(thumb)
			//sizes <- info.Size()
			//fmt.Println("for:==>")
			sizes <- rand.Intn(10)
		}(f)
	}

	//fmt.Println("exec for finish")
	// closer, 所有goroutines执行完后再close掉sizes
	//go func() {
	//	//fmt.Println("wait sizes")
	//	wg.Wait() //等待上面for中的所有goroutine执行完毕
	//	//fmt.Println("close sizes")
	//	close(sizes) //关闭chan，好让main goroutine正常退出for
	//}()

	//放这里等待关闭也可以，就是必须等待所有goroutine执行完毕再执行下面，不像上面，可以同步进行，main中goroutine可以执行一部分
	//wg.Wait()
	//close(sizes)

	var total int
	for size := range sizes {
		total += size
	}

	return total
}

func main() {
	//for i:=1;i<10;i++{
	//	fmt.Println(rand.Intn(10) )
	//}

	ch := make(chan string, 3)
	for i := 0; i < 2; i++ {
		ch <- strconv.Itoa(i)
	}
	close(ch) //因为makeThumbnails4中的for f := range filenames是同一个goroutine，所以不关闭的化for会一直阻塞
	makeThumbnails4(ch)
}
