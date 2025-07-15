#!/bin/bash
# TimeLine 桌面计时器卸载脚本

set -e

echo "=== TimeLine 桌面计时器卸载脚本 ==="
echo

# 设置目录
INSTALL_DIR="$HOME/.local/share/timeline"
BIN_DIR="$HOME/.local/bin"
DESKTOP_DIR="$HOME/.local/share/applications"
CONFIG_DIR="$HOME/.config/timeline"

echo "将要删除以下文件和目录:"
echo "- 应用程序目录: $INSTALL_DIR"
echo "- 启动脚本: $BIN_DIR/timeline"
echo "- 桌面文件: $DESKTOP_DIR/timeline.desktop"
echo "- 配置目录: $CONFIG_DIR (包含用户数据)"
echo

read -p "确定要卸载 TimeLine 计时器吗? (y/N): " -n 1 -r
echo

if [[ ! $REPLY =~ ^[Yy]$ ]]; then
    echo "取消卸载"
    exit 0
fi

echo "开始卸载..."

# 删除应用程序目录
if [ -d "$INSTALL_DIR" ]; then
    echo "删除应用程序目录..."
    rm -rf "$INSTALL_DIR"
fi

# 删除启动脚本
if [ -f "$BIN_DIR/timeline" ]; then
    echo "删除启动脚本..."
    rm -f "$BIN_DIR/timeline"
fi

# 删除桌面文件
if [ -f "$DESKTOP_DIR/timeline.desktop" ]; then
    echo "删除桌面文件..."
    rm -f "$DESKTOP_DIR/timeline.desktop"
fi

# 询问是否删除配置目录
if [ -d "$CONFIG_DIR" ]; then
    echo
    read -p "是否删除配置和用户数据目录? ($CONFIG_DIR) (y/N): " -n 1 -r
    echo
    if [[ $REPLY =~ ^[Yy]$ ]]; then
        echo "删除配置目录..."
        rm -rf "$CONFIG_DIR"
    else
        echo "保留配置目录: $CONFIG_DIR"
    fi
fi

# 更新桌面数据库
if command -v update-desktop-database &> /dev/null; then
    echo "更新桌面数据库..."
    update-desktop-database "$DESKTOP_DIR" 2>/dev/null || true
fi

echo
echo "=== 卸载完成! ==="
echo "TimeLine 计时器已从系统中移除。"

if [ -d "$CONFIG_DIR" ]; then
    echo "配置目录已保留: $CONFIG_DIR"
    echo "如需完全清理，请手动删除该目录。"
fi

echo
