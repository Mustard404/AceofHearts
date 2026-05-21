package models

import (
	"AceofHearts/config"
	"AceofHearts/utils"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"time"

	"github.com/StackExchange/wmi"
)

type win32OS struct {
	CSName           string
	Caption          string
	Version          string
	OSArchitecture   string
	WindowsDirectory string
	Description      string
}

// Sysinfo 获取系统信息并验证管理员权限
func Sysinfo() {
	utils.Title("系统信息获取")

	var osInfo []win32OS
	if err := wmi.Query("SELECT CSName, Caption, Version, OSArchitecture, WindowsDirectory, Description FROM Win32_OperatingSystem", &osInfo); err != nil {
		fmt.Println(config.Red("获取系统信息失败:"), err)
		return
	}

	currentUser, _ := user.Current()

	if len(osInfo) > 0 {
		info := osInfo[0]
		fields := [][2]string{
			{"电脑名称", info.CSName},
			{"操作系统", info.Caption},
			{"系统版本", info.Version},
			{"系统位数", info.OSArchitecture},
			{"系统路径", info.WindowsDirectory},
			{"激活状态", info.Description},
			{"当前用户", currentUser.Username},
		}
		for _, f := range fields {
			fmt.Printf("%-8s: %s\n", f[0], f[1])
			utils.Log("-\t%s: %s\n", f[0], f[1])
		}
	}

	if isAdmin() {
		fmt.Println("管理员权限: " + config.Green("是"))
	} else {
		fmt.Println("管理员权限: " + config.Red("否"))
		fmt.Println(config.Red("请以管理员身份重新运行!"))
		time.Sleep(5 * time.Second)
		os.Exit(1)
	}
}

func isAdmin() bool {
	return exec.Command("net", "session").Run() == nil
}
