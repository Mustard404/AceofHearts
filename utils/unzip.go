package utils

import (
	"AceofHearts/config"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func Unzip(zipFilePath string, outputDir string, move bool) {
	if move {
		tempDir := filepath.Dir(outputDir)
		cmdPath := "Bandizip.exe"
		cmdArgs := []string{"bx", "-y", fmt.Sprintf("-o:%s", tempDir), zipFilePath}
		Cmd(cmdPath, cmdArgs...)
		// 让子弹飞一会。。。
		time.Sleep(3 * time.Second)
		os.RemoveAll(outputDir)
		baseName := filepath.Base(zipFilePath)
		baseName = strings.TrimSuffix(baseName, filepath.Ext(baseName))
		tempDir = filepath.Join(tempDir, baseName)
		err := os.Rename(tempDir, outputDir)
		if err != nil {
			fmt.Println(config.Red("重命名文件夹失败:"), err)
			return
		}
	} else {
		cmdPath := "Bandizip.exe"
		cmdArgs := []string{"bx", "-y", fmt.Sprintf("-o:%s", outputDir), zipFilePath}
		Cmd(cmdPath, cmdArgs...)
	}
}
