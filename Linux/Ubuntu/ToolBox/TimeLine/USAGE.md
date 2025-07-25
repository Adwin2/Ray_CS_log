# TimeLine 使用指南

## 快速开始

### 1. 运行应用程序

```bash
# 直接运行
python3 timeline.py

# 或者安装后运行
./install.sh
timeline
```

### 2. 界面说明

应用程序启动后会显示一个半透明的小窗口，包含以下元素：

- **时间显示**: 当前系统时间 (HH:MM:SS)
- **时间段信息**: 显示当前是否在学习时间段内
- **进度条**: 显示当前学习时间段的进度和剩余时间
- **状态按钮**: 手动切换学习/休息模式

### 3. 学习时间段

应用程序预设了三个学习时间段：

- **上午学习**: 08:30 - 11:30 (3小时)
- **下午学习**: 14:30 - 17:30 (3小时)  
- **晚上学习**: 19:30 - 22:30 (3小时)

### 4. 状态模式

- **📚 学习模式**: 绿色按钮，专注学习时间
- **☕ 休息模式**: 橙色按钮，放松休息时间

## 操作说明

### 基本操作

1. **移动窗口**: 点击窗口任意位置拖拽
2. **切换状态**: 点击状态按钮在学习/休息模式间切换
3. **查看进度**: 进度条显示当前时间段完成百分比和剩余时间

### 快捷键

| 快捷键 | 功能 |
|--------|------|
| `Ctrl+M` | 最小化/恢复窗口 |
| `Ctrl+T` | 切换窗口置顶 |
| `Ctrl+O` | 切换透明度 |
| `Escape` | 重置窗口位置 |

### 右键菜单

右键点击窗口可以访问以下功能：

- **最小化/恢复**: 控制窗口显示状态
- **窗口置顶/取消置顶**: 控制窗口层级
- **切换透明度**: 在不同透明度间切换
- **重置位置**: 恢复到默认窗口位置
- **退出**: 关闭应用程序

## 进度显示说明

### 在学习时间段内

- 进度条显示当前时间段的完成百分比
- 文本显示格式: `XX.X% | 剩余 HH:MM`
- 时间段信息显示: `当前时间段: [时间段名称]`

### 在休息时间

- 进度条为空
- 显示 "休息时间"
- 如果还有后续学习时间段，会显示下一时间段开始时间

## 配置说明

### 配置文件位置

- **应用配置**: `config/settings.json`
- **用户配置**: `~/.config/timeline/`
- **窗口位置**: `~/.config/timeline/window_position.json`
- **状态历史**: `~/.config/timeline/state.json`

### 自定义时间段

编辑 `config/settings.json` 文件：

```json
{
    "time_periods": [
        {
            "name": "自定义时间段",
            "start": "09:00",
            "end": "12:00"
        }
    ]
}
```

### 自定义颜色

```json
{
    "colors": {
        "study_mode": "#4CAF50",    // 学习模式颜色
        "rest_mode": "#FF9800",     // 休息模式颜色
        "progress_bar": "#2196F3"   // 进度条颜色
    }
}
```

## 常见问题

### Q: 窗口位置不对怎么办？
A: 按 `Escape` 键重置窗口位置，或者删除 `~/.config/timeline/window_position.json` 文件。

### Q: 如何修改学习时间段？
A: 编辑 `config/settings.json` 文件中的 `time_periods` 部分。

### Q: 窗口太大/太小怎么调整？
A: 修改 `config/settings.json` 中的 `window.width` 和 `window.height` 值。

### Q: 如何让窗口更透明？
A: 修改 `config/settings.json` 中的 `window.opacity` 值（0.0-1.0）。

### Q: 应用程序无法启动？
A: 检查是否安装了必要的依赖：
```bash
sudo apt install python3-gi python3-gi-cairo gir1.2-gtk-3.0
```

## 技术特性

- **轻量级**: 占用内存小，CPU使用率低
- **半透明**: 不干扰其他应用程序的使用
- **自动保存**: 窗口位置和状态自动保存
- **GNOME集成**: 完美融入GNOME桌面环境
- **实时更新**: 每秒更新时间和进度信息

## 卸载

如需卸载应用程序：

```bash
./uninstall.sh
```

卸载脚本会询问是否删除用户配置和数据。
