package models

import (
	"AceofHearts/config"
	"AceofHearts/utils"
	"fmt"
	"github.com/StackExchange/wmi"
	"os"
	"os/exec"
	"os/user"
	"time"
)

type Win32OperatingSystem struct {
	CSName           string
	Caption          string
	Version          string
	OSArchitecture   string
	WindowsDirectory string
	Description      string
}

func isAdmin() (bool, error) {
	cmd := exec.Command("net", "session")
	err := cmd.Run()
	if err != nil {
		return false, nil
	}
	return true, nil
}

func Sysinfo() {

	utils.Title("系统信息获取")

	var osInfo []Win32OperatingSystem
	err := wmi.Query("SELECT CSName, Caption, Version, OSArchitecture, WindowsDirectory, Description FROM Win32_OperatingSystem", &osInfo)
	if err != nil {
		fmt.Println(config.Red("获取系统信息失败:"), err)
		return
	}

	currentUser, err := user.Current()
	if err != nil {
		fmt.Println(config.Red("获取用户信息失败:"), err)
		return
	}

	isAdmin, err := isAdmin()
	if err != nil {
		fmt.Println(config.Red("获取用户权限失败:"), err)
		return
	}

	if len(osInfo) > 0 {
		SystemOs := osInfo[0]
		fmt.Printf("电脑名称: %s\n", SystemOs.CSName)
		fmt.Printf("操作系统名称: %s\n", SystemOs.Caption)
		fmt.Printf("操作系统版本: %s\n", SystemOs.Version)
		fmt.Printf("操作系统位数: %s\n", SystemOs.OSArchitecture)
		fmt.Printf("操作系统路径: %s\n", SystemOs.WindowsDirectory)
		fmt.Printf("系统激活状态: %s\n", SystemOs.Description)
		fmt.Printf("当前用户名: %s\n", currentUser.Username)
		utils.Log("-\t电脑名称: %s\t\n", SystemOs.CSName)
		utils.Log("-\t操作系统名称: %s\t\n", SystemOs.Caption)
		utils.Log("-\t操作系统版本: %s\t\n", SystemOs.Version)
		utils.Log("-\t操作系统位数: %s\t\n", SystemOs.OSArchitecture)
		utils.Log("-\t操作系统路径: %s\t\n", SystemOs.WindowsDirectory)
		utils.Log("-\t系统激活状态: %s\t\n", SystemOs.Description)
		utils.Log("-\t当前用户名: %s\t\n", currentUser.Username)
	}

	if isAdmin {
		fmt.Println("管理员权限：" + config.Green("是"))
	} else {
		fmt.Println("管理员权限：" + config.Red("否"))
		fmt.Println(config.Red("请重新以管理员身份运行..."))
		time.Sleep(time.Minute)
		os.Exit(0)
	}
}
