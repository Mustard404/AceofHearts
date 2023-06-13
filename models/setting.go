package models

import (
	"AceofHearts/config"
	"AceofHearts/utils"
	"bufio"
	"encoding/base64"
	"fmt"
	"gopkg.in/ini.v1"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func Setting() {
	GenerateAccessToken()
	SetDirectory()
}

func GenerateAccessToken() {
	utils.Title("读取Github API")

	conFine := filepath.Join(config.ExecutablePath, "config.ini")
	cfg, err := ini.Load(conFine)
	if err != nil {
		log.Fatal("Error loading config file:", err)
	}

	// 获取配置值
	clientID := cfg.Section("Github").Key("clientID").String()
	clientSecrets := cfg.Section("Github").Key("clientSecrets").String()

	fmt.Println(fmt.Sprintf("clientID: %s", clientID))
	fmt.Println(fmt.Sprintf("clientSecrets: %s", clientSecrets))

	if len(clientID) == 0 || len(clientSecrets) == 0 {
		fmt.Print(config.Red("Client ID和Client Secrets不可为空"))
		os.Exit(0)
	}

	accessToken := fmt.Sprintf("%s:%s", clientID, clientSecrets)
	accessTokenBytes := []byte(accessToken)
	encodedAccessToken := base64.StdEncoding.EncodeToString(accessTokenBytes)
	rateLimit := checkRateLimit(encodedAccessToken)

	if rateLimit < 100 {
		fmt.Println(fmt.Sprintf("此AccessToken速率剩余: %s次！", config.Red(rateLimit)))
		fmt.Println(config.Red("此AccessToken速率剩余较低，请替换或过一小时后重试！"))
	} else {
		fmt.Println(fmt.Sprintf("此AccessToken速率剩余: %s次！", config.Green(rateLimit)))
	}

	config.AccessToken = encodedAccessToken
}

func checkRateLimit(encodedAccessToken string) int {
	req, _ := http.NewRequest("GET", "https://api.github.com/users/Mustard404", nil)
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", encodedAccessToken))
	client := http.Client{}
	resp, _ := client.Do(req)
	rateLimitStr := resp.Header.Get("X-RateLimit-Remaining")
	rateLimit, err := strconv.Atoi(rateLimitStr)
	if err != nil {
		fmt.Println(config.Red("无法获取API速率!"))
	}
	return rateLimit
}

func SetDirectory() {
	utils.Title("设置部署路径")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入工具包部署目录(默认为 C:\\Tools)：")
	toolsDir, _ := reader.ReadString('\n')
	toolsDir = strings.TrimSpace(toolsDir)
	if len(toolsDir) == 0 {
		toolsDir = "C:\\Tools"
	}
	_, err := os.Stat(toolsDir)
	if err == nil {
		message := fmt.Sprintf("注: 目录 %s 已存在!\n", toolsDir)
		fmt.Printf(config.Red(message))
	} else {
		err = os.Mkdir(toolsDir, 0755)
		if err != nil {
			fmt.Println(config.Red("无法创建目录:", err))
		}
	}
	fmt.Printf("设置部署路径为: %s\n", config.Magenta(toolsDir))
	utils.Log("-\t设置部署路径为: %s\n", toolsDir)
	fmt.Print("按任意键继续...")
	_, _ = reader.ReadString('\n')

	folders := map[string]string{
		"ROOT":                              "",
		"TEMP":                              "\\TEMP",
		"Doskey":                            "\\Doskey",
		"Env":                               "\\Env",
		"Git":                               "\\Env\\Git",
		"Go":                                "\\Env\\Go",
		"Java8":                             "\\Env\\Java\\Java8",
		"Laragon":                           "\\Env\\Laragon",
		"Node":                              "\\Env\\Node",
		"Python2":                           "\\Env\\Python\\Python2",
		"Python3":                           "\\Env\\Python\\Python3",
		"Ruby":                              "\\Env\\Ruby",
		"Productivity":                      "\\Productivity",
		"Bandizip":                          "\\Productivity\\Bandizip",
		"Bitwarden":                         "\\Productivity\\Bitwarden",
		"Everything":                        "\\Productivity\\Everything",
		"Firefox":                           "\\Productivity\\Firefox",
		"Sublime_Text":                      "\\Productivity\\Sublime_Text",
		"Typora":                            "\\Productivity\\Typora",
		"Development":                       "\\Development",
		"Mobaxterm":                         "\\Development\\Mobaxterm",
		"Navicat_Premium":                   "\\Development\\Navicat_Premium",
		"Vscode":                            "\\Development\\Vscode",
		"Winscp":                            "\\Development\\Winscp",
		"Bypass_AV":                         "\\Bypass_AV",
		"AV_Evasion_Tool":                   "\\Bypass_AV\\AV_Evasion_Tool",
		"BypassAntiVirus":                   "\\Bypass_AV\\BypassAntiVirus",
		"GBByPass":                          "\\Bypass_AV\\GBByPass",
		"MateuszEx":                         "\\Bypass_AV\\MateuszEx",
		"Shellcodeloader":                   "\\Bypass_AV\\Shellcodeloader",
		"C2_Framework":                      "\\C2_Framework",
		"Cobalt_Strike":                     "\\C2_Framework\\Cobalt_Strike",
		"OLa":                               "\\C2_Framework\\Cobalt_Strike\\scripts",
		"Sliver":                            "\\C2_Framework\\Sliver",
		"Dir_Scanner":                       "\\Dir_Scanner",
		"Dirmap":                            "\\Dir_Scanner\\Dirmap",
		"Dirsearch":                         "\\Dir_Scanner\\Dirsearch",
		"Ffuf":                              "\\Dir_Scanner\\Ffuf",
		"JSFinder":                          "\\Dir_Scanner\\JSFinder",
		"WebPathBrute":                      "\\Dir_Scanner\\WebPathBrute",
		"EXP":                               "\\EXP",
		"Advanced_SQL_Injection_Cheatsheet": "\\EXP\\Advanced_SQL_Injection_Cheatsheet",
		"CF":                                "\\EXP\\CF",
		"Commix":                            "\\EXP\\Commix",
		"Dalfox":                            "\\EXP\\Dalfox",
		"Fuxploider":                        "\\EXP\\Fuxploider",
		"JNDIExploit":                       "\\EXP\\JNDIExploit",
		"LFISuite":                          "\\EXP\\LFISuite",
		"ShiroExploit":                      "\\EXP\\ShiroExploit",
		"ShiroAttack2":                      "\\EXP\\ShiroAttack2",
		"Sqlmap":                            "\\EXP\\Sqlmap",
		"SuperSQLInjection":                 "\\EXP\\SuperSQLInjection",
		"XSStrike":                          "\\EXP\\XSStrike",
		"XXEinjector":                       "\\EXP\\XXEinjector",
		"Xssor2":                            "\\EXP\\Xssor2",
		"WeblogicTool":                      "\\EXP\\WeblogicTool",
		"Ysomap":                            "\\EXP\\Ysomap",
		"Fingerprinting":                    "\\Fingerprinting",
		"CMSeeK":                            "\\Fingerprinting\\CMSeeK",
		"EHole":                             "\\Fingerprinting\\EHole",
		"ObserverWard":                      "\\Fingerprinting\\ObserverWard",
		"TideFinger":                        "\\Fingerprinting\\TideFinger",
		"Info_Gathering":                    "\\Info_Gathering",
		"Amass":                             "\\Info_Gathering\\Amass",
		"Ct_Exposer":                        "\\Info_Gathering\\Ct_Exposer",
		"Dnsx":                              "\\Info_Gathering\\Dnsx",
		"Ds_store_exp":                      "\\Info_Gathering\\Ds_store_exp",
		"Fofa_Viewer":                       "\\Info_Gathering\\Fofa_Viewer",
		"Gitgot":                            "\\Info_Gathering\\GitGot",
		"Gitgraber":                         "\\Info_Gathering\\GitGraber",
		"Githack":                           "\\Info_Gathering\\GitHack",
		"Gitminer":                          "\\Info_Gathering\\GitMiner",
		"Gitrob":                            "\\Info_Gathering\\Gitrob",
		"Gobuster":                          "\\Info_Gathering\\Gobuster",
		"Ksubdomain":                        "\\Info_Gathering\\Ksubdomain",
		"Layer":                             "\\Info_Gathering\\Layer",
		"Oneforall":                         "\\Info_Gathering\\OneForAll",
		"Spiderfoot":                        "\\Info_Gathering\\Spiderfoot",
		"Subfinder":                         "\\Info_Gathering\\Subfinder",
		"Svnexploit":                        "\\Info_Gathering\\SvnExploit",
		"Theharvester":                      "\\Info_Gathering\\TheHarvester",
		"Maintain_Access":                   "\\Maintain_Access",
		"Girsh":                             "\\Maintain_Access\\Girsh",
		"Govenom":                           "\\Maintain_Access\\Govenom",
		"Platypus":                          "\\Maintain_Access\\Platypus",
		"Reverse_ssh":                       "\\Maintain_Access\\Reverse_ssh",
		"Pass_Attack":                       "\\Pass_Attack",
		"HackBrowserData":                   "\\Pass_Attack\\HackBrowserData",
		"Hashcat":                           "\\Pass_Attack\\Hashcat",
		"John":                              "\\Pass_Attack\\John",
		"Johnny":                            "\\Pass_Attack\\Johnny",
		"Patator":                           "\\Pass_Attack\\Patator",
		"Probable_Wordlists":                "\\Pass_Attack\\Probable_Wordlists",
		"Psudohash":                         "\\Pass_Attack\\Psudohash",
		"Thc_Hydra":                         "\\Pass_Attack\\Thc_Hydra",
		"Penetration_Testing_Framework":     "\\Penetration_Testing_Framework",
		"Metasploit":                        "\\Penetration_Testing_Framework\\Metasploit",
		"Pocsuite3":                         "\\Penetration_Testing_Framework\\Pocsuite3",
		"Port_Forwarding":                   "\\Port_Forwarding",
		"Frp":                               "\\Port_Forwarding\\Frp",
		"Iox":                               "\\Port_Forwarding\\Iox",
		"Neo_reGeorg":                       "\\Port_Forwarding\\Neo_reGeorg",
		"Nps":                               "\\Port_Forwarding\\Nps",
		"reGeorg":                           "\\Port_Forwarding\\reGeorg",
		"Port_Scanner":                      "\\Port_Scanner",
		"GoScan":                            "\\Port_Scanner\\GoScan",
		"Masscan":                           "\\Port_Scanner\\Masscan",
		"Naabu":                             "\\Port_Scanner\\Naabu",
		"NimScan":                           "\\Port_Scanner\\NimScan",
		"Nmap":                              "\\Port_Scanner\\Nmap",
		"RouterScan":                        "\\Port_Scanner\\RouterScan",
		"Scaninfo":                          "\\Port_Scanner\\Scaninfo",
		"Sx":                                "\\Port_Scanner\\Sx",
		"TXPortMap":                         "\\Port_Scanner\\TXPortMap",
		"Yujian":                            "\\Port_Scanner\\Yujian",
		"Privilege_Escalation":              "\\Privilege_Escalation",
		"Windows_Kernel_Exploits":           "\\Privilege_Escalation\\Windows_Kernel_Exploits",
		"Linux_Kernel_Exploits":             "\\Privilege_Escalation\\Linux_Kernel_Exploits",
		"SharpSQLToolsGUI":                  "\\Privilege_Escalation\\SharpSQLToolsGUI",
		"Proxy":                             "\\Proxy",
		"BurpSuite":                         "\\Proxy\\BurpSuite",
		"Proxifier":                         "\\Proxy\\Proxifier",
		"SocksCap64":                        "\\Proxy\\SocksCap64",
		"Wireshark":                         "\\Proxy\\Wireshark",
		"Yakit":                             "\\Proxy\\yakit",
		"Vuln_Scanner":                      "\\Vuln_Scanner",
		"Afrog":                             "\\Vuln_Scanner\\Afrog",
		"Nuclei":                            "\\Vuln_Scanner\\Nuclei",
		"OSV_Scanner":                       "\\Vuln_Scanner\\OSV_Scanner",
		"Goby":                              "\\Vuln_Scanner\\Goby",
		"Vscan":                             "\\Vuln_Scanner\\Vscan",
		"Xray":                              "\\Vuln_Scanner\\Xray",
		"Super_Xray":                        "\\Vuln_Scanner\\Super_Xray",
		"Vuln_Search":                       "\\Vuln_Search",
		"Webshell":                          "\\Webshell",
		"Antsword":                          "\\Webshell\\Antsword",
		"Behinder":                          "\\Webshell\\Behinder",
		"Godzilla":                          "\\Webshell\\Godzilla",
		"Skyscorpion":                       "\\Webshell\\Skyscorpion",
		"Webshell_Script":                   "\\Webshell\\Webshell_Script",
		"Dawn_Launcher":                     "\\Dawn_Launcher",
	}

	for key, value := range folders {
		config.ToolDirectory[key] = filepath.Join(toolsDir, value)
	}

}
