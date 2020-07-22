package main

import (
	"fmt"
	"strings"
)

//bytes、strings、strconv和unicode包对字符串处理

// basename removes directory components and a .suffix.
// e.g., a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
func basename(s string) string {
	//Discard last '/' and everything before.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	// Preserve everything before last '.'.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

//使用strings.LastIndex
func basename1(s string) string {
	slash := strings.LastIndex(s, "/") // -1 if "/" not found
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

func main() {
	//fmt.Println(basename("a/b/c.go")) // "c"
	//fmt.Println(basename("c.d.go"))   // "c.d"
	//fmt.Println(basename("abc"))      // "abc"

	fmt.Println(basename1("a/b/c.go")) // "c"
	fmt.Println(basename1("c.d.go"))   // "c.d"
	fmt.Println(basename1("abc"))      // "abc"
}
