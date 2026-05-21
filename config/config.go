package config

import (
	"fmt"

	"github.com/fatih/color"
)

var (
	// Version 程序版本号
	Version = "2.0"

	// AccessToken GitHub API 认证 Token (Base64编码)
	AccessToken string

	// ExecutablePath 程序所在目录
	ExecutablePath string

	// ToolsDir 工具部署根目录
	ToolsDir string

	// XlsxPath 工具清单 Excel 路径
	XlsxPath string

	// ProxyURL 代理地址
	ProxyURL string

	// GithubMirror GitHub 镜像站地址
	GithubMirror string
)

// 颜色输出函数
var (
	Cyan    = color.New(color.FgCyan).SprintFunc()
	Green   = color.New(color.FgGreen).SprintFunc()
	Magenta = color.New(color.FgMagenta).SprintFunc()
	Red     = color.New(color.FgRed).SprintFunc()
	Yellow  = color.New(color.FgYellow).SprintFunc()
)

// Warn 输出黄色警告信息
func Warn(format string, a ...interface{}) {
	fmt.Println(Yellow("[警告]"), fmt.Sprintf(format, a...))
}
