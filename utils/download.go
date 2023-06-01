package utils

import (
	"fmt"
	"github.com/cheggaaa/pb/v3"
	"io"
	"net/http"
	"os"
	"strings"
)

func Download(link string, directory string) string {
	err := os.MkdirAll(directory, 0755)
	if err != nil {
		fmt.Errorf("无法创建目录：%s", err)
	}

	resp, err := http.Get(link)
	if err != nil {
		fmt.Fprintf(os.Stderr, "下载失败：%v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	filename := getFilenameFromURL(link, resp.Header.Get("Content-Disposition"))
	if filename == "" {
		filename = getFilenameFromURL(link, "")
	}
	fmt.Printf("Downloading %s...\n", filename)
	filename = fmt.Sprintf("%s\\%s", directory, filename)

	file, err := os.Create(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "创建文件失败：%v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	contentLength := resp.ContentLength

	// 使用进度条
	bar := pb.Full.Start64(contentLength)
	bar.Set(pb.Bytes, true)
	bar.SetTemplateString(`{{string . "prefix"}}{{counters . }} {{bar . }} {{percent . }} {{string . "suffix"}}`)
	defer bar.Finish()

	reader := bar.NewProxyReader(resp.Body)

	_, err = io.Copy(file, reader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "下载失败：%v\n", err)
		os.Exit(1)
	}

	return filename
}

func getFilenameFromURL(link string, contentDisposition string) string {
	// 从 Content-Disposition 中提取文件名
	if contentDisposition != "" {
		startIndex := strings.Index(contentDisposition, "filename=")
		if startIndex != -1 {
			startIndex += len("filename=")
			endIndex := strings.Index(contentDisposition[startIndex:], ";")
			if endIndex == -1 {
				endIndex = len(contentDisposition)
			} else {
				endIndex += startIndex
			}
			filename := contentDisposition[startIndex:endIndex]
			filename = strings.Trim(filename, "\"'")
			if filename != "" {
				return filename
			}
		}
	}

	// 从 URL 中提取文件名
	tokens := strings.Split(link, "/")
	return tokens[len(tokens)-1]
}
