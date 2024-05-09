#include "utility.h"

int main(int argc, int argv[]) {
    //服务器IP + port
    struct sockaddr_in serverAddr;
    serverAddr.sin_family = PF_INET;
    serverAddr.sin_port = htons(SERVER_PORT);
    serverAddr.sin_addr.s_addr = inet_addr(SERVER_IP);
    //创建监听socket
    int listener = socket(PF_INET, SOCK_STREAM, 0);
    if(listener < 0) {
        perror("listener error");
        exit(-1);
    }
    printf("listener socket created");
    //绑定地址
    if(bind(listener, (struct sockaddr *)&serverAddr, sizeof(serverAddr)) < 0) {
        perror("bind error");
        exit(-1);
    }
    //监听
    int ret = listen(listener, 5);
    if(ret < 0) {
        perror("listen error");
        exit(-1);
    }
    printf("start to listen: %s\n", SERVER_IP);
    //在内核中创建事件表
    int epfd = epoll_create(EPOLL_SIZE);
    if(epfd < 0) {
        perror("epfd error");
        exit(-1);
    }
    printf("epoll created, epollfd = %d\n",epfd);
    static struct epoll_event events[EPOLL_SIZE];
    //往内核事件表里添加事件
    addfd(epfd, listener, true);
    //主循环
    while(1) {
        //epoll_events_count表示就绪事件的数目
        int epoll_events_count = epoll_wait(epfd, events, EPOLL_SIZE, -1);
        if(epoll_events_count < 0) {
            perror("epoll failure");
            break;
        }
    printf("epoll_events_count = %d\n",epoll_events_count);
        //处理这epoll_events_count个就绪事件
        for(int i = 0; i < epoll_events_count; ++i) {
            int sockfd = events[i].data.fd;
            //新用户连接
            if(sockfd == listener) {
                struct sockaddr_in client_address;
                socklen_t client_addrLength = sizeof(struct sockaddr_in);
                int clientfd = accept(listener, (struct sockaddr*)&client_address, &client_addrLength);
                printf("client connection from: %s : %d(IP:port),clientfd = %d\n", inet_ntoa(client_address.sin_addr),ntohs(client_address.sin_port),clientfd);
                
                //将该新的客户端添加到内核事件列表
                addfd(epfd, clientfd, true);

                //服务端用list保存用户连接
                clients_list.push_back(clientfd);
                printf("Add new clientfd = %d to epoll\n", clientfd);
                printf("Now there are %d clients in the chat room\n",(int)clients_list.size());

                //服务端发送欢迎消息
                printf("Welcome!\n");
                char message[BUF_SIZE];
                bzero(message, BUF_SIZE);
                sprintf(message, SERVER_WELCOME, clientfd);
                int ret = send(clientfd, message, BUF_SIZE, 0);
                if(ret < 0) {
                    perror("send error");
                    exit(-1);
                }
            }
            /*
            唤醒客户端，
            处理用户发来的消息并广播
            使其他用户收到消息
            */
            else {
                int ret = sendBroadcastmessage(sockfd);
                if(ret < 0) {
                    perror("wall error");
                    exit(-1);
                }
            }
        }
    }
    //关闭socket
    close(listener);
    //关闭内核，停止监控注册事件的发生
    close(epfd);
    return 0;
}