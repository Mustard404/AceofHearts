package utils

import (
	"AceofHearts/config"
	"fmt"
	"net/http"
	"strings"
	"time"
)

var mirrorAvailable bool

// CheckMirror 预检镜像站是否可用，在 Setting 后调用一次
func CheckMirror() {
	if config.GithubMirror == "" {
		return
	}

	testURL := strings.TrimRight(config.GithubMirror, "/")
	// 镜像预检不走代理（镜像本身就是替代直连的加速方案）
	client := &http.Client{
		Timeout:   10 * time.Second,
		Transport: &http.Transport{},
	}
	resp, err := client.Head(testURL)
	if err != nil || resp.StatusCode >= 500 {
		config.Warn("镜像站 %s 不可用，将直连 GitHub", config.GithubMirror)
		mirrorAvailable = false
		return
	}

	mirrorAvailable = true
	fmt.Printf("镜像站可用: %s\n", config.Green(config.GithubMirror))
}

// MirrorURL 将 GitHub URL 替换为镜像站地址
func MirrorURL(originalURL string) string {
	if config.GithubMirror == "" || !mirrorAvailable {
		return originalURL
	}
	mirror := strings.TrimRight(config.GithubMirror, "/") + "/"

	if strings.Contains(originalURL, "github.com") ||
		strings.Contains(originalURL, "codeload.github.com") {
		return mirror + originalURL
	}
	return originalURL
}

// MirrorDownload 先尝试镜像站下载，失败后回退到原始地址
func MirrorDownload(originalURL string, directory string) (string, error) {
	mirrorURL := MirrorURL(originalURL)

	if mirrorURL != originalURL {
		filename, err := DownloadWithRetry(mirrorURL, directory)
		if err == nil {
			return filename, nil
		}
		config.Warn("镜像站下载失败，回退到直连: %s", originalURL)
		mirrorAvailable = false
	}
	return DownloadWithRetry(originalURL, directory)
}
