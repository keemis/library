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

// go build -ldflags "-X github.com/keemis/library/flag.GitVersion=36fe168"
func init() {
	flag.Parse()
	if *VerOpt {
		stamp := "unknown"
		t, _ := strconv.ParseInt(BuildStamp, 10, 64)
		if t > 0 {
			stamp = time.Unix(t, 0).Format("2006-01-02 15:04:05")
		}
		fmt.Printf("\tGitVersion: %v \n", GitVersion)
		fmt.Printf("\tGitBranch: %v \n", GitBranch)
		fmt.Printf("\tBuildStamp: %v \n", stamp)
		fmt.Printf("\tGoVersion: %v \n", GoVersion)
		fmt.Printf("\tSysUname: %v \n", SysUName)
		os.Exit(0)
	}
}
