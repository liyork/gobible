package main

import (
	"fmt"
	"github.com/liyork/bible/com/wolf/bible/test/testexecseq/version"
)

func init() {
	fmt.Println("app/fetch-version.go ==> init()")
}
func fetchVersion() string {
	fmt.Println("app/fetch-version.go ==> fetchVersion()")
	return version.Version
}
