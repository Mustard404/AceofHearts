package utils

import (
	"AceofHearts/config"
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func Log(format string, args ...interface{}) {

	logfile := filepath.Join(config.ExecutablePath, "LOG.MD")
	file, err := os.OpenFile(logfile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("无法打开日志文件:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	log := fmt.Sprintf(format, args...)

	_, err = writer.WriteString(log)
	if err != nil {
		fmt.Println("无法写入日志:", err)
		return
	}

	err = writer.Flush()
	if err != nil {
		fmt.Println("无法刷新日志缓冲区:", err)
		return
	}
}
