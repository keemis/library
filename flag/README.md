# Golang 程序版本显示版本信息

1、引入：
 
```
import (
    _ "github.com/keemis/library/flag" // 版本信息
)
```
 
2、导入：
  
 ```
go build -ldflags "-X github.com/keemis/library/flag.GitVersion=36fe168"
 ```
 
3、输出：
 
```
./app -v
 
Version Information: 
         GitVersion: c53b81e 
         GitBranch: develop 
         BuildStamp: 2020-03-23 13:37:40 
         GoVersion: go version go1.13.4 linux/amd64 
         SysUname: Linux testvm01 x86_64 
```
 