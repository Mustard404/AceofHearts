package utils

import (
	"AceofHearts/config"
	"fmt"
	"os/exec"
	"strings"
)

// GitClone 克隆仓库到指定目录，支持镜像加速
func GitClone(repoURL string, directory string) error {
	cloneURL := MirrorURL(repoURL)

	fmt.Println(config.Magenta("[AceofHearts]#:"), "git clone --depth 1", cloneURL, directory)
	cmd := exec.Command("git", "clone", "--depth", "1", cloneURL, directory)
	output, err := cmd.CombinedOutput()

	if err != nil && cloneURL != repoURL {
		// 镜像失败，回退到原始地址
		config.Warn("镜像站克隆失败，尝试直连...")
		cmd = exec.Command("git", "clone", "--depth", "1", repoURL, directory)
		output, err = cmd.CombinedOutput()
	}

	if err != nil {
		return fmt.Errorf("git clone 失败: %v\n%s", err, strings.TrimSpace(string(output)))
	}
	fmt.Println(config.Green("Git clone 完成!"))
	return nil
}
