#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
TimeLine 轻量级版本
专注于最小内存占用的桌面计时器
"""

import gi
gi.require_version('Gtk', '3.0')

from gi.repository import Gtk, GLib
from datetime import datetime
import gc

class LightweightTimeLine:
    """轻量级TimeLine应用"""
    
    def __init__(self):
        # 最小化配置
        self.time_periods = [
            ("上午学习", (8, 30), (11, 30)),
            ("下午学习", (14, 30), (17, 30)),
            ("晚上学习", (19, 30), (22, 30))
        ]
        self.current_state = "study"
        self.setup_window()
        self.setup_ui()
        
        # 启动定时器（每30秒更新一次）
        GLib.timeout_add_seconds(30, self.update_display)
        self.update_display()
    
    def setup_window(self):
        """设置窗口"""
        self.window = Gtk.Window()
        self.window.set_title("📚 学习")
        self.window.set_default_size(240, 80)
        self.window.set_decorated(True)
        self.window.set_resizable(False)
        self.window.set_skip_taskbar_hint(True)
        self.window.set_keep_above(True)
        self.window.set_opacity(0.85)
        self.window.connect("destroy", Gtk.main_quit)
        
        # 简单的拖拽支持
        self.window.add_events(gi.repository.Gdk.EventMask.BUTTON_PRESS_MASK)
        self.window.connect("button-press-event", self.on_button_press)
    
    def setup_ui(self):
        """设置UI"""
        vbox = Gtk.Box(orientation=Gtk.Orientation.VERTICAL, spacing=4)
        vbox.set_margin_left(8)
        vbox.set_margin_right(8)
        vbox.set_margin_top(6)
        vbox.set_margin_bottom(6)
        
        # 时间显示
        self.time_label = Gtk.Label()
        self.time_label.set_markup("<span size='large' weight='bold'>--:--</span>")
        vbox.pack_start(self.time_label, False, False, 0)
        
        # 状态信息
        self.status_label = Gtk.Label()
        self.status_label.set_markup("<span size='small'>状态信息</span>")
        vbox.pack_start(self.status_label, False, False, 0)
        
        # 状态按钮
        self.state_button = Gtk.Button()
        self.state_button.set_label("📚 学习")
        self.state_button.connect("clicked", self.toggle_state)
        vbox.pack_start(self.state_button, False, False, 0)
        
        self.window.add(vbox)
    
    def get_current_period(self):
        """获取当前时间段"""
        now = datetime.now()
        current_hour = now.hour
        current_minute = now.minute
        current_total_minutes = current_hour * 60 + current_minute
        
        for name, (start_h, start_m), (end_h, end_m) in self.time_periods:
            start_minutes = start_h * 60 + start_m
            end_minutes = end_h * 60 + end_m
            
            if start_minutes <= current_total_minutes <= end_minutes:
                progress = (current_total_minutes - start_minutes) / (end_minutes - start_minutes)
                remaining_minutes = end_minutes - current_total_minutes
                return {
                    'name': name,
                    'progress': progress,
                    'remaining_hours': remaining_minutes // 60,
                    'remaining_minutes': remaining_minutes % 60
                }
        return None
    
    def update_display(self):
        """更新显示"""
        now = datetime.now()
        time_str = now.strftime("%H:%M")
        self.time_label.set_markup(f"<span size='large' weight='bold'>{time_str}</span>")
        
        current_period = self.get_current_period()
        if current_period:
            status_text = f"{current_period['name']} | 剩余 {current_period['remaining_hours']:02d}:{current_period['remaining_minutes']:02d}"
        else:
            status_text = "休息中"
        
        self.status_label.set_markup(f"<span size='small'>{status_text}</span>")
        
        # 强制垃圾回收
        gc.collect()
        return True
    
    def toggle_state(self, button):
        """切换状态"""
        if self.current_state == "study":
            self.current_state = "rest"
            self.state_button.set_label("☕ 休息")
            self.window.set_title("☕ 休息")
        else:
            self.current_state = "study"
            self.state_button.set_label("📚 学习")
            self.window.set_title("📚 学习")
    
    def on_button_press(self, widget, event):
        """处理拖拽"""
        if event.button == 1:
            widget.begin_move_drag(event.button, event.x_root, event.y_root, event.time)
        return True
    
    def run(self):
        """运行应用"""
        self.window.show_all()
        Gtk.main()

def main():
    """主函数"""
    app = LightweightTimeLine()
    app.run()

if __name__ == "__main__":
    main()
