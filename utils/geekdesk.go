package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// WriteGeekDeskData 生成 C# 脚本并编译执行，将已安装工具写入 GeekDesk 数据文件
func WriteGeekDeskData(geekdeskDir, xlsxPath, toolsRoot string) error {
	tools, err := LoadTools(xlsxPath)
	if err != nil {
		return err
	}

	// 按分类分组
	groups := make(map[string][]Tool)
	var categoryOrder []string
	for _, t := range tools {
		if t.Status == "已安装" {
			if _, exists := groups[t.Category]; !exists {
				categoryOrder = append(categoryOrder, t.Category)
			}
			groups[t.Category] = append(groups[t.Category], t)
		}
	}

	if len(groups) == 0 {
		return fmt.Errorf("没有已安装的工具")
	}

	// 生成 C# 代码
	csCode := generateCSharpCode(groups, categoryOrder, toolsRoot, geekdeskDir)

	// 写入 C# 文件
	csFile := filepath.Join(geekdeskDir, "WriteData.cs")
	if err := os.WriteFile(csFile, []byte(csCode), 0644); err != nil {
		return fmt.Errorf("写入 C# 文件失败: %v", err)
	}

	// 编译
	exeFile := filepath.Join(geekdeskDir, "WriteData.exe")
	cscPath := findCSC()
	if cscPath == "" {
		return fmt.Errorf("未找到 csc.exe，请确认已安装 .NET Framework")
	}
	geekdeskExe := filepath.Join(geekdeskDir, "GeekDesk.exe")

	fmt.Println("  编译数据写入工具...")
	if !Cmd(cscPath, "/nologo", "/out:"+exeFile,
		"/reference:"+geekdeskExe,
		"/reference:"+filepath.Join(geekdeskDir, "lib", "Newtonsoft.Json.dll"),
		csFile) {
		return fmt.Errorf("编译 C# 失败")
	}

	// 执行
	fmt.Println("  写入工具数据...")
	if !Cmd(exeFile) {
		return fmt.Errorf("写入 GeekDesk 数据失败")
	}

	// 清理临时文件
	os.Remove(csFile)
	os.Remove(exeFile)

	fmt.Printf("  已写入 %d 个分类、%d 个工具到 GeekDesk\n", len(groups), countInstalled(tools))
	return nil
}

func countInstalled(tools []Tool) int {
	n := 0
	for _, t := range tools {
		if t.Status == "已安装" {
			n++
		}
	}
	return n
}

func generateCSharpCode(groups map[string][]Tool, categoryOrder []string, toolsRoot, geekdeskDir string) string {
	var sb strings.Builder

	sb.WriteString(`using System;
using System.Collections.ObjectModel;
using System.IO;
using System.Runtime.Serialization.Formatters.Binary;
using GeekDesk.ViewModel;

class WriteData
{
    static void Main()
    {
        var appData = new AppData();
        appData.menuList = new ObservableCollection<MenuInfo>();
`)

	for _, category := range categoryOrder {
		tools := groups[category]
		sb.WriteString(fmt.Sprintf(`
        // %s
        {
            var menu = new MenuInfo();
            menu.MenuName = "%s";
            menu.MenuId = Guid.NewGuid().ToString();
            menu.IconList = new ObservableCollection<IconInfo>();
`, category, category))

		for _, t := range tools {
			toolPath := filepath.Join(toolsRoot, t.Path)
			escapedPath := strings.ReplaceAll(toolPath, `\`, `\\`)
			escapedName := strings.ReplaceAll(t.Name, `"`, `\"`)

			sb.WriteString(fmt.Sprintf(`
            {
                var icon = new IconInfo();
                icon.Name = "%s";
                icon.Path = @"%s";
                menu.IconList.Add(icon);
            }
`, escapedName, escapedPath))
		}

		sb.WriteString(`            appData.menuList.Add(menu);
        }
`)
	}

	dataFile := filepath.Join(geekdeskDir, "Data")
	escapedDataFile := strings.ReplaceAll(dataFile, `\`, `\\`)

	sb.WriteString(fmt.Sprintf(`
        // 序列化写入
        var formatter = new BinaryFormatter();
        using (var fs = new FileStream("%s", FileMode.Create))
        {
            formatter.Serialize(fs, appData);
        }
        Console.WriteLine("GeekDesk Data 写入成功!");
    }
}
`, escapedDataFile))

	return sb.String()
}

// findCSC 自动探测可用的 csc.exe 路径，优先 Framework64
func findCSC() string {
	bases := []string{
		`C:\Windows\Microsoft.NET\Framework64`,
		`C:\Windows\Microsoft.NET\Framework`,
	}
	for _, base := range bases {
		entries, err := os.ReadDir(base)
		if err != nil {
			continue
		}
		// 倒序遍历，优先使用高版本
		for i := len(entries) - 1; i >= 0; i-- {
			csc := filepath.Join(base, entries[i].Name(), "csc.exe")
			if _, err := os.Stat(csc); err == nil {
				return csc
			}
		}
	}
	return ""
}
