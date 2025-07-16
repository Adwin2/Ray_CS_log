#!/usr/bin/env python3
"""
TimeLine 轻量级启动脚本
最小内存占用版本
"""

import sys
import os

# 添加src目录到Python路径
sys.path.insert(0, os.path.join(os.path.dirname(__file__), "src"))

if __name__ == "__main__":
    try:
        from lightweight_main import main
        main()
    except ImportError as e:
        print(f"导入错误: {e}")
        print("请确保已安装GTK3依赖")
        sys.exit(1)
    except Exception as e:
        print(f"应用程序启动失败: {e}")
        sys.exit(1)
