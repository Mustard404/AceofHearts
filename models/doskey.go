package models

import (
	"AceofHearts/config"
	"AceofHearts/utils"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Doskey() {
	utils.Title("设置Doskey (Windows Alias)")
	doskey := `
@echo off
doskey ls = dir /b $*
doskey rm = del $*
doskey -rf = /s $*
doskey mk = md $*
doskey clear = cls $*
doskey ll = dir $*
doskey cat = type $*
doskey which = where $*
rem Bypass_AV
doskey mateuszex = $ToolDirectory\Bypass_AV\MateuszEx\MazteuszEx.exe $*
doskey shellcode = $ToolDirectory\Bypass_AV\Shellcodeloader\shellcodeLoader.exe $*
doskey shellcodeloader = $ToolDirectory\Bypass_AV\Shellcodeloader\shellcodeLoader.exe $*
rem C2_Framework
doskey cobaltstrike = cd  $ToolDirectory\C2_Framework\Cobalt_Strike\ ^& start_win.bat $*
doskey sliver = $ToolDirectory\C2_Framework\Sliver\sliver-client_windows.exe $*
rem Dir_Scanner
doskey dirmap = py -3 $ToolDirectory\Dir_Scanner\Dirmap\dirmap.py $*
doskey dirsearch = py -3 $ToolDirectory\Dir_Scanner\Dirsearch\dirsearch.py $*
doskey ffuf = $ToolDirectory\Dir_Scanner\Ffuf\ffuf.exe $*
doskey jsfinder = py -3 $ToolDirectory\Dir_Scanner\JSFinder\JSFinder.py $*
doskey webpatchbrute = $ToolDirectory\Dir_Scanner\WebPathBrute\7kbscan-WebPathBrute.exe $*
rem EXP
doskey commix = py -3 $ToolDirectory\EXP\Commix\commix.py $*
doskey dalfox = $ToolDirectory\EXP\Dalfox\dalfox.exe $*
doskey commix = py -3 $ToolDirectory\EXP\Fuxploider\fuxploider.py $*
doskey jndiexploit = java -jar $ToolDirectory\EXP\JNDIExploit\JNDIExploit-1.4-SNAPSHOT.jar $*
doskey lfisuite = py -3 $ToolDirectory\EXP\LFISuite\lfisuite.py $*
doskey sqlmap = py -3 $ToolDirectory\EXP\Sqlmap\sqlmap.py $*
doskey xxeinjector = ruby $ToolDirectory\EXP\XXEinjector\XXEinjector.rb $*
doskey ysomap = java -jar $ToolDirectory\EXP\Ysomap\ysomap.jar $*
rem Info_Gathering
doskey amass = $ToolDirectory\Info_Gathering\Amass\amass.exe $*
doskey ct = py -3 $ToolDirectory\Info_Gathering\Ct_Exposer\ct-exposer.py $*
doskey dnsx = $ToolDirectory\Info_Gathering\Dnsx\dnsx.exe $*
doskey fofaviewer = java -jar $ToolDirectory\Info_Gathering\Fofa_Viewer\fofaviewer.jar $*
doskey layer = $ToolDirectory\Info_Gathering\Layer\Layer.exe $*
doskey oneforall = py -3 $ToolDirectory\Info_Gathering\OneForAll\oneforall.py $*
rem Maintain_Access
doskey girsh = py -3 $ToolDirectory\Maintain_Access\Girsh\Girsh.exe $*
doskey govenom = $ToolDirectory\Maintain_Access\Govenom\govenom.exe $*
doskey platypus = $ToolDirectory\Maintain_Access\Platypus\Platypus_windows_amd64.exe $*
doskey reversessh = $ToolDirectory\Maintain_Access\Reverse_ssh\upx_reverse-sshx64.exe $*
rem Pass_Attack
doskey hackbrowserdata = $ToolDirectory\Pass_Attack\HackBrowserData\hack-browser-data-windows-64bit.exe $*
doskey hashcat = $ToolDirectory\Pass_Attack\Hashcat\hashcat.exe $*
doskey psudohash = py -3 $ToolDirectory\Pass_Attack\Psudohash\psudohash.py
doskey hydra = $ToolDirectory\Pass_Attack\Thc_Hydra\hydra.exe $*
rem Penetration_Testing_Framework
doskey msfconsole = cd  $ToolDirectory\Penetration_Testing_Framework\Metasploit\bin ^& msfconsole.bat $*
rem Port_Scanner
doskey masscan = $ToolDirectory\Port_Scanner\Masscan\masscan.exe $*
doskey nmap = $ToolDirectory\Port_Scanner\Nmap\nmap.exe $*
doskey txportmap = $ToolDirectory\Port_Scanner\TXPortMap\TxPortMap_windows_x64.exe $*
rem Vuln_Scanner
doskey afrog = $ToolDirectory\Vuln_Scanner\Afrog\afrog.exe $*
doskey nuclei = $ToolDirectory\Vuln_Scanner\Nuclei\nuclei.exe $*
doskey vscan = $ToolDirectory\Vuln_Scanner\Vscan\vscan.exe $*
doskey xray = $ToolDirectory\Vuln_Scanner\Xray\xray_windows_amd64.exe $*
	`
	doskey = strings.Replace(doskey, "$ToolDirectory", config.ToolDirectory["ROOT"], -1)
	os.Mkdir(config.ToolDirectory["Doskey"], 0755)
	doskeyFile := filepath.Join(config.ToolDirectory["Doskey"], "doskey.bat")
	file, err := os.OpenFile(doskeyFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("无法创建批处理文件:", err)
		os.Exit(1)
	}
	defer file.Close()

	_, err = file.WriteString(doskey)
	if err != nil {
		fmt.Println("无法写入批处理文件:", err)
		os.Exit(1)
	}

	fmt.Println("批处理文件已保存:", config.ToolDirectory["Doskey"])

	result := utils.Cmd("reg", "add", "HKCU\\Software\\Microsoft\\Command Processor", "/v", "AutoRun", "/t", "REG_SZ", "/d", doskeyFile, "/f")
	if result {
		fmt.Println(config.Green("Doskey添加到启动！\n"))
		utils.Log(fmt.Sprintf("-\tDoskey添加到启动！\n\t-\t安装路径: %s\n", config.ToolDirectory["Doskey"]))
	} else {
		fmt.Println(config.Red("Doskey添加到启动失败！\n"))
		utils.Log("-\tDoskey添加到启动失败！\n")
	}

	clinkCommand := []string{""}
	utils.ChocoInstall("Clink", "clink-maintained", "", clinkCommand)
}
