package main

import (
	"database/sql"
	"fmt"
)

//Go语言查询一个SQL数据库的API会干净地将查询中固定的部分和变化的部分分开
func listTracks(db sql.DB, artist string, minYear, maxYear int) {
	result, err := db.Exec( //避免sql注入攻击
		"SELECT * FROM tracks WHERE artist = ? AND ? <= year AND year <= ?",
		artist, minYear, maxYear)
	if err != nil {
		fmt.Println(result)
	}
}

//sql执行内部可能有这样基于类型判断
func sqlQuote(x interface{}) string {
	if x == nil {
		return "NULL"
	} else if _, ok := x.(int); ok {
		return fmt.Sprintf("%d", x)
	} else if _, ok := x.(uint); ok {
		return fmt.Sprintf("%d", x)
	} else if b, ok := x.(bool); ok {
		if b {
			return "TRUE"
		}
		return "FALSE"
	} else if s, ok := x.(string); ok {
		return "sqlQuoteString(s)" + s // (not shown)
	} else {
		panic(fmt.Sprintf("unexpected type %T: %v", x, x))
	}
}

func main() {
	//基于类型判断
	//switch x.(type) {
	//case nil:       // ...
	//case int, uint: // ...
	//case bool:      // ...
	//case string:    // ...
	//default:        // ...
	//}

	//当一个或多个case类型是接口时，case的顺序就会变得很重要

	//sqlQuote中对于bool和string需要通过类型断言访问提取的值。类型开关语句扩展形式，将提取的值绑定到一个在每个case范围内的新变量
	var x interface{}
	switch x := x.(type) { //一个类型开关隐式的创建了一个语言块，因此新变量x的定义不会和外面块中的x变量冲突
	case nil:
		fmt.Println("NULL")
	case int:
		fmt.Println("int")
	case bool:
		fmt.Println(x) //x是bool类型
	case string: //每一个case也会隐式的创建一个单独的语言块。
		fmt.Println(x) //x是string类型
	default:
		panic(fmt.Sprintf("unexpected type %T:%v", x, x))
	}
}
