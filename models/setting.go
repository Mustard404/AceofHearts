package models

import (
	"AceofHearts/config"
	"AceofHearts/utils"
	"bufio"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// Setting 初始化配置：GitHub Token、代理、部署路径
func Setting() {
	configureProxy()
	loadGitHubToken()
	setDeployPath()
}

// 项目内置默认 GitHub OAuth App 凭证
const (
	defaultClientID      = "16489800309c385117b2"
	defaultClientSecrets = "1d14c4bfd8ae74e3a64c35f3ed8dfdf81312ae68"
)

// loadGitHubToken 优先使用内置 Token，失败或速率不足时交互输入
func loadGitHubToken() {
	utils.Title("配置 GitHub API")

	// 先尝试内置 Token
	token := base64.StdEncoding.EncodeToString([]byte(defaultClientID + ":" + defaultClientSecrets))
	remaining := checkRateLimit(token)

	if remaining > 100 {
		fmt.Printf("使用内置 Token，API 速率剩余: %s 次\n", config.Green(remaining))
		config.AccessToken = token
		return
	}

	if remaining >= 0 {
		fmt.Printf("内置 Token 速率剩余: %s 次，较低\n", config.Red(remaining))
	} else {
		fmt.Println(config.Yellow("内置 Token 验证失败"))
	}

	// 内置 Token 不可用，交互输入
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\n请提供自己的 GitHub OAuth App 凭证 (速率 5000次/小时)")
	fmt.Println("获取方式: GitHub → Settings → Developer settings → OAuth Apps")
	fmt.Print("Client ID (留空使用无认证模式): ")
	clientID, _ := reader.ReadString('\n')
	clientID = strings.TrimSpace(clientID)

	if clientID == "" {
		fmt.Println(config.Yellow("使用无认证模式 (60次/小时)"))
		return
	}

	fmt.Print("Client Secrets: ")
	clientSecrets, _ := reader.ReadString('\n')
	clientSecrets = strings.TrimSpace(clientSecrets)

	if clientSecrets == "" {
		fmt.Println(config.Yellow("使用无认证模式"))
		return
	}

	token = base64.StdEncoding.EncodeToString([]byte(clientID + ":" + clientSecrets))
	remaining = checkRateLimit(token)

	if remaining < 0 {
		fmt.Println(config.Red("Token 验证失败，使用无认证模式"))
		return
	}

	fmt.Printf("API 速率剩余: %s 次\n", config.Green(remaining))
	config.AccessToken = token
}

func checkRateLimit(token string) int {
	req, _ := http.NewRequest("GET", "https://api.github.com/rate_limit", nil)
	req.Header.Set("Authorization", "Basic "+token)

	resp, err := utils.GetHTTPClient().Do(req)
	if err != nil {
		fmt.Println(config.Red("无法连接 GitHub API"))
		return -1
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return -1
	}
	remaining, _ := strconv.Atoi(resp.Header.Get("X-RateLimit-Remaining"))
	return remaining
}

// configureProxy 交互式代理和镜像配置
func configureProxy() {
	utils.Title("代理配置")
	reader := bufio.NewReader(os.Stdin)

	// 代理
	fmt.Print("请输入代理地址 (例: http://127.0.0.1:7890，留空跳过): ")
	proxy, _ := reader.ReadString('\n')
	config.ProxyURL = strings.TrimSpace(proxy)

	if config.ProxyURL != "" {
		os.Setenv("http_proxy", config.ProxyURL)
		os.Setenv("https_proxy", config.ProxyURL)
		os.Setenv("HTTP_PROXY", config.ProxyURL)
		os.Setenv("HTTPS_PROXY", config.ProxyURL)
		os.Setenv("all_proxy", config.ProxyURL)
		fmt.Printf("代理已设置: %s\n", config.Green(config.ProxyURL))
		utils.Log("-\t代理: %s\n", config.ProxyURL)
	} else {
		fmt.Println("未使用代理")
	}

	// 镜像站
	fmt.Print("请输入 GitHub 镜像站 (例: https://ghfast.top/，留空跳过): ")
	mirror, _ := reader.ReadString('\n')
	config.GithubMirror = strings.TrimSpace(mirror)

	if config.GithubMirror != "" {
		fmt.Printf("镜像站已设置: %s\n", config.Green(config.GithubMirror))
		utils.Log("-\t镜像站: %s\n", config.GithubMirror)
	} else {
		fmt.Println("未使用镜像站")
	}
}

// setDeployPath 设置工具部署目录
func setDeployPath() {
	utils.Title("设置部署路径")
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("请输入工具部署目录 (默认 C:\\Tools): ")
	dir, _ := reader.ReadString('\n')
	dir = strings.TrimSpace(dir)
	if dir == "" {
		dir = `C:\Tools`
	}

	if _, err := os.Stat(dir); err == nil {
		fmt.Printf(config.Yellow("目录 %s 已存在\n"), dir)
	} else {
		os.MkdirAll(dir, 0755)
	}

	config.ToolsDir = dir
	config.XlsxPath = filepath.Join(config.ExecutablePath, "tools.xlsx")

	fmt.Printf("部署路径: %s\n", config.Magenta(dir))
	fmt.Printf("工具清单: %s\n", config.Magenta(config.XlsxPath))
	utils.Log("-\t部署路径: %s\n", dir)

	os.MkdirAll(filepath.Join(dir, "TEMP"), 0755)

	fmt.Print("按回车继续...")
	reader.ReadString('\n')
}
