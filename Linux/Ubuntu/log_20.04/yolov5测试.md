安装Anaconda

下载安装包 安装 配置环境变量 source ~/.bashrc 输入conda list 以测试

添加Anaconda国内镜像配置

conda config --add channels https://mirrors.tuna.tsinghua.edu.cn/anaconda/pkgs/free

​      ··································          https://mirrors.tuna.tsinghua.edu.cn/anaconda/pkgs/main

conda config --set show_channel_urls yes



安装pytorch （目前最低1.7版本）

`conda create -n pytorch1.7 python=3.8`

安装成功后激活pytorch1.7环境

`source activate pytorch1.7`（激活使用source命令，conda无效）

在创建的环境下安装pytorch1.7

conda install pytorch torchvision cudatoolkit=10.2  -c python

[^10.2]: cuda版本号

编辑.bashrc文件，设置使用1.7环境下的python3.8

alias python='////(文件路径)'

yolo项目克隆与安装

git clone //（仓库地址）

（在pytorch环境中）安装所需库 在yolov5 目录中

`pip install -r requirements.txt -i https://pypi.tuna.tsinghua.edu.cn/simple`

下载预训练权重文件放在weights文件夹里

使用方法：

离线：`python detect.py --source ./inferenec/images/ --weights weights/yolov5s.pt --conf 0.4`

摄像头：`python detect.py --source 0 --conf 0.4`



来源：https://www.bilibili.com/video/BV1rt4y1W7Dc/?spm_id_from=333.1007.top_right_bar_window_default_collection.content.click
