# TimeLine 桌面计时器

一个专为Ubuntu 20.04.6 LTS (GNOME 3.36.8) 设计的个人学习时间管理桌面应用程序。

## 功能特性

### 🎯 核心功能
- **固定学习时间段**: 自动识别三个预设学习时间段
  - 上午: 08:30-11:30
  - 下午: 14:30-17:30  
  - 晚上: 19:30-22:30
- **实时进度显示**: 显示当前时间段的学习进度和剩余时间
- **状态管理**: 手动切换学习/休息模式，带有视觉反馈
- **智能提醒**: 自动识别当前是否在学习时间段内

### 🎨 界面设计
- **GNOME风格**: 完美融入GNOME桌面环境
- **半透明窗口**: 优雅的半透明效果，不干扰工作
- **紧凑布局**: 小巧精致，占用最少桌面空间
- **状态指示**: 不同颜色和图标区分学习/休息状态

### 🖱️ 交互功能
- **窗口拖拽**: 点击任意位置拖拽移动窗口
- **右键菜单**: 丰富的上下文菜单选项
- **键盘快捷键**: 便捷的快捷键操作
- **位置记忆**: 自动保存和恢复窗口位置

## 系统要求

- **操作系统**: Ubuntu 20.04.6 LTS
- **桌面环境**: GNOME 3.36.8
- **窗口系统**: X11
- **Python版本**: Python 3.8+
- **依赖包**: GTK3, PyGObject

## 快速安装

### 自动安装（推荐）

```bash
# 克隆项目
git clone <repository-url>
cd TimeLine

# 运行安装脚本
./install.sh
```

安装脚本会自动：
- 检查并安装系统依赖
- 复制应用程序文件到用户目录
- 创建启动脚本和桌面文件
- 配置应用程序菜单

### 手动安装

1. **安装系统依赖**:
```bash
sudo apt update
sudo apt install python3-gi python3-gi-cairo gir1.2-gtk-3.0
```

2. **运行应用程序**:
```bash
python3 timeline.py
```

## 使用方法

### 启动应用程序

安装后可通过以下方式启动：

1. **命令行**: `timeline`
2. **应用程序菜单**: 搜索 "TimeLine"
3. **直接运行**: `~/.local/bin/timeline`

### 基本操作

- **查看进度**: 应用程序会自动显示当前时间段的学习进度
- **切换状态**: 点击状态按钮在学习/休息模式间切换
- **移动窗口**: 点击窗口任意位置拖拽移动
- **右键菜单**: 右键点击窗口显示更多选项

### 键盘快捷键

| 快捷键 | 功能 |
|--------|------|
| `Ctrl+M` | 最小化/恢复窗口 |
| `Ctrl+T` | 切换窗口置顶 |
| `Ctrl+O` | 切换透明度 |
| `Escape` | 重置窗口位置 |

### 右键菜单选项

- **最小化/恢复**: 控制窗口显示状态
- **窗口置顶/取消置顶**: 控制窗口层级
- **切换透明度**: 在不同透明度间切换
- **重置位置**: 恢复到默认窗口位置
- **退出**: 关闭应用程序

## 配置说明

### 配置文件位置

- **应用配置**: `~/.config/timeline/`
- **窗口位置**: `~/.config/timeline/window_position.json`
- **状态历史**: `~/.config/timeline/state.json`
- **使用历史**: `~/.config/timeline/history.json`

### 自定义配置

编辑 `config/settings.json` 可自定义：

```json
{
    "window": {
        "width": 300,
        "height": 120,
        "opacity": 0.9,
        "always_on_top": true
    },
    "time_periods": [
        {
            "name": "上午学习",
            "start": "08:30",
            "end": "11:30"
        }
    ],
    "colors": {
        "study_mode": "#4CAF50",
        "rest_mode": "#FF9800",
        "progress_bar": "#2196F3"
    }
}
```

## 项目结构

```
TimeLine/
├── src/                    # 源代码目录
│   ├── main.py            # 主应用程序
│   ├── time_manager.py    # 时间管理模块
│   ├── state_manager.py   # 状态管理模块
│   ├── window_manager.py  # 窗口管理模块
│   └── ui_components.py   # UI组件模块
├── config/                # 配置文件目录
│   └── settings.json      # 应用程序配置
├── assets/                # 资源文件目录
│   └── icons/             # 图标资源
├── timeline.py            # 启动脚本
├── timeline.desktop       # 桌面文件
├── install.sh             # 安装脚本
├── uninstall.sh           # 卸载脚本
├── requirements.txt       # Python依赖
└── README.md              # 说明文档
```

## 卸载

运行卸载脚本：

```bash
./uninstall.sh
```

卸载脚本会：
- 删除应用程序文件
- 删除启动脚本和桌面文件
- 询问是否删除用户配置和数据

## 故障排除

### 常见问题

1. **应用程序无法启动**
   - 检查Python3是否已安装: `python3 --version`
   - 检查GTK依赖: `python3 -c "import gi"`
   - 重新安装依赖: `sudo apt install python3-gi python3-gi-cairo gir1.2-gtk-3.0`

2. **窗口显示异常**
   - 重置窗口位置: 按 `Escape` 键
   - 删除位置配置: `rm ~/.config/timeline/window_position.json`

3. **透明度不生效**
   - 确保使用X11窗口系统
   - 检查桌面环境是否支持窗口透明度

### 日志和调试

应用程序会在终端输出调试信息，如遇问题可通过命令行启动查看：

```bash
python3 timeline.py
```

## 技术实现

- **GUI框架**: GTK3 + PyGObject
- **编程语言**: Python 3.8+
- **设计模式**: 模块化架构，分离关注点
- **配置管理**: JSON格式配置文件
- **状态持久化**: 本地文件存储

## 贡献

欢迎提交问题报告和功能建议！

## 许可证

本项目采用 MIT 许可证。

## 更新日志

### v1.0.0 (2025-07-16)
- 初始版本发布
- 实现基本时间管理功能
- 支持状态切换和窗口交互
- 完整的安装和配置系统
