package main

import (
	"AceofHearts/config"
	"AceofHearts/install"
	"AceofHearts/models"
	"AceofHearts/utils"
	"os"
	"path/filepath"
)

func main() {
	// 初始化路径
	exe, _ := os.Executable()
	config.ExecutablePath = filepath.Dir(exe)

	// 初始化日志
	utils.InitLog()

	// 启动流程
	models.Banner()
	models.Setting()
	utils.CheckMirror()
	models.Sysinfo()
	models.Sysoptimization()

	// 从 Excel 清单安装（支持断点续装）
	install.InstallFromXlsx(config.XlsxPath)

	// 启动器 & Doskey
	install.Launcher()
	install.Doskey()
}
