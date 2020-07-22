package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

//sort包内置的提供了根据一些排序函数来对任何序列排序的功能
//使用了一个接口类型sort.Interface来指定通用的排序算法和可能被排序到的序列类型之间的约定
//这个接口的实现由序列的具体表示和它希望排序的元素决定，序列的表示经常是一个切片。

func testSort() {
	names := []string{"1", "5", "3"}
	sort.Sort(sort.StringSlice(names))
	sort.Strings(names)
	// 简化
	fmt.Println(names)
}

//每个track都是单独的一行，每一列都是属性
type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

//定义新的符合sort.Interface类型的切片类型
type byArtist []*Track

func (x byArtist) Len() int           { return len(x) }
func (x byArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
func (x byArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

//打印成表格
func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "_____", "______", "_____", "____", "______")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush()
}

//指向Track对象的指针，指针是一个机器字码长度而Track对象可能是八个或更多字节。便于快速交换数据
var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func testSortTrack() {
	sort.Sort(byArtist(tracks))
	printTracks(tracks)

	//内部reverse组合了Interface,Less使用了Interface.Less不过交换了下标
	sort.Sort(sort.Reverse(byArtist(tracks)))
	printTracks(tracks)
}

//对于byArtist/byYear，仅仅Less不一样，所以再进行抽取，仅仅需要提供less
type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

func testCustomSortTrack() {
	sort.Sort(customSort{tracks, func(x, y *Track) bool {
		if x.Title != y.Title {
			return x.Title < y.Title
		}
		if x.Year != y.Year {
			return x.Year < y.Year
		}
		if x.Length != y.Length {
			return x.Length < y.Length
		}
		return false
	}})
	printTracks(tracks)
}

func testIntsAreSorted() {
	values := []int{3, 1, 4, 1}
	fmt.Println(sort.IntsAreSorted(values))
	sort.Ints(values)
	fmt.Println(values)
	fmt.Println(sort.IntsAreSorted(values))
	sort.Sort(sort.Reverse(sort.IntSlice(values)))
	fmt.Println(values)
	fmt.Println(sort.IntsAreSorted(values))
}

func main() {
	//testSort()
	//testSortTrack()
	//testCustomSortTrack()
	testIntsAreSorted()
}
