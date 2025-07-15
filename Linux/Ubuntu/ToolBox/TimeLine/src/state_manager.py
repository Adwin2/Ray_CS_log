#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
状态管理模块
负责管理应用程序的学习/休息状态，提供状态持久化和通知功能
"""

import json
import os
from datetime import datetime, timedelta
from pathlib import Path
from typing import Dict, List, Optional, Callable

class StateManager:
    """状态管理器类"""
    
    def __init__(self, config_dir: Optional[str] = None):
        """
        初始化状态管理器
        
        Args:
            config_dir: 配置目录路径，如果为None则使用默认路径
        """
        self.config_dir = Path(config_dir) if config_dir else Path.home() / ".config" / "timeline"
        self.state_file = self.config_dir / "state.json"
        self.history_file = self.config_dir / "history.json"
        
        # 确保配置目录存在
        self.config_dir.mkdir(parents=True, exist_ok=True)
        
        # 当前状态
        self.current_state = "study"  # study 或 rest
        self.state_start_time = datetime.now()
        self.manual_override = False  # 是否手动覆盖了自动状态
        
        # 状态变化回调函数列表
        self.state_change_callbacks: List[Callable[[str, str], None]] = []
        
        # 加载保存的状态
        self.load_state()
    
    def add_state_change_callback(self, callback: Callable[[str, str], None]):
        """
        添加状态变化回调函数
        
        Args:
            callback: 回调函数，接收 (old_state, new_state) 参数
        """
        self.state_change_callbacks.append(callback)
    
    def remove_state_change_callback(self, callback: Callable[[str, str], None]):
        """移除状态变化回调函数"""
        if callback in self.state_change_callbacks:
            self.state_change_callbacks.remove(callback)
    
    def set_state(self, new_state: str, manual: bool = True):
        """
        设置新状态
        
        Args:
            new_state: 新状态 ("study" 或 "rest")
            manual: 是否为手动设置
        """
        if new_state not in ["study", "rest"]:
            raise ValueError("状态必须是 'study' 或 'rest'")
        
        old_state = self.current_state
        
        if old_state != new_state:
            # 记录状态变化历史
            self._record_state_change(old_state, new_state, manual)
            
            # 更新状态
            self.current_state = new_state
            self.state_start_time = datetime.now()
            self.manual_override = manual
            
            # 保存状态
            self.save_state()
            
            # 触发回调
            for callback in self.state_change_callbacks:
                try:
                    callback(old_state, new_state)
                except Exception as e:
                    print(f"状态变化回调执行失败: {e}")
    
    def toggle_state(self):
        """切换状态"""
        new_state = "rest" if self.current_state == "study" else "study"
        self.set_state(new_state, manual=True)
    
    def get_current_state(self) -> str:
        """获取当前状态"""
        return self.current_state
    
    def get_state_duration(self) -> timedelta:
        """获取当前状态持续时间"""
        return datetime.now() - self.state_start_time
    
    def is_manual_override(self) -> bool:
        """检查当前状态是否为手动覆盖"""
        return self.manual_override
    
    def clear_manual_override(self):
        """清除手动覆盖标志"""
        self.manual_override = False
        self.save_state()
    
    def get_state_info(self) -> Dict:
        """获取完整的状态信息"""
        duration = self.get_state_duration()
        return {
            "state": self.current_state,
            "start_time": self.state_start_time.isoformat(),
            "duration_seconds": int(duration.total_seconds()),
            "duration_minutes": int(duration.total_seconds() / 60),
            "manual_override": self.manual_override
        }
    
    def save_state(self):
        """保存当前状态到文件"""
        state_data = {
            "current_state": self.current_state,
            "state_start_time": self.state_start_time.isoformat(),
            "manual_override": self.manual_override,
            "last_updated": datetime.now().isoformat()
        }
        
        try:
            with open(self.state_file, 'w', encoding='utf-8') as f:
                json.dump(state_data, f, indent=2, ensure_ascii=False)
        except Exception as e:
            print(f"保存状态失败: {e}")
    
    def load_state(self):
        """从文件加载状态"""
        if not self.state_file.exists():
            return
        
        try:
            with open(self.state_file, 'r', encoding='utf-8') as f:
                state_data = json.load(f)
            
            self.current_state = state_data.get("current_state", "study")
            self.manual_override = state_data.get("manual_override", False)
            
            # 解析开始时间
            start_time_str = state_data.get("state_start_time")
            if start_time_str:
                self.state_start_time = datetime.fromisoformat(start_time_str)
            else:
                self.state_start_time = datetime.now()
                
        except Exception as e:
            print(f"加载状态失败: {e}")
            # 使用默认状态
            self.current_state = "study"
            self.state_start_time = datetime.now()
            self.manual_override = False
    
    def _record_state_change(self, old_state: str, new_state: str, manual: bool):
        """记录状态变化历史"""
        history_entry = {
            "timestamp": datetime.now().isoformat(),
            "old_state": old_state,
            "new_state": new_state,
            "manual": manual,
            "duration_seconds": int(self.get_state_duration().total_seconds())
        }
        
        # 加载现有历史
        history = self._load_history()
        history.append(history_entry)
        
        # 只保留最近100条记录
        if len(history) > 100:
            history = history[-100:]
        
        # 保存历史
        try:
            with open(self.history_file, 'w', encoding='utf-8') as f:
                json.dump(history, f, indent=2, ensure_ascii=False)
        except Exception as e:
            print(f"保存历史记录失败: {e}")
    
    def _load_history(self) -> List[Dict]:
        """加载状态变化历史"""
        if not self.history_file.exists():
            return []
        
        try:
            with open(self.history_file, 'r', encoding='utf-8') as f:
                return json.load(f)
        except Exception as e:
            print(f"加载历史记录失败: {e}")
            return []
    
    def get_today_summary(self) -> Dict:
        """获取今日状态总结"""
        history = self._load_history()
        today = datetime.now().date()
        
        study_time = 0
        rest_time = 0
        state_changes = 0
        
        for entry in history:
            entry_date = datetime.fromisoformat(entry["timestamp"]).date()
            if entry_date == today:
                state_changes += 1
                if entry["old_state"] == "study":
                    study_time += entry["duration_seconds"]
                elif entry["old_state"] == "rest":
                    rest_time += entry["duration_seconds"]
        
        # 添加当前状态的时间
        current_duration = int(self.get_state_duration().total_seconds())
        if self.current_state == "study":
            study_time += current_duration
        else:
            rest_time += current_duration
        
        return {
            "date": today.isoformat(),
            "study_time_seconds": study_time,
            "study_time_minutes": study_time // 60,
            "rest_time_seconds": rest_time,
            "rest_time_minutes": rest_time // 60,
            "state_changes": state_changes,
            "current_state": self.current_state,
            "current_duration_minutes": current_duration // 60
        }
    
    def get_visual_state_config(self) -> Dict:
        """获取当前状态的视觉配置"""
        if self.current_state == "study":
            return {
                "color": "#4CAF50",
                "icon": "📚",
                "label": "学习模式",
                "description": "专注学习时间"
            }
        else:
            return {
                "color": "#FF9800", 
                "icon": "☕",
                "label": "休息模式",
                "description": "放松休息时间"
            }
