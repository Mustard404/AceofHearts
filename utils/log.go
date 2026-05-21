package utils

import (
	"AceofHearts/config"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// Log 追加写入日志到 LOG.MD
func Log(format string, args ...interface{}) {
	logfile := filepath.Join(config.ExecutablePath, "LOG.MD")
	f, err := os.OpenFile(logfile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return
	}
	defer f.Close()
	fmt.Fprintf(f, format, args...)
}

// InitLog 初始化日志文件
func InitLog() {
	logfile := filepath.Join(config.ExecutablePath, "LOG.MD")
	os.Remove(logfile)
	Log("# AceofHearts 部署日志\n")
	Log("- 时间: %s\n", time.Now().Format("2006-01-02 15:04:05"))
	Log("- 版本: %s\n\n", config.Version)
}
