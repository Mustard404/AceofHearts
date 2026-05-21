# AceofHearts

![Languages](https://img.shields.io/github/languages/top/Mustard404/AceofHearts?color=yellow)
![Tag](https://img.shields.io/github/v/tag/Mustard404/AceofHearts?color=red)
![Stars](https://img.shields.io/github/stars/Mustard404/AceofHearts.svg)

红桃A(AceofHearts)是一款专为渗透测试人员设计的实用工具，旨在简化渗透测试环境的搭建过程并提供便捷的部署解决方案。

## 功能特点

- **突破环境限制**：支持在物理机、VMware、VirtualBox、Parallels Desktop 等虚拟化环境中部署，灵活适应各种场景。
- **多系统支持**：支持 Windows 10 / Windows 11 操作系统。
- **实时更新**：通过 Chocolatey 和 GitHub Releases 获取最新版本工具，保持工具与最新开发进展同步。
- **GitHub 镜像加速**：内置镜像站支持（如 ghfast.top），自动预检可用性，国内网络也能流畅部署。
- **断点续装**：基于 Excel 工具清单驱动安装，中断后重新运行自动跳过已安装工具。
- **下载重试**：网络波动时自动重试 3 次，提高部署成功率。
- **代码开源**：杜绝恶意代码注入风险，支持自行编译和审查。

## 项目声明

- **项目依赖**：部分 Python 项目需手动安装依赖（pip install）。
- **破解激活**：考虑到安全性及合法性，仅将破解程序下载到 Tools 目录下，请自行处理。
- **工具来源**：绝大部分工具来自项目官方，尽可能保证安全性，但不排除开源项目中可能包含恶意代码。
- **版权说明**：项目中引用多数开源项目，如造成侵权请提 Issues 进行删除。

## 免责声明

- 渗透测试工具仅用于合法的安全评估和授权测试目的。
- 未经合法授权，禁止对未授权目标进行测试。
- 使用者应理解使用本工具可能涉及的风险，并自行承担全部责任。
- 开发者不对使用者的行为和后果负责，不承担任何明示或暗示的保证。

## 安装

### 准备

- **系统代理**：大部分工具来自 GitHub，建议配置系统代理提高部署速度。
- **GitHub 镜像**（可选）：无代理环境下可使用 GitHub 镜像站加速。
- **管理员权限**：需以管理员身份运行。

### 使用

1. 通过 [Releases](https://github.com/Mustard404/AceofHearts/releases) 下载最新版本
2. 将程序移动到目标系统中（建议使用全新系统环境）
3. 以**管理员身份**运行：

```
AceofHearts.exe
```

4. 按提示配置：
   - **代理地址**（例：`http://127.0.0.1:7890`，留空跳过）
   - **GitHub 镜像站**（例：`https://ghfast.top/`，留空跳过）
   - **工具部署目录**（默认 `C:\Tools`）

5. 程序将自动完成：系统优化 → Chocolatey 部署 → 工具安装 → 启动器部署 → 命令别名配置

### GitHub API 速率

程序内置 GitHub OAuth App 凭证（5000 次/小时），也可自行注册：

1. 访问 [GitHub Developer Settings](https://github.com/settings/developers)
2. OAuth Apps → New OAuth App
3. 获取 Client ID 和 Client Secrets

## 部署流程

```
AceofHearts.exe
  ├─ 代理 & 镜像配置
  ├─ GitHub API Token 配置
  ├─ 部署路径设置
  ├─ 镜像可用性预检
  ├─ 系统信息获取 & 管理员验证
  ├─ 系统优化
  │   ├─ Windows Defender 白名单
  │   ├─ 显示文件扩展名
  │   ├─ 开启 Telnet 客户端
  │   └─ Chocolatey 部署
  ├─ 工具安装（Excel 清单驱动，支持断点续装）
  │   ├─ Chocolatey 安装（自定义安装路径）
  │   ├─ GitHub Releases 下载（镜像加速 + 重试）
  │   ├─ GitHub Clone（镜像加速）
  │   └─ URL 直接下载
  ├─ GeekDesk 启动器部署
  └─ Doskey 命令别名配置
```

## 工具清单

工具列表通过 `tools.xlsx` 配置，支持以下安装方式：

| 安装方式 | 说明 |
|---------|------|
| Chocolatey | 通过 Chocolatey 包管理器安装，支持自定义安装路径 |
| GitHub Releases | 从 GitHub Releases 下载最新版本 |
| GitHub GitClone | 浅克隆 GitHub 仓库 |
| Download | 从指定 URL 直接下载 |

### 编程环境

Git、Go、Java8、Python3、Python2、Node.js、Ruby、Laragon

### 生产力工具

Bandizip、Bitwarden、Chrome、Everything、Firefox、Sublime Text、Typora

### 开发工具

MobaXterm、Navicat Premium、VSCode、WinSCP、Clink

### 代理抓包

BurpSuite、Proxifier、SocksCap64、Wireshark、Yakit

### 端口扫描

Masscan、Naabu、NimScan、Nmap、Scaninfo、TXPortMap

### 渗透测试工具

包含免杀、C2、目录扫描、漏洞利用、指纹识别、信息收集、权限维持、密码攻击、渗透框架、端口转发、提权辅助、漏洞扫描、Webshell 等分类，共 190+ 工具。

详见 `tools.xlsx` 工具清单。

## 自定义工具清单

编辑 `tools.xlsx` 可自定义安装工具：

| 列 | 说明 |
|----|------|
| ID | 工具编号 |
| Name | 工具名称 |
| Path | 安装子路径（相对于部署目录） |
| InstallMethod | 安装方式：Chocolatey / GitHub / Download |
| Category | 分类名称 |
| URL | 下载地址（Download 方式） |
| InstallType | GitHub 类型：Releases / GitClone / Packages |
| ChocoPkg | Chocolatey 包名 |
| ChocoArgs | Chocolatey 自定义参数 |
| GitHubUser | GitHub 用户名 |
| GitHubRepo | GitHub 仓库名 |
| ReleaseKey | Release 资产匹配关键词 |
| Status | 安装状态：未安装 / 已安装 / 失败 |

## Star History

<a href="https://star-history.com/#Mustard404/AceofHearts&Date">
  <picture>
    <source media="(prefers-color-scheme: dark)" srcset="https://api.star-history.com/svg?repos=Mustard404/AceofHearts&type=Date&theme=dark" />
    <source media="(prefers-color-scheme: light)" srcset="https://api.star-history.com/svg?repos=Mustard404/AceofHearts&type=Date" />
    <img alt="Star History Chart" src="https://api.star-history.com/svg?repos=Mustard404/AceofHearts&type=Date" />
  </picture>
</a>
