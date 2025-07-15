#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
TimeLine Desktop Timer Application
桌面计时器应用程序主入口

作者: 用户定制
版本: 1.0.0
适用系统: Ubuntu 20.04.6 LTS (GNOME 3.36.8)
"""

import gi
gi.require_version('Gtk', '3.0')
gi.require_version('Gdk', '3.0')

from gi.repository import Gtk, Gdk, GLib
import json
import os
import sys
from datetime import datetime, time
from pathlib import Path

from time_manager import TimeManager
from state_manager import StateManager
from window_manager import WindowManager
from ui_components import (
    CompactWindow, MainContainer, TimeDisplay,
    InfoLabel, StyledProgressBar, StateButton
)

class TimeLineApp:
    """桌面计时器主应用程序类"""

    def __init__(self):
        """初始化应用程序"""
        self.load_config()
        self.time_manager = TimeManager(self.config["time_periods"])
        self.state_manager = StateManager()

        # 添加状态变化回调
        self.state_manager.add_state_change_callback(self.on_state_changed)

        self.setup_window()
        self.setup_ui()
        self.setup_window_manager()
        self.setup_timer()
        
    def load_config(self):
        """加载配置文件"""
        config_path = Path(__file__).parent.parent / "config" / "settings.json"
        try:
            with open(config_path, 'r', encoding='utf-8') as f:
                self.config = json.load(f)
        except FileNotFoundError:
            # 默认配置
            self.config = {
                "window": {"width": 300, "height": 120, "opacity": 0.9, "always_on_top": True},
                "time_periods": [
                    {"name": "上午学习", "start": "08:30", "end": "11:30"},
                    {"name": "下午学习", "start": "14:30", "end": "17:30"},
                    {"name": "晚上学习", "start": "19:30", "end": "22:30"}
                ],
                "colors": {"study_mode": "#4CAF50", "rest_mode": "#FF9800", "progress_bar": "#2196F3"}
            }
    
    def setup_window(self):
        """设置主窗口"""
        # 初始状态显示
        initial_state = self.state_manager.get_visual_state_config()
        self.window = CompactWindow(f"{initial_state['icon']} {initial_state['label']}", self.config)

        # 连接关闭事件
        self.window.connect("destroy", Gtk.main_quit)

        # 设置窗口位置
        if "position" in self.config["window"]:
            self.window.move(
                self.config["window"]["position"]["x"],
                self.config["window"]["position"]["y"]
            )
    
    def setup_ui(self):
        """设置用户界面"""
        # 主容器
        main_container = MainContainer(spacing=4)

        # 当前时间显示
        self.time_display = TimeDisplay()
        self.time_display.set_time("--:--:--")
        main_container.pack_start(self.time_display, False, False, 0)

        # 时间段信息
        self.period_label = InfoLabel("当前时间段: 无")
        main_container.pack_start(self.period_label, False, False, 0)

        # 进度条
        progress_color = self.config.get("colors", {}).get("progress_bar", "#2196F3")
        self.progress_bar = StyledProgressBar(progress_color)
        main_container.pack_start(self.progress_bar, False, False, 0)

        # 状态切换按钮
        study_color = self.config.get("colors", {}).get("study_mode", "#4CAF50")
        rest_color = self.config.get("colors", {}).get("rest_mode", "#FF9800")
        self.state_button = StateButton(study_color, rest_color)
        self.state_button.connect("clicked", self.toggle_state)
        main_container.pack_start(self.state_button, False, False, 0)

        self.window.add(main_container)

    def setup_window_manager(self):
        """设置窗口管理器"""
        self.window_manager = WindowManager(self.window, self.config)
    
    def setup_timer(self):
        """设置定时器"""
        # 每秒更新一次
        GLib.timeout_add_seconds(1, self.update_display)
        # 立即更新一次
        self.update_display()
    
    def update_display(self):
        """更新显示内容"""
        now = datetime.now()

        # 更新时间显示
        time_str = now.strftime("%H:%M:%S")
        self.time_display.set_time(time_str)

        # 使用时间管理器获取当前时间段
        current_period = self.time_manager.get_current_period(now.time())

        if current_period:
            # 在学习时间段内
            period_name = current_period["name"]
            progress = self.time_manager.calculate_progress(now.time(), current_period)

            # 获取剩余时间
            remaining_hours, remaining_minutes = self.time_manager.get_remaining_time(now.time(), current_period)

            self.period_label.set_info_text(f"{period_name}")
            self.progress_bar.set_fraction(progress)

            # 显示进度和剩余时间
            start_str = current_period["start"].strftime("%H:%M")
            end_str = current_period["end"].strftime("%H:%M")
            progress_text = f"{progress*100:.0f}% | 剩余 {remaining_hours:02d}:{remaining_minutes:02d}"
            self.progress_bar.set_text(progress_text)
        else:
            # 不在学习时间段内
            next_period = self.time_manager.get_next_period(now.time())
            if next_period:
                next_start = next_period["start"].strftime("%H:%M")
                self.period_label.set_info_text(f"休息中 | 下一时段 {next_start}")
            else:
                self.period_label.set_info_text("休息中 | 今日已结束")

            self.progress_bar.set_fraction(0)
            self.progress_bar.set_text("休息中")

        return True  # 继续定时器
    

    
    def toggle_state(self, button):
        """切换学习/休息状态"""
        self.state_manager.toggle_state()

    def on_state_changed(self, old_state: str, new_state: str):
        """状态变化回调"""
        self.state_button.set_state(new_state)

        # 更新窗口标题只显示状态
        visual_config = self.state_manager.get_visual_state_config()
        self.window.set_title(f"{visual_config['icon']} {visual_config['label']}")

        print(f"状态已切换: {old_state} -> {new_state}")
    
    def run(self):
        """运行应用程序"""
        self.window.show_all()
        Gtk.main()

def main():
    """主函数"""
    app = TimeLineApp()
    app.run()

if __name__ == "__main__":
    main()
