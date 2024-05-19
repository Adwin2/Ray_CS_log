

总结 完全跟着做的  没有解包功能

ICMP校验和是干什么用的：

1. 校验和是为了检测数据是否被篡改或被破坏过。（类似伞列（哈希）算法）

2. 校验和的计算是通过将数据分成两部分，一部分是校验和本身，一部分是数据本身，然后将两部分相加，再取反，得到校验和。

# 拓展

- 数据传输类型

TCP 协议 SOCK_STREAM 网络数据流

UDP 协议 SOCK_DGRAM 网络报文

ICMP 协议 SOCK_RAW



- 网络描述符 socketfd     接口

- windows linux 标准库 exit（）都是返回-1表示失败 0表示成功但没有什么特殊的返回值  正整数表示成功并返回有特殊意义的值

- IP地址   点分十进制  数点格式的字符串

IP地址在网络服务器上 属于  unsigned int 类型



sendto == connect + send     |     send --- 网络中的  write操作 写入文件 写入sfd

















