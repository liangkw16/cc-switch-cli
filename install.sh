#!/bin/bash
set -e

# CCS 源码部署安装命令:
# curl -fsSL https://raw.githubusercontent.com/liangkw16/cc-switch-cli/main/install.sh | sudo sh

# 检查 Go 环境
if ! command -v go &>/dev/null 2>&1; then
    echo "错误: 未安装 Go，请先安装 Go"
    echo "  访问 https://go.dev/dl/ 下载"
    exit 1
fi

# 获取系统信息
ARCH=$(uname -m)
OS=$(uname -s)

# 创建临时目录
TMP_DIR=$(mktemp -d)
trap "rm -rf $TMP_DIR" EXIT

echo "正在克隆源码..."
git clone --depth 1 https://github.com/liangkw16/cc-switch-cli.git "$TMP_DIR"

echo "正在编译..."
cd "$TMP_DIR/cc-switch-cli"

# 编译
go build -o ccs

# 安装
echo "正在安装到 /usr/local/bin/ccs..."
sudo install -m 755 ccs /usr/local/bin/ccs

# 清理
rm -rf "$TMP_DIR"

echo "CCS 安装完成！"
echo "使用 'ccs --help' 查看命令"
