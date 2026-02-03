#!/bin/bash
set -e

# CCS 一行安装命令:
# curl -fsSL https://raw.githubusercontent.com/liangkw16/cc-switch-cli/main/install.sh | sudo sh

# 检测系统架构
ARCH=$(uname -m)
OS=$(uname -s)

# 处理架构名称映射
case $ARCH in
    x86_64)
        ARCH="amd64"
        ;;
    aarch64)
        ARCH="arm64"
        ;;
esac

case $OS in
    Darwin*)
        FILE="ccs-darwin-amd64"
        ;;
    Linux*)
        FILE="ccs-linux-$ARCH"
        ;;
    *)
        echo "不支持的操作系统: $OS"
        exit 1
        ;;
esac

# 获取最新版本号（使用 jq 更可靠）
echo "正在获取最新版本..."
if command -v jq &>/dev/null 2>&1; then
    VERSION=$(curl -fsSL https://api.github.com/repos/liangkw16/cc-switch-cli/releases/latest | jq -r '.tag_name')
else
    # 如果没有 jq，使用简单的 grep + cut 方式
    VERSION=$(curl -fsSL https://api.github.com/repos/liangkw16/cc-switch-cli/releases/latest | grep -o '"tag_name":"[^"]*"' | cut -d'"' -f4)
fi

# 确保版本号有 v 前缀（如果没有则添加）
if [[ "$VERSION" != v* ]]; then
    VERSION="v$VERSION"
fi

# 下载
echo "正在下载 CCS $VERSION..."
LATEST_URL="https://github.com/liangkw16/cc-switch-cli/releases/download/$VERSION/$FILE"
curl -fsSL "$LATEST_URL" -o /tmp/ccs

# 安装
echo "正在安装到 /usr/local/bin/ccs..."
sudo install -m 755 /tmp/ccs /usr/local/bin/ccs

# 清理
rm -f /tmp/ccs

echo "CCS $VERSION 安装完成！"
echo "使用 'ccs --help' 查看命令"
