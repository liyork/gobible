package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

//演示请求github返回json数据，然后编解码

const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

// SearchIssues queries the GitHub issue tracker
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	join := strings.Join(terms, " ")
	q := url.QueryEscape(join) //对查询中的特殊字符进行转义操作
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	//基于流式的解码器json.Decoder，可以从一个输入流解码JSON数据
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return &result, nil
}

func main() {
	//go build json-github.go
	//./json-github repo:golang/go is:open json decoder
	//testManualFormat()

	//go build texttemplate.go json-github.go
	//./texttemplate repo:golang/go is:open json decoder
	//testTextTemplateFormat()

	//go build htmltemplate.go json-github.go
	//./htmltemplate repo:golang/go commenter:gopherbot json encoder >issues.html
	//浏览器打开issues.html
	//标题中含有&和<字符的issue：./htmltemplate repo:golang/go 3133 10535 >issues2.html
	//
	//testHtmlTemplateFormat()

	trustHTML()
}

func testManualFormat() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}

func trustHTML() {
	const templ = `<p>A: {{.A}}</p><p>B: {{.B}}</p>`
	t := template.Must(template.New("escape").Parse(templ))

	var data struct {
		A string        // untrusted plain text
		B template.HTML // trusted HTML，抑制自动转义的行为
	}
	data.A = "<b>Hello!</b>" //转义失效
	data.B = "<b>Hello!</b>"

	if err := t.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}
}
