package utils

import (
	"AceofHearts/config"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

// GitReleases 从 GitHub Releases 下载工具，返回 (文件路径, 版本, 描述, 错误)
func GitReleases(user, repo, key, directory string) (string, string, string, error) {
	api := fmt.Sprintf("https://api.github.com/repos/%s/%s", user, repo)
	data, err := fetchReleaseInfo(api)
	if err != nil {
		return "", "", "", fmt.Errorf("获取 Release 信息失败 (%s/%s): %w", user, repo, err)
	}

	downloadURL, err := resolveDownloadURL(user, repo, key, data)
	if err != nil {
		return "", "", "", err
	}
	if downloadURL == "" {
		return "", "", "", fmt.Errorf("未找到下载地址: %s/%s (key=%s)", user, repo, key)
	}

	version := "暂无版本号"
	if v, ok := data["tag_name"].(string); ok {
		version = v
	}
	description := "暂无描述"
	if d, ok := data["body"].(string); ok && d != "" {
		if idx := strings.Index(d, "\n"); idx > 0 {
			description = d[:idx]
		} else {
			description = d
		}
	}

	filename, err := MirrorDownload(downloadURL, directory)
	if err != nil {
		return "", "", "", fmt.Errorf("下载失败: %w", err)
	}
	return filename, version, description, nil
}

// fetchReleaseInfo 获取最新 Release 信息
func fetchReleaseInfo(api string) (map[string]interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", api+"/releases/latest", nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %w", err)
	}
	if config.AccessToken != "" {
		req.Header.Set("Authorization", "Basic "+config.AccessToken)
	}

	resp, err := GetHTTPClient().Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusForbidden {
		fmt.Println(config.Red("超出 GitHub API 速率限制，请更换 AccessToken！"))
		os.Exit(1)
	}

	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("解析 JSON 失败: %w", err)
	}
	return data, nil
}

// resolveDownloadURL 根据 key 匹配 Release 资产的下载地址
func resolveDownloadURL(user, repo, key string, data map[string]interface{}) (string, error) {
	switch key {
	case "main":
		return fmt.Sprintf("https://github.com/%s/%s/archive/refs/heads/main.zip", user, repo), nil
	case "master":
		return fmt.Sprintf("https://codeload.github.com/%s/%s/zip/refs/heads/master", user, repo), nil
	}

	assets, ok := data["assets"].([]interface{})
	if !ok {
		return "", fmt.Errorf("未找到 Release 资产: %s/%s", user, repo)
	}
	for _, asset := range assets {
		m, ok := asset.(map[string]interface{})
		if !ok {
			continue
		}
		url, ok := m["browser_download_url"].(string)
		if !ok {
			continue
		}
		if strings.Contains(url, key) {
			return url, nil
		}
	}

	config.Warn("未匹配到关键词 [%s]: %s/%s", key, user, repo)
	return "", nil
}
