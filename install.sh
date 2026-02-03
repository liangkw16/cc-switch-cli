#!/bin/bash
set -e

# CCS 一行安装命令:
# curl -fsSL https://raw.githubusercontent.com/liangkw16/cc-switch-cli/main/install.sh | sudo sh

# 检测系统架构
ARCH=$(uname -m)
OS=$(uname -s)

case $OS in
    Darwin*)
        FILE="ccs-darwin-$ARCH"
        ;;
    Linux*)
        FILE="ccs-linux-$ARCH"
        ;;
    *)
        echo "不支持的操作系统: $OS"
        exit 1
        ;;
esac

# 下载最新版本
echo "正在下载 CCS..."
LATEST_URL="https://github.com/liangkw16/cc-switch-cli/releases/latest/download/ccs"
curl -fsSL "$LATEST_URL" -o /tmp/ccs

# 安装
echo "正在安装到 /usr/local/bin/ccs..."
sudo install -m 755 /tmp/ccs /usr/local/bin/ccs

# 清理
rm -f /tmp/ccs

echo "CCS 安装完成！"
echo "使用 'ccs --help' 查看命令"
