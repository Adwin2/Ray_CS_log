#!/usr/bin/env python3
"""
内存使用测试脚本
"""
import psutil
import subprocess
import time
import os

def get_memory_usage(pid):
    """获取进程内存使用情况"""
    try:
        process = psutil.Process(pid)
        memory_info = process.memory_info()
        return {
            'rss': memory_info.rss / 1024 / 1024,  # MB
            'vms': memory_info.vms / 1024 / 1024,  # MB
            'percent': process.memory_percent()
        }
    except psutil.NoSuchProcess:
        return None

def main():
    print("启动TimeLine应用程序...")
    
    # 启动轻量级版本
    process = subprocess.Popen(['python3', 'timeline_lite.py'],
                              stdout=subprocess.PIPE,
                              stderr=subprocess.PIPE)
    
    # 等待应用程序启动
    time.sleep(3)
    
    print(f"应用程序PID: {process.pid}")
    
    # 监控内存使用
    for i in range(10):
        memory = get_memory_usage(process.pid)
        if memory:
            print(f"第{i+1}次测量:")
            print(f"  RSS内存: {memory['rss']:.1f} MB")
            print(f"  虚拟内存: {memory['vms']:.1f} MB")
            print(f"  内存占用率: {memory['percent']:.1f}%")
            print()
        time.sleep(2)
    
    # 终止应用程序
    process.terminate()
    process.wait()
    print("测试完成")

if __name__ == "__main__":
    main()
