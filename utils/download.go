package utils

import (
	"AceofHearts/config"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/cheggaaa/pb/v3"
)

const maxRetries = 3

// DownloadWithRetry 下载文件，最多重试 3 次
func DownloadWithRetry(link string, directory string) (string, error) {
	if err := os.MkdirAll(directory, 0755); err != nil {
		return "", fmt.Errorf("无法创建目录: %v", err)
	}

	var lastErr error
	for attempt := 1; attempt <= maxRetries; attempt++ {
		filename, err := downloadOnce(link, directory)
		if err == nil {
			return filename, nil
		}
		lastErr = err
		if attempt < maxRetries {
			wait := time.Duration(attempt*2) * time.Second
			config.Warn("下载失败 (第%d次)，%v后重试: %v", attempt, wait, err)
			time.Sleep(wait)
		}
	}
	return "", fmt.Errorf("重试 %d 次后仍失败: %w", maxRetries, lastErr)
}

func downloadOnce(link string, directory string) (string, error) {
	client := GetHTTPClient()
	resp, err := client.Get(link)
	if err != nil {
		return "", fmt.Errorf("请求失败: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("HTTP %d: %s", resp.StatusCode, link)
	}

	filename := parseFilename(link, resp.Header.Get("Content-Disposition"))
	filePath := filepath.Join(directory, filename)
	fmt.Printf("Downloading %s ...\n", filename)

	f, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("创建文件失败: %v", err)
	}
	defer f.Close()

	bar := pb.Full.Start64(resp.ContentLength)
	bar.Set(pb.Bytes, true)
	bar.SetTemplateString(`{{string . "prefix"}}{{counters . }} {{bar . }} {{percent . }} {{string . "suffix"}}`)
	defer bar.Finish()

	if _, err := io.Copy(f, bar.NewProxyReader(resp.Body)); err != nil {
		return "", fmt.Errorf("写入失败: %v", err)
	}
	return filePath, nil
}

// parseFilename 从 URL 或 Content-Disposition 中提取文件名
func parseFilename(link string, disposition string) string {
	if disposition != "" {
		if idx := strings.Index(disposition, "filename="); idx != -1 {
			name := disposition[idx+len("filename="):]
			if end := strings.Index(name, ";"); end != -1 {
				name = name[:end]
			}
			name = strings.Trim(name, "\"'")
			if name != "" {
				return name
			}
		}
	}
	parts := strings.Split(link, "/")
	name := parts[len(parts)-1]
	if idx := strings.Index(name, "?"); idx != -1 {
		name = name[:idx]
	}
	return name
}
