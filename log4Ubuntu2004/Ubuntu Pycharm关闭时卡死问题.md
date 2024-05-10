# Ubuntu Pycharm关闭时卡死问题

## （一）使用kill杀死进程

执行下面的命令：

```ps -e | grep java```

输出

```bash
4046 ?        00:06:04 java
```

最前面就是java的pid号，也就是 4046

执行命令：

```bash
kill -s 9 4046  # 将4046替换为自己的java pid号
```

执行以后就可以退出pycharm了

## （二）使用系统监视器

在工具面板中搜索 system monitor 即可找到系统监视器

如果打不开系统监视器，需要执行下面两行命令：

```
snap remove gnome-system-monitor  # 使用snap移除掉软件
sudo apt-get install gnome-system-monitor  # apt安装
```

安装以后就可以打开系统监视器了

打开以后，在最上方的三个选项卡中，选择进程

使用Ctrl + f 打开搜索，输入java

就可以找到java的进程，右键点击java，选择杀死