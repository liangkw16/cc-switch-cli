# CCS - Claude Code Switcher

一个简单的 CLI 工具，用于管理和切换 Claude Code 配置档案。

## 功能

- **CLI 命令**：`add`、`list`、`use`、`remove` 操作配置
- **TUI 界面**：交互式分屏界面（左侧档案列表 + 右侧配置预览）
- **自动备份**：切换前自动备份，保留最近 5 个备份

## 安装

### 一行安装命令

```bash
curl -fsSL https://raw.githubusercontent.com/liangkw16/cc-switch-cli/main/install.sh | sudo sh
```

### 一行卸载命令

```bash
curl -fsSL https://raw.githubusercontent.com/liangkw16/cc-switch-cli/main/uninstall.sh | sudo sh
```

### 手动安装

```bash
git clone git@github.com:liangkw16/cc-switch-cli.git
cd cc-switch-cli
sudo install -m 755 ccs /usr/local/bin/ccs
```

## 使用

### 添加配置档案

```bash
ccs add glm
```

按提示输入配置：

```
ANTHROPIC_AUTH_TOKEN: liangkaiwen
ANTHROPIC_BASE_URL: https://glm47.codedancer.bytedance.net
ANTHROPIC_DEFAULT_HAIKU_MODEL: GPT-5.2-Codex
ANTHROPIC_DEFAULT_OPUS_MODEL: GPT-5.2-Codex
ANTHROPIC_DEFAULT_SONNET_MODEL: GPT-5.2-Codex
ANTHROPIC_MODEL: GPT-5.2-Codex
```

### 列出所有档案

```bash
ccs list
```

### 切换档案

```bash
ccs use glm
```

切换后重启终端或 Claude Code 即可生效。

### 删除档案

```bash
ccs remove glm
```

### 启动 TUI

```bash
ccs ui
```

## TUI 界面

```
┌────────────────────────────────────────────────────────────┐
│  CCS - Claude Code Switcher                        │
├──────────────────────┬─────────────────────────────────┤
│                      │                             │
│  Profiles            │  Profile Details              │
│                      │                             │
│  > glm (active)      │  API Token                  │
│    official          │    lian***aiwen            │
│                      │                             │
│                      │  Base URL                   │
│                      │    https://glm47...          │
│                      │                             │
│                      │  Default Model              │
│                      │    GPT-5.2-Codex           │
│                      │                             │
│  [Enter] 切换  [r] 删除  [q] 退出        │
└──────────────────────┴─────────────────────────────────┘
```

**快捷键：**
- `↑/↓` 或 `j/k` - 上下选择
- `Enter` - 切换到选中的档案
- `r` - 删除选中的档案
- `q` - 退出

## 配置字段

| 字段 | 说明 |
|-----|------|
| `ANTHROPIC_AUTH_TOKEN` | API 认证令牌 |
| `ANTHROPIC_BASE_URL` | 自定义 API 地址 |
| `ANTHROPIC_DEFAULT_HAIKU_MODEL` | Haiku 模型名称 |
| `ANTHROPIC_DEFAULT_OPUS_MODEL` | Opus 模型名称 |
| `ANTHROPIC_DEFAULT_SONNET_MODEL` | Sonnet 模型名称 |
| `ANTHROPIC_MODEL` | 默认/回退模型 |

## 命令

| 命令 | 说明 |
|-----|------|
| `ccs add <name>` | 添加新档案 |
| `ccs list` | 列出所有档案 |
| `ccs use <name>` | 切换到指定档案 |
| `ccs remove <name>` | 删除档案 |
| `ccs ui` | 启动交互界面 |

## 配置文件

- **配置存储**：`~/.ccs/profiles.json`
- **Claude 配置**：`~/.claude/settings.json`（或 `claude.json`）
- **备份目录**：`~/.ccs/backups/`

## 开发

```bash
# 克隆仓库
git clone git@github.com:liangkw16/cc-switch-cli.git
cd cc-switch-cli

# 运行
go run .

# 构建
go build -o ccs
```

## License

MIT
