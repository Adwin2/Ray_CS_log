#ifndef UTILITY_H_INCLUDED
#define UTILITY_H_INCLUDED

#include<iostream>
#include<list>
#include<sys/types.h>
#include<sys/socket.h>
#include<netinet/in.h>
#include<arpa/inet.h>
#include<sys/epoll.h>
#include<fcntl.h>
#include<errno.h>
#include<unistd.h>
#include<stdio.h>
#include<stdlib.h>
#include<string.h>

using namespace std;

//save all the clents`socket  已链表形式储存 准备好的
list<int> clients_list;


//server ip 
#define SERVER_IP "127.0.0.1"

//server port 
#define SERVER_PORT 8888

//epoll size 
#define EPOLL_SIZE 5000

//信息缓冲区大小
#define BUF_SIZE 0xFFFF

#define SERVER_WELCOME "Welcome"

#define SERVER_MESSAGE "ClientID %d say >> %s"

//exit
#define EXIT "EXIT"

#define CAUTION "There is only one in the chat room"


//设置非阻塞模式
int setnonblocking(int sockfd) {
    fcntl(sockfd,F_SETFL, fcntl(sockfd, F_GETFD, 0)|O_NONBLOCK);
    return 0;
}


void addfd(int epollfd, int fd, bool enable_et) {
    struct epoll_event ev; //<epoll.h>
    ev.data.fd = fd;
    ev.events = EPOLLIN;
    if( enable_et) 
        ev.events = EPOLLIN | EPOLLET;
    epoll_ctl(epollfd, EPOLL_CTL_ADD, fd, &ev);
    setnonblocking(fd);
    printf("fd added to epoll\n \n");
}



int sendBroadcastmessage(int clientfd) {
    //buf 接收消息  ，message保存消息
    char buf[BUF_SIZE], message[BUF_SIZE];
    bzero(buf, BUF_SIZE);
    bzero(message, BUF_SIZE);


    printf("read from client(clientID = %d)\n",clientfd);
    int len = recv(clientfd,buf,BUF_SIZE,0);

    //len=0-->客户端已关闭连接
    if(len <= 0) {
        close(clientfd);
        clients_list.remove(clientfd);
        printf("ClientID = %d closed.\n now there are %d client in the chat room\n",clientfd,(int)clients_list.size());
    }
    else {
        //chatroom只有一个连接 
        if(clients_list.size() == 1) {
            send(clientfd, CAUTION, strlen(CAUTION),0);
            return len;
        }
        //广播消息格式
        sprintf(message, SERVER_MESSAGE, clientfd, buf);
        //改用了C++11 循环方式
        for(auto it:clients_list) {
            if(send(it, message, BUF_SIZE, 0) < 0) {
                perror("error");
                exit(-1);
            }
        }
    }
    return len;
}
#endif //UTILITY_H_INCLUDED