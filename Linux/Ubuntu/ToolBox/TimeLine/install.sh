#!/bin/bash
# TimeLine 桌面计时器安装脚本
# 适用于 Ubuntu 20.04.6 LTS

set -e

echo "=== TimeLine 桌面计时器安装脚本 ==="
echo "适用系统: Ubuntu 20.04.6 LTS (GNOME 3.36.8)"
echo

# 检查系统
if ! command -v python3 &> /dev/null; then
    echo "错误: 未找到 Python3，请先安装 Python3"
    exit 1
fi

# 检查GTK依赖
echo "检查系统依赖..."
MISSING_DEPS=()

if ! python3 -c "import gi" 2>/dev/null; then
    MISSING_DEPS+=("python3-gi")
fi

if ! python3 -c "import gi; gi.require_version('Gtk', '3.0'); from gi.repository import Gtk" 2>/dev/null; then
    MISSING_DEPS+=("python3-gi-cairo" "gir1.2-gtk-3.0")
fi

if [ ${#MISSING_DEPS[@]} -ne 0 ]; then
    echo "需要安装以下依赖包:"
    printf '%s\n' "${MISSING_DEPS[@]}"
    echo
    read -p "是否现在安装这些依赖? (y/N): " -n 1 -r
    echo
    if [[ $REPLY =~ ^[Yy]$ ]]; then
        echo "安装依赖包..."
        sudo apt update
        sudo apt install -y "${MISSING_DEPS[@]}"
    else
        echo "请手动安装依赖包后重新运行安装脚本"
        echo "命令: sudo apt install ${MISSING_DEPS[*]}"
        exit 1
    fi
fi

# 设置安装目录
INSTALL_DIR="$HOME/.local/share/timeline"
BIN_DIR="$HOME/.local/bin"
DESKTOP_DIR="$HOME/.local/share/applications"

echo "创建安装目录..."
mkdir -p "$INSTALL_DIR"
mkdir -p "$BIN_DIR"
mkdir -p "$DESKTOP_DIR"

# 复制应用程序文件
echo "复制应用程序文件..."
cp -r src/ "$INSTALL_DIR/"
cp -r config/ "$INSTALL_DIR/"
cp -r assets/ "$INSTALL_DIR/"
cp timeline.py "$INSTALL_DIR/"
cp requirements.txt "$INSTALL_DIR/"

# 创建启动脚本
echo "创建启动脚本..."
cat > "$BIN_DIR/timeline" << EOF
#!/bin/bash
cd "$INSTALL_DIR"
python3 timeline.py "\$@"
EOF

chmod +x "$BIN_DIR/timeline"

# 创建桌面文件
echo "创建桌面文件..."
cat > "$DESKTOP_DIR/timeline.desktop" << EOF
[Desktop Entry]
Version=1.0
Type=Application
Name=TimeLine 计时器
Name[en]=TimeLine Timer
Comment=个人学习时间管理桌面计时器
Comment[en]=Personal study time management desktop timer
GenericName=Timer
GenericName[en]=Timer
Exec=$BIN_DIR/timeline
Icon=appointment-soon
Terminal=false
StartupNotify=true
Categories=Utility;Office;GTK;
Keywords=timer;study;productivity;time;management;
MimeType=
StartupWMClass=timeline
EOF

# 更新桌面数据库
if command -v update-desktop-database &> /dev/null; then
    echo "更新桌面数据库..."
    update-desktop-database "$DESKTOP_DIR"
fi

echo
echo "=== 安装完成! ==="
echo "应用程序已安装到: $INSTALL_DIR"
echo "启动脚本位置: $BIN_DIR/timeline"
echo "桌面文件位置: $DESKTOP_DIR/timeline.desktop"
echo
echo "使用方法:"
echo "1. 命令行启动: timeline"
echo "2. 从应用程序菜单启动: 搜索 'TimeLine'"
echo "3. 直接运行: $BIN_DIR/timeline"
echo
echo "快捷键:"
echo "- Ctrl+M: 最小化/恢复窗口"
echo "- Ctrl+T: 切换窗口置顶"
echo "- Ctrl+O: 切换透明度"
echo "- Escape: 重置窗口位置"
echo "- 右键点击: 显示上下文菜单"
echo
echo "配置文件位置: ~/.config/timeline/"
echo

# 检查PATH
if [[ ":$PATH:" != *":$BIN_DIR:"* ]]; then
    echo "注意: $BIN_DIR 不在您的 PATH 中"
    echo "请将以下行添加到您的 ~/.bashrc 或 ~/.profile 文件中:"
    echo "export PATH=\"\$PATH:$BIN_DIR\""
    echo "然后运行: source ~/.bashrc"
    echo
fi

echo "安装完成! 您现在可以启动 TimeLine 计时器了。"
