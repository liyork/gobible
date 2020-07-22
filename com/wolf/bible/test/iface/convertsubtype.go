package main

import (
	"io"
)

// io包中有相同实现，
// writeString writes s to w.
// if w has a WriteString method, it is invoked instead of w.Write.
func writesString(w io.Writer, s string) (n int, err error) {
	//定义内部接口，为了下面断言是否有相同方法
	//定义一个特定类型的方法隐式地获取了对特定行为的协约。
	type stringWriter interface {
		WriteString(string) (n int, err error)
	}
	//这里涉及接口具体化，是否有真实必要?可能具体子类实现时需要进行特定处理。
	if sw, ok := w.(stringWriter); ok {
		return sw.WriteString(s) // avoid a copy
	}
	bytes := []byte(s) //转换需要分配内存并做拷贝，很快就不用了(临时拷贝)，可能影响性能
	return w.Write(bytes)
}

func writeHeader(w io.Writer, contentType string) error {
	if _, err := writesString(w, "Content-Type"); err != nil {
		return err
	}
	if _, err := writesString(w, contentType); err != nil {
		return err
	}
	// ...
	return nil
}

func main() {
	writeHeader(nil, "")
}
