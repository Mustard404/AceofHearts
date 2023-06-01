package utils

import (
	"AceofHearts/config"
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"log"
	"net/http"
	"os"
	"strings"
)

func GitReleases(user string, repo string, key string, directory string) (string, string, string) {

	api := fmt.Sprintf("https://api.github.com/repos/%s/%s", user, repo)

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/releases/latest", api), nil)
	if err != nil {
		log.Fatal("创建请求失败:", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", config.AccessToken))
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("请求失败:", err)
	}

	if resp.StatusCode == http.StatusForbidden {
		log.Println(config.Red("超出Github Api速率限制, 请更换accessToken！"))
		resp.Body.Close()
		os.Exit(0)
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Fatal("解析 JSON 失败:", err)
	}

	downloadUrl := ""
	if key == "main" {
		downloadUrl = fmt.Sprintf("https://github.com/%s/%s/archive/refs/heads/main.zip", user, repo)
	} else if key == "master" {
		downloadUrl = fmt.Sprintf("https://codeload.github.com/%s/%s/zip/refs/heads/master", user, repo)
	} else {
		for _, asset := range data["assets"].([]interface{}) {
			assetMap := asset.(map[string]interface{})
			browserDownloadURL := assetMap["browser_download_url"].(string)
			if strings.Contains(browserDownloadURL, key) {
				downloadUrl = browserDownloadURL
			}
		}
	}

	var Version string

	if data["tag_name"] != nil {
		Version = data["tag_name"].(string)
	} else {
		Version = "暂无版本号！"
	}
	Description := GitDescription(api, config.AccessToken)
	Filename := Download(downloadUrl, directory)
	return Filename, Version, Description

}

func GitDescription(Api string, accessToken string) string {
	red := color.New(color.FgRed).SprintFunc()

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/releases/latest", Api), nil)
	if err != nil {
		log.Fatal("创建请求失败:", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", accessToken))
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("请求失败:", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusForbidden {
		log.Println(red("超出Github Api速率限制, 请更换accessToken！"))
		resp.Body.Close()
		os.Exit(0)
	}

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Fatal("解析 JSON 失败:", err)
	}

	description, ok := data["description"].(string)
	if !ok {
		description = "暂无描述"
	}

	return description
}
