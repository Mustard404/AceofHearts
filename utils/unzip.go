package utils

import (
	"AceofHearts/config"
	"archive/zip"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Unzip 解压文件到指定目录
// move=true 时先解压到父目录再重命名（用于源码包解压后多一层目录的情况）
func Unzip(zipFilePath, outputDir string, move bool) {
	if move {
		tempDir := filepath.Dir(outputDir)
		unzipTo(zipFilePath, tempDir)

		os.RemoveAll(outputDir)
		baseName := strings.TrimSuffix(filepath.Base(zipFilePath), filepath.Ext(zipFilePath))
		src := filepath.Join(tempDir, baseName)
		if err := os.Rename(src, outputDir); err != nil {
			fmt.Println(config.Red("重命名文件夹失败:"), err)
		}
	} else {
		unzipTo(zipFilePath, outputDir)
	}
}

// unzipTo 优先使用 Bandizip，不可用时回退到 Go 原生解压
func unzipTo(zipFilePath, outputDir string) {
	if _, err := exec.LookPath("Bandizip.exe"); err == nil {
		Cmd("Bandizip.exe", "bx", "-y", fmt.Sprintf("-o:%s", outputDir), zipFilePath)
		return
	}

	// Go 原生 zip 解压兜底（仅支持 .zip 格式）
	if strings.HasSuffix(strings.ToLower(zipFilePath), ".zip") {
		if err := unzipNative(zipFilePath, outputDir); err != nil {
			fmt.Println(config.Red("原生解压失败:"), err)
		}
		return
	}

	fmt.Println(config.Red("未找到 Bandizip 且文件非 .zip 格式，无法解压: " + zipFilePath))
}

// unzipNative 使用 Go 标准库解压 zip 文件
func unzipNative(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		fpath := filepath.Join(dest, f.Name)

		// 防止 zip slip 攻击
		if !strings.HasPrefix(filepath.Clean(fpath), filepath.Clean(dest)+string(os.PathSeparator)) {
			return fmt.Errorf("非法文件路径: %s", f.Name)
		}

		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, 0755)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(fpath), 0755); err != nil {
			return err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}

		rc, err := f.Open()
		if err != nil {
			outFile.Close()
			return err
		}

		_, err = io.Copy(outFile, rc)
		rc.Close()
		outFile.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
