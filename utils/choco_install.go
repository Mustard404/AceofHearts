package utils

import (
	"AceofHearts/config"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func ChocoInstall(name string, packageName string, directory string, command []string, args ...string) {
	cmdArgs := []string{"install", packageName, "-y"}
	cmdArgs = append(cmdArgs, command...)
	cmdArgs = append(cmdArgs, args...)
	result := exec.Command("choco", cmdArgs...)

	cmdString := strings.Join(result.Args, " ")
	fmt.Println(config.Magenta("[AceofHearts]#:"), cmdString)

	out, err := result.Output()
	if err != nil {
		message := fmt.Sprintf("%s 部署失败！", name)
		fmt.Println(config.Red(message))
		fmt.Println(string(out))
		log.Fatal(err)
	}
	message := fmt.Sprintf("%s 部署成功！", name)
	fmt.Println(config.Green(message))
	Version := checkVersion(packageName)
	Log("-\t%s部署成功！\t\n\t-\t版本:%s\t\n\t-\t路径: %s\n", name, Version, directory)
}

func checkVersion(packageName string) string {

	cmd := exec.Command("choco", "list")
	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.Contains(strings.ToLower(line), strings.ToLower(packageName)) {
			fields := strings.Fields(line)
			if len(fields) >= 2 {
				return fields[1]
			}
		}
	}

	return fmt.Sprintf("无法找到包名为 %s 的版本号", packageName)
}
