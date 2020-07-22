package main

import (
	"fmt"
	"golang.org/x/net/html"
)

//通常来说，不应该对panic异常做任何处理，但有时，也许我们可以从异常中恢复，至少我们可以在程序崩溃前，做一些操作。如关闭连接等资源

//若出现异常，不会让解析器崩溃，而是将panic异常当做普通的解析错误处理，并附加额外信息返回
func Parse(input string) (s string, err error) {
	defer func() {
		if p := recover(); p != nil { //recover会使程序从panic中恢复并返回panic value，未发生panic则recover返回nil
			err = fmt.Errorf("internal error:%v", p)
		}
	}()
	//parse string...
	var x = 0
	_ = 1 / x
	fmt.Println("normal flow..") //有异常不被执行
	return "", nil
}

//不加区分的恢复所有的panic异常，不是可取的做法；因为在panic之后，无法保证包级变量的状态仍然和我们预期一致。
// 比如，对数据结构的一次重要更新没有被完整完成、文件或者网络连接没有被关闭、获得的锁没有被释放。
// 此外，如果写日志时产生的panic被不加区分的恢复，可能会导致漏洞被忽略。
//不应该试图去恢复其他包引起的panic
//公有的API应该将函数的运行失败作为error返回，而不是panic
//只恢复应该被恢复的panic异常。这些异常所占的比例应该尽可能的低
//有些情况下，无法恢复。某些致命错误会导致Go在运行时终止程序，如内存不足。

func testParse() {
	_, err := Parse("sss")
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("main end..")
	//若没有recover，那么Parse产生异常，此行执行不了
}

//演示区分panic异常
// soleTitle returns the text of the first non-empty title element in doc,
// and an error if there was not exactly one.
func soleTitle(doc *html.Node) (title string, err error) {
	type bailout struct{}
	defer func() {
		switch p := recover(); p {
		case nil: //no panic
		case bailout{}: //"expected" panic
			err = fmt.Errorf("multiple tile elements")
		default:
			panic(p) // unexpected panic; carry on panicking
		}
	}()
	// Bail out of recursion if we find more than one nonempty title.
	forEachNode(doc, func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			if title != "" {
				panic(bailout{}) // multiple title elements，只为了演示机制，对可预期的错误不应采用panic
			}
			title = n.FirstChild.Data
		}
	}, nil)
	if title == "" {
		return "", fmt.Errorf("no title element")
	}
	return title, nil
}

func main() {
	//testParse()
	soleTitle(nil)
}
