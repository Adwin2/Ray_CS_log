# Pycharm import cv2报错解决--1.30

（一）报错原因：python版本下多了，导致python解释器定位错误

（一）解决方法：到设置 ->项目 ->Python解释器 中把路径调整到/usr/bin/python3

（二）这个时候发现cv2还是import失败

报错原因：cv2没有放在python3路径下

`sudo cp -r 原路径   /usr/bin/python3`

即可解决

附：[终端默认python版本更改教程](https://blog.csdn.net/qq_34994476/article/details/121932064)

