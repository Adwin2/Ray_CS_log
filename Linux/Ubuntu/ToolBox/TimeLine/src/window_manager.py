#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
窗口管理模块
负责处理窗口的拖拽、定位、最小化等交互功能
"""

import gi
gi.require_version('Gtk', '3.0')
gi.require_version('Gdk', '3.0')

from gi.repository import Gtk, Gdk, GLib
import json
from pathlib import Path
from typing import Dict, Optional, Tuple

class WindowManager:
    """窗口管理器类"""
    
    def __init__(self, window: Gtk.Window, config: Optional[Dict] = None):
        """
        初始化窗口管理器
        
        Args:
            window: GTK窗口对象
            config: 配置字典
        """
        self.window = window
        self.config = config or {}
        
        # 拖拽相关变量
        self.is_dragging = False
        self.drag_start_x = 0
        self.drag_start_y = 0
        self.window_start_x = 0
        self.window_start_y = 0
        
        # 窗口状态
        self.is_minimized = False
        self.saved_position = None
        self.saved_size = None
        
        # 配置文件路径
        self.config_dir = Path.home() / ".config" / "timeline"
        self.position_file = self.config_dir / "window_position.json"
        
        # 确保配置目录存在
        self.config_dir.mkdir(parents=True, exist_ok=True)
        
        self.setup_window_events()
        self.load_window_position()
    
    def setup_window_events(self):
        """设置窗口事件处理"""
        # 鼠标事件
        self.window.add_events(
            Gdk.EventMask.BUTTON_PRESS_MASK |
            Gdk.EventMask.BUTTON_RELEASE_MASK |
            Gdk.EventMask.POINTER_MOTION_MASK |
            Gdk.EventMask.ENTER_NOTIFY_MASK |
            Gdk.EventMask.LEAVE_NOTIFY_MASK
        )
        
        # 连接事件处理器
        self.window.connect("button-press-event", self.on_button_press)
        self.window.connect("button-release-event", self.on_button_release)
        self.window.connect("motion-notify-event", self.on_motion_notify)
        self.window.connect("configure-event", self.on_configure)
        self.window.connect("window-state-event", self.on_window_state)
        self.window.connect("delete-event", self.on_delete_event)
        
        # 键盘快捷键
        self.window.connect("key-press-event", self.on_key_press)
    
    def on_button_press(self, widget, event):
        """鼠标按下事件"""
        if event.button == 1:  # 左键
            self.is_dragging = True
            self.drag_start_x = event.x_root
            self.drag_start_y = event.y_root
            
            # 获取窗口当前位置
            self.window_start_x, self.window_start_y = self.window.get_position()
            
            # 设置鼠标光标
            cursor = Gdk.Cursor.new(Gdk.CursorType.FLEUR)
            event.window.set_cursor(cursor)
            
            return True
        elif event.button == 3:  # 右键
            self.show_context_menu(event)
            return True
        
        return False
    
    def on_button_release(self, widget, event):
        """鼠标释放事件"""
        if event.button == 1 and self.is_dragging:
            self.is_dragging = False
            
            # 恢复默认光标
            event.window.set_cursor(None)
            
            # 保存新位置
            self.save_window_position()
            
            return True
        
        return False
    
    def on_motion_notify(self, widget, event):
        """鼠标移动事件"""
        if self.is_dragging:
            # 计算新位置
            new_x = self.window_start_x + (event.x_root - self.drag_start_x)
            new_y = self.window_start_y + (event.y_root - self.drag_start_y)
            
            # 确保窗口不会移出屏幕
            new_x, new_y = self.constrain_to_screen(new_x, new_y)
            
            # 移动窗口
            self.window.move(int(new_x), int(new_y))
            
            return True
        
        return False
    
    def on_configure(self, widget, event):
        """窗口配置变化事件"""
        # 延迟保存位置，避免频繁写入
        GLib.timeout_add(500, self.save_window_position)
        return False
    
    def on_window_state(self, widget, event):
        """窗口状态变化事件"""
        if event.new_window_state & Gdk.WindowState.ICONIFIED:
            self.is_minimized = True
        else:
            self.is_minimized = False
        
        return False
    
    def on_delete_event(self, widget, event):
        """窗口关闭事件"""
        self.save_window_position()
        return False
    
    def on_key_press(self, widget, event):
        """键盘按键事件"""
        # Ctrl+M: 最小化/恢复
        if event.state & Gdk.ModifierType.CONTROL_MASK and event.keyval == Gdk.KEY_m:
            self.toggle_minimize()
            return True
        
        # Ctrl+T: 切换置顶
        if event.state & Gdk.ModifierType.CONTROL_MASK and event.keyval == Gdk.KEY_t:
            self.toggle_always_on_top()
            return True
        
        # Ctrl+O: 切换透明度
        if event.state & Gdk.ModifierType.CONTROL_MASK and event.keyval == Gdk.KEY_o:
            self.toggle_opacity()
            return True
        
        # Escape: 重置位置
        if event.keyval == Gdk.KEY_Escape:
            self.reset_position()
            return True
        
        return False
    
    def constrain_to_screen(self, x: float, y: float) -> Tuple[int, int]:
        """将窗口位置限制在屏幕范围内"""
        screen = self.window.get_screen()
        monitor = screen.get_monitor_at_window(self.window.get_window())
        geometry = screen.get_monitor_geometry(monitor)
        
        window_width, window_height = self.window.get_size()
        
        # 确保窗口不会完全移出屏幕
        min_x = geometry.x - window_width + 50  # 至少保留50像素可见
        max_x = geometry.x + geometry.width - 50
        min_y = geometry.y
        max_y = geometry.y + geometry.height - 50
        
        x = max(min_x, min(max_x, x))
        y = max(min_y, min(max_y, y))
        
        return int(x), int(y)
    
    def show_context_menu(self, event):
        """显示右键菜单"""
        menu = Gtk.Menu()
        
        # 最小化/恢复
        minimize_item = Gtk.MenuItem.new_with_label("最小化" if not self.is_minimized else "恢复")
        minimize_item.connect("activate", lambda x: self.toggle_minimize())
        menu.append(minimize_item)
        
        # 分隔符
        menu.append(Gtk.SeparatorMenuItem())
        
        # 置顶切换
        try:
            always_on_top = self.window.get_keep_above()
        except AttributeError:
            always_on_top = True  # 默认值
        top_item = Gtk.MenuItem.new_with_label("取消置顶" if always_on_top else "窗口置顶")
        top_item.connect("activate", lambda x: self.toggle_always_on_top())
        menu.append(top_item)
        
        # 透明度切换
        opacity_item = Gtk.MenuItem.new_with_label("切换透明度")
        opacity_item.connect("activate", lambda x: self.toggle_opacity())
        menu.append(opacity_item)
        
        # 分隔符
        menu.append(Gtk.SeparatorMenuItem())
        
        # 重置位置
        reset_item = Gtk.MenuItem.new_with_label("重置位置")
        reset_item.connect("activate", lambda x: self.reset_position())
        menu.append(reset_item)
        
        # 分隔符
        menu.append(Gtk.SeparatorMenuItem())
        
        # 退出
        quit_item = Gtk.MenuItem.new_with_label("退出")
        quit_item.connect("activate", lambda x: Gtk.main_quit())
        menu.append(quit_item)
        
        menu.show_all()
        menu.popup(None, None, None, None, event.button, event.time)
    
    def toggle_minimize(self):
        """切换最小化状态"""
        if self.is_minimized:
            self.window.deiconify()
        else:
            self.window.iconify()
    
    def toggle_always_on_top(self):
        """切换窗口置顶状态"""
        try:
            current_state = self.window.get_keep_above()
            self.window.set_keep_above(not current_state)
        except AttributeError:
            # 如果方法不存在，尝试其他方式
            print("窗口置顶功能在当前GTK版本中不可用")
    
    def toggle_opacity(self):
        """切换透明度"""
        current_opacity = self.window.get_opacity()
        if current_opacity > 0.9:
            self.window.set_opacity(0.7)
        elif current_opacity > 0.7:
            self.window.set_opacity(0.5)
        else:
            self.window.set_opacity(0.95)
    
    def reset_position(self):
        """重置窗口位置到默认位置"""
        default_x = self.config.get("window", {}).get("position", {}).get("x", 100)
        default_y = self.config.get("window", {}).get("position", {}).get("y", 50)
        
        self.window.move(default_x, default_y)
        self.save_window_position()
    
    def save_window_position(self):
        """保存窗口位置到配置文件"""
        if self.window.get_window() is None:
            return False  # 窗口还未实现
        
        x, y = self.window.get_position()
        width, height = self.window.get_size()
        opacity = self.window.get_opacity()
        # GTK3兼容性：使用正确的方法名
        try:
            always_on_top = self.window.get_keep_above()
        except AttributeError:
            # 如果方法不存在，使用默认值
            always_on_top = self.config.get("window", {}).get("always_on_top", True)
        
        position_data = {
            "x": x,
            "y": y,
            "width": width,
            "height": height,
            "opacity": opacity,
            "always_on_top": always_on_top
        }
        
        try:
            with open(self.position_file, 'w', encoding='utf-8') as f:
                json.dump(position_data, f, indent=2)
        except Exception as e:
            print(f"保存窗口位置失败: {e}")
        
        return False  # 不重复调用
    
    def load_window_position(self):
        """从配置文件加载窗口位置"""
        if not self.position_file.exists():
            return
        
        try:
            with open(self.position_file, 'r', encoding='utf-8') as f:
                position_data = json.load(f)
            
            # 应用保存的位置和设置
            x = position_data.get("x", 100)
            y = position_data.get("y", 50)
            
            # 确保位置在屏幕范围内
            x, y = self.constrain_to_screen(x, y)
            
            # 延迟应用位置，确保窗口已经实现
            GLib.idle_add(self._apply_saved_position, x, y, position_data)
            
        except Exception as e:
            print(f"加载窗口位置失败: {e}")
    
    def _apply_saved_position(self, x: int, y: int, position_data: Dict):
        """应用保存的窗口位置和设置"""
        self.window.move(x, y)
        
        # 应用其他设置
        if "opacity" in position_data:
            self.window.set_opacity(position_data["opacity"])
        
        if "always_on_top" in position_data:
            self.window.set_keep_above(position_data["always_on_top"])
        
        return False  # 只执行一次
