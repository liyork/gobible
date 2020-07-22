package main

import "fmt"

func init() {
	fmt.Println("app/entry.go ==> init()")
}

var myVersion = fetchVersion()

//cd /test/testexecseq
//go run app/*.go
//├── 被执行的主包
//├── 初始化所有被导入的包
//|  ├── 初始化所有被导入的包 ( 递归定义 )
//|  ├── 初始化所有全局变量
//|  └── INIt 函数以字母序被调用
//└── 初始化主包
//     ├── 初始化所有全局变量
//     └── INIt 函数以字母序被调用

//本例中：
//先编译entry.go没有依赖，再编译fetch-version.go依赖test/testexecseq/version，编译entry.go无依赖，get-version.go无依赖
//初始化entry.go变量Version
//version/entry.go ==> getLocalVersion()
//调用到version/get-version.go
//version/get-version.go ==> getVersion()
//get-version.go无变量，则开始调用包version的init
//version/entry.go ==> init()
//version/get-version.go ==> init()
//version依赖初始完成，初始化包app中的变量
//app/fetch-version.go ==> fetchVersion()
//app中无变量，调用init方法
//app/entry.go ==> init()
//app/fetch-version.go ==> init()
//app/entry.go ==> main()
//version ===>  1.0.0

func main() {
	fmt.Println("app/entry.go ==> main()")
	fmt.Println("version ===> ", myVersion)
}
