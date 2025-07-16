#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
UI组件模块
提供符合GNOME设计语言的自定义UI组件
"""

import gi
gi.require_version('Gtk', '3.0')
gi.require_version('Gdk', '3.0')

from gi.repository import Gtk, Gdk, Pango
from typing import Dict, Optional

class StyledProgressBar(Gtk.ProgressBar):
    """自定义样式的进度条"""
    
    def __init__(self, color: str = "#2196F3"):
        super().__init__()
        self.color = color
        self.setup_style()
    
    def setup_style(self):
        """设置进度条样式"""
        self.set_show_text(True)
        
        # 创建CSS样式
        css_provider = Gtk.CssProvider()
        css_data = f"""
        progressbar {{
            min-height: 20px;
            border-radius: 10px;
            background-color: rgba(0, 0, 0, 0.1);
        }}
        progressbar progress {{
            background-color: {self.color};
            border-radius: 10px;
            min-height: 20px;
        }}
        """
        css_provider.load_from_data(css_data.encode())
        
        context = self.get_style_context()
        context.add_provider(css_provider, Gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)

class StateButton(Gtk.Button):
    """状态切换按钮"""

    def __init__(self, study_color: str = "#4CAF50", rest_color: str = "#FF9800"):
        super().__init__()
        self.study_color = study_color
        self.rest_color = rest_color
        self.current_state = "study"
        self.css_provider = None  # 缓存CSS提供者
        self.setup_style()
        self.update_appearance()
    
    def setup_style(self):
        """设置按钮基础样式"""
        self.set_size_request(-1, 28)

        # 设置字体
        font_desc = Pango.FontDescription()
        font_desc.set_weight(Pango.Weight.NORMAL)
        font_desc.set_size(Pango.SCALE * 10)
        self.override_font(font_desc)
    
    def set_state(self, state: str):
        """设置按钮状态"""
        if state in ["study", "rest"]:
            self.current_state = state
            self.update_appearance()
    
    def update_appearance(self):
        """更新按钮外观"""
        if self.current_state == "study":
            self.set_label("📚 学习")
            self._apply_color_style(self.study_color)
        else:
            self.set_label("☕ 休息")
            self._apply_color_style(self.rest_color)
    
    def _apply_color_style(self, color: str):
        """应用颜色样式"""
        # 重用CSS提供者，减少内存分配
        if self.css_provider is None:
            self.css_provider = Gtk.CssProvider()
            context = self.get_style_context()
            context.add_provider(self.css_provider, Gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)

        css_data = f"""
        button {{
            background-color: {color};
            color: white;
            border-radius: 6px;
            border: none;
            padding: 4px 12px;
            font-weight: normal;
            font-size: 11px;
        }}
        button:hover {{
            background-color: {self._darken_color(color)};
        }}
        button:active {{
            background-color: {self._darken_color(color, 0.3)};
        }}
        """
        self.css_provider.load_from_data(css_data.encode())
    
    def _darken_color(self, hex_color: str, factor: float = 0.2) -> str:
        """使颜色变暗"""
        # 简单的颜色变暗算法
        hex_color = hex_color.lstrip('#')
        rgb = tuple(int(hex_color[i:i+2], 16) for i in (0, 2, 4))
        darkened = tuple(int(c * (1 - factor)) for c in rgb)
        return f"#{darkened[0]:02x}{darkened[1]:02x}{darkened[2]:02x}"

class TimeDisplay(Gtk.Label):
    """时间显示组件"""
    
    def __init__(self, font_size: str = "x-large"):
        super().__init__()
        self.font_size = font_size
        self.setup_style()
    
    def setup_style(self):
        """设置时间显示样式"""
        # 设置字体
        font_desc = Pango.FontDescription()
        font_desc.set_family("monospace")
        font_desc.set_weight(Pango.Weight.BOLD)
        font_desc.set_size(Pango.SCALE * 14)
        self.override_font(font_desc)

        # 设置对齐
        self.set_halign(Gtk.Align.CENTER)
        self.set_valign(Gtk.Align.CENTER)
    
    def set_time(self, time_str: str):
        """设置时间文本"""
        self.set_markup(f"<span size='{self.font_size}' weight='bold' font_family='monospace' color='#2c3e50'>{time_str}</span>")

class InfoLabel(Gtk.Label):
    """信息标签组件"""
    
    def __init__(self, text: str = "", secondary: bool = False):
        super().__init__()
        self.secondary = secondary
        self.setup_style()
        if text:
            self.set_text(text)
    
    def setup_style(self):
        """设置标签样式"""
        self.set_halign(Gtk.Align.CENTER)
        self.set_valign(Gtk.Align.CENTER)

        if self.secondary:
            # 次要信息样式
            self.set_markup(f"<span color='#666666' size='x-small'>{self.get_text()}</span>")
        else:
            # 主要信息样式
            font_desc = Pango.FontDescription()
            font_desc.set_weight(Pango.Weight.NORMAL)
            font_desc.set_size(Pango.SCALE * 10)
            self.override_font(font_desc)
    
    def set_info_text(self, text: str):
        """设置信息文本"""
        if self.secondary:
            self.set_markup(f"<span color='#666666' size='x-small'>{text}</span>")
        else:
            self.set_text(text)

class MainContainer(Gtk.Box):
    """主容器组件"""
    
    def __init__(self, spacing: int = 10):
        super().__init__(orientation=Gtk.Orientation.VERTICAL, spacing=spacing)
        self.setup_style()
    
    def setup_style(self):
        """设置容器样式"""
        self.set_margin_left(10)
        self.set_margin_right(10)
        self.set_margin_top(6)
        self.set_margin_bottom(6)

        # 设置半透明背景样式
        css_provider = Gtk.CssProvider()
        css_data = """
        box {
            background-color: rgba(240, 240, 240, 0.7);
            border-radius: 8px;
            border: 1px solid rgba(200, 200, 200, 0.3);
        }
        """
        css_provider.load_from_data(css_data.encode())

        context = self.get_style_context()
        context.add_provider(css_provider, Gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)

class CompactWindow(Gtk.Window):
    """紧凑型窗口"""
    
    def __init__(self, title: str = "TimeLine", config: Optional[Dict] = None):
        super().__init__()
        self.config = config or {}
        self.setup_window(title)
    
    def setup_window(self, title: str):
        """设置窗口属性"""
        self.set_title(title)
        
        # 从配置获取窗口大小
        width = self.config.get("window", {}).get("width", 320)
        height = self.config.get("window", {}).get("height", 140)
        self.set_default_size(width, height)
        
        # 窗口属性
        self.set_decorated(True)
        self.set_resizable(False)
        self.set_skip_taskbar_hint(True)
        self.set_skip_pager_hint(True)
        
        # 透明度
        opacity = self.config.get("window", {}).get("opacity", 0.95)
        self.set_opacity(opacity)
        
        # 置顶
        if self.config.get("window", {}).get("always_on_top", True):
            self.set_keep_above(True)
        
        # 设置窗口图标（如果有的话）
        try:
            self.set_icon_name("appointment-soon")
        except:
            pass  # 如果图标不存在就忽略
        
        # 设置窗口样式
        self.setup_window_style()
    
    def setup_window_style(self):
        """设置窗口样式"""
        css_provider = Gtk.CssProvider()
        css_data = """
        window {
            background-color: rgba(240, 240, 240, 0.1);
            border-radius: 8px;
        }
        """
        css_provider.load_from_data(css_data.encode())

        context = self.get_style_context()
        context.add_provider(css_provider, Gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)

        # 应用到整个应用程序
        Gtk.StyleContext.add_provider_for_screen(
            Gdk.Screen.get_default(),
            css_provider,
            Gtk.STYLE_PROVIDER_PRIORITY_APPLICATION
        )
