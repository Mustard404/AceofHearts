package install

import (
	"AceofHearts/config"
	"AceofHearts/utils"
	"fmt"
	"os"
	"path/filepath"
)

// Launcher 部署 GeekDesk 启动器并写入工具数据
func Launcher() {
	utils.Title("启动器部署 (GeekDesk)")

	geekdeskDir := filepath.Join(config.ToolsDir, "GeekDesk")
	tempDir := filepath.Join(config.ToolsDir, "TEMP")
	os.MkdirAll(geekdeskDir, 0755)

	filename, version, _, err := utils.GitReleases("BookerLiu", "GeekDesk", ".zip", tempDir)
	if err != nil {
		fmt.Println(config.Red("下载 GeekDesk 失败:"), err)
		return
	}
	utils.Unzip(filename, geekdeskDir, false)
	fmt.Println(config.Green("GeekDesk 部署完成!"))
	utils.Log("-\tGeekDesk v%s → %s\n", version, geekdeskDir)

	// 自动写入已安装工具到 GeekDesk 数据文件
	if err := utils.WriteGeekDeskData(geekdeskDir, config.XlsxPath, config.ToolsDir); err != nil {
		fmt.Println(config.Red("写入 GeekDesk 数据失败:"), err)
	} else {
		fmt.Println(config.Green("已自动写入工具列表到 GeekDesk!"))
	}

	// 桌面快捷方式
	home, _ := os.UserHomeDir()
	shortcut := filepath.Join(home, "Desktop", "GeekDesk")
	target := filepath.Join(geekdeskDir, "GeekDesk.exe")
	utils.Cmd("cmd", "/c", "mklink", shortcut, target)
}
