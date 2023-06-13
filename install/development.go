package install

import (
	"AceofHearts/config"
	"AceofHearts/utils"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
	"os"
)

func Development() {

	utils.Title("开发环境部署")
	os.Mkdir(config.ToolDirectory["Development"], 0755)

	mobaxtermCommand := []string{"--installargs", fmt.Sprintf("INSTALLDIR=%s", config.ToolDirectory["Mobaxterm"])}
	utils.ChocoInstall("Mobaxterm", "mobaxterm", config.ToolDirectory["Mobaxterm"], mobaxtermCommand)

	fileName := utils.Download(getNavicatPremium(), config.ToolDirectory["TEMP"])
	result := utils.Cmd(fileName, "/SILENT", "/VERYSILENT", fmt.Sprintf("/DIR=%s", config.ToolDirectory["Navicat_Premium"]))
	if result {
		fmt.Println(config.Green("Navicat Premium部署成功! \n"))
		utils.Log(fmt.Sprintf("-\tNavicat Premium部署成功! \t\n\t-\t版本: 16\t\n\t-\t路径: %s \n", config.ToolDirectory["Navicat_Premium"]))
	} else {
		fmt.Println(config.Red("Navicat Premium部署失败! \n"))
		utils.Log("-\tNavicat Premium部署失败! \n")
	}

	vscodeCommand := []string{fmt.Sprintf("--install-arguments='/DIR=\"%s\"'", config.ToolDirectory["Vscode"])}
	utils.ChocoInstall("Vscode", "vscode", config.ToolDirectory["Vscode"], vscodeCommand)

	winscpCommand := []string{fmt.Sprintf("--install-arguments='/DIR=\"%s\"'", config.ToolDirectory["Winscp"])}
	utils.ChocoInstall("Winscp", "winscp", config.ToolDirectory["Winscp"], winscpCommand)

}

type Response struct {
	DownloadLink string `json:"download_link"`
	BannerURL    string `json:"banner_url"`
}

func getNavicatPremium() string {

	url := "https://www.navicat.com.cn/includes/Navicat/direct_download.php"
	payload := map[string]string{
		"product":    "navicat_premium_cs_x64.exe",
		"location":   "1",
		"support":    "",
		"linux_dist": "",
	}

	client := resty.New()
	resp, err := client.R().
		SetFormData(payload).
		Post(url)

	if err != nil {
		log.Fatal("请求失败:", err)
	}

	if resp.StatusCode() == 200 {
		var response Response
		if err := json.Unmarshal(resp.Body(), &response); err != nil {
			log.Fatal("解析 Navicat Premium 下载地址失败! :", err)
		}
		return "http://" + response.DownloadLink
	} else {
		fmt.Println("请求Navicat Premium 下载地址失败! ，状态码:", resp.StatusCode())
	}
	return ""
}
