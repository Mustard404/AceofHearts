package main

import (
	"AceofHearts/config"
	"AceofHearts/install"
	"AceofHearts/models"
	"os"
	"path/filepath"
)

func main() {

	executablePath, _ := os.Executable()
	config.ExecutablePath = filepath.Dir(executablePath)
	logFile := filepath.Join(config.ExecutablePath, "LOG.MD")

	os.RemoveAll(logFile)
	models.Banner()
	models.Setting()
	models.Sysinfo()
	models.Sysoptimization()
	install.Env()
	install.Productivity()
	install.Development()
	install.PenetrationTestingTools()
	install.Launcher()
	models.Doskey()

}
