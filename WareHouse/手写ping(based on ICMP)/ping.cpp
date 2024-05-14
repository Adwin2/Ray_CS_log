//手写ping程序，基于ICMP协议
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <netdb.h>
#include <unistd.h>
#include <time.h>
#include <sys/time.h>
#include <signal.h>
#include <errno.h>

#define MAX_PACKET_SIZE 65535
#define MAX_TTL 64
#define MAX_WAIT_TIME 1000
#define MAX_TRIES 3

int main(int argc, char argv[]) {
    //准备SOCKET  创建 + 设置 SOCKET 设置接收机的IP地址
    int sockfd;
    struct sockaddr_in dest_addr;
    //打包
    
    //发包
    
    //收包

    //解包

    //输出统计信息





    return 0;
}