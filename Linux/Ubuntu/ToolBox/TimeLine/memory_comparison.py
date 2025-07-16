#!/usr/bin/env python3
"""
内存使用对比测试
"""
import psutil
import subprocess
import time

def test_version(script_name, version_name):
    """测试指定版本的内存使用"""
    print(f"\n=== 测试 {version_name} ===")
    
    # 启动应用程序
    process = subprocess.Popen(['python3', script_name], 
                              stdout=subprocess.PIPE, 
                              stderr=subprocess.PIPE)
    
    # 等待启动
    time.sleep(3)
    
    try:
        proc = psutil.Process(process.pid)
        memory_info = proc.memory_info()
        
        rss_mb = memory_info.rss / 1024 / 1024
        vms_mb = memory_info.vms / 1024 / 1024
        
        print(f"RSS内存: {rss_mb:.1f} MB")
        print(f"虚拟内存: {vms_mb:.1f} MB")
        print(f"内存占用率: {proc.memory_percent():.1f}%")
        
        # 终止进程
        process.terminate()
        process.wait()
        
        return rss_mb, vms_mb
        
    except Exception as e:
        print(f"测试失败: {e}")
        process.terminate()
        return 0, 0

def main():
    print("TimeLine 内存使用对比测试")
    print("=" * 40)
    
    # 测试原版本
    original_rss, original_vms = test_version('timeline.py', '优化版本')
    
    # 测试轻量级版本
    lite_rss, lite_vms = test_version('timeline_lite.py', '轻量级版本')
    
    # 对比结果
    print(f"\n=== 对比结果 ===")
    print(f"优化版本:     RSS {original_rss:.1f} MB, VMS {original_vms:.1f} MB")
    print(f"轻量级版本:   RSS {lite_rss:.1f} MB, VMS {lite_vms:.1f} MB")
    
    if original_rss > 0 and lite_rss > 0:
        rss_reduction = ((original_rss - lite_rss) / original_rss) * 100
        vms_reduction = ((original_vms - lite_vms) / original_vms) * 100
        
        print(f"\n内存优化效果:")
        print(f"RSS内存减少: {rss_reduction:.1f}%")
        print(f"虚拟内存减少: {vms_reduction:.1f}%")
        print(f"绝对减少: {original_rss - lite_rss:.1f} MB")

if __name__ == "__main__":
    main()
