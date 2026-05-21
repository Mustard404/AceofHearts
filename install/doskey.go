package install

import (
	"AceofHearts/config"
	"AceofHearts/utils"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Doskey 设置 Windows 命令别名
func Doskey() {
	utils.Title("设置 Doskey (Windows Alias)")

	doskeyContent := `@echo off
doskey ls = dir /b $*
doskey rm = del $*
doskey mk = md $*
doskey clear = cls $*
doskey ll = dir $*
doskey cat = type $*
doskey which = where $*
rem ── 渗透工具快捷命令 ──
doskey sqlmap = py -3 $ROOT\EXP\Sqlmap\sqlmap.py $*
doskey nmap = $ROOT\Port_Scanner\Nmap\nmap.exe $*
doskey masscan = $ROOT\Port_Scanner\Masscan\masscan.exe $*
doskey fscan = $ROOT\Port_Scanner\Fscan\fscan.exe $*
doskey nuclei = $ROOT\Vuln_Scanner\Nuclei\nuclei.exe $*
doskey xray = $ROOT\Vuln_Scanner\Xray\xray_windows_amd64.exe $*
doskey afrog = $ROOT\Vuln_Scanner\Afrog\afrog.exe $*
doskey ffuf = $ROOT\Dir_Scanner\Ffuf\ffuf.exe $*
doskey httpx = $ROOT\Info_Gathering\httpx\httpx.exe $*
doskey subfinder = $ROOT\Info_Gathering\Subfinder\subfinder.exe $*
doskey dnsx = $ROOT\Info_Gathering\Dnsx\dnsx.exe $*
doskey katana = $ROOT\Info_Gathering\katana\katana.exe $*
doskey dalfox = $ROOT\EXP\Dalfox\dalfox.exe $*
doskey hydra = $ROOT\Pass_Attack\Thc_Hydra\hydra.exe $*
doskey hashcat = $ROOT\Pass_Attack\Hashcat\hashcat.exe $*
doskey frpc = $ROOT\Port_Forwarding\Frp\frpc.exe $*
doskey frps = $ROOT\Port_Forwarding\Frp\frps.exe $*
doskey dirsearch = py -3 $ROOT\Dir_Scanner\Dirsearch\dirsearch.py $*
doskey gobuster = $ROOT\Port_Scanner\Gobuster\gobuster.exe $*
`
	doskeyContent = strings.ReplaceAll(doskeyContent, "$ROOT", config.ToolsDir)

	doskeyDir := filepath.Join(config.ToolsDir, "Doskey")
	os.MkdirAll(doskeyDir, 0755)
	doskeyFile := filepath.Join(doskeyDir, "doskey.bat")

	if err := os.WriteFile(doskeyFile, []byte(doskeyContent), 0644); err != nil {
		fmt.Println(config.Red("创建 doskey.bat 失败:"), err)
		return
	}
	fmt.Println("doskey.bat 已保存:", config.Magenta(doskeyFile))

	// 注册到启动项
	if utils.Cmd("reg", "add", `HKCU\Software\Microsoft\Command Processor`,
		"/v", "AutoRun", "/t", "REG_SZ", "/d", doskeyFile, "/f") {
		fmt.Println(config.Green("Doskey 添加到启动!"))
		utils.Log("-\tDoskey → %s\n", doskeyDir)
	}

	// Clink 增强
	utils.ChocoInstall("Clink", "clink-maintained", "", "")
}
