package utils

import (
	"AceofHearts/config"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func Cmd(command string, args ...string) bool {
	cmd := exec.Command(command, args...)

	cmdString := strings.Join(cmd.Args, " ")
	fmt.Println(config.Magenta("[AceofHearts]#:"), cmdString)

	_, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
		return false
	}

	return true
}
