# errs

errs 是一个替代golang标准库errors的第三方包

## 快速使用

#### 下载安装

    go get github.com/keemis/library

#### 使用例子
```go
package main

import (
	"fmt"

	"github.com/keemis/library/errs"
)

func main() {
	// NEW
	err := errs.New("a php exception occurred")
	fmt.Println("1、", err)
	fmt.Printf("2、 %v \n", err)
	fmt.Printf("3、 %+v \n", err)

	// Errorf
	err = errs.Errorf("a %v exception occurred", "golang")
	fmt.Println("4、", err.Error())

	// Wrap
	err = errs.Wrap(err, "main")
	fmt.Println("5、", err)

	// Special
	err = errs.NewWithOption(errs.WithCode(501), errs.WithMsg("%v high-level exceptions", "two"), errs.WithData("stop"))
	fmt.Println("6、", err)
	fmt.Printf("7、 code: %v , msg: %v, data: %v \n", errs.GetCode(err), errs.GetMsg(err), errs.GetData(err))
	fmt.Printf("8、 stack: %v \n", errs.GetStack(err))
}
```

#### 运行结果

    go run main.go
    
    1、 a php exception occurred
    2、 a php exception occurred 
    3、 a php exception occurred
    main.main
            /usr/local/develop/src/test/main.go:11
    runtime.main
            /usr/local/go/src/runtime/proc.go:203
    runtime.goexit
            /usr/local/go/src/runtime/asm_amd64.s:1373 
    4、 a golang exception occurred
    5、 main: a golang exception occurred
    6、 two high-level exceptions
    7、 code: 501 , msg: two high-level exceptions, data: stop 
    8、 stack: 
    main.main
            /usr/local/develop/src/test/main.go:25
    runtime.main
            /usr/local/go/src/runtime/proc.go:203
    runtime.goexit
            /usr/local/go/src/runtime/asm_amd64.s:1373 
