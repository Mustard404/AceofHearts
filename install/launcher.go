package install

import (
	"AceofHearts/config"
	"AceofHearts/utils"
	"fmt"
	"os"
	"path/filepath"
)

func Launcher() {

	utils.Title("启动器部署")
	os.Mkdir(config.ToolDirectory["Dawn_Launcher"], 0755)
	filename, version, description := utils.GitReleases("AceofHearts404", "Dawn_Launcher", ".zip", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["Dawn_Launcher"], false)
	fmt.Println(config.Green("Dawn_Launcher 部署完成！"))
	utils.Log(fmt.Sprintf("-\tDawn_Launcher\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["Dawn_Launcher"]))

	fmt.Println("创建启动器桌面快捷方式")
	homeDirectory, _ := os.UserHomeDir()
	homeDirectory = filepath.Join(homeDirectory, "Desktop\\Dawn Launcher")
	launcherDirectory := fmt.Sprintf("%s\\Dawn Launcher.exe", config.ToolDirectory["Dawn_Launcher"])
	utils.Cmd("cmd", "/k", "mklink", homeDirectory, launcherDirectory)

}
