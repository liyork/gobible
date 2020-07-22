package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

//panic是来自被调函数的信号，表示发生了某个已知的bug。一个良好的程序永远不应该发生panic异常。
//对于大部分函数而言，永远无法确保能否成功运行
//当本该可信的操作出乎意料的失败后，我们必须弄清楚导致失败的原因
//错误是软件包API和应用程序用户界面的一个重要组成部分
//对于那些将运行失败看作是预期结果的函数，它们会返回一个额外的返回值，通常是最后一个，来传递错误信息。
//如果导致失败的原因只有一个，额外的返回值可以是一个布尔值，通常被命名为ok
//如果导致失败的原因不止一种，额外的返回值是error类型,error值可能是nil或者non-nil
//通常，当函数返回non-nil的error时，其他的返回值是未定义的(undefined),这些未定义的返回值应该被忽略

//在Go中，函数运行失败时会返回错误信息，这些错误信息被认为是一种预期的值而非异常（exception），
// 虽然Go有各种异常机制，但这些机制仅被使用在处理那些未被预料到的错误，即bug，而不是那些在健壮程序中应该被避免的程序错误。
//Go这样设计的原因是由于对于某个应该在控制流程中处理的错误而言，将这个错误以异常的形式抛出会混乱对错误的描述，这通常会导致一些糟糕的后果。
// 当某个程序错误被当作异常处理后，这个错误会将堆栈根据信息返回给终端用户，这些信息复杂且无用，无法帮助定位错误。
//正因此，Go使用控制流机制（如if和return）处理异常，这使得编码人员能更多的关注错误处理。

//直接将错误返回给调用者
func errProcess1() (string, error) {
	_, err := http.Get("url")
	if err != nil {
		return "", err
	}
	return "", nil
}

//当对html.Parse的调用失败时，errProcess2不会直接返回html.Parse的错误，因为缺少两条重要信息：
// 1、错误发生在解析器；2、url已经被解析。
// 这些信息有助于错误的处理，errProcess2会构造新的错误信息返回给调用者：
//由于错误信息经常是以链式组合在一起的，所以错误信息中应避免大写和换行符
//编写错误信息时，我们要确保错误信息对问题细节的描述是详尽的。尤其是要注意错误信息表达的一致性，即相同的函数或同包内的同一组函数返回的错误在构成和处理方式上是相似的
//一般而言，被调函数f(x)会将调用信息和参数信息作为发生错误时的上下文放在错误信息中并返回给调用者，调用者需要添加一些错误信息中不包含的信息，比如添加url到html.Parse返回的错误中
func errProcess2() (string, error) {
	var url string
	var body io.Reader
	_, err := html.Parse(body)
	if err != nil {
		//添加额外的上下文信息到原始错误信息
		return "", fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	return "", nil
}

//如果错误的发生是偶然性的，或由不可预知的问题导致的。一个明智的选择是重新尝试失败的操作。
// 在重试时，我们需要限制重试的时间间隔或重试的次数，防止无限制的重试。
func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil // success
		}
		log.Printf("server not responding (%s);retrying...", err)
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}

//如果错误发生后，程序无法继续运行，输出错误信息并结束程序，这种策略只应在main中执行
//对库函数而言，应仅向上传播错误，除非该错误意味着程序内部包含不一致性，即遇到了bug，才能在库函数中结束程序。
func main() {
	var url string
	if err := WaitForServer(url); err != nil {
		log.Fatalf("Site is down:%v\n", err)
	}
}

//有时，我们只需要输出错误信息就足够了，不需要中断程序的运行
//if err := Ping(); err != nil {
//log包中的所有函数会为没有换行符的字符串增加换行符。
//    log.Printf("ping failed: %v; networking disabled",err)
//}
//或者标准错误流输出错误信息。
//if err := Ping(); err != nil {
//    fmt.Fprintf(os.Stderr, "ping failed: %v; networking disabled\n", err)
//}

//我们可以直接忽略掉错误
//我们应该在每次函数调用后，都养成考虑错误处理的习惯，当你决定忽略某个错误时，你应该在清晰的记录下你的意图。

//Go中大部分函数的代码结构几乎相同，首先是一系列的初始检查，防止错误发生，之后是函数的实际逻辑。
func processReadErr() error {
	in := bufio.NewReader(os.Stdin)
	for {
		r, _, err := in.ReadRune()
		if err == io.EOF {
			break // finished reading
		}
		if err != nil {
			return fmt.Errorf("read failed:%v", err)
		}
		fmt.Println(r)
	}
	return nil
}
