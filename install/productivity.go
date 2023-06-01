package install

import (
	"AceofHearts/config"
	"AceofHearts/utils"
	"fmt"
	"golang.org/x/sys/windows/registry"
	"os"
	"runtime"
)

func Productivity() {

	utils.Title("生产力工具部署")
	os.Mkdir(config.ToolDirectory["Productivity"], 0755)

	bandizipCommand := []string{fmt.Sprintf("--install-arguments=/D:\"%s\"", config.ToolDirectory["Bandizip"])}
	utils.ChocoInstall("Bandizip", "bandizip", config.ToolDirectory["Bandizip"], bandizipCommand)

	PATH := os.Getenv("PATH")
	updatedPath := config.ToolDirectory["Bandizip"] + string(os.PathListSeparator) + PATH
	os.Setenv("Path", updatedPath)

	bitwardenCommand := []string{"--ia", fmt.Sprintf("/D=\"%s\"", config.ToolDirectory["Bitwarden"])}
	utils.ChocoInstall("Bitwarden", "bitwarden", config.ToolDirectory["Bitwarden"], bitwardenCommand)

	// 无法通过命令安装指定目录
	//config.ToolDirectory["chrome"] = r"C:\Program Files\Google\\"
	var chromeCommand []string
	utils.ChocoInstall("Chrome", "googlechrome", config.ToolDirectory["Chrome"], chromeCommand)

	extensions := map[string]string{
		"Bitwarden":                "nngceckbapebfimnlniiiahkandclblb",
		"FOFA Pro View":            "dobbfkjhgbkmmcooahlnllfopfmhcoln",
		"Hack-Tools":               "cmbndhnoonmghfofefkcccljbkdpamhi",
		"HackBar":                  "ginpbkfigcoaokgflihfhhmglmbchinc",
		"IP, DNS & Security Tools": "phjkepckmcnjohilmbjlcoblenhgpjmo",
		"Proxy SwitchyOmega":       "padekgcemlokbadohgkifijomclgjgif",
		"Wappalyzer":               "gppongmhjkpfnbhagpmjfkannfbllamg",
	}

	for name, id := range extensions {
		chromeExtensions(name, id)
	}

	everythingCommand := []string{"--installargs", fmt.Sprintf("/D=%s", config.ToolDirectory["Everything"])}
	utils.ChocoInstall("Everything", "everything", config.ToolDirectory["Everything"], everythingCommand)

	firefoxCommand := []string{"--params", fmt.Sprintf("/InstallDir:%s", config.ToolDirectory["Firefox"])}
	utils.ChocoInstall("Firefox", "firefox", config.ToolDirectory["Firefox"], firefoxCommand)

	sublimetextCommand := []string{fmt.Sprintf("--install-arguments='/VERYSILENT /SUPPRESSMSGBOXES /NORESTART /SP- /DIR=%s'", config.ToolDirectory["Sublime_Text"])}
	utils.ChocoInstall("SublimeText", "sublimetext4", config.ToolDirectory["Sublime_Text"], sublimetextCommand)

	typoraCommand := []string{fmt.Sprintf("--install-arguments='/DIR=%s'", config.ToolDirectory["Typora"])}
	utils.ChocoInstall("Typora", "typora", config.ToolDirectory["Typora"], typoraCommand)
}

func chromeExtensions(name string, id string) {

	var regKeyPath string
	updateURL := "https://clients2.google.com/service/update2/crx"

	if runtime.GOARCH == "amd64" {
		regKeyPath = "Software\\Wow6432Node\\Google\\Chrome\\Extensions"
	} else if runtime.GOARCH == "386" {
		regKeyPath = "Software\\Google\\Chrome\\Extensions"
	} else {
		fmt.Println("非Windows系统")
		return
	}

	regKeyPath = regKeyPath + "\\" + id
	extKey, _, err := registry.CreateKey(registry.LOCAL_MACHINE, regKeyPath, registry.CREATE_SUB_KEY|registry.SET_VALUE)
	if err != nil {
		fmt.Println("无法创建或打开子键:", err)
		return
	}

	defer extKey.Close()

	err = extKey.SetStringValue("update_url", updateURL)
	if err != nil {
		fmt.Println("无法写入注册表值:", err)
		return
	}

	fmt.Println("扩展程序", name, "已添加到Chrome中！")
}
