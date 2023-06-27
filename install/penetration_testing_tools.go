package install

import (
	"AceofHearts/config"
	"AceofHearts/utils"
	"fmt"
	"os"
	"path/filepath"
)

func PenetrationTestingTools() {

	utils.Title("渗透测试工具部署")
	utils.Title2("渗透测试工具-免杀")
	os.Mkdir(config.ToolDirectory["Bypass_AV"], 0755)
	bypassAV()
	utils.Title2("渗透测试工具-C2")
	os.Mkdir(config.ToolDirectory["C2_Framework"], 0755)
	c2Framework()
	utils.Title2("渗透测试工具-目录扫描")
	os.Mkdir(config.ToolDirectory["Dir_Scanner"], 0755)
	dirScanner()
	utils.Title2("渗透测试工具-漏洞利用")
	os.Mkdir(config.ToolDirectory["EXP"], 0755)
	EXP()
	utils.Title2("渗透测试工具-指纹识别")
	os.Mkdir(config.ToolDirectory["Fingerprinting"], 0755)
	fingerprinting()
	utils.Title2("渗透测试工具-信息收集")
	os.Mkdir(config.ToolDirectory["Info_Gathering"], 0755)
	infoGathering()
	utils.Title2("渗透测试工具-权限维持")
	os.Mkdir(config.ToolDirectory["Maintain_Access"], 0755)
	maintainAccess()
	utils.Title2("渗透测试工具-密码攻击")
	os.Mkdir(config.ToolDirectory["Pass_Attack"], 0755)
	passAttack()
	utils.Title2("渗透测试工具-渗透测试框架")
	os.Mkdir(config.ToolDirectory["Penetration_Testing_Framework"], 0755)
	penetrationTestingFramework()
	utils.Title2("渗透测试工具-端口转发")
	os.Mkdir(config.ToolDirectory["Port_Forwarding"], 0755)
	portForwarding()
	utils.Title2("渗透测试工具-端口扫描")
	os.Mkdir(config.ToolDirectory["Port_Scanner"], 0755)
	portScanner()
	utils.Title2("渗透测试工具-提权辅助")
	os.Mkdir(config.ToolDirectory["Privilege_Escalation"], 0755)
	privilegeEscalation()
	utils.Title2("渗透测试工具-代理")
	os.Mkdir(config.ToolDirectory["Proxy"], 0755)
	proxy()
	utils.Title2("渗透测试工具-漏洞扫描")
	os.Mkdir(config.ToolDirectory["Vuln_Scanner"], 0755)
	vulnScanner()
	utils.Title2("渗透测试工具-漏洞检索")
	os.Mkdir(config.ToolDirectory["Vuln_Search"], 0755)
	vulnSearch()
	utils.Title2("渗透测试工具-Webshell")
	os.Mkdir(config.ToolDirectory["Webshell"], 0755)
	webshell()
}

func bypassAV() {

	filename, version, description := utils.GitReleases("1y0n", "AV_Evasion_Tool", ".zip", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["AV_Evasion_Tool"], false)
	fmt.Println(config.Green("掩日部署完成！"))
	utils.Log(fmt.Sprintf("-\t掩日\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["AV_Evasion_Tool"]))

	filename, version, description = utils.GitReleases("TideSec", "BypassAntiVirus", "master", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["BypassAntiVirus"], true)
	fmt.Println(config.Green("BypassAntiVirus 部署完成！"))
	utils.Log(fmt.Sprintf("-\tBypassAntiVirus\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["BypassAntiVirus"]))

	filename, version, description = utils.GitReleases("AceofHearts404", "GBByPass", ".jar", config.ToolDirectory["GBByPass"])
	fmt.Println(config.Green("GBByPass 部署完成！"))
	utils.Log(fmt.Sprintf("-\tGBByPass\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["GBByPass"]))

	filename, version, description = utils.GitReleases("sairson", "MateuszEx", ".exe", config.ToolDirectory["MateuszEx"])
	fmt.Println(config.Green("MateuszEx 部署完成！"))
	utils.Log(fmt.Sprintf("-\tMateuszEx\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["MateuszEx"]))

	filename, version, description = utils.GitReleases("knownsec", "shellcodeloader", ".7z", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["Shellcodeloader"], false)
	fmt.Println(config.Green("Shellcodeloader部署完成！"))
	utils.Log(fmt.Sprintf("-\tShellcodeloader\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["Shellcodeloader"]))

}

func c2Framework() {

	filename, version, description := utils.GitReleases("AceofHearts404", "cs_open_jar", "jar.zip", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["Cobalt_Strike"], true)
	fmt.Println(config.Green("Cobalt_Strike 部署完成！"))
	utils.Log(fmt.Sprintf("-\tCobalt Strike\n\t-\tVersion: 4.5\n\t-\t简介: SafetyTeam 魔改版，请关注该公众号！\n\t-\t安装路径: %s\n", config.ToolDirectory["Cobalt_Strike"]))

	filename, version, description = utils.GitReleases("d3ckx1", "OLa", ".7z", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["OLa"], false)
	fmt.Println(config.Green("OLa 部署完成！"))
	utils.Log(fmt.Sprintf("-\t\tOLa(请手动加载)\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["Cobalt_Strike"]))

	filename, version, description = utils.GitReleases("BishopFox", "sliver", "client_windows.exe", config.ToolDirectory["Sliver"])
	filename, version, description = utils.GitReleases("BishopFox", "sliver", "server_windows.exe", config.ToolDirectory["Sliver"])
	fmt.Println(config.Green("Sliver 部署完成！"))
	utils.Log(fmt.Sprintf("-\tSliver\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["Sliver"]))

}

func dirScanner() {

	filename, version, description := utils.GitReleases("H4ckForJob", "dirmap", "master", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["Dirmap"], true)
	fmt.Println(config.Green("Dirmap 部署完成！"))
	utils.Log(fmt.Sprintf("-\tDirmap\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["Dirmap"]))

	filename, version, description = utils.GitReleases("maurosoria", "dirsearch", "master", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["Dirsearch"], true)
	fmt.Println(config.Green("Dirsearch 部署完成！"))
	utils.Log(fmt.Sprintf("-\tDirsearch\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["Dirsearch"]))

	filename, version, description = utils.GitReleases("ffuf", "ffuf", "windows_amd64.zip", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["Ffuf"], false)
	fmt.Println(config.Green("Ffuf 部署完成！"))
	utils.Log(fmt.Sprintf("-\tFfuf\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["Ffuf"]))

	filename, version, description = utils.GitReleases("Threezh1", "JSFinder", "master", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["JSFinder"], true)
	fmt.Println(config.Green("JSFinder 部署完成！"))
	utils.Log(fmt.Sprintf("-\tJSFinder\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["JSFinder"]))

	filename, version, description = utils.GitReleases("7kbstorm", "7kbscan-WebPathBrute", ".zip", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["WebPathBrute"], false)
	fmt.Println(config.Green("WebPathBrute 部署完成！"))
	utils.Log(fmt.Sprintf("-\tWebPathBrute\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["WebPathBrute"]))

}

func EXP() {

	filename, version, description := utils.GitReleases("kleiton0x00", "Advanced-SQL-Injection-Cheatsheet", "main", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["Advanced_SQL_Injection_Cheatsheet"], true)
	fmt.Println(config.Green("Advanced-SQL-Injection-Cheatsheet 部署完成！"))
	utils.Log(fmt.Sprintf("-\tAdvanced-SQL-Injection-Cheatsheet\n\t-\tVersion: 无\n\t-\t简介: 此存储库包含所有类型的 SQL 注入的高级方法。\n\t-\t安装路径: %s\n", config.ToolDirectory["Advanced_SQL_Injection_Cheatsheet"]))

	filename, version, description = utils.GitReleases("teamssix", "cf", "windows_amd64.zip", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["CF"], false)
	fmt.Println(config.Green("CF 部署完成！"))
	utils.Log(fmt.Sprintf("-\tCF\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["CF"]))

	filename, version, description = utils.GitReleases("commixproject", "commix", "master", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["Commix"], true)
	fmt.Println(config.Green("Commix 部署完成！"))
	utils.Log(fmt.Sprintf("-\tCommix\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["Commix"]))

	filename, version, description = utils.GitReleases("hahwul", "dalfox", "windows_amd64.tar.gz", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["Dalfox"], false)
	fmt.Println(config.Green("Dalfox 部署完成！"))
	utils.Log(fmt.Sprintf("-\tDalfox\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["Dalfox"]))

	filename, version, description = utils.GitReleases("almandin", "fuxploider", "master", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["Fuxploider"], true)
	fmt.Println(config.Green("Fuxploider 部署完成！"))
	utils.Log(fmt.Sprintf("-\tFuxploider\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["Fuxploider"]))

	filename, version, description = utils.GitReleases("AceofHearts404", "JNDIExploit", ".zip", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["JNDIExploit"], true)
	fmt.Println(config.Green("JNDIExploit 部署完成！"))
	utils.Log(fmt.Sprintf("-\tJNDIExploit\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["JNDIExploit"]))

	filename, version, description = utils.GitReleases("D35m0nd142", "LFISuite", "master", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["LFISuite"], true)
	fmt.Println(config.Green("LFISuite 部署完成！"))
	utils.Log(fmt.Sprintf("-\tLFISuite\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["LFISuite"]))

	filename, version, description = utils.GitReleases("feihong-cs", "ShiroExploit-Deprecated", ".7z", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["ShiroExploit"], true)
	fmt.Println(config.Green("ShiroExploit 部署完成！"))
	utils.Log(fmt.Sprintf("-\tShiroExploit\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["ShiroExploit"]))

	filename, version, description = utils.GitReleases("SummerSec", "ShiroAttack2", ".zip", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["ShiroAttack2"], false)
	fmt.Println(config.Green("ShiroAttack2 部署完成！"))
	utils.Log(fmt.Sprintf("-\tShiroAttack2\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["ShiroAttack2"]))

	filename, version, description = utils.GitReleases("sqlmapproject", "sqlmap", "master", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["Sqlmap"], true)
	fmt.Println(config.Green("Sqlmap 部署完成！"))
	utils.Log(fmt.Sprintf("-\tSqlmap\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["Sqlmap"]))

	filename, version, description = utils.GitReleases("shack2", "SuperSQLInjectionV1", ".zip", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["SuperSQLInjection"], false)
	fmt.Println(config.Green("SuperSQLInjection 部署完成！"))
	utils.Log(fmt.Sprintf("-\tSuperSQLInjection\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["SuperSQLInjection"]))

	filename, version, description = utils.GitReleases("s0md3v", "XSStrike", "master", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["XSStrike"], true)
	fmt.Println(config.Green("XSStrike 部署完成！"))
	utils.Log(fmt.Sprintf("-\tXSStrike\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["XSStrike"]))

	filename, version, description = utils.GitReleases("evilcos", "xssor2", "master", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["Xssor2"], true)
	fmt.Println(config.Green("Xssor2 部署完成！"))
	utils.Log(fmt.Sprintf("-\tXssor2\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["Xssor2"]))

	filename, version, description = utils.GitReleases("enjoiz", "XXEinjector", "master", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["XXEinjector"], true)
	fmt.Println(config.Green("XXEinjector 部署完成！"))
	utils.Log(fmt.Sprintf("-\tXXEinjector\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["XXEinjector"]))

	filename, version, description = utils.GitReleases("KimJun1010", "WeblogicTool", ".jar", config.ToolDirectory["WeblogicTool"])
	fmt.Println(config.Green("WeblogicTool 部署完成！"))
	utils.Log(fmt.Sprintf("-\tWeblogicTool\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["WeblogicTool"]))

	filename, version, description = utils.GitReleases("wh1t3p1g", "ysomap", ".jar", config.ToolDirectory["Ysomap"])
	fmt.Println(config.Green("Ysomap 部署完成！"))
	utils.Log(fmt.Sprintf("-\tYsomap\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["Ysomap"]))
}

func fingerprinting() {

	filename, version, description := utils.GitReleases("Tuhinshubhra", "CMSeeK", "master", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["CMSeeK"], true)
	fmt.Println(config.Green("CMSeeK 部署完成！"))
	utils.Log(fmt.Sprintf("-\tCMSeeK\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["CMSeeK"]))

	filename, version, description = utils.GitReleases("EdgeSecurityTeam", "EHole", "windows", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["EHole"], false)
	fmt.Println(config.Green("EHole 部署完成！"))
	utils.Log(fmt.Sprintf("-\tCMSeeK\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["EHole"]))

	filename, version, description = utils.GitReleases("0x727", "ObserverWard", "windows-msvc.zip", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["ObserverWard"], false)
	fmt.Println(config.Green("ObserverWard 部署完成！"))
	utils.Log(fmt.Sprintf("-\tObserverWard\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["ObserverWard"]))

	filename, version, description = utils.GitReleases("TideSec", "TideFinger", "master", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["TideFinger"], false)
	fmt.Println(config.Green("TideFinger 部署完成！"))
	utils.Log(fmt.Sprintf("-\tTideFinger\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["TideFinger"]))

}

func infoGathering() {

	filename, version, description := utils.GitReleases("owasp-amass", "amass", "amass_Windows_amd64.zip", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["Amass"], true)
	fmt.Println(config.Green("Amass 部署完成！"))
	utils.Log(fmt.Sprintf("-\tAmass\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["Amass"]))

	filename, version, description = utils.GitReleases("chris408", "ct-exposer", "master", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["Ct_Exposer"], true)
	fmt.Println(config.Green("Ct_Exposer 部署完成！"))
	utils.Log(fmt.Sprintf("-\tCt_Exposer\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["Ct_Exposer"]))

	filename, version, description = utils.GitReleases("projectdiscovery", "dnsx", "windows_amd64.zip", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["Dnsx"], false)
	fmt.Println(config.Green("Dnsx 部署完成！"))
	utils.Log(fmt.Sprintf("-\tDnsx\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["Dnsx"]))

	filename, version, description = utils.GitReleases("lijiejie", "ds_store_exp", "master", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["Ds_store_exp"], true)
	fmt.Println(config.Green("Ds_store_exp 部署完成！"))
	utils.Log(fmt.Sprintf("-\tDs_store_exp\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["Ds_store_exp"]))

	filename, version, description = utils.GitReleases("wgpsec", "fofa_viewer", "JDK8.zip", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["Fofa_Viewer"], false)
	fmt.Println(config.Green("Fofa_Viewer 部署完成！"))
	utils.Log(fmt.Sprintf("-\tFofa_Viewer\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["Fofa_Viewer"]))

	filename, version, description = utils.GitReleases("hisxo", "gitGraber", "master", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["Gitgraber"], true)
	fmt.Println(config.Green("Gitgraber 部署完成！"))
	utils.Log(fmt.Sprintf("-\tGitgraber\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["Gitgraber"]))

	filename, version, description = utils.GitReleases("lijiejie", "GitHack", "master", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["Githack"], true)
	fmt.Println(config.Green("Githack 部署完成！"))
	utils.Log(fmt.Sprintf("-\tGithack\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["Githack"]))

	filename, version, description = utils.GitReleases("UnkL4b", "GitMiner", "master", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["Gitminer"], true)
	fmt.Println(config.Green("Gitminer 部署完成！"))
	utils.Log(fmt.Sprintf("-\tGitminer\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["Gitminer"]))

	// 未发布正式版本，暂不调取
	filename = utils.Download("https://github.com/michenriksen/gitrob/releases/download/v2.0.0-beta/gitrob_windows_amd64_2.0.0-beta.zip", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["Gitrob"], false)
	fmt.Println(config.Green("Gitrob 部署完成！"))
	utils.Log(fmt.Sprintf("-\tGitminer\n\t-\tVersion: v2.0.0-beta\n\t-\t简介: Gitrob 是一种工具，可帮助查找推送到 Github 上公共存储库的潜在敏感文件。\n\t-\t安装路径: %s\n", config.ToolDirectory["Gitminer"]))

	filename, version, description = utils.GitReleases("OJ", "gobuster", "Windows_x86_64.zip", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["Gobuster"], false)
	fmt.Println(config.Green("Gobuster 部署完成！"))
	utils.Log(fmt.Sprintf("-\tGobuster\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["Gobuster"]))

	filename, version, description = utils.GitReleases("boy-hack", "ksubdomain", "windows.tar", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["Ksubdomain"], false)
	fmt.Println(config.Green("Ksubdomain 部署完成！"))
	utils.Log(fmt.Sprintf("-\tKsubdomain\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["Ksubdomain"]))

	filename, version, description = utils.GitReleases("euphrat1ca", "LayerDomainFinder", ".zip", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["Layer"], false)
	fmt.Println(config.Green("Layer 部署完成！"))
	utils.Log(fmt.Sprintf("-\tLayer\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["Layer"]))

	filename, version, description = utils.GitReleases("shmilylty", "OneForAll", "master", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["Oneforall"], true)
	fmt.Println(config.Green("Oneforall 部署完成！"))
	utils.Log(fmt.Sprintf("-\tOneforall\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["Oneforall"]))

	filename, version, description = utils.GitReleases("smicallef", "spiderfoot", "master", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["Spiderfoot"], true)
	fmt.Println(config.Green("Spiderfoot 部署完成！"))
	utils.Log(fmt.Sprintf("-\tSpiderfoot\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["Spiderfoot"]))

	filename, version, description = utils.GitReleases("projectdiscovery", "subfinder", "windows_amd64.zip", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["Subfinder"], false)
	fmt.Println(config.Green("Subfinder 部署完成！"))
	utils.Log(fmt.Sprintf("-\tSubfinder\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["Subfinder"]))

	filename, version, description = utils.GitReleases("admintony", "svnExploit", "master", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["Svnexploit"], true)
	fmt.Println(config.Green("Svnexploit 部署完成！"))
	utils.Log(fmt.Sprintf("-\tSvnexploit\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["Svnexploit"]))

	filename, version, description = utils.GitReleases("laramies", "theharvester", "master", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["Theharvester"], true)
	fmt.Println(config.Green("Theharvester 部署完成！"))
	utils.Log(fmt.Sprintf("-\tTheharvester\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["Theharvester"]))
}

func maintainAccess() {

	filename, version, description := utils.GitReleases("nodauf", "Girsh", "windows_amd64.zip", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["Girsh"], false)
	fmt.Println(config.Green("Girsh 部署完成！"))
	utils.Log(fmt.Sprintf("-\tGirsh\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["Girsh"]))

	filename, version, description = utils.GitReleases("erikgeiser", "govenom", "windows_x64.zip", config.ToolDirectory["Govenom"])
	utils.Unzip(filename, config.ToolDirectory["Govenom"], false)
	fmt.Println(config.Green("Govenom 部署完成！"))
	utils.Log(fmt.Sprintf("-\tGovenom\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["Govenom"]))

	filename, version, description = utils.GitReleases("WangYihang", "Platypus", "windows_amd64.exe", config.ToolDirectory["Platypus"])
	fmt.Println(config.Green("Platypus 部署完成！"))
	utils.Log(fmt.Sprintf("-\tPlatypus\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["Platypus"]))

	filename, version, description = utils.GitReleases("Fahrj", "reverse-ssh", "x64.exe", config.ToolDirectory["Reverse_ssh"])
	fmt.Println(config.Green("Reverse_ssh 部署完成！"))
	utils.Log(fmt.Sprintf("-\tReverse_ssh\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["Reverse_ssh"]))

}

func passAttack() {

	filename, version, description := utils.GitReleases("moonD4rk", "HackBrowserData", "windows-64bit.zip", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["HackBrowserData"], false)
	fmt.Println(config.Green("HackBrowserData 部署完成！"))
	utils.Log(fmt.Sprintf("-\tHackBrowserData\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["HackBrowserData"]))

	filename, version, description = utils.GitReleases("hashcat", "hashcat", ".7z", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["Hashcat"], true)
	fmt.Println(config.Green("Hashcat 部署完成！"))
	utils.Log(fmt.Sprintf("-\tHashcat\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["Hashcat"]))

	filename, version, description = utils.GitReleases("maaaaz", "thc-hydra-windows", ".zip", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["Thc_Hydra"], false)
	fmt.Println(config.Green("Thc_Hydra 部署完成！"))
	utils.Log(fmt.Sprintf("-\tThc_Hydra\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["Thc_Hydra"]))

	filename = utils.Download("https://www.openwall.com/john/k/john-1.9.0-jumbo-1-win64.7z", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["John"], true)
	fmt.Println(config.Green("John 部署完成！"))
	utils.Log(fmt.Sprintf("-\tJohn\n\t-\tVersion: v1.9.0 \n\t-\t简介: John the Ripper 是一种开源密码安全审核和密码恢复工具，可用于许多操作系统\n\t-\t安装路径: %s\n", config.ToolDirectory["John"]))

	filename = utils.Download("https://openwall.info/wiki/_media/john/johnny/johnny_2.2_win.zip", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["Johnny"], false)
	fmt.Println(config.Green("Johnny 部署完成！"))
	utils.Log(fmt.Sprintf("-\tJohnny\n\t-\tVersion: v2.2\n\t-\t简介: Johnny 是流行的密码破解者 John the Ripper 的 跨平台开源GUI前端。\n\t-\t安装路径: %s\n", config.ToolDirectory["Johnny"]))

	filename, version, description = utils.GitReleases("lanjelot", "patator", "master", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["Patator"], true)
	fmt.Println(config.Green("Patator 部署完成！"))
	utils.Log(fmt.Sprintf("-\tPatator\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["Patator"]))

	filename, version, description = utils.GitReleases("berzerk0", "Probable-Wordlists", "master", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["Probable_Wordlists"], true)
	fmt.Println(config.Green("Probable_Wordlists 部署完成！"))
	utils.Log(fmt.Sprintf("-\tProbable_Wordlists\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["Probable_Wordlists"]))

	filename, version, description = utils.GitReleases("t3l3machus", "psudohash", "main", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["Psudohash"], true)
	fmt.Println(config.Green("Psudohash 部署完成！"))
	utils.Log(fmt.Sprintf("-\tPsudohash\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["Psudohash"]))

}

func penetrationTestingFramework() {

	filename := utils.Download("https://windows.metasploit.com/metasploitframework-latest.msi", config.ToolDirectory["TEMP"])
	result := utils.Cmd("msiexec", "/i", filename, "/qn", fmt.Sprintf("INSTALLLOCATION=%s", config.ToolDirectory["Penetration_Testing_Framework"]))
	oldFilename := filepath.Join(config.ToolDirectory["Penetration_Testing_Framework"], "\\metasploit-framework")
	os.Rename(oldFilename, config.ToolDirectory["Metasploit"])
	if result {
		fmt.Println(config.Green("部署Metasploit成功！\n"))
		utils.Log(fmt.Sprintf("-\t部署Metasploit成功！\n\t-\t安装路径: %s\n", config.ToolDirectory["Metasploit"]))
	} else {
		fmt.Println(config.Red("部署Metasploit失败！\n"))
		utils.Log("-\t部署Metasploit失败！\n")
	}

	filename, version, description := utils.GitReleases("knownsec", "pocsuite3", "master", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["Pocsuite3"], true)
	fmt.Println(config.Green("Pocsuite3部署完成！"))
	utils.Log(fmt.Sprintf("-\tPocsuite3\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["Pocsuite3"]))
}

func portForwarding() {

	filename, version, description := utils.GitReleases("fatedier", "frp", "windows_amd64.zip", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["Frp"], true)
	fmt.Println(config.Green("Frp 部署完成！"))
	utils.Log(fmt.Sprintf("-\tFrp\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["Frp"]))

	filename, version, description = utils.GitReleases("EddieIvan01", "iox", "Windows_x86_64.tar.gz", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["Iox"], false)
	fmt.Println(config.Green("Iox 部署完成！"))
	utils.Log(fmt.Sprintf("-\tIox\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["Iox"]))

	filename, version, description = utils.GitReleases("L-codes", "Neo-reGeorg", "master", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["Neo_reGeorg"], true)
	fmt.Println(config.Green("Neo_reGeorg 部署完成！"))
	utils.Log(fmt.Sprintf("-\tNeo_reGeorg\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["Neo_reGeorg"]))

	filename, version, description = utils.GitReleases("ehang-io", "nps", "windows_amd64_server.tar.gz", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["Nps"], false)
	fmt.Println(config.Green("Nps 部署完成！"))
	utils.Log(fmt.Sprintf("-\tNps\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["Nps"]))

	filename, version, description = utils.GitReleases("sensepost", "reGeorg", "master", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["reGeorg"], true)
	fmt.Println(config.Green("reGeorg 部署完成！"))
	utils.Log(fmt.Sprintf("-\treGeorg\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["reGeorg"]))

}

func portScanner() {

	filename, version, description := utils.GitReleases("AceofHearts404", "masscan", "masscan.exe", config.ToolDirectory["Masscan"])
	fmt.Println(config.Green("Masscan 部署完成！"))
	utils.Log(fmt.Sprintf("-\tMasscan\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["Masscan"]))

	filename, version, description = utils.GitReleases("projectdiscovery", "naabu", "windows_amd64.zip", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["Naabu"], false)
	fmt.Println(config.Green("Naabu 部署完成！"))
	utils.Log(fmt.Sprintf("-\tNaabu\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["Naabu"]))

	filename, version, description = utils.GitReleases("elddy", "NimScan", "NimScan.exe", config.ToolDirectory["NimScan"])
	fmt.Println(config.Green("NimScan 部署完成！"))
	utils.Log(fmt.Sprintf("-\tNimScan\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["NimScan"]))

	nmapCommand := []string{""}
	utils.ChocoInstall("Nmap", "nmap", config.ToolDirectory["Nmap"], nmapCommand)
	os.Rename("C:\\Program Files (x86)\\Nmap", config.ToolDirectory["Nmap"])

	filename, version, description = utils.GitReleases("AceofHearts404", "RouterScan", ".zip", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["RouterScan"], false)
	fmt.Println(config.Green("RouterScan 部署完成！"))
	utils.Log(fmt.Sprintf("-\tRouterScan\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["RouterScan"]))

	filename, version, description = utils.GitReleases("redtoolskobe", "scaninfo", "windows_x64.exe", config.ToolDirectory["Scaninfo"])
	fmt.Println(config.Green("Scaninfo 部署完成！"))
	utils.Log(fmt.Sprintf("-\tScaninfo\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["Scaninfo"]))

	filename, version, description = utils.GitReleases("4dogs-cn", "TXPortMap", "windows_x64.exe", config.ToolDirectory["TXPortMap"])
	fmt.Println(config.Green("TXPortMap 部署完成！"))
	utils.Log(fmt.Sprintf("-\tTXPortMap\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["TXPortMap"]))

}

func privilegeEscalation() {

	filename, version, description := utils.GitReleases("SecWiki", "windows-kernel-exploits", "master", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["Windows_Kernel_Exploits"], true)
	fmt.Println(config.Green("Windows_Kernel_Exploits 部署完成！"))
	utils.Log(fmt.Sprintf("-\tWindows_Kernel_Exploits\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["Windows_Kernel_Exploits"]))

	filename, version, description = utils.GitReleases("SecWiki", "linux-kernel-exploits", "master", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["Linux_Kernel_Exploits"], true)
	fmt.Println(config.Green("Linux_Kernel_Exploits 部署完成！"))
	utils.Log(fmt.Sprintf("-\tLinux_Kernel_Exploits\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["Linux_Kernel_Exploits"]))

	filename, version, description = utils.GitReleases("RowTeam", "SharpSQLTools", "SharpSQLToolsGUI.zip", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["SharpSQLToolsGUI"], false)
	fmt.Println(config.Green("SharpSQLToolsGUI 部署完成！"))
	utils.Log(fmt.Sprintf("-\tSharpSQLToolsGUI\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["SharpSQLToolsGUI"]))
}

func proxy() {

	//burpsuiteCommand := []string{fmt.Sprintf("--install-arguments=\"-dir %s\"", config.ToolDirectory["BurpSuite"])}
	//utils.ChocoInstall("BurpSuite", "burp-suite-pro-edition", config.ToolDirectory["BurpSuite"], burpsuiteCommand)
	filename, version, description := utils.GitReleases("AceofHearts404", "bp", ".zip", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["BurpSuite"], false)
	fmt.Println(config.Green("BurpSuite 部署完成！"))
	utils.Log(fmt.Sprintf("-\tBurpSuite\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["BurpSuite"]))

	proxifierCommand := []string{fmt.Sprintf("--install-arguments=\"/DIR=%s\"", config.ToolDirectory["Proxifier"])}
	utils.ChocoInstall("Proxifier", "proxifier", config.ToolDirectory["Proxifier"], proxifierCommand)

	filename, version, description = utils.GitReleases("AceofHearts404", "SocksCap64", ".zip", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["SocksCap64"], false)
	fmt.Println(config.Green("SocksCap64 部署完成！"))
	utils.Log(fmt.Sprintf("-\tSocksCap64\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["SocksCap64"]))

	wiresharkCommand := []string{fmt.Sprintf("--install-arguments=\"/D=%s\"", config.ToolDirectory["Wireshark"])}
	utils.ChocoInstall("Wireshark", "wireshark", config.ToolDirectory["Wireshark"], wiresharkCommand)

	filename, version, description = utils.GitReleases("yaklang", "yakit", ".exe", config.ToolDirectory["TEMP"])
	utils.Cmd(filename)
	yakitDirectory := filepath.Join(os.Getenv("USERPROFILE"), "\\AppData\\Local\\Programs\\yakit")
	err := os.Rename(yakitDirectory, config.ToolDirectory["Proxy"])
	if err != nil {
		fmt.Println(config.Red("重命名文件夹失败:"), err)
		return
	}
	fmt.Println(config.Green("Yakit 部署完成！"))
	utils.Log(fmt.Sprintf("-\tYakit\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["Yakit"]))
}

func vulnScanner() {
	filename, version, description := utils.GitReleases("zan8in", "afrog", "windows_amd64.zip", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["Afrog"], false)
	fmt.Println(config.Green("Afrog 部署完成！"))
	utils.Log(fmt.Sprintf("-\tAfrog\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["Afrog"]))

	filename, version, description = utils.GitReleases("projectdiscovery", "nuclei", "windows_amd64.zip", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["Nuclei"], false)
	fmt.Println(config.Green("Nuclei 部署完成！"))
	utils.Log(fmt.Sprintf("-\tNuclei\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["Nuclei"]))

	filename, version, description = utils.GitReleases("google", "osv-scanner", "windows_arm64.exe", config.ToolDirectory["OSV_Scanner"])
	fmt.Println(config.Green("OSV Scanner 部署完成！"))
	utils.Log(fmt.Sprintf("-\tOSV Scanner\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["OSV_Scanner"]))

	filename, version, description = utils.GitReleases("gobysec", "Goby", "goby-win", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["Goby"], true)
	fmt.Println(config.Green("Goby 部署完成！"))
	utils.Log(fmt.Sprintf("-\tGoby\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["Goby"]))

	filename, version, description = utils.GitReleases("veo", "vscan", "windows_amd64.zip", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["Vscan"], false)
	fmt.Println(config.Green("Vscan 部署完成！"))
	utils.Log(fmt.Sprintf("-\tVscan\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["Vscan"]))

	filename, version, description = utils.GitReleases("chaitin", "Xray", "windows_amd64.exe.zip", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["Xray"], false)
	fmt.Println(config.Green("Xray 部署完成！"))
	utils.Log(fmt.Sprintf("-\tXray\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["Xray"]))

	filename, version, description = utils.GitReleases("4ra1n", "super-xray", ".jar", config.ToolDirectory["Super_Xray"])
	fmt.Println(config.Green("Super_Xray 部署完成！"))
	utils.Log(fmt.Sprintf("-\tSuper_Xray\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["Super_Xray"]))
}

func vulnSearch() {

}

func webshell() {
	filename, version, description := utils.GitReleases("AntSwordProject", "AntSword-Loader", "win32-x64.zip", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["Antsword"], true)
	fmt.Println(config.Green("Antsword 部署完成！"))
	utils.Log(fmt.Sprintf("-\tAntSword\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["Antsword"]))

	filename, version, description = utils.GitReleases("rebeyond", "Behinder", ".zip", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["Behinder"], false)
	fmt.Println(config.Green("Behinder 部署完成！"))
	utils.Log(fmt.Sprintf("-\tBehinder\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["Behinder"]))

	filename, version, description = utils.GitReleases("BeichenDream", "Godzilla", ".jar", config.ToolDirectory["Godzilla"])
	fmt.Println(config.Green("Godzilla 部署完成！"))
	utils.Log(fmt.Sprintf("-\tGodzilla\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["Godzilla"]))

	filename, version, description = utils.GitReleases("shack2", "skyscorpion", ".zip", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["Skyscorpion"], false)
	fmt.Println(config.Green("Skyscorpion 部署完成！"))
	utils.Log(fmt.Sprintf("-\tSkyscorpion\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["Skyscorpion"]))

	filename, version, description = utils.GitReleases("tennc", "webshell", "master", config.ToolDirectory["TEMP"])
	utils.Unzip(filename, config.ToolDirectory["Webshell_Script"], true)
	fmt.Println(config.Green("Webshell_Script 部署完成！"))
	utils.Log(fmt.Sprintf("-\tWebshell_Script\n\t-\tVersion: %s\n\t-\t简介: %s\n\t-\t安装路径: %s\n", version, description, config.ToolDirectory["Webshell_Script"]))
}
