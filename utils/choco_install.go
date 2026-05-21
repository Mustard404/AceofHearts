package utils

import (
	"AceofHearts/config"
	"context"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

const chocoTimeout = 10 * time.Minute

// ChocoInstall 使用 Chocolatey 安装软件包，超时自动跳过
func ChocoInstall(name, packageName, directory, extraArgs string) error {
	cmdArgs := []string{"install", packageName, "-y"}
	if extraArgs != "" {
		cmdArgs = append(cmdArgs, splitArgs(extraArgs)...)
	}

	ctx, cancel := context.WithTimeout(context.Background(), chocoTimeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, "choco", cmdArgs...)
	fmt.Println(config.Magenta("[AceofHearts]#:"), strings.Join(cmd.Args, " "))

	output, err := cmd.CombinedOutput()
	if ctx.Err() == context.DeadlineExceeded {
		return fmt.Errorf("%s 安装超时 (%v)", name, chocoTimeout)
	}
	if err != nil {
		fmt.Println(string(output))
		return fmt.Errorf("%s 部署失败: %v", name, err)
	}

	fmt.Println(config.Green(fmt.Sprintf("%s 部署成功!", name)))
	version := chocoVersion(packageName)
	Log("-\t%s 部署成功!\n\t-\t版本: %s\n\t-\t路径: %s\n", name, version, directory)
	return nil
}

// splitArgs 按空格分割参数，但保留引号内的内容作为整体
func splitArgs(s string) []string {
	var args []string
	var current strings.Builder
	inQuote := false
	quoteChar := byte(0)

	for i := 0; i < len(s); i++ {
		c := s[i]
		if !inQuote && (c == '"' || c == '\'') {
			inQuote = true
			quoteChar = c
			current.WriteByte(c)
		} else if inQuote && c == quoteChar {
			inQuote = false
			current.WriteByte(c)
		} else if !inQuote && c == ' ' {
			if current.Len() > 0 {
				args = append(args, current.String())
				current.Reset()
			}
		} else {
			current.WriteByte(c)
		}
	}
	if current.Len() > 0 {
		args = append(args, current.String())
	}
	return args
}

// chocoVersion 查询已安装的 Chocolatey 包版本
func chocoVersion(packageName string) string {
	output, err := exec.Command("choco", "list").Output()
	if err != nil {
		return "未知"
	}
	for _, line := range strings.Split(string(output), "\n") {
		if strings.Contains(strings.ToLower(line), strings.ToLower(packageName)) {
			fields := strings.Fields(line)
			if len(fields) >= 2 {
				return fields[1]
			}
		}
	}
	return "未知"
}
