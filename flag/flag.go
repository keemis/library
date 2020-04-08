package flag

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/keemis/library/timer"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/process"
)

var (
	GitVersion = "unknown"
	GitBranch  = "unknown"
	BuildStamp = "unknown"
	GoVersion  = "unknown"
	SysUname   = "unknown"

	VerOpt = flag.Bool("v", false, "Print application version")
	SysOpt = flag.Bool("s", false, "Print system information")
)

// 初始化，入口
func Init() {
	if !flag.Parsed() {
		flag.Parse()
	}

	if *VerOpt {
		printVer()
		os.Exit(0)
	}

	if *SysOpt {
		printSystem()
		os.Exit(0)
	}
}

// go build -ldflags "-X github.com/keemis/library/flag.GitVersion=36fe168"
func printVer() {
	stamp := "unknown"
	t, _ := strconv.ParseInt(BuildStamp, 10, 64)
	if t > 0 {
		stamp = time.Unix(t, 0).Format(timer.GoTimeFormat)
	}
	fmt.Printf("Version Information: \n")
	fmt.Printf("\tGitVersion: %v \n", GitVersion)
	fmt.Printf("\tGitBranch: %v \n", GitBranch)
	fmt.Printf("\tBuildStamp: %v \n", stamp)
	fmt.Printf("\tGoVersion: %v \n", GoVersion)
	fmt.Printf("\tSysUname: %v \n", SysUname)
}

// 打印当前系统信息
func printSystem() {
	// Memory
	fmt.Printf("Memory Information: \n")
	if v, err := mem.VirtualMemory(); err == nil {
		fmt.Printf("\tTotal: %.2f GB \n", float64(v.Total)/1024/1024/1024)
		fmt.Printf("\tAvailable: %.2f GB \n", float64(v.Available)/1024/1024/1024)
		fmt.Printf("\tUsed Percent: %.2f%% \n", v.UsedPercent)
	}
	fmt.Println()
	// Disk
	fmt.Printf("Disk Information: \n")
	if v, err := disk.Usage("/"); err == nil {
		fmt.Printf("\tTotal: %.2f GB \n", float64(v.Total)/1024/1024/1024)
		fmt.Printf("\tFree: %.2f GB \n", float64(v.Free)/1024/1024/1024)
		fmt.Printf("\tUsed Percent: %.2f%% \n", 100-float64(v.Free)/float64(v.Total)*100)
	}
	fmt.Println()
	// CPU
	fmt.Printf("CPU Information: \n")
	if vs, err := cpu.Info(); err == nil {
		for _, v := range vs {
			fmt.Printf("\tModel: %v \n", v.ModelName)
			break
		}
	}
	if v, err := cpu.Counts(false); err == nil {
		fmt.Printf("\tPhysical Cores: %v \n", v)
	}
	if v, err := cpu.Counts(true); err == nil {
		fmt.Printf("\tLogical Cores: %v \n", v)
	}
	if vs, err := cpu.Percent(time.Millisecond*250, false); err == nil {
		for _, v := range vs {
			fmt.Printf("\tTotal Percent: %.2f%% \n", v)
			break
		}
	}
	fmt.Println()
	// Working
	fmt.Printf("Working Information: \n")
	if vs, err := host.Users(); err == nil {
		fmt.Print("\tUser: ")
		for _, user := range vs {
			fmt.Print(user.User, "  ")
		}
		fmt.Print("\n")
	}
	if vs, err := process.Processes(); err == nil {
		fmt.Printf("\tProcess Count: %v \n", len(vs))
	}
	if vs, err := net.Connections(""); err == nil {
		fmt.Printf("\tConnection Count: %v \n", len(vs))
	}
	fmt.Println()
	// System
	fmt.Printf("System Information: \n")
	if v, err := host.BootTime(); err == nil {
		fmt.Printf("\tBoot Time: %v \n", time.Unix(int64(v), 0).Local().Format(timer.GoTimeFormat))
	}
	if platform, family, version, err := host.PlatformInformation(); err == nil {
		fmt.Printf("\tPlatform: %v \n", platform)
		fmt.Printf("\tFamily: %v \n", family)
		fmt.Printf("\tVersion: %v \n", version)
	}
	if v, err := host.KernelVersion(); err == nil {
		fmt.Printf("\tKernel Version: %v \n", v)
	}
	fmt.Println()
}
