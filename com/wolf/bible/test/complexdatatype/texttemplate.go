package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

//复杂的打印格式，一般需要将格式化代码分离出来以便更安全地修改。这些功能是由text/template和html/template等模板包提供的
//里面包含了一个或多个由双花括号包含的{{action}}对象
//对于每一个action，都有一个当前值的概念，对应点操作符，写作“.”。当前值“.”最初被初始化为调用模板时的参数
//action中，|操作符表示将前一个表达式的结果作为后一个函数的输入

const templ = `{{.TotalCount}} issues:
{{range .Items}}----------------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{end}}`

//time.Time类型对应的JSON值是一个标准时间格式的字符串
func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

var report *template.Template

func init() {
	//创建并分析模板(一般只需要执行一次),模板通常在编译时就测试好了
	report = template.Must( //检查模板必须存在
		template.New("issuelist"). //创建模板
			Funcs(template.FuncMap{"daysAgo": daysAgo}). //将daysAgo等自定义函数注册到模板中
			Parse(templ), //分析模板
	)
}

func testTextTemplateFormat() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

//go doc text/template