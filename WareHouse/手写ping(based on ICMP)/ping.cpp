//手写ping程序，基于ICMP协议
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <netinet/ip.h>
#include <netinet/ip_icmp.h>
#include <arpa/inet.h>
#include <netdb.h>
#include <unistd.h>
#include <sys/time.h>
#include <signal.h>
#include <errno.h>

const int PACKET_SIZE = 4096;

const int MAX_PACKET_NUM = 4;

const int MAX_DELAY_TIME = 5;

//要发送的包
char sendPacket[PACKET_SIZE] = {0};

//要接受的包
char recvPacket[PACKET_SIZE] = {0};

int nsend = 0;
int nrecv = 0;

struct sockaddr_in fromAddr = {0};
int fromlen = sizeof(fromlen); 

//计算校验和
u_short cal_cksum(u_short* icmp, int packsize);

//打包
int pack(int nsend);

//发包
void send_Packet(int sfd, struct sockaddr_in* pDestAddr, int len);

//收包
void recv_Packet(int sfd);

//解包
//int unpack(int len, int uid);

//定时器信号处理函数
void signalrmHand(int signo);

//计算收发过程中的时间差
void tv_sub(struct timeval* out, struct timeval* in);

int main(int argc, char* argv[]) {

    if(argc < 2) {
        printf("请输入对方IP\n");
        return -1;
    }

    //准备SOCKET  创建 + 设置 SOCKET 设置接收机的IP地址
    int sockfd;
    struct sockaddr_in sock_addr;
    //创建socket
    struct protoent* protocal = getprotobyname("icmp");
    if( protocal == NULL )  return -1;
    
    int sfd = socket(AF_INET,SOCK_RAW,protocal->p_proto);
    if(sfd == -1) printf("创建socket失败：%m\n"),exit(-1);
    printf("创建socket成功！\n");
    //设置socket
    int size = 1024*50;
    int r = setsockopt(sfd, SOL_SOCKET,SO_RCVBUF, &size,sizeof(size));
    if(r == -1) printf("设置SOCKET失败：%m\n"),close(sfd),exit(-1);
    printf("设置IP地址成功！\n");
    //接收机的IP地址设置
    struct sockaddr_in destADDR = {0};
    destADDR.sin_family = AF_INET;
    struct hostent* host = NULL; 
    in_addr_t addr = inet_addr(argv[1]);   //将IP地址转换为数字
    if (addr == INADDR_NONE) //不是IP地址
    {
        host = gethostbyname(argv[1]);//将域名转换为IP地址
        if(host == NULL) {
            printf("错误IP地址或无法解析域名：%s\n",argv[1]);
            close(sfd);
            return -1;
        }
        memcpy(&(destADDR.sin_addr),host->h_addr,host->h_length);
        //printf("IP地址解析成功：%s -> %s\n",argv[1],inet_ntoa(destADDR.sin_addr));
    }else
    {
        destADDR.sin_addr.s_addr = addr;
    }
    //打包
    pack(nsend);
    //发包
    send_Packet(sfd,&destADDR,sizeof(destADDR));
    //收包
    recv_Packet(sfd);
    //解包
    //unpack(nrecv, getuid());
    //输出统计信息
    printf("发送了%d个包，接收了%d个包，%%%f丢失！\n",nsend,nrecv,
    (nsend - nrecv)*1.0/nsend*100);
    //关闭socket
    close(sfd);
    return 0;
}

void send_Packet(int sfd, struct sockaddr_in* pDestAddr, int len) {
    int packSize;
    int r;
    while( nsend < MAX_PACKET_NUM) {
        nsend ++;
        packSize = pack(nsend);

        r = sendto(sfd, sendPacket, packSize, 0,
            (struct sockaddr*)pDestAddr, len);
        if(r < 0) {
            printf("第%d个包发送失败：%m\n",nsend);
            continue;
        }
        printf("第%d个包发送成功！\n",nsend);
        sleep(1);
    }
}

//打包
int pack(int nsend) {
    struct icmp* pIcmp =  (struct icmp*) sendPacket;
    pIcmp->icmp_type = ICMP_ECHO;
    pIcmp->icmp_code = 0;
    pIcmp->icmp_cksum = 0;
    pIcmp->icmp_id = getuid();
    pIcmp->icmp_seq = nsend;

    int packSize = 8 + 56;//留作区分

    struct timeval* tval = (struct timeval*)pIcmp->icmp_data;

    gettimeofday(tval, NULL);//设置发送时间

    pIcmp->icmp_cksum = cal_cksum((u_short*)pIcmp, packSize);//校验和计算
    return packSize;
}

//计算校验和
u_short cal_cksum(u_short* icmp, int packsize) {
    int nleft = packsize;
    int sum = 0;
    u_short* w = icmp;
    u_short answer = 0;


    while(nleft > 1) {
        sum += *w++;
        nleft -= 2;
    }

    if(nleft == 1) {
        *(unsigned char*)(&answer) = *(unsigned char*)w;
        sum += answer;
    }

    sum = (sum>>16) + (sum&0xffff);
    sum += (sum>>16);
    answer = ~sum;

    return answer; 
}

//收包
void recv_Packet(int sfd) {
    signal(SIGALRM, signalrmHand);
    int r;
    while(nrecv < nsend) {
        //设置定时器
        alarm(MAX_DELAY_TIME);

        r = recvfrom(sfd,recvPacket,PACKET_SIZE,0,(struct sockaddr*)&fromAddr, (socklen_t*)&fromlen);
        if(r < 0) {
            if(errno == EINTR) continue;
            printf("接收失败：%m\n");
            continue;
        }
    }   printf("recvPacket:%s\n",recvPacket);
}

//解包
// int unpack(int len, int uid) {
    
// }


//计算收发过程中的时间差


//定时器信号处理函数
void signalrmHand(int s) {
    if(s == SIGALRM) {
        printf("----------------------------myPing-------------------------");
        printf("发送了%d个包，接收了%d个包，%%%f丢失！\n",nsend,nrecv,
        (nsend - nrecv)*1.0/nsend*100);
        printf("----------------------------myPing-------------------------");
        exit(-1);
    }
}
