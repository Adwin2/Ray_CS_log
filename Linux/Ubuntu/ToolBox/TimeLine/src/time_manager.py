#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
时间管理模块
负责处理学习时间段的识别、进度计算等时间相关逻辑
"""

from datetime import datetime, time, timedelta
from typing import Dict, List, Optional, Tuple

class TimeManager:
    """时间管理器类"""
    
    def __init__(self, time_periods: List[Dict]):
        """
        初始化时间管理器
        
        Args:
            time_periods: 时间段配置列表
        """
        self.time_periods = time_periods
        self._parse_time_periods()
    
    def _parse_time_periods(self):
        """解析时间段配置"""
        self.parsed_periods = []
        for period in self.time_periods:
            start_time = datetime.strptime(period["start"], "%H:%M").time()
            end_time = datetime.strptime(period["end"], "%H:%M").time()
            self.parsed_periods.append({
                "name": period["name"],
                "start": start_time,
                "end": end_time,
                "duration_minutes": self._calculate_duration(start_time, end_time)
            })
    
    def _calculate_duration(self, start_time: time, end_time: time) -> int:
        """计算时间段持续时间（分钟）"""
        start_minutes = start_time.hour * 60 + start_time.minute
        end_minutes = end_time.hour * 60 + end_time.minute
        return end_minutes - start_minutes
    
    def get_current_period(self, current_time: Optional[time] = None) -> Optional[Dict]:
        """
        获取当前时间段
        
        Args:
            current_time: 当前时间，如果为None则使用系统当前时间
            
        Returns:
            当前时间段信息，如果不在任何学习时间段内则返回None
        """
        if current_time is None:
            current_time = datetime.now().time()
        
        for period in self.parsed_periods:
            if period["start"] <= current_time <= period["end"]:
                return period
        return None
    
    def calculate_progress(self, current_time: Optional[time] = None, 
                          period: Optional[Dict] = None) -> float:
        """
        计算当前时间段的进度
        
        Args:
            current_time: 当前时间
            period: 时间段信息，如果为None则自动获取当前时间段
            
        Returns:
            进度百分比 (0.0 - 1.0)
        """
        if current_time is None:
            current_time = datetime.now().time()
        
        if period is None:
            period = self.get_current_period(current_time)
        
        if period is None:
            return 0.0
        
        current_minutes = current_time.hour * 60 + current_time.minute
        start_minutes = period["start"].hour * 60 + period["start"].minute
        end_minutes = period["end"].hour * 60 + period["end"].minute
        
        if current_minutes < start_minutes:
            return 0.0
        elif current_minutes > end_minutes:
            return 1.0
        else:
            total_duration = end_minutes - start_minutes
            elapsed = current_minutes - start_minutes
            return elapsed / total_duration if total_duration > 0 else 0.0
    
    def get_next_period(self, current_time: Optional[time] = None) -> Optional[Dict]:
        """
        获取下一个学习时间段
        
        Args:
            current_time: 当前时间
            
        Returns:
            下一个时间段信息，如果今天没有更多时间段则返回None
        """
        if current_time is None:
            current_time = datetime.now().time()
        
        current_minutes = current_time.hour * 60 + current_time.minute
        
        for period in self.parsed_periods:
            start_minutes = period["start"].hour * 60 + period["start"].minute
            if start_minutes > current_minutes:
                return period
        return None
    
    def get_remaining_time(self, current_time: Optional[time] = None,
                          period: Optional[Dict] = None) -> Tuple[int, int]:
        """
        获取当前时间段剩余时间
        
        Args:
            current_time: 当前时间
            period: 时间段信息
            
        Returns:
            (剩余小时, 剩余分钟)
        """
        if current_time is None:
            current_time = datetime.now().time()
        
        if period is None:
            period = self.get_current_period(current_time)
        
        if period is None:
            return (0, 0)
        
        current_minutes = current_time.hour * 60 + current_time.minute
        end_minutes = period["end"].hour * 60 + period["end"].minute
        
        remaining_minutes = max(0, end_minutes - current_minutes)
        hours = remaining_minutes // 60
        minutes = remaining_minutes % 60
        
        return (hours, minutes)
    
    def get_elapsed_time(self, current_time: Optional[time] = None,
                        period: Optional[Dict] = None) -> Tuple[int, int]:
        """
        获取当前时间段已过时间
        
        Args:
            current_time: 当前时间
            period: 时间段信息
            
        Returns:
            (已过小时, 已过分钟)
        """
        if current_time is None:
            current_time = datetime.now().time()
        
        if period is None:
            period = self.get_current_period(current_time)
        
        if period is None:
            return (0, 0)
        
        current_minutes = current_time.hour * 60 + current_time.minute
        start_minutes = period["start"].hour * 60 + period["start"].minute
        
        elapsed_minutes = max(0, current_minutes - start_minutes)
        hours = elapsed_minutes // 60
        minutes = elapsed_minutes % 60
        
        return (hours, minutes)
    
    def is_study_time(self, current_time: Optional[time] = None) -> bool:
        """
        判断当前是否为学习时间
        
        Args:
            current_time: 当前时间
            
        Returns:
            是否为学习时间
        """
        return self.get_current_period(current_time) is not None
    
    def get_daily_summary(self, current_time: Optional[time] = None) -> Dict:
        """
        获取今日学习时间总结
        
        Args:
            current_time: 当前时间
            
        Returns:
            包含总学习时间、已完成时间段等信息的字典
        """
        if current_time is None:
            current_time = datetime.now().time()
        
        current_minutes = current_time.hour * 60 + current_time.minute
        total_study_minutes = 0
        completed_periods = 0
        current_period = None
        
        for period in self.parsed_periods:
            total_study_minutes += period["duration_minutes"]
            end_minutes = period["end"].hour * 60 + period["end"].minute
            
            if current_minutes > end_minutes:
                completed_periods += 1
            elif period["start"] <= current_time <= period["end"]:
                current_period = period
        
        return {
            "total_study_minutes": total_study_minutes,
            "total_study_hours": total_study_minutes / 60,
            "completed_periods": completed_periods,
            "total_periods": len(self.parsed_periods),
            "current_period": current_period,
            "completion_rate": completed_periods / len(self.parsed_periods) if self.parsed_periods else 0
        }
