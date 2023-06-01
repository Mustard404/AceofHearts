package install

import (
	"AceofHearts/config"
	"AceofHearts/utils"
	"fmt"
	"os"
	"strings"
)

func Env() {

	utils.Title("编程环境部署")
	os.Mkdir(config.ToolDirectory["Env"], 0755)

	gitCommand := []string{fmt.Sprintf("--install-arguments=\"/VERYSILENT /SUPPRESSMSGBOXES /NORESTART /SP- /DIR=%s\"", config.ToolDirectory["Git"])}
	utils.ChocoInstall("Git", "git", config.ToolDirectory["Git"], gitCommand)

	goCommand := []string{fmt.Sprintf("--install-arguments=\"INSTALLDIR=%s /qn /norestart\"", config.ToolDirectory["Go"])}
	utils.ChocoInstall("Go", "golang", config.ToolDirectory["Go"], goCommand)

	config.ToolDirectory["Java8"] = strings.Replace(config.ToolDirectory["Java8"], "\\", "\\\\", -1)
	java8Command := []string{"-params", fmt.Sprintf("installdir=%s", config.ToolDirectory["Java8"])}
	utils.ChocoInstall("Java8", "jdk8", config.ToolDirectory["Java8"], java8Command)

	laragonCommand := []string{fmt.Sprintf("--install-arguments=\"/SILENT /VERYSILENT /DIR=%s\"", config.ToolDirectory["Laragon"])}
	utils.ChocoInstall("Laragon", "laragon", config.ToolDirectory["Laragon"], laragonCommand)

	nodeCommand := []string{fmt.Sprintf("--install-arguments='/qn INSTALLDIR=%s'", config.ToolDirectory["Node"])}
	utils.ChocoInstall("Node", "nodejs-lts", config.ToolDirectory["Node"], nodeCommand)

	python2Command := []string{fmt.Sprintf("--params='/InstallDir:%s'", config.ToolDirectory["Python2"])}
	utils.ChocoInstall("Python2", "python2", config.ToolDirectory["Python2"], python2Command)

	python3Command := []string{fmt.Sprintf("--params='/InstallDir:%s'", config.ToolDirectory["Python3"])}
	utils.ChocoInstall("Python3", "python311", config.ToolDirectory["Python3"], python3Command)

	rubyCommand := []string{fmt.Sprintf("--install-arguments=\"/Dir=%s\"", config.ToolDirectory["Ruby"])}
	utils.ChocoInstall("Ruby", "ruby", config.ToolDirectory["Ruby"], rubyCommand)
}
