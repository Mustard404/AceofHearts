package models

import (
	"AceofHearts/config"
	"AceofHearts/utils"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Sysoptimization 系统优化：白名单、Chocolatey、代理设置
// 通过标记文件判断是否已执行过，断点续装时跳过
func Sysoptimization() {
	utils.Title("系统优化")

	markFile := filepath.Join(config.ExecutablePath, ".sysopt_done")
	if _, err := os.Stat(markFile); err == nil {
		fmt.Println(config.Green("系统优化已完成，跳过"))
		// 仍需确保 Chocolatey 在 PATH 中
		ensureChocoPath()
		return
	}

	execPath, _ := os.Executable()

	// Windows Defender 白名单
	fmt.Println("添加 AceofHearts 到 Windows Defender 白名单")
	utils.Cmd("powershell", "-Command", "Add-MpPreference", "-ExclusionPath", execPath)

	fmt.Printf("添加 %s 到 Windows Defender 白名单\n", config.ToolsDir)
	if utils.Cmd("powershell", "-Command", "Add-MpPreference", "-ExclusionPath", config.ToolsDir) {
		fmt.Println(config.Green("设置白名单成功!"))
		utils.Log("-\t设置 Windows Defender 白名单成功\n")
	}

	// 显示文件扩展名
	fmt.Println("显示文件扩展名")
	utils.Cmd("reg", "add", `HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\Explorer\Advanced`,
		"/v", "HideFileExt", "/t", "REG_DWORD", "/d", "0", "/f")

	// 开启 Telnet 客户端
	fmt.Println("开启 Telnet 客户端")
	utils.Cmd("dism", "/Online", "/Enable-Feature", "/FeatureName:TelnetClient")

	// 部署 Chocolatey
	fmt.Println("部署 Chocolatey")
	if utils.Cmd("powershell", "-Command",
		"Set-ExecutionPolicy Bypass -Scope Process -Force; iex ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'))") {
		fmt.Println(config.Green("Chocolatey 部署成功!"))
		utils.Log("-\tChocolatey 部署成功\n")
	}

	ensureChocoPath()

	// 设置 Chocolatey 代理
	if config.ProxyURL != "" {
		fmt.Printf("设置 Chocolatey 代理: %s\n", config.ProxyURL)
		utils.Cmd("choco", "config", "set", "proxy", config.ProxyURL)
		utils.Log("-\tChocolatey 代理: %s\n", config.ProxyURL)
	}

	// 标记完成
	os.WriteFile(markFile, []byte("done"), 0644)
}

func ensureChocoPath() {
	chocoPath := filepath.Join(os.Getenv("ALLUSERSPROFILE"), "chocolatey", "bin")
	if _, err := os.Stat(chocoPath); err != nil {
		// Chocolatey 未安装，跳过
		return
	}

	currentPATH := os.Getenv("PATH")
	if !strings.Contains(strings.ToLower(currentPATH), strings.ToLower(chocoPath)) {
		newPATH := chocoPath + string(os.PathListSeparator) + currentPATH
		os.Setenv("PATH", newPATH)
		os.Setenv("Path", newPATH)
	}
}
