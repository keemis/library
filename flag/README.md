# Golang 程序显示版本信息

**1、引入：**
 
```
import (
    _ "github.com/keemis/library/flag" // 版本信息
)
```
 
**2、导入：**
  
 ```
go build -ldflags "-X 'github.com/keemis/library/flag.GitVersion=c53b81e' -X 'github.com/keemis/library/flag.GitBranch=develop' -X 'github.com/keemis/library/flag.BuildStamp=1584942580' -X 'github.com/keemis/library/flag.GoVersion=go version go1.13.4 linux/amd64' -X 'github.com/keemis/library/flag.SysUName=Linux testvm01 x86_64' "
 ```
 
**3、输出：**
 
```
./app -v
 
Version Information: 
         GitVersion: c53b81e 
         GitBranch: develop 
         BuildStamp: 2020-03-23 13:49:40 
         GoVersion: go version go1.13.4 linux/amd64 
         SysUname: Linux testvm01 x86_64 
```
 
 