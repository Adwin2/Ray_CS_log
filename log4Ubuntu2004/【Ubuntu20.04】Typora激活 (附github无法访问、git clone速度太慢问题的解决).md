# 【Ubuntu20.04】Typora激活 (附github无法访问、git clone速度太慢问题的解决)(请先阅读补充文档！)

## 绪（一）、前因 

前因是想在ubuntu激活typora，然后在github上找到一个[Node_inject的项目](https://github.com/DiamondHunters/NodeInject)，

还有一国内大佬对这个项目编译整合成了一片[完整的中文教程](https://github.com/hazukieq/Yporaject/blob/master/README.md)，

但是存在github无法访问和git clone速度太太太慢的问题，

所以从 github，git clone到 typora激活 一顿搞，作此**文档记录**一下。

## 绪（二）、梗概

**整合**了三部分教程，略微**简化**部分操作过程。

> 文章内容基本就是本人解决标题所述问题的步骤，亲测可行。
>
> 参考网页链接已附在文章末尾。



------

# >问题解决部分



# **一、首先解决github无法访问的问题**

### 1.进入`/etc/hosts`文件

终端输入

```
sudo vim /etc/hosts
```

### 2.如果存在与github.com相关的内容，先删除，然后文件末尾添加以下内容

```
140.82.114.25                 alive.github.com 
140.82.112.25                 live.github.com 
185.199.108.154               github.githubassets.com 
140.82.112.22                 central.github.com 
185.199.108.133               desktop.githubusercontent.com 
185.199.108.153               assets-cdn.github.com 
185.199.108.133               camo.githubusercontent.com 
185.199.108.133               github.map.fastly.net 
199.232.69.194                github.global.ssl.fastly.net 
140.82.112.4                  gist.github.com 
185.199.108.153               github.io 
140.82.114.4                  github.com 
192.0.66.2                    github.blog 
140.82.112.6                  api.github.com 
185.199.108.133               raw.githubusercontent.com 
185.199.108.133               user-images.githubusercontent.com 
185.199.108.133               favicons.githubusercontent.com 
185.199.108.133               avatars5.githubusercontent.com 
185.199.108.133               avatars4.githubusercontent.com 
185.199.108.133               avatars3.githubusercontent.com 
185.199.108.133               avatars2.githubusercontent.com 
185.199.108.133               avatars1.githubusercontent.com 
185.199.108.133               avatars0.githubusercontent.com 
185.199.108.133               avatars.githubusercontent.com 
140.82.112.10                 codeload.github.com 
52.217.223.17                 github-cloud.s3.amazonaws.com 
52.217.199.41                 github-com.s3.amazonaws.com 
52.217.93.164                 github-production-release-asset-2e65be.s3.amazonaws.com 
52.217.174.129                github-production-user-asset-6210df.s3.amazonaws.com 
52.217.129.153                github-production-repository-file-5c1aeb.s3.amazonaws.com 
185.199.108.153               githubstatus.com 
64.71.144.202                 github.community 
23.100.27.125                 github.dev 
185.199.108.133               media.githubusercontent.com
```

> 具体步骤：按i进入插入模式，移动光标到文档末尾并cv上述内容，
>
> 最后点击‘’esc ‘’+ ‘’:wq‘’保存退出

### 3.安装nscd

输入

```
sudo apt install nscd
```

> 注意显示内容，没有报错下一步

### 4.重启nscd服务

输入

```
sudo /etc/init.d/nscd restart
```

> **显示**`Restarting nscd (via systemctl): nscd.service.`即可

### 注意：这里亲测有可能会出现重新连接网络后，又打不开github的问题

`在终端再次重启nscd服务`（即再次在终端输入4中需要输入的内容）可以解决（亲测），

然后注意教程后访问也并不是百分百可以访问，有时候耐心刷新几下也行

> 原理大概是对指定域名配置指定的IP地址，可以在访问github及相关网页，解析时跳转使用指定的可访问的节点上。

# 二、解决git clone“过早EOF”的问题

由于下一步就是**克隆Yporaject项目**，需要用到`git clone `命令，这个时候会出现如图的报错

![](./QQ20240130-115611.png)

### 解决办法：增加缓冲区大小

格式`git config --global http.postBuffer <buffer_size>`

<buffer_size>部分是设置的缓冲区大小，设置一个较大的值，104857600（即100MB）

即终端输入

```
git config --global http.postBuffer 104857600
```

> 这个办法本人亲测有效
>
> 还有更多办法见文章底部相关教程

------

# >Typora安装并激活部分

# 三、*前往[Typora官网](https://typoraio.cn/releases/all)下载最新版本（目前1.8.9）（64位.deb类型）

终端输入

```
sudo dpkg -i  .deb文件路径
```

直接安装即可

# 四、[激活Typora部分](https://github.com/hazukieq/Yporaject/blob/master/README.md)

> （前面完成的话，这里可以直接看本标题链接的Github大佬原教程，一些原理解释可能会漏掉）

## 1.克隆Yporaject项目

终端输入

```
git clone https://github.com/hazukieq/Yporaject.git --depth=1
```

## 2.配置RUST编译环境

### （1）运行官方脚本安装

```
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
```

这里会有显示如下内容

![](./11.png)

### （2）接着按照这个内容指示终端输入

```
source "$HOME/.cargo/env"
```

### （3）检查cargo

终端输入

```
cargo --version
```

显示如图内容即可

<img src="./22.png"  />

## 3.编译Yporaject项目

```
# 进入 Yporaject 项目
cd Yporaject
# 运行编译命令
cargo build
# 查看二进制是否生成,程序名称为 node_inject！！！！很重要
ls target/debug
```
> 务必确认target/debug 下 是否生成了 node_inject 二进制程序

依次输入运行这四行命令

最终显示如下，即正常

![](./33.png)

## 4.复制二进制程序到安装目录下

```
# 复制二进制程序到Typora目录下
sudo cp target/debug/node_inject /usr/share/typora

# 进入相关目录
cd /usr/share/typora

# 给予二进制程序执行权限
sudo chmod +x node_inject

sudo ./node_inject
```

依次输入运行这四行命令

------

# >Congrats ! 到这里所有的相关配置都已结束



## 5.获取许可证激活码

回到Yporaject目录，终端输入

```
cd Yporaject/

```

打开license-gen 文件夹

```
cd license-gen
```

编译

```
cargo build
```

运行

```
cargo run
```

最后出现

![](./44.png)

圈住的即是生成的激活码

## 6.激活软件

打开Typora

点击“帮助”

点击“我的许可证”

![](./55.png)

然后就会看到输入邮箱和激活码的页面，邮箱随便填，激活码cv即可

耐心等待后出现如图界面

![](./QQ20240130-104225.png)

# 最后，就全部完成啦！！！Congratulation！



# 附：参考原教程

[Ubuntu解决Github无法访问的问题](https://cloud.tencent.com/developer/article/2144993)

[Yporaject项目官方教程 -- github](https://github.com/hazukieq/Yporaject/blob/master/README.md)

[NodeInject项目（不是教程，只是源码）--github](https://github.com/DiamondHunters/NodeInject_Hook_example)

[git clone early EOF解决办法](https://cloud.tencent.com/developer/article/2369133?areaId=106001)

[知乎上搬运的Github大佬教程](https://zhuanlan.zhihu.com/p/636193675)

