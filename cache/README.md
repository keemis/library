# cache

cache 是一个操作缓存的第三方包

## 快速使用

#### 下载安装

    go get github.com/keemis/library

#### 使用例子
```go
package main

import (
	"fmt"

	"github.com/keemis/library/cache"
)

func main() {
	// init
	cache.Init(5 * 1024 * 1024)

	// set
	cache.Client().Set([]byte("key"), []byte("val_xxx"), 0)

	// get
	v, err := cache.Client().Get([]byte("key"))
	fmt.Printf("value: %v, err: %v \n", string(v), err)

	// del
	affected := cache.Client().Del([]byte("key"))
	fmt.Printf("affected: %v \n", affected)

	// get
	v, err = cache.Client().Get([]byte("key"))
	fmt.Printf("value: %v, err: %v \n", string(v), err)
}
```

#### 运行结果

    go run main.go 
    
    value: val_xxx, err: <nil> 
    affected: true 
    value: , err: Entry not found 
     
     