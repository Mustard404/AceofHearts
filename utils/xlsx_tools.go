package utils

import (
	"fmt"
	"strconv"

	"github.com/xuri/excelize/v2"
)

// Tool 工具定义结构体
type Tool struct {
	ID            int
	Name          string
	Path          string
	InstallMethod string // Chocolatey / GitHub / Download
	Category      string
	URL           string
	InstallType   string // Releases / Packages / GitClone
	ChocoPkg      string
	ChocoArgs     string
	GitHubUser    string
	GitHubRepo    string
	ReleaseKey    string
	Status        string // 未安装 / 已安装 / 失败
}

const sheetName = "工具清单"

// LoadTools 从 xlsx 文件读取工具列表
func LoadTools(xlsxPath string) ([]Tool, error) {
	f, err := excelize.OpenFile(xlsxPath)
	if err != nil {
		return nil, fmt.Errorf("无法打开工具清单: %v", err)
	}
	defer f.Close()

	rows, err := f.GetRows(sheetName)
	if err != nil {
		return nil, fmt.Errorf("无法读取工作表: %v", err)
	}

	var tools []Tool
	for i, row := range rows {
		if i == 0 || len(row) < 13 {
			continue // 跳过表头或不完整行
		}
		id, err := strconv.Atoi(row[0])
		if err != nil {
			fmt.Printf("第 %d 行 ID 格式错误，已跳过: %s\n", i+1, row[0])
			continue
		}
		tools = append(tools, Tool{
			ID:            id,
			Name:          row[1],
			Path:          row[2],
			InstallMethod: row[3],
			Category:      row[4],
			URL:           row[5],
			InstallType:   row[6],
			ChocoPkg:      row[7],
			ChocoArgs:     row[8],
			GitHubUser:    row[9],
			GitHubRepo:    row[10],
			ReleaseKey:    row[11],
			Status:        row[12],
		})
	}
	return tools, nil
}

// UpdateToolStatus 更新 xlsx 中指定工具的状态
func UpdateToolStatus(xlsxPath string, toolID int, status string) error {
	f, err := excelize.OpenFile(xlsxPath)
	if err != nil {
		return err
	}
	defer f.Close()

	rows, err := f.GetRows(sheetName)
	if err != nil {
		return err
	}

	// 状态在第13列 (M列)
	for i, row := range rows {
		if i == 0 || len(row) < 1 {
			continue
		}
		id, err := strconv.Atoi(row[0])
		if err != nil {
			continue
		}
		if id == toolID {
			cell := fmt.Sprintf("M%d", i+1)
			f.SetCellValue(sheetName, cell, status)
			break
		}
	}
	return f.Save()
}
