# flag

flag 是一个打印显示程序版本信息的第三方包

## 快速使用

#### 下载安装

    go get github.com/keemis/library

#### 使用例子
```go
import (
    _ "github.com/keemis/library/flag" // 版本信息
)
```

#### 运行结果

    go build -ldflags "-X 'github.com/keemis/library/flag.GitVersion=c53b81e' -X 'github.com/keemis/library/flag.GitBranch=develop' -X 'github.com/keemis/library/flag.BuildStamp=1584942580' -X 'github.com/keemis/library/flag.GoVersion=go version go1.13.4 linux/amd64' -X 'github.com/keemis/library/flag.SysUname=Linux testvm01 x86_64' "
    
    ./app -v
     
    Version Information: 
             GitVersion: c53b81e 
             GitBranch: develop 
             BuildStamp: 2020-03-23 13:49:40 
             GoVersion: go version go1.13.4 linux/amd64 
             SysUname: Linux testvm01 x86_64 
