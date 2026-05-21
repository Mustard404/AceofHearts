package install

import (
	"AceofHearts/config"
	"AceofHearts/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

// InstallFromXlsx 从 Excel 工具清单驱动安装，支持断点续装
func InstallFromXlsx(xlsxPath string) {
	utils.Title("工具安装")

	tools, err := utils.LoadTools(xlsxPath)
	if err != nil {
		fmt.Println(config.Red("加载工具清单失败:"), err)
		os.Exit(1)
	}

	pending, skipped := 0, 0
	for _, t := range tools {
		if t.Status == "已安装" {
			skipped++
		} else {
			pending++
		}
	}
	fmt.Printf("工具总数: %d | 待安装: %s | 已跳过: %s\n\n",
		len(tools), config.Magenta(pending), config.Green(skipped))

	installed, failed, idx := 0, 0, 0
	lastCategory := ""

	for _, tool := range tools {
		if tool.Status == "已安装" {
			continue
		}
		idx++

		if tool.Category != lastCategory {
			utils.Title2(tool.Category)
			lastCategory = tool.Category
		}

		toolPath := filepath.Join(config.ToolsDir, tool.Path)
		fmt.Printf("\n[%d/%d] %s (%s)\n", idx, pending, config.Cyan(tool.Name), tool.InstallMethod)

		if err := installTool(tool, toolPath); err != nil {
			fmt.Println(config.Red(fmt.Sprintf("  ✗ %s: %v", tool.Name, err)))
			utils.Log("-\t✗ %s 失败: %v\n", tool.Name, err)
			utils.UpdateToolStatus(xlsxPath, tool.ID, "失败")
			failed++
		} else {
			fmt.Println(config.Green(fmt.Sprintf("  ✓ %s 安装成功", tool.Name)))
			utils.Log("-\t✓ %s 安装成功 → %s\n", tool.Name, toolPath)
			utils.UpdateToolStatus(xlsxPath, tool.ID, "已安装")
			installed++
			postInstallHook(tool)
		}
	}

	fmt.Printf("\n安装完成! 成功: %s | 失败: %s | 跳过: %d\n",
		config.Green(installed), config.Red(failed), skipped)
	utils.Log("\n安装完成: 成功 %d, 失败 %d, 跳过 %d\n", installed, failed, skipped)
}

func installTool(tool utils.Tool, toolPath string) error {
	os.MkdirAll(toolPath, 0755)

	switch tool.InstallMethod {
	case "Chocolatey":
		return installChoco(tool, toolPath)
	case "GitHub":
		return installFromGitHub(tool, toolPath)
	case "Download":
		return installFromURL(tool, toolPath)
	default:
		return fmt.Errorf("未知安装方式: %s", tool.InstallMethod)
	}
}

// chocoInstallArgs 为 Chocolatey 包生成自定义安装路径参数
var chocoInstallArgs = map[string]func(dir string) string{
	"git":              func(dir string) string { return fmt.Sprintf(`--install-arguments="/VERYSILENT /SUPPRESSMSGBOXES /NORESTART /SP- /DIR=%s"`, dir) },
	"golang":           func(dir string) string { return fmt.Sprintf(`--install-arguments="INSTALLDIR=%s /qn /norestart"`, dir) },
	"jdk8":             func(dir string) string { return fmt.Sprintf(`-params "installdir=%s"`, strings.ReplaceAll(dir, `\`, `\\`)) },
	"laragon":          func(dir string) string { return fmt.Sprintf(`--install-arguments="/SILENT /VERYSILENT /DIR=%s"`, dir) },
	"nodejs-lts":       func(dir string) string { return fmt.Sprintf(`--install-arguments="/qn INSTALLDIR=%s"`, dir) },
	"python2":          func(dir string) string { return fmt.Sprintf(`--params "/InstallDir:%s"`, dir) },
	"python311":        func(dir string) string { return fmt.Sprintf(`--params "/InstallDir:%s"`, dir) },
	"ruby":             func(dir string) string { return fmt.Sprintf(`--install-arguments="/Dir=%s"`, dir) },
	"bandizip":         func(dir string) string { return fmt.Sprintf(`--install-arguments="/D:%s"`, dir) },
	"bitwarden":        func(dir string) string { return fmt.Sprintf(`--ia "/D=%s"`, dir) },
	"everything":       func(dir string) string { return fmt.Sprintf(`--installargs "/D=%s"`, dir) },
	"firefox":          func(dir string) string { return fmt.Sprintf(`--params "/InstallDir:%s"`, dir) },
	"sublimetext4":     func(dir string) string { return fmt.Sprintf(`--install-arguments="/VERYSILENT /SUPPRESSMSGBOXES /NORESTART /SP- /DIR=%s"`, dir) },
	"typora":           func(dir string) string { return fmt.Sprintf(`--install-arguments="/DIR=%s"`, dir) },
	"mobaxterm":        func(dir string) string { return fmt.Sprintf(`--installargs "INSTALLDIR=%s"`, dir) },
	"vscode":           func(dir string) string { return fmt.Sprintf(`--install-arguments="/DIR=%s"`, dir) },
	"winscp":           func(dir string) string { return fmt.Sprintf(`--install-arguments="/DIR=%s"`, dir) },
	"proxifier":        func(dir string) string { return fmt.Sprintf(`--install-arguments="/DIR=%s"`, dir) },
	"wireshark":        func(dir string) string { return fmt.Sprintf(`--install-arguments="/D=%s"`, dir) },
}

func installChoco(tool utils.Tool, toolPath string) error {
	args := tool.ChocoArgs
	// 如果 xlsx 中没有自定义参数，使用内置映射
	if args == "" {
		if fn, ok := chocoInstallArgs[tool.ChocoPkg]; ok {
			args = fn(toolPath)
		}
	}
	return utils.ChocoInstall(tool.Name, tool.ChocoPkg, toolPath, args)
}

func installFromGitHub(tool utils.Tool, toolPath string) error {
	tempDir := filepath.Join(config.ToolsDir, "TEMP")
	os.MkdirAll(tempDir, 0755)

	switch tool.InstallType {
	case "GitClone":
		repoURL := fmt.Sprintf("https://github.com/%s/%s.git", tool.GitHubUser, tool.GitHubRepo)
		return utils.GitClone(repoURL, toolPath)

	case "Releases":
		filename, version, _, err := utils.GitReleases(tool.GitHubUser, tool.GitHubRepo, tool.ReleaseKey, tempDir)
		if err != nil {
			return err
		}
		if isArchive(filename) {
			utils.Unzip(filename, toolPath, false)
		} else {
			os.MkdirAll(toolPath, 0755)
			if err := os.Rename(filename, filepath.Join(toolPath, filepath.Base(filename))); err != nil {
				return fmt.Errorf("移动文件失败: %w", err)
			}
		}
		fmt.Printf("  版本: %s\n", version)
		return nil

	case "Packages":
		_, version, _, err := utils.GitReleases(tool.GitHubUser, tool.GitHubRepo, tool.ReleaseKey, toolPath)
		if err != nil {
			return err
		}
		fmt.Printf("  版本: %s\n", version)
		return nil

	default:
		return fmt.Errorf("未知安装类型: %s", tool.InstallType)
	}
}

func installFromURL(tool utils.Tool, toolPath string) error {
	os.MkdirAll(toolPath, 0755)

	downloadURL := tool.URL

	// Navicat 特殊处理：通过 API 获取真实下载链接
	if strings.Contains(tool.URL, "navicat.com") {
		realURL, err := getNavicatDownloadURL()
		if err != nil {
			return fmt.Errorf("获取 Navicat 下载链接失败: %v", err)
		}
		downloadURL = realURL
	}

	filename, err := utils.DownloadWithRetry(downloadURL, toolPath)
	if err != nil {
		return fmt.Errorf("下载失败: %w", err)
	}
	if isArchive(filename) {
		utils.Unzip(filename, toolPath, false)
	}
	return nil
}

// getNavicatDownloadURL 通过 Navicat 官网 API 获取真实下载链接
func getNavicatDownloadURL() (string, error) {
	apiURL := "https://www.navicat.com.cn/includes/Navicat/direct_download.php"
	data := url.Values{
		"product":    {"navicat_premium_cs_x64.exe"},
		"location":   {"1"},
		"support":    {""},
		"linux_dist": {""},
	}

	client := utils.GetHTTPClient()
	resp, err := client.PostForm(apiURL, data)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var result struct {
		DownloadLink string `json:"download_link"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}
	return "http://" + result.DownloadLink, nil
}

func isArchive(name string) bool {
	lower := strings.ToLower(name)
	exts := []string{".zip", ".7z", ".tar.gz", ".tar", ".rar", ".tgz"}
	for _, ext := range exts {
		if strings.HasSuffix(lower, ext) {
			return true
		}
	}
	return false
}

// addToPATH 将路径添加到当前进程的 PATH 环境变量
func addToPATH(dirs ...string) {
	path := os.Getenv("PATH")
	for _, dir := range dirs {
		if _, err := os.Stat(dir); err == nil {
			path = dir + string(os.PathListSeparator) + path
		}
	}
	os.Setenv("Path", path)
}

// postInstallHook 工具安装完成后的钩子
func postInstallHook(tool utils.Tool) {
	switch tool.Name {
	case "Git":
		addToPATH(
			`C:\Program Files\Git\cmd`,
			`C:\Program Files\Git\bin`,
			`C:\Program Files (x86)\Git\cmd`,
		)
		if config.ProxyURL != "" {
			fmt.Println("  设置 Git 代理...")
			utils.Cmd("git", "config", "--global", "http.proxy", config.ProxyURL)
			utils.Cmd("git", "config", "--global", "https.proxy", config.ProxyURL)
		}
	case "Bandizip":
		addToPATH(filepath.Join(config.ToolsDir, tool.Path))
	case "Nmap":
		addToPATH(`C:\Program Files (x86)\Nmap`)
	}
}
