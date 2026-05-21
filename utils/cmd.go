package utils

import (
	"AceofHearts/config"
	"fmt"
	"os/exec"
	"strings"
)

// Cmd 执行系统命令，返回是否成功
func Cmd(command string, args ...string) bool {
	cmd := exec.Command(command, args...)
	fmt.Println(config.Magenta("[AceofHearts]#:"), strings.Join(cmd.Args, " "))

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("%s\n%s\n", config.Red(err), string(output))
		return false
	}
	return true
}

// CmdOutput 执行命令并返回输出
func CmdOutput(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	output, err := cmd.CombinedOutput()
	return string(output), err
}
