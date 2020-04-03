# copier

copier 是一个复制struct对象的第三方包

## 快速使用

#### 下载安装

    go get github.com/jinzhu/copier

#### 使用例子
```go
package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/copier"
)

type From struct {
	ID         int
	Name       string
	Age        int
	CreateTime time.Time
}

type To struct {
	ID     int
	Name   string
	Age    int
	Role   string
	Gender int
}

func main() {
	from := From{
		ID:         11,
		Name:       "simon",
		Age:        23,
		CreateTime: time.Now(),
	}
	to := To{}
	_ = copier.Copy(&to, &from)
	fmt.Printf("from: %+v \n", from)
	fmt.Printf("to: %+v \n", to)
}
```

#### 运行结果

    go run main.go 
    
	from: {ID:11 Name:simon Age:23 CreateTime:2020-04-03 16:19:43.987703 +0800 CST m=+0.000114347} 
	to: {ID:11 Name:simon Age:23 Role: Gender:0} 

        