package config

import "github.com/fatih/color"

var AccessToken string
var ExecutablePath string
var ToolDirectory map[string]string

func init() {
	ToolDirectory = make(map[string]string)
}

var Cyan = color.New(color.FgCyan).SprintFunc()
var Green = color.New(color.FgGreen).SprintFunc()
var Magenta = color.New(color.FgMagenta).SprintFunc()
var Red = color.New(color.FgRed).SprintFunc()
