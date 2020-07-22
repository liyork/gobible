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

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		//sleep1Second()
		if entry.IsDir() {
			n.Add(1)
			//sleep1Second()
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, n, fileSizes)
		} else {
			//sleep1Second()
			fileSizes <- entry.Size()
		}
	}
}

func sleep1Second() {
	time.Sleep(1 * time.Second)
}

// sema is a counting semaphore for limiting concurrency in dirents.
var sema = make(chan struct{}, 20)

// dirents return the entries of directory dir.
func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}        // acquire token
	defer func() { <-sema }() // release token
	entries, err := ioutil.ReadDir(dir)
	sleep1Second()
	if err != nil {
		fmt.Fprint(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

var verbose = flag.Bool("v", false, "show verbose progress messges")

//go run traveldir.go -v $GOPATH/src/test
func main() {
	traverseDir()
}

func traverseDir() {
	// Determine the initial directories.
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	// Traverse each root of the file tree in paralle.
	fileSizes := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, fileSizes) // 由于慢在访问磁盘上，所以开启新goroutine对每个root
	}
	go func() { //新开goroutine进行等待所有walkDir的goroutien执行完毕，关闭chan
		n.Wait()
		close(fileSizes)
	}()
	// Print the results periodically.
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(1000 * time.Microsecond)
	}
	var nfiles, nbytes int64
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes was closed
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage("tick:", nfiles, nbytes)
		}
	}
	printDiskUsage("total:", nfiles, nbytes)
	// final totals
}

func printDiskUsage(prefix string, nfiles, nbytes int64) {
	fmt.Printf(prefix+"%d files  %.1f KB\n", nfiles, float64(nbytes)/1024)
}
