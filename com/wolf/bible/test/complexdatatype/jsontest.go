package main

import (
	"encoding/json"
	"fmt"
	"log"
)

//JavaScript对象表示法（JSON）是一种用于发送和接收结构化信息的标准协议。
//由于简洁性、可读性和流行程度等原因，JSON是应用最广泛的一个
//encoding/json、encoding/xml、encoding/asn1
//JSON是对JavaScript中各种类型的值——字符串、数字、布尔值和对象——Unicode本文编码
//基本的JSON类型有数字（十进制或科学记数法）、布尔值（true或false）、字符串(以双引号包含的Unicode字符序列)
//boolean         true
//number          -273.15
//string          "She said \"Hello, BF\""
//array           ["gold", "silver", "bronze"]
//object          {"year": 1980,
//                 "event": "archery",
//                 "medals": ["gold", "silver", "bronze"]}

type Movie struct {
	Title string
	//字符串面值是结构体成员Tag，编译阶段关联到该成员的元信息字符串
	//通常是一系列用空格分隔的key:"value"键值对序列。因为值中含义双引号字符，因此一般用原生字符串面值的形式书写
	//json开头键名对应的值用于控制encoding/json包的编码和解码的行为
	//成员Tag中json对应值的第一部分用于指定JSON对象的名字,omitempty表示当Go语言结构体成员为空或零值时不生成JSON对象
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

func main() {
	//testMarshal()
	testUnmarshal()
}

func testMarshal() {
	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false,
			Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true,
			Actors: []string{"Paul Newman"}},
		{Title: "Bullitt", Year: 1968, Color: true,
			Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	}
	//在编码时，默认使用Go语言结构体的成员名字作为JSON的对象。只有导出的结构体成员才会被编码
	data, err := json.Marshal(movies) //marshaling,返回编码后的字节slice,没有空白缩进紧凑。
	//data, err := json.MarshalIndent(movies, "", "    ")
	//便于阅读
	if err != nil {
		log.Fatalf("JSON marshaling failed:%s", err)
	}
	fmt.Printf("%s\n", data)
}

func testUnmarshal() {
	data := []byte(`[{"Title":"Casablanca","released":1942,"Actors":["Humphrey Bogart","Ingrid Bergman"]},{"Title":"Cool Hand Luke","released":1967,"color":true,"Actors":["Paul Newman"]},{"Title":"Bullitt","released":1968,"color":true,"Actors":["Steve McQueen","Jacqueline Bisset"]}]`)
	var titles []struct{ Title string }
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Println(titles)
}
