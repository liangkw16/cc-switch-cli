#!/bin/bash
set -e

# CCS 一行卸载命令:
# curl -fsSL https://raw.githubusercontent.com/liangkw16/cc-switch-cli/main/uninstall.sh | sudo sh

echo "正在卸载 CCS..."

# 删除二进制
sudo rm -f /usr/local/bin/ccs

# 删除配置
rm -rf ~/.config/ccs ~/.config/cc-switch

echo "CCS 已卸载！"
