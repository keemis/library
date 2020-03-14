package flag

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

var (
	GitVersion = "unknown"
	GitBranch  = "unknown"
	BuildStamp = "unknown"
	GoVersion  = "unknown"
	SysUName   = "unknown"
	VerOpt     = flag.Bool("v", false, "Print application version")
)

// init 初始化
// go build -ldflags "-X main.GitVersion=36fe168 -X main.GitBranch=develop"
func init() {
	flag.Parse()
	if *VerOpt {
		t, _ := strconv.ParseInt(BuildStamp, 10, 64)
		fmt.Printf("\tGitVersion: %v \n", GitVersion)
		fmt.Printf("\tGitBranch: %v \n", GitBranch)
		fmt.Printf("\tBuildStamp: %v \n", time.Unix(t, 0).Format("2006-01-02 15:04:05"))
		fmt.Printf("\tGoVersion: %v \n", GoVersion)
		fmt.Printf("\tSysUname: %v \n", SysUName)
		os.Exit(0)
	}
}
