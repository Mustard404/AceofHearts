package models

import (
	"AceofHearts/config"
	"AceofHearts/utils"
	"fmt"
	"os"
)

func Sysoptimization() {

	utils.Title("系统优化")

	executablePath, _ := os.Executable()
	fmt.Printf("添加AceofHearts到Windows Defender白名单\n")
	result := utils.Cmd("powershell", "-Command", "Add-MpPreference", "-ExclusionPath", executablePath)
	if result {
		fmt.Println(config.Green("设置AceofHearts白名单成功！\n"))
	} else {
		fmt.Println(config.Red("设置AceofHearts白名单失败！\n"))
	}

	fmt.Printf("添加目录%s到Windows Defender白名单\n", config.ToolDirectory["Tools"])
	result = utils.Cmd("powershell", "-Command", "Add-MpPreference", "-ExclusionPath", config.ToolDirectory["ROOT"])
	if result {
		fmt.Println(config.Green("设置Windows Defender白名单成功！\n"))
		utils.Log("-\t设置Windows Defender白名单成功！\n")
	} else {
		fmt.Println(config.Red("设置Windows Defender白名单失败！\n"))
		utils.Log("-\t设置Windows Defender白名单失败！\n")
	}

	fmt.Printf("显示文件扩展名\n")
	result = utils.Cmd("reg", "add", "HKEY_CURRENT_USER\\Software\\Microsoft\\Windows\\CurrentVersion\\Explorer\\Advanced", "/v", "HideFileExt", "/t", "REG_DWORD", "/d", "0", "/f")
	if result {
		fmt.Println(config.Green("显示文件扩展名成功！\n"))
		utils.Log("-\t显示文件扩展名成功！\n")
	} else {
		fmt.Println(config.Red("显示文件扩展名失败！\n"))
		utils.Log("-\t显示文件扩展名失败！\n")
	}

	fmt.Printf("开启Telnet客户端\n")
	result = utils.Cmd("dism", "/Online", "/Enable-Feature", "/FeatureName:TelnetClient")
	if result {
		fmt.Println(config.Green("开启Telnet客户端成功！\n"))
		utils.Log("-\t开启Telnet客户端成功！\n")
	} else {
		fmt.Println(config.Red("开启Telnet客户端失败！\n"))
		utils.Log("-\t开启Telnet客户端失败！\n")
	}

	fmt.Printf("部署Chocolatey\n")
	result = utils.Cmd("powershell", "-Command", "Set-ExecutionPolicy Bypass -Scope Process -Force; iex ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'))")
	if result {
		fmt.Println(config.Green("部署Chocolatey成功！\n"))
		utils.Log("-\t部署Chocolatey成功！\n")
	} else {
		fmt.Println(config.Red("部署Chocolatey失败！\n"))
		utils.Log("-\t部署Chocolatey失败！\n")
	}
	// 添加Chocolatey到环境变量中
	PATH := os.Getenv("PATH")
	allusersprofile := os.Getenv("ALLUSERSPROFILE")
	updatedPath := allusersprofile + "\\chocolatey\\bin" + string(os.PathListSeparator) + PATH
	os.Setenv("Path", updatedPath)

}
