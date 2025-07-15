# TimeLine 开发者文档

本文档面向希望了解、修改或扩展 TimeLine 桌面计时器的开发者。

## 架构设计

### 模块化架构

TimeLine 采用模块化设计，将不同功能分离到独立模块中：

```
src/
├── main.py            # 主应用程序入口和协调器
├── time_manager.py    # 时间段管理和进度计算
├── state_manager.py   # 学习/休息状态管理
├── window_manager.py  # 窗口交互和位置管理
└── ui_components.py   # 自定义UI组件
```

### 设计原则

1. **单一职责**: 每个模块负责特定功能
2. **松耦合**: 模块间通过明确接口通信
3. **可扩展**: 易于添加新功能和修改现有功能
4. **配置驱动**: 行为通过配置文件控制

## 核心模块详解

### 1. TimeManager (时间管理器)

**职责**: 处理学习时间段的识别、进度计算等时间相关逻辑

**主要方法**:
- `get_current_period()`: 获取当前时间段
- `calculate_progress()`: 计算进度百分比
- `get_remaining_time()`: 获取剩余时间
- `get_daily_summary()`: 获取每日总结

**使用示例**:
```python
time_manager = TimeManager(config["time_periods"])
current_period = time_manager.get_current_period()
progress = time_manager.calculate_progress()
```

### 2. StateManager (状态管理器)

**职责**: 管理学习/休息状态，提供状态持久化和历史记录

**主要方法**:
- `set_state()`: 设置新状态
- `toggle_state()`: 切换状态
- `get_state_info()`: 获取状态信息
- `add_state_change_callback()`: 添加状态变化回调

**状态持久化**:
- 状态保存到 `~/.config/timeline/state.json`
- 历史记录保存到 `~/.config/timeline/history.json`

### 3. WindowManager (窗口管理器)

**职责**: 处理窗口拖拽、定位、快捷键等交互功能

**主要功能**:
- 窗口拖拽移动
- 右键上下文菜单
- 键盘快捷键处理
- 窗口位置持久化

**事件处理**:
```python
window_manager = WindowManager(window, config)
# 自动设置所有事件处理器
```

### 4. UI Components (UI组件)

**职责**: 提供符合GNOME设计语言的自定义UI组件

**组件列表**:
- `CompactWindow`: 紧凑型主窗口
- `TimeDisplay`: 时间显示组件
- `StyledProgressBar`: 自定义进度条
- `StateButton`: 状态切换按钮
- `InfoLabel`: 信息标签

## 开发环境设置

### 1. 环境要求

```bash
# Ubuntu 20.04.6 LTS
sudo apt update
sudo apt install python3-dev python3-gi-dev python3-gi-cairo gir1.2-gtk-3.0

# 开发工具（可选）
sudo apt install python3-pip git
pip3 install pylint black
```

### 2. 项目设置

```bash
# 克隆项目
git clone <repository-url>
cd TimeLine

# 直接运行（开发模式）
python3 timeline.py

# 或者从src目录运行
cd src
python3 main.py
```

### 3. 代码规范

- **Python版本**: 3.8+
- **代码风格**: PEP 8
- **文档字符串**: Google风格
- **类型提示**: 推荐使用

**示例代码**:
```python
def calculate_progress(self, current_time: Optional[time] = None, 
                      period: Optional[Dict] = None) -> float:
    """
    计算当前时间段的进度
    
    Args:
        current_time: 当前时间
        period: 时间段信息
        
    Returns:
        进度百分比 (0.0 - 1.0)
    """
    # 实现代码...
```

## 扩展开发

### 添加新的时间段类型

1. 修改 `config/settings.json`:
```json
{
    "time_periods": [
        {
            "name": "新时间段",
            "start": "13:00",
            "end": "14:00",
            "type": "break"  // 新增类型字段
        }
    ]
}
```

2. 更新 `TimeManager` 类以支持新类型

### 添加新的UI主题

1. 在 `ui_components.py` 中添加新的颜色方案:
```python
def get_theme_colors(theme_name: str) -> Dict:
    themes = {
        "dark": {"study": "#2E7D32", "rest": "#F57C00"},
        "light": {"study": "#4CAF50", "rest": "#FF9800"}
    }
    return themes.get(theme_name, themes["light"])
```

2. 更新配置文件支持主题选择

### 添加通知功能

1. 安装通知依赖:
```bash
sudo apt install libnotify-dev
pip3 install plyer
```

2. 在 `StateManager` 中添加通知回调:
```python
def send_notification(self, message: str):
    """发送桌面通知"""
    from plyer import notification
    notification.notify(
        title="TimeLine",
        message=message,
        timeout=5
    )
```

## 调试和测试

### 调试模式

设置环境变量启用调试输出:
```bash
export TIMELINE_DEBUG=1
python3 timeline.py
```

### 单元测试

创建测试文件 `tests/test_time_manager.py`:
```python
import unittest
from datetime import time
from src.time_manager import TimeManager

class TestTimeManager(unittest.TestCase):
    def setUp(self):
        self.config = [
            {"name": "测试", "start": "09:00", "end": "10:00"}
        ]
        self.manager = TimeManager(self.config)
    
    def test_current_period(self):
        test_time = time(9, 30)
        period = self.manager.get_current_period(test_time)
        self.assertIsNotNone(period)
        self.assertEqual(period["name"], "测试")

if __name__ == "__main__":
    unittest.main()
```

### 性能分析

使用 `cProfile` 分析性能:
```bash
python3 -m cProfile -o profile.stats timeline.py
python3 -c "import pstats; pstats.Stats('profile.stats').sort_stats('cumulative').print_stats(10)"
```

## 打包和分发

### 创建可执行文件

使用 PyInstaller:
```bash
pip3 install pyinstaller
pyinstaller --onefile --windowed timeline.py
```

### 创建 DEB 包

1. 安装打包工具:
```bash
sudo apt install debhelper dh-python
```

2. 创建 `debian/` 目录结构
3. 使用 `dpkg-buildpackage` 构建包

### AppImage 打包

使用 python-appimage:
```bash
pip3 install python-appimage
python-appimage build timeline.py
```

## 贡献指南

### 提交代码

1. Fork 项目
2. 创建功能分支: `git checkout -b feature/new-feature`
3. 提交更改: `git commit -am 'Add new feature'`
4. 推送分支: `git push origin feature/new-feature`
5. 创建 Pull Request

### 代码审查

- 确保代码符合PEP 8规范
- 添加适当的文档字符串
- 包含必要的测试用例
- 更新相关文档

### 问题报告

使用 GitHub Issues 报告问题，请包含:
- 系统信息 (Ubuntu版本、Python版本)
- 错误信息和堆栈跟踪
- 重现步骤
- 预期行为

## 常见开发问题

### GTK 相关问题

1. **导入错误**: 确保安装了 `python3-gi`
2. **显示问题**: 检查 X11 环境变量
3. **主题问题**: 确保GTK主题兼容

### 窗口管理问题

1. **拖拽不工作**: 检查事件掩码设置
2. **位置保存失败**: 检查配置目录权限
3. **透明度问题**: 确认窗口管理器支持

### 性能优化

1. **减少定时器频率**: 根据需要调整更新间隔
2. **优化UI更新**: 只在必要时更新界面
3. **内存管理**: 及时清理不需要的对象

## 未来发展方向

### 计划功能

- [ ] 多语言支持
- [ ] 自定义主题系统
- [ ] 统计和报告功能
- [ ] 云同步支持
- [ ] 插件系统

### 技术改进

- [ ] 迁移到GTK4
- [ ] 添加Wayland支持
- [ ] 改进配置管理
- [ ] 增强测试覆盖率

---

如有开发相关问题，请查阅源代码注释或提交Issue。
