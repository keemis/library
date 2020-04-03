# gjson

gjson 是一个类似XPath的路径解析JSON的第三方包

## 快速使用

#### 下载安装

    go get github.com/tidwall/gjson

#### 使用例子
```go
package main

import (
	"fmt"

	"github.com/tidwall/gjson"
)

const jsonStr = `
	{
	  "name": {"first": "Tom", "last": "Anderson"},
	  "age":37,
	  "children": ["Sara","Alex","Jack","Cat"],
	  "fav.movie": "Deer Hunter",
	  "friends": [
		{"first": "Dale", "last": "Murphy", "age": 44, "nets": ["ig", "fb", "tw"]},
		{"first": "Roger", "last": "Craig", "age": 68, "nets": ["fb", "tw"]},
		{"first": "Jane", "last": "Murphy", "age": 47, "nets": ["ig", "tw"]}
	  ]
	}
`

func main() {
	val := gjson.Get(jsonStr, "name.first")
	fmt.Printf("name.first: %+v \n", val.String())

	val = gjson.Get(jsonStr, "age")
	fmt.Printf("age: %+v \n", val.String())

	val = gjson.Get(jsonStr, "friends.#")
	fmt.Printf("friends.#: %+v \n", val.String())

	val = gjson.Get(jsonStr, "friends.1.first")
	fmt.Printf("friends.1.first: %+v \n", val.String())

	val = gjson.Get(jsonStr, "friends.#.first")
	fmt.Printf("friends.#.first: %+v \n", val.Array()[0].String())
}
```

#### 运行结果

    go run main.go 
    
	name.first: Tom 
	age: 37 
	friends.#: 3 
	friends.1.first: Roger 
	friends.#.first: Dale 
