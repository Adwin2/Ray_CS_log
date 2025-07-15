#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
TimeLine Desktop Timer Application
桌面计时器应用程序启动脚本

使用方法:
    python3 timeline.py
    或
    ./timeline.py

作者: 用户定制
版本: 1.0.0
适用系统: Ubuntu 20.04.6 LTS (GNOME 3.36.8)
"""

import sys
import os
from pathlib import Path

# 添加src目录到Python路径
src_path = Path(__file__).parent / "src"
sys.path.insert(0, str(src_path))

try:
    from src.main import main
    
    if __name__ == "__main__":
        main()
        
except ImportError as e:
    print(f"导入错误: {e}")
    print("请确保已安装所需的依赖包:")
    print("  sudo apt install python3-gi python3-gi-cairo gir1.2-gtk-3.0")
    print("  或")
    print("  pip3 install -r requirements.txt")
    sys.exit(1)
    
except Exception as e:
    print(f"应用程序启动失败: {e}")
    sys.exit(1)
