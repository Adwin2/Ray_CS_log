# [Linux系统编程](https://blog.csdn.net/qq_47944751/article/details/131568915 ) 

## 一、文件编程

### 1.open 和 fopen 的区别

#### 1.1 来源

· open 是UNIX系统调用函数  返回的是文件描述符（文件在文件描述表中的索引）。

· fopen 是ANSIC标准中的C语言库函数，在不同系统中调用不同的系统API。返回的是一个指向文件结构的指针。

>（注）C语言库函数需要调用系统API实现
>
>同样可以看出 fopen的移植性要好于 open

#### 1.2 适用范围 与 文件IO层次

· open 返回文件描述符 UUNIX下的一切都以文件的形式操作 包括从操作普通正规文件（Regular File）

· fopen 操作 Regular File

从文件IO层次看  fopen  属于用户态（高级文件io） 

 			    open 属于内核态（低级文件io）

#### 1.3 缓冲系统	

##### 缓冲文件系统 

> `fopen, fclose, fread, fwrite, fgetc, fgets, fputc, fputs, freopen, fseek, ftell, rewind`等。

· 特点 ： 内存开辟一个缓冲区 ， 为程序的每一个文件使用；执行读操作时 从磁盘文件将数据先读入内存缓冲区，装满后再从缓冲区依次读出。写操作同理。

> 由此可见，缓冲区的大小决定操作外存的次数。一般来说，缓冲区的大小由机器定。

--> 缓冲文件系统是借助==文件结构体指针==来对文件进行管理，通过文件指针来对文件进行访问，可以读写字符 字符串 格式化数据 二进制数据。

非缓冲系统 依赖于操作系统 对文件进行读写 是系统级的输入输出 ，只能读写二进制文件 但效率高 （ANSIC标准不再包括非缓冲文件系统，勿用。）

##### 非缓冲文件系统 

> `open, close, read, write, getc, getchar, putc, putchar`等

### 2.文件操作

#### 2.1打开  创建 写入 读取 

> 分为标准C库实现 和 open write等非缓冲文件系统类型操作实现

#### 2.2光标移动操作

2.2.1包含的头文件
 #include <sys/types.h>
 #include <unistd.h>

2.2.2函数原型
off_t lseek(int fd, off_t offset, int whence);

2.2.3函数参数说明
int fd ：文件描述符
off_t offset：偏移多少个字节
int whence：光标偏移位置

【整个函数的意思是：将文件读写指针相对whence移动offset个字节位置】

2.2.4lseek函数描述
给whence参数设定偏移位置：

SEEK_SET：光标偏移到头

SEEK_CUR：光标为当前位置

SEEK_END：光标偏移到末尾

2.2.5函数返回值
成功完成后，Iseek()返回从文件开始的偏移位置(以字节为单位)【就是返回偏移了多少个字节】。发生错误时，返回值(off_t) -1，并设置errno来指示错误。

> 注：write操作后光标处于文件末尾 ==注意光标位置==

3.Linux文件操作步骤

1、在Linux中要操作一个文件，一般是先open打开一个文件，得到文件描述符，然后对文件进行读写操作(或其他操作)，最后是close关闭文件即可。

2、强调一点：我们对文件进行操作时，一定要先打开文件，只有**打开成功后才能操作**，如果打开失败，就不用进行后面的操作了，最后读写完成后，一定要**关闭文件**，否则会造成文件损坏。

3、文件平时是存放在块设备中的文件系统文件中的，我们把这种文件叫**静态文件**，当我们去open打开一个文件时，Linux内核做的操作包括：内核在进程中建立一个打开的文件的数据结构，记录下我们打开的这个文件；内核在内存中申请一段内存，并且将静态文件的内容从块设备中读取到内核中特定的地址管理存放(叫动态文件)。

4、打开文件后，以后对这个文件的读写操作，都是针对**内存中的这一份动态文件**的，而不是针对静态文件的。当我们对动态文件进行读写以后，此时内存中的动态文件和块设备文件中的静态文件就不同步了，当我们close关闭动态文件时，**close内部内核将内存中的动态文件的内容去更新(同步)的块设备中的静态文件**。

> 块设备本身读写非常不灵活 按块读写  。
>
> 而内存是按字节操作的，并且可以随机操作，很灵活。

3.`main` 函数参数的意义

```c
#include <stdio.h>

int main(int argc, char **argv) {
	printf("params num : %d\n",argc);
	for(int i = 0; i < argc; i ++) {
		printf("%s\n",argv[i]);
	}
}
```

argc ：代表的是 ./a.out ''文件1'' ''文件2'' 这三个参数的个数

argv[0] ：代表第一个参数./a.out

argv[1] ：代表第二个参数 ''文件1''

argv[2] ：代表第二个参数 ''文件2''

由此可见argv是存放 字符型数组 的数组

#### 2.3 标准C库部分函数格式

2.3.1 fputc 

写入字符、字符串 到文件  (fp 指向文件 str 存放字符串的指针) 

`fputc('c',fp);`

`fpuc(*str,fp);`

2.3.2 fgetc, feof

· fgetc ：意为从文件指针stream指向的文件中读取一个字符，读取一个字节后，光标位置后移一个字节。

返回值，是返回所读取的一个字节。如果读到文件末尾或者读取出错时返回EOF。虽然返回一个字节，但返回值不为unsigned char的原因为，返回值要能表示-1（即为EOF）。

· feof ：其功能是检测流上的文件结束符，如果文件结束，则返回非0值，否则返回0

注意：feof 判断文件结束是通过读取函数fread/fscanf等返回错误来识别的，故而判断文件是否结束应该是在读取函数之后进行判断。比如，在while循环读取一个文件时，如果是在读取函数之前进行判断，则如果文件最后一行是空白行，可能会造成内存错误。

## 二、进程

### 1.查看

`ps -aux | grep XXX`

或  `top` 

#### 1.1进程标识符 （pid）

非负整数表示的唯一ID 代表每一个进程

作用： 进程的调度，系统的初始化

> 存在父进程与子进程的关系

### 2.`fork` 的使用及其 与 `vfork`的区别

pid_t   fork( void ) ;

#### 2.1 `fork` 返回值

返回值为0，代表当前进程是子进程； 

返回值是正数，代表当前进程为父进程； 返回值即为子进程的pid。

调用失败，返回 -1。

#### 2.2 fork 与 vfork 的区别 （ 勘误原文 ）

vfork是**在父进程的地址空间中**创建一个子进程，子进程会暂时运行在父进程的空间中，直到调用exec或exit。在这期间，子进程对数据的修改会影响父进程。

另外，虽然fork创建的子进程会拷贝父进程的地址空间，但在Copy-On-Write（写时复制）的机制下，并不会立即复制整个地址空间，只有在子进程或父进程尝试修改共享的内存页时，才会进行复制。因此，父子进程在开始阶段共享地址空间，但在修改共享数据时会发生复制。

（ 1 ） vfork 不拷贝 直接使用父进程的存储空间 （共享资源）

（ 2 ）vfork 保证子进程先运行，子进程exit后 父进程才执行。

### 3. 进程的正常与异常退出

#### 3.1 正常退出

· main 末尾调用 return

· 进程 调用 exit( )  标准C库 通常写作 exit( 0 );

· 进程 调用 _exit( ) 或 _Exit( )  系统调用  

· 进程最后一个线程返回 调用 pthread_exit( );

#### 3.2 异常退出

· 调用 abort

· 进程收到信号  如 ctrl + c

· 最后一个线程对取消请求作出响应

>​	如果这个线程是最后一个活动的线程，整个进程也会因为没有活动线程而异常退出。这种情况可能会导致一些资源**没有被正确释放**，或者一些操作没有完成，因此需要在设计多线程程序时注意处理取消请求的情况，确保程序能够正确退出并**保持数据的一致性。**如：互斥锁 信号量等 处理共享资源

### 4. 异常进程

#### 4.1 僵尸进程

​	子进程终止状态不被父进程收集，导致终止运行的子进程的进程描述符仍然存在。僵尸进程会占用系统资源，如果系统中存在大量的僵尸进程，可能会导致系统资源耗尽。

​	父进程调用wait()或waitpid()等系统调用来获取子进程的终止状态，这样子进程的进程描述符就会被释放，不再是僵尸进程。

-> 处理方法 ： 调用 wait 等函数

```c
#include <sys/types.h>
#include <sys/wait.h>

   pid_t wait(int *status);

   pid_t waitpid(pid_t pid, int *status, int options);

   int waitid(idtype_t idtype, id_t id, siginfo_t *infop, int options);
```

在等待的过程中：

如果其所有子进程都还在运行，则阻塞。
如果一个子进程已终止，正等待父进程获取其终止状态，则取得该子进程的终止状态立即返回。
如果它没有任何子进程，则立即出错返回。
status参数是一个整型数指针

如果非空，子进程退出状态==存放在它指向的地址中==。

如果空，不关心退出状态。wait(NULL)

检查wait和waitpid所返回的终止状态的宏，来解析状态码。宏 说明 ：

```c
 WIFEXITED(status)	    //若为正常终止子进程返回的状态，则为真。对于这种情况可执行WEXITSTATUS(status)，取子进程传送给exit、_exit或_Exit参数的低8位
 WIFSIGNALED(status)	   //若为异常终止子进程返回的状态，则为真（接到一个不捕捉的信号）。对于这种情况，可执行WTERMSIG(status)，取使子进程终止的信号编号。另外，有些实现定义宏WCOREDUMP(status)//，若已产生终止进程的core文件，则它返回真
WIFSTOPPED(status)	   //若为当前暂停子进程的返回状态，则为真。对于这种情况，可执行WSTOPSIG(status)//，取使子进程暂停的信号编号
WIFCONTINUED(status)	   //若在作业控制暂停后已经继续的子进程返回了状态，则为真。（POSIX.1的XSI扩展；仅用于waitpid。）
```

注：父进程waitpid(pid,&status,WNOHANG);即为等待不阻塞 子进程仍是僵尸进程，俩进程都在运行

#### 4.2 孤儿进程

父进程在子进程之前终止，此时子进程叫做孤儿进程

**Linux避免系统存在过多的孤儿进程，init进程（系统的一个初始化进程，它的pid号为1）收留孤儿进程，变成孤儿进程的父进程**

### 5.  exec族函数

#### 5.1exec族函数的作用

我们用fork函数创建新进程后，经常会在新进程中调用exec函数去执行另外一个程序。当进程调用exec函数时，该进程被完全替换为新程序。因为调用exec函数并不创建新进程，所以前后进程的ID并没有改变

#### 5.2exec族函数功能

在调用进程内部执行一个可执行文件。可执行文件既可以是二进制文件，也可以是任何Linux下可执行的脚本文件。

#### 5.3函数族

exec函数族分别是：execl, execlp, execle, execv, execvp, execvpe

#### 5.4函数原型

```c
#include <unistd.h>
extern char **environ;

int execl(const char *path, const char *arg, ...);
int execlp(const char *file, const char *arg, ...);
int execle(const char *path, const char *arg,..., char * const envp[]);
int execv(const char *path, char *const argv[]);
int execvp(const char *file, char *const argv[]);
int execvpe(const char *file, char *const argv[],char *const envp[]);
```

#### 5.5返回值

*exec函数族的函数执行成功后不会返回，调用失败时，会设置errno并返回-1，然后从原程序的调用点接着往下执行。*
参数说明：
path：可执行文件的路径名字
arg：可执行程序所带的参数，第一个参数为可执行文件名字，没有带路径且arg必须以NULL结束
file：如果参数file中包含/，则就将其视为路径名，否则就按 PATH环境变量，在它所指定的各目录中搜寻可执行文件。

exec族函数参数极难记忆和分辨，函数名中的字符会给我们一些帮助：
l : 使用参数列表
p：使用文件名，并从PATH环境进行寻找可执行文件
v：应先构造一个指向各参数的指针数组，然后将该数组的地址作为这些函数的参数。
e：多了envp[]数组，使用新的环境变量代替调用进程的环境变量

#### 5.6以execl函数为例子来编写代码说明：

​	带l的一类exac函数（l表示list），包括execl、execlp、execle，要求将新程序的每个命令行参数都说明为 一个单独的参数。这种参数表**以空指针结尾**。

//文件execl.c

```c
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
//函数原型：int execl(const char *path, const char *arg, ...);

int main(void)
{
    printf("before execl\n");
    if(execl("./echoarg","echoarg","abc",NULL) == -1)
    {
        printf("execl failed!\n"); 

        perror("why");     //如果execl返回出错，返回了一个error，可以被这perror( )解析出来 
    }
    printf("after execl\n");
    return 0;

}
```



```c
//文件echoarg.c
#include <stdio.h>

int main(int argc,char *argv[])
{
    int i = 0;
    for(i = 0; i < argc; i++)
    {
        printf("argv[%d]: %s\n",i,argv[i]); 
    }
    return 0;
}
```

实验结果：

ubuntu:~/test/exec_test$ ./execl
before execl****
argv[0]: echoarg
argv[1]: abc

实验说明：

我们先用gcc编译echoarg.c，生成可执行文件echoarg并放在当前路径目录下。文件echoarg的作用是打印命令行参数。然后再编译execl.c并执行execl可执行文件。用execl 找到并执行echoarg，将当前进程main替换掉，所以”after execl” 没有在终端被打印出来。

#### 5.7活用execl族函数来查找系统时间

1. 首先用命令 whereis date ，找到系统时间 date 的绝对路径：

2. 

```C
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
//函数原型：int execl(const char *path, const char *arg, ...);

int main(void)
{
    printf("this pro get system data\n");
    if(execl("/bin/date","date",NULL,NULL) == -1)
    {
        printf("execl failed!\n");

        perror("why");

   }
    printf("after execl\n");
    return 0;
}
```



#### 5.8execlp函数

```c
#include<stdio.h>
#include<unistd.h>

int main(void)
{
   printf("before execl\n");

   if(execlp("ps","ps",NULL,NULL)==-1)
   {
        printf("execl failed\n");
        perror("why");
   } 

   printf("afther execl\n");

   return 0;
}
```



#### 5.9execvp

```C
#include<stdio.h>
#include<unistd.h>

int main(void)
{
   printf("before execl\n");

   char *argv[]={"ps",NULL,NULL};

   if(execvp("ps",argv)==-1)
   {
        printf("execl failed\n");
        perror("why");
   } 

   printf("afther execl\n");

   return 0;
}
```



### 6. linux 下修改环境变量配置绝对路径

修改环境变量的好处就是，可以直接将要加路径如 ./ 才能运行的可执行文件，直接就可以用名字就能运行。

用命令 pwd 找 找出当前路径
用命令 echo P A T H 找出环境变量（按图中的方法，这一步可省略，直接第 3 步即可）。用命令 e x p o r t P A T H = PATH 找出环境变量（按图中的方法，这一步可省略，直接第3步即可）。 用命令 export PATH=PATH找出环境变量（按图中的方法，这一步可省略，直接第3步即可）。用命令exportPATH=PATH:当前路径 ，将环境变量和当前路径连在一起 。这样修改环境变量就完成了

#### 6.1  Linux下exec配合fork使用

实现功能，当父进程检测到输入为1的时候，创建子进程把配置文件的字段值修改掉。

被修改的字段的配置文件config.txt

//config.txt

SPEED=5
LENG=9
SCORE=90
LEVEL=95

修改字段的文件 changData.c

//changData.c

```c
#include <sys/types.h>
#include <sys/stat.h>
#include <fcntl.h>
#include <stdio.h>
#include <unistd.h>
#include <string.h>
#include <stdlib.h>

int main(int argc,char **argv)
{
        int fdSrc;
        char *readBuf = NULL;
    if(argc != 2){
            printf("parames error");
            exit(-1);
    }
 
    fdSrc = open(argv[1],O_RDWR);            /1.打开文件
    int size = lseek(fdSrc,0,SEEK_END);
    lseek(fdSrc,0,SEEK_SET);
 
    readBuf =(char *)malloc(sizeof(char)*size +8);
    int n_read = read(fdSrc,readBuf,size);    //2.读文件
 
    char *p = strstr(readBuf,"LENG=");        //3.找到要修改的地方       
    if(p==NULL){                        
            printf("not found\n");
            exit(-1);
    }
    p = p + strlen("LENG=");
    *p = '5';
 
    lseek(fdSrc,0,SEEK_SET);
    int n_write =  write(fdSrc,readBuf,strlen(readBuf));    4.改了之后写入文件
 
    close(fdSrc);
 
    return 0;
}
```


将修改字段的文件changData.c ，（ gcc changData.c -o changData ） ，生成可执行文件 changData

由下面 execl( ) 函数配合 fork( ) 函数使用的代码，让 exexl( ) 函数调用可执行文件 changData，来修改配置文件

//demo.c

```c
#include <stdio.h>
#include <sys/types.h>
#include <unistd.h>
#include <sys/stat.h>
#include <fcntl.h>
#include <string.h>
#include <stdlib.h>

int main()
{
        pid_t pid;
        int data = 10;
    while(1){
            printf("please input your data:\n");
            scanf("%d",&data);
            if(data == 1){
                    pid = fork();
 
                    if(pid > 0){
                            wait(NULL);
                    }
 
                    if(pid == 0){
                            while(1){
 
                                    execl("./changData","changData","TEST.config",NULL);    //execl 调用 修改文件changData
 
                            }
                    }
            }
            else{
                    printf("wait , do nothing!\n");
            }
    }
    return 0;
```
}
配置文件被修改后：将 LENG=9 改成了 LENG=5

//config.txt

SPEED=5
LENG=5
SCORE=90
LEVEL=95

### 7. linux下system函数

#### 7.1 system()函数原型

NAME
system - execute a shell command

```C

#include <stdlib.h>

 int system(const char *command);
```



#### 7.2 ststem()函数返回值

成功，则返回**进程的状态值**；

当sh不能执行时，返回127；

失败返回 -1；

#### 7.3 system()函数源码

```C
int system(const char * cmdstring)
{
    pid_t pid;
    int status;
    if(cmdstring == NULL)
    {
        return (1); //如果cmdstring为空，返回非零值，一般为1
    }
    if((pid = fork())<0)
    {
        status = -1; //fork失败，返回-1
    }
    else if(pid == 0)
    {
        execl("/bin/sh", "sh", "-c", cmdstring, (char *)0);
        _exit(127); /* exec执行失败返回127，注意exec只在失败时才返回现在的进程，成功的话现在的
        进程就不存在*/
    }
    else //父进程
    {
        while(waitpid(pid, &status, 0) < 0)
        {
            if(errno != EINTR)
            {
                status = -1; //如果waitpid被信号中断，则返回-1
                break;
            }
        }
    }
    return status; //如果waitpid成功，则返回子进程的返回状态
```

#### 7.4 system()函数小应用代码demo

实现小功能，执行 vim 中的 ps - l 命令

```c
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
//函数原型：int system(const char *command);

int main(void)
{
    printf("this pro get system date:\n");
    if(system("ps -l") == -1)
    {
        printf("system failed!\n");
        perror("why");
   }
    printf("after system!!!\n");
    return 0;

}
```



通过运行结果可以看出：system( ) 函数调用完之后，代码还会往下走。 而exec族函数则不会往下走。

system( ) 函数的参数书写的规律是，可执行文件怎么执行，就怎么写：比如 system(“./a.out aa bb”)；

### 8. linux下popen函数

#### 8.1 popen原型

```c
#include <stdio.h>
FILE *popen(const char *command, const char *type);
 
   int pclose(FILE *stream);
```


​    
#### 8.2 函数说明

​	popen()函数通过创建一个管道，调用fork()产生一个子进程，执行一个shell以运行命令来开启一个进程。这个管道必须由pclose()函数关闭，而不是fclose()函数。	           	pclose()函数关闭标准I/O流，等待命令执行结束，然后返回shell的终止状态。如果shell不能被执行，则pclose()返回的终止状态与shell已执行exit一样。type参数只能是读或者写中的一种，得到的返回值（标准I/O流）也具有和type相应的只读或只写类型。

​	如果type是"r"则文件指针连接到command的标准输出；如果type是"w"则文件指针连接到command的标准输入。

​	command参数是一个指向以NULL结束的shell命令字符串的指针。这行命令将被传到bin/sh并使用-c标志，shell将执行这个命令。

​	popen()的返回值是个标准I/O流，必须由pclose来终止。前面提到这个流是单向的（只能用于读或写）。

​	向这个流写内容相当于写入该命令的标准输入，命令的标准输出和调用popen()的进程相同；与之相反的，从流中读数据相当于读取命令的标准输出，命令的标准输入和调用popen()的进程相同。

#### 8.3 返回值

如果调用 fork() 或 pipe() 失败，或者不能分配内存将返回NULL，否则返回标准I/O流。popen() 没有为内存分配失败设置errno值。如果调用fork()或pipe()时出现错误，errno被设为相应的错误类型。如果type参数不合法，errno将返回EINVAL。

#### 8.4 system()函数在应用中的好处

可以获取运行的结果

```c
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

//FILE *popen(const char *command, const char *type);

//size_t fread(void *ptr, size_t size, size_t nmemb, FILE *stream);

int main(void)
{
        FILE *fp;
        char ret[1024]={0};

        fp = popen("ps","r");   
     
        int nread = fread(ret,1,1024,fp);
     
        printf("read ret %d byte,ret = %s \n",nread,ret);
     
        return 0;

}
```

代码popen中的 ps 是linux中的 ps 指令。

### 9. **进程间通信**

### 介绍

进程间通信（IPC，InterProcess Communication）是指在不同的进程之间传播或交换信息。

IPC的方式有管道（包括无名管道和命名管道）、消息队列、信号、信号量、共享内存、Socket（套接字）、Streams等。其中**Socket**和**Streams**支持不同主机上的两个进程间通信。

#### 9.1 管道

##### 9.1.1 无名管道

· 半双工 ，即数据智能在一个方向上流动，具有固定的读端和写端。

· 只能用于父子进程之间的通信

· 管道创建在**内存**中，进程结束空间释放，管道销毁。可以用read、write对其进行读写。

-> 函数原型

```C
#include<unistd.h>

int pipe(int pipefd[2]);
```

->返回值

成功返回 0，  失败返回 -1

>创建成功会创建俩文件描述符，fd[0]为读而打开，fd[1]为写而打开。

demo

```C
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
 
//int pipe(int pipefd[2]);
//ssize_t write(int fd, const void *buf, size_t count);
//ssize_t read(int fd, void *buf, size_t count);
 
int main()
{
        int fd[2];
        int pid;
        char buf[128];
 
        if(pipe(fd) == -1){
                printf("creat pipe fail\n");
        }
 
        pid = fork();
 
        if(pid < 0 ){
                printf("creat child fail\n");
        }
        else if(pid > 0){
                sleep(3);
 
                printf("this is father\n");
                close(fd[0]);
                write(fd[1],"hello from father",strlen("hello from father"));
 
                wait();
        }
        else{
                printf("this is child\n");
                close(fd[1]);
                read(fd[0],buf,128);
                printf("read = %s \n",buf);
 
                exit(0);
        }
 
        return 0;
}

```

> 管道的特性：无数据时 阻塞

##### 9.1.2 命名管道

· 可以在无关的进程之间交换数据

· 有路径名与之关联，它以一种特殊设备文件形式存在与文件系统中。

-> 函数原型

```C
#include <sys/types.h>
#include <sys/stat.h>

int mkfifo(const char *pathname, mode_t mode);
```

其中mode参数与open函数中的mode相同。一旦创建了一个FIFO，就可以用一般的文件I/O函数来操作。

当open 一个FIFO时，是否设置非阻塞标志（O_NONBLOCK）的区别：

若没有指定O_NONBLOCK（默认），只读open要阻塞到某个其他进程为写而打开FIFO。类似的，只写open要阻塞到某个其他进程为读而打开它。（一般选择默认）
若是指定O_NONBLOCK，则只读open立即返回。而只写open将出错返回-1，如果没有进程已经为读而打开该FIFO，其errno为ENXIO。
FIFO的通信方式类似于在进程中使用文件来传输数据，只不过FIFO类型文件同时具有管道的特性。在数据读出时，FIFO管道中同时清楚数据，并且”先进先出“。

->命名管道demo

read.c

```c
#include <sys/types.h>
#include <sys/stat.h>
#include <stdio.h>
#include <errno.h>
#include <fcntl.h>
#include <unistd.h>
#include <string.h>
 
//ssize_t read(int fd, void *buf, size_t count);
//int mkfifo(const char *pathname, mode_t mode);
//int open(const char *pathname, int flags);
 
int main()
{
        int fd;
        char buf[128] = {0};
 
        if(mkfifo("./file",0600) == -1 && errno != EEXIST){
                printf("mkfifo fail\n");
                perror("why");
        }
 
        fd = open("./file",O_RDONLY);
        printf("read open success\n");
 
        while(1){
                int n_read = read(fd,buf,128);
                printf("read %d byte,contxt = %s \n",n_read,buf);
        }
 
        close(fd);
 
        return 0;
}
```

write.c

```c
#include <sys/types.h>
#include <sys/stat.h>
#include <stdio.h>
#include <errno.h>
#include <fcntl.h>
#include <unistd.h>
#include <string.h>
 
//ssize_t write(int fd, const void *buf, size_t count);
//int mkfifo(const char *pathname, mode_t mode);
//int open(const char *pathname, int flags);
 
int main()
{
        int fd;
        char *str = "message from fifo";
        int cnt = 0;
 
        fd = open("./file",O_WRONLY);
        printf("write open success\n");
 
        while(1){
                write(fd,str,strlen(str));
 
                sleep(1);
        }
        close(fd);
 
        return 0;
}

```



#### 9.2 消息队列

##### 9.2.1 特点

· 消息队列用于记录，其中的消息具有特定的格式以及特定的优先级。

· 消息队列独立于发送与接收进程。进程不影响消息队列及其中的内容。

· 可以实现信息的随机查询，可以按消息的类型读取。

##### 9.2.2 函数原型及参数

```C
#include <sys/types.h>
#include <sys/ipc.h>
#include <sys/msg.h>

//创建或打开消息队列 成功返回队列ID 失败返回-1
int msgget(key_t key, int msgflg);
//添加消息 ：成功返回0 失败返回 -1
int msgsnd(int msqid, const void *msgp, size_t msgsz, int msgflg);
//读取消息： 成功返回消息长度，失败返回 -1
ssize_t msgrcv(int msqid, void *msgp, size_t msgsz, long msgtyp, int msgflg);
//控制消息队列 : 成功返回0 失败返回 -1
int msgctl(int msqid, int cmd, struct msqid_ds *buf);
```

key值相当于一个索引。进程通过key在linux内核中找到相应队列。

在以下两种情况中，msgget 将创建一个新的消息队列：

--》如果没有与键值相对应的消息队列，并且flag中包含了 IPC_CREAT 标志位。
--》key 参数为 IPC_PRIVATE。

函数msgrcv在读取消息队列时，type参数有下面几种情况

​	type == 0 ，返回队列中的第一个消息；
​	type > 0 ，返回队列中消息类型为 type 的第一个消息；
​	type < 0 ，返回队列中消息类型值小于或等于 type 绝对值的消息，如果有多个，则取类型值最小的消息。

#### 9.3 信号

##### 9.3.1介绍

对于linux 来说 信号即软中断。信号为linux提供了一种处理异步事件的方法。 如ctrl + c 等

##### 9.3.2 信号的名字和编号

每个信号都有一个名字和编号，这些名字都以“SIG”开头，例如“SIGIO ”、“SIGCHLD”等等。
信号定义在**signal.h**头文件中，信号名都定义为正整数。
具体的信号名称可以使用kill -l来查看信号的名字以及序号，信号是从1开始编号的，不存在0号信号。kill对于信号0又特殊的应用。

##### 9.3.3  信号的处理

信号的处理有三种方法，分别是：忽略、捕捉 和 默认动作
· 忽略信号，大多数信号可以使用这个方式来处理，但是有两种信号不能被忽略（分别是 SIGKILL和SIGSTOP）。因为他们向内核和超级用户提供了进程终止和停止的可靠方法，如果忽略了，那么这个进程就变成了没人能管理的的进程，显然是内核设计者不希望看到的场景
· 捕捉信号，需要告诉内核，用户希望如何处理某一种信号，说白了就是写一个信号处理函数，然后将这个函数告诉内核。当该信号产生时，由内核来调用用户自定义的函数，以此来实现某种信号的处理。
· 系统默认动作，对于每个信号来说，系统都对应由默认的处理动作，当发生了该信号，系统会自动执行。不过，对系统来说，大部分的处理方式都比较粗暴，就是直接杀死该进程。
具体的信号默认动作可以使用man 7 signal来查看系统的具体定义。可以参考 《UNIX 环境高级编程（第三部）》的 P251——P256中间对于每个信号有详细的说明。

##### 9.3.4 信号处理函数的注册

1. 入门版：函数 **`signal`**
2. 高级版：函数 **`sigaction`**

demo.c  for signal

```C
#include <signal.h>
#include <stdio.h>
 
//       typedef void (*sighandler_t)(int);
 
//       sighandler_t signal(int signum, sighandler_t handler);
 
void handler(int signum)
{
        printf("get signum = %d\n",signum);
        switch(signum){
                case 2:
                        printf("SIGINT\n");
                        break;
                case 9:
                        printf("SIGKILL\n");
                        break;
                case 10:
                        printf("SIGUSR1\n");
                        break;
        }
}
 
int main()
{
        signal(SIGINT,handler);
        signal(SIGKILL,handler);
        signal(SIGUSR1,handler);
 
        while(1);
        return 0;
}
```

origin for sigaction

```C
#include <signal.h>
 
int sigaction(int signum, const struct sigaction *act,struct sigaction *oldact);
 
struct sigaction {
   void  (*sa_handler)(int); //信号处理程序，不接受额外数据，SIG_IGN 为忽略，SIG_DFL 为默认动作
   void  (*sa_sigaction)(int, siginfo_t *, void *); //信号处理程序，能够接受额外数据和sigqueue配合使用
   sigset_t  sa_mask;//阻塞关键字的信号集，可以再调用捕捉函数之前，把信号添加到信号阻塞字，信号捕捉函数返回之前恢复为原先的值。
   int  sa_flags;//影响信号的行为SA_SIGINFO表示能够接受数据
 };
//回调函数句柄sa_handler、sa_sigaction只能任选其一
```



##### 9.3.5 信号处理发送函数

  1.入门版：**kill**   //无法携带数据
  2.高级版：**sigqueue**

origin for kill

> 除了kill函数还可以用sprintf()配合system()函数来做到同样的信号处理发送

```C
#include <sys/types.h>
#include <signal.h>
 
int kill(pid_t pid, int sig);
```

origin for sigqueue

```C
#include <signal.h>
int sigqueue(pid_t pid, int sig, const union sigval value);
union sigval {
   int   sival_int;
   void *sival_ptr;
 };
```

sigqueue 函数只能把信号发送给单个进程，可以使用 value 参数向信号处理程序传递整数值或者指针值。

sigqueue 函数不但可以发送额外的数据，还可以让信号进行排队（操作系统必须实现了 POSIX.1的实时扩展），对于设置了阻塞的信号，使用 sigqueue 发送多个同一信号，在解除阻塞时，接受者会接收到发送的信号队列中的信号，而不是直接收到一次。

但是，信号不能无限的排队，信号排队的最大值受到SIGQUEUE_MAX的限制，达到最大限制后，sigqueue 会失败，errno 会被设置为 EAGAIN。

#### 9.4 信号量

##### 9.4.1 介绍

[信号量](https://so.csdn.net/so/search?q=信号量&spm=1001.2101.3001.7020)（semaphore）与已经介绍过的 IPC 结构不同，它是一个计数器。信号量用于实现进程间的互斥与同步，而不是用于存储进程间通信数据

##### 9.4.2 特点

· 信号量用于进程间同步，若要在进程间传递胡数据需要结合共享内存

· 信号量基于操作系统的 PV 操作，程序对信号量的操作都是原子操作。（P操作:拿锁。V操作：放回锁）

· 每次对信号量的 PV 操作不仅限于对信号量值加 1 或 减1 ，而且可以加加减任意正整数。

· 支持<u>信号量组</u> 

> 信号量组是一种用于控制并发访问的机制，可以用来保护共享资源，防止多个线程同时访问。支持信号量组意味着系统或程序可以对信号量组进行操作，例如创建、初始化、增加或减少信号量值等。这样可以更有效地管理并发访问，提高系统的稳定性和性能。

##### 9.4.3 函数原型

最简单的信号量是只能取 0 和 1 的变量，这也是信号量最常见的一种形式，叫做**二值信号量**（Binary Semaphore）。而可以取多个正整数的信号量被称为**通用信号量**。

Linux 下的信号量函数都是在通用的信号量数组上进行操作，而不是在一个单一的二值信号量上进行操作。

```C
#include <sys/types.h>
#include <sys/ipc.h>
#include <sys/sem.h>

//创建或获取一个信号量组，成功会返回信号量集 ID (即semid)，失败返回 -1
int semget(key_t key, int nsems, int semflg);
//对信号量组进行操作，改变信号量的值，成功返回 0，失败返回 -1 （用于 PV 操作）
int semop(int semid, struct sembuf *sops, unsigned nsops);    
//控制信号量的相关信息 （用于给信号量初始化）
int semctl(int semid, int semnum, int cmd, ...);
```

-> semop 中的 sembuf 

```C
指向一个结构数组的指针，每个数组元素至少包含以下几个成员：
 
struct sembuf{
   short sem_num; //信号量编号，除非使用一组信号量，否则它的取值为0
   short sem_op;  //信号量在一次操作中需要改变的数值。通常用到两个值，-1，也就是p操作，表示拿锁；+1，也就是V操作，表示放回锁。
   short sem_flg; //通过被设置为SEM_UNDO。表示操作系统会跟踪当前进程对这个信号量的修改情况，如果这个进程在没有释放该信号量的情况下终止，操作系统将自动释放该进程持有的信号量，防止其他进程一直处于等待状态。 
};  
```

-> semctl

系统调用semctl用来执行在信号量集上的控制操作。这和在消息队列中的系统调用msgctl是十分相似的

参数 cmd 可以使用的命令

·IPC_STAT 读取一个信号量集的数据结构semid_ds，并将其存储在semun中的buf参数中。
·IPC_SET 设置信号量集的数据结构semid_ds中的元素ipc_perm，其值取自semun中的buf参数。
·IPC_RMID 将信号量集从内存中删除。
·GETALL 用于读取信号量集中的所有信号量的值。
·GETNCNT 返回正在等待资源的进程数目。
·GETPID 返回最后一个执行semop操作的进程的PID。
·GETVAL 返回信号量集中的一个单个的信号量的值。
·GETZCNT 返回正在等待完全空闲的资源的进程数目。
·SETALL 设置信号量集中的所有的信号量的值。
·SETVAL 设置信号量集中的一个单独的信号量的值。【一般用这个】

##### 9.4.4 信号量demo .c

```C
#include <stdio.h>
#include <sys/types.h>
#include <sys/ipc.h>
#include <sys/sem.h>
#include <stdlib.h>
 
//联合体，用于semctl初始化
union semun
{
	int              val;
	struct semid_ds *buf;
	unsigned short *array;
};	
 
//初始化信号量
int init_sem(int sem_id,int value)
{
	union semun tmp;
	tmp.val = value;
 
	if(semctl(sem_id,0,SETVAL,tmp) == -1){
		perror("init semaphore error");
		return -1;
	}
	return 0;
}
 
// P 操作
// 若信号量值为 1，获取资源并将信号量值 -1
// 若信号量值为 0，进程挂起等待
int sem_p(int sem_id)
{
        struct sembuf sbuf;
        sbuf.sem_num = 0;
        sbuf.sem_op = -1;
        sbuf.sem_flg = SEM_UNDO;
 
        if(semop(sem_id,&sbuf,1) == -1){
                perror("P operation error");
                return -1;
        }
        
	return 0;
}
 
// V 操作
// 释放资源并将信号量值+1
// 如果有进程正在挂起等待，则唤醒他们
int sem_v(int sem_id)
{
	struct sembuf sbuf;
	sbuf.sem_num = 0;	//序号
	sbuf.sem_op = 1;	// V 操作
	sbuf.sem_flg = SEM_UNDO;
 
	if(semop(sem_id,&sbuf,1) == -1){
		perror("V operation error");
		return -1;
	}
	
	return 0;
}
 
//删除信号量集
int del_sem(int sem_id)
{
	union semun tmp;
 
	if(semctl(sem_id,0,IPC_RMID,tmp) == -1){
		perror("delete semaphore erroe");
		return -1;
	}
	return 0;
}
 
int main()
{
	int   sem_id; //信号量集 ID
	key_t key;
	pid_t pid;
 
	//获取key值
	if((key = ftok(".",'z')) < 0){
		perror("ftok error");
		exit(1);
	}
 
	//创建信号量集，其中只有一个信号量
	if(sem_id = semget(key,1,IPC_CREAT|0666) == -1){
		perror("semget error");
		exit(1);
	}
 
	//初始化：初值设为 0 资源被占用
	init_sem(sem_id,0);
 	pid = fork()；
	if(pid == -1){
		perror("Fork error");		
	}
	else if(pid == 0){	//子进程
		sleep(2);
		printf("process child: pid = %d\n",getpid());
		sem_v(sem_id);	//释放资源
	}
	else{ //父进程
		sem_p(sem_id);	//等待资源
		printf("process father: pid = %d\n",getpid());
		sem_v(sem_id);	//释放资源
		del_sem(sem_id);//删除信号量集
	}
 
	return 0;
}
```

#### 9.5 共享内存

> 一般搭配信号量使用

##### 9.5.1 介绍

共享内存（shared memory），指两个或多个进程共享一个给定的存储区。

**查看共享内存的命令：ipcs -m （在共享内存创建之后，并在断开连接之前，可以加exit(0)退出程序，此时可以用查看命令看到创建的共享内存有哪些）**

**删除共享内存的命令：ipcs -m id号**

##### 9.5.2 特点

共享内存是最快的一种IPC，因为进程是直接对内存进行存储，而不需要任何数据的拷贝。

它有一个特性：只能单独一个进程写或读，如果A和B进程同时写，会造成数据的混乱，（所以需要搭配信号量来使用 异步）

##### 9.5.3 函数原型

```C
//创建或获取一个共享内存：成功返回共享内存ID，失败返回 -1
int shmget(key_t key, size_t size, int shmflg);
//连接共享内存到当前进程的地址空间:成功返回指向共享内存的指针，失败返回 -1
void *shmat(int shmid, const void *shmaddr, int shmflg);
//断开与共享内存的连接：成功返回0，失败返回 -1
int shmdt(const void *shmaddr);
//控制共享内存的相关信息：成功返回0，失败返回 -1
int shmctl(int shmid, int cmd, struct shmid_ds *buf);
```

->    当 **shmget**函数创建一段共享内存的时候，必须指定其size，而且必须是以兆对齐的（一兆空间是：1024）；而如果引用一个已存在的共享内存，则将size指定为0。

​	参数3：一般是，创建+权限（可读可写），IPC_CREAT|0600

- 当一段共享内存被创建以后，它并不能被任何进程访问。必须使用 **shmat**函数连接该共享内存到当前进程的地址空间，连接成功后把共享内存区映射到调用进程的地址空间，随后可像本地空间一样访问。  					

  参数1：共享内存id ，参数2：一般写0，默认是由linux内核自行安排共享内存， 参数3：一般写0，代表映射的共享内存是可读可写。若是指定了 SHM_RDONLY ，则是以只读的方

- **shmdt函数**函数是用来断开shmat建立的连接的。注意，这并不是从系统中删除该共享内存，只是当前进程不能访问该共享内存而已。

- shmctl函数可以对共享内存执行多种操作，根据参数 cmd 执行相应的操作。常用的是IPC_RMID（从系统中删除该共享内存）。

​	shmctl ，参数2：删除指令，一般写IPC_RMID ，参数3：	一般写0，因为我们一般不关心共享内存中的数据信息。

##### 9.5.4 编程思路

1. 创建/打开共享内存
2. 连接映射共享内存
3. 写入数据 strcpy
4. 断开共享内存连接
5. 销毁共享内存

##### 9.5.5 代码demo

shm_write.c

```C
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/ipc.h>
#include <sys/shm.h>
 
//int shmget(key_t key, size_t size, int shmflg);
//void *shmat(int shmid, const void *shmaddr, int shmflg);
 
int main()
{
        int shmid;
        char *shmaddr;
 
        key_t key;
        key = ftok(".",1);    //“.” 代表当前路径 ，第二个参数随意数字
        
        //创建共享内存
        shmid = shmget(key,1024*4,IPC_CREAT|0600);
        if(shmid == -1){
                printf("creat shm fail\n");
                exit(-1);
        }
 
        //连接映射共享内存
        shmaddr = shmat(shmid,0,0);
 
        printf("shmat OK\n");
 
        //将数据拷贝到共享内存
        strcpy(shmaddr,"hello world\n");
 
        sleep(5);            //等待5秒，避免一下子断开连接。等待另外一个进程读完。
 
        //断开共享内存连接
        shmdt(shmaddr);
        //删除共享内存
        shmctl(shmid,IPC_RMID,0);
 
        printf("quit\n");
 
        return 0;
}
```

shm_get.c

```C
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/ipc.h>
#include <sys/shm.h>
 
//int shmget(key_t key, size_t size, int shmflg);
//void *shmat(int shmid, const void *shmaddr, int shmflg);
 
int main()
{
        int shmid;
        char *shmaddr;
 
        key_t key;
        key = ftok(".",1);    //“.” 代表当前路径 ，第二个参数随意数字
 
        //打开共享内存    
        shmid = shmget(key,1024*4,0);    //以打开方式的时候，第三个参数写0，直接获取，不创建
        if(shmid == -1){
                printf("creat shm fail\n");
                exit(-1);
        }
 
        //连接并映射共享内存
        shmaddr = shmat(shmid,0,0);
 
        printf("get from shm_write message is : %s",shmaddr);
 
        //断开共享内存连接
        shmdt(shmaddr);
 
        //删除共享内存
        shmctl(shmid,IPC_RMID,0);
 
        printf("quit\n");
 
        return 0;
}
```

#### 9.6 结合消息队列 信号量 共享内存

代码demo.c

get.c

```C
#include<stdio.h>
#include<stdlib.h>
#include<sys/shm.h>  // shared memory
#include<sys/sem.h>  // semaphore
#include<sys/msg.h>  // message queue
#include<string.h>   // memcpy
 
// 消息队列结构
struct msg_form {
    long mtype;
    char mtext;
};
 
// 联合体，用于semctl初始化
union semun
{
    int              val; /*for SETVAL*/
    struct semid_ds *buf;
    unsigned short  *array;
};
 
// 初始化信号量
int init_sem(int sem_id, int value)
{
    union semun tmp;
    tmp.val = value;
    if(semctl(sem_id, 0, SETVAL, tmp) == -1)
    {
        perror("Init Semaphore Error");
        return -1;
    }
    return 0;
}
 
// P操作:
//  若信号量值为1，获取资源并将信号量值-1 
//  若信号量值为0，进程挂起等待
int sem_p(int sem_id)
{
    struct sembuf sbuf;
    sbuf.sem_num = 0; /*序号*/
    sbuf.sem_op = -1; /*P操作*/
    sbuf.sem_flg = SEM_UNDO;
 
    if(semop(sem_id, &sbuf, 1) == -1)
    {
        perror("P operation Error");
        return -1;
    }
    return 0;
}
 
// V操作：
//  释放资源并将信号量值+1
//  如果有进程正在挂起等待，则唤醒它们
int sem_v(int sem_id)
{
    struct sembuf sbuf;
    sbuf.sem_num = 0; /*序号*/
    sbuf.sem_op = 1;  /*V操作*/
    sbuf.sem_flg = SEM_UNDO;
 
    if(semop(sem_id, &sbuf, 1) == -1)
    {
        perror("V operation Error");
        return -1;
    }
    return 0;
}
 
// 删除信号量集
int del_sem(int sem_id)
{
    union semun tmp;
    if(semctl(sem_id, 0, IPC_RMID, tmp) == -1)
    {
        perror("Delete Semaphore Error");
        return -1;
    }
    return 0;
}
 
// 创建一个信号量集
int creat_sem(key_t key)
{
    int sem_id;
    if((sem_id = semget(key, 1, IPC_CREAT|0666)) == -1)
    {
        perror("semget error");
        exit(-1);
    }
    init_sem(sem_id, 1);  /*初值设为1资源未占用*/
    return sem_id;
}
 
 
int main()
{
    key_t key;
    int shmid, semid, msqid;
    char *shm;
    char data[] = "this is server";
    struct shmid_ds buf1;  /*用于删除共享内存*/
    struct msqid_ds buf2;  /*用于删除消息队列*/
    struct msg_form msg;  /*消息队列用于通知对方更新了共享内存*/
 
    // 获取key值
    if((key = ftok(".", 'z')) < 0)
    {
        perror("ftok error");
        exit(1);
    }
 
    // 创建共享内存
    if((shmid = shmget(key, 1024, IPC_CREAT|0666)) == -1)
    {
        perror("Create Shared Memory Error");
        exit(1);
    }
 
    // 连接共享内存
    shm = (char*)shmat(shmid, 0, 0);
    if((int)shm == -1)
    {
        perror("Attach Shared Memory Error");
        exit(1);
    }
 
 
    // 创建消息队列
    if ((msqid = msgget(key, IPC_CREAT|0777)) == -1)
    {
        perror("msgget error");
        exit(1);
    }
 
    // 创建信号量
    semid = creat_sem(key);
    
    // 读数据
    while(1)
    {
        msgrcv(msqid, &msg, 1, 888, 0); /*读取类型为888的消息*/
        if(msg.mtext == 'q')  /*quit - 跳出循环*/ 
            break;
        if(msg.mtext == 'r')  /*read - 读共享内存*/
        {
            sem_p(semid);
            printf("%s\n",shm);
            sem_v(semid);
        }
    }
 
    // 断开连接
    shmdt(shm);
 
    /*删除共享内存、消息队列、信号量*/
    shmctl(shmid, IPC_RMID, &buf1);
    msgctl(msqid, IPC_RMID, &buf2);
    del_sem(semid);
    return 0;
}
```

send.c

```C
#include<stdio.h>
#include<stdlib.h>
#include<sys/shm.h>  // shared memory
#include<sys/sem.h>  // semaphore
#include<sys/msg.h>  // message queue
#include<string.h>   // memcpy
 
// 消息队列结构
struct msg_form {
    long mtype;
    char mtext;
};
 
// 联合体，用于semctl初始化
union semun
{
    int              val; /*for SETVAL*/
    struct semid_ds *buf;
    unsigned short  *array;
};
 
// P操作:
//  若信号量值为1，获取资源并将信号量值-1 
//  若信号量值为0，进程挂起等待
int sem_p(int sem_id)
{
    struct sembuf sbuf;
    sbuf.sem_num = 0; /*序号*/
    sbuf.sem_op = -1; /*P操作*/
    sbuf.sem_flg = SEM_UNDO;
 
    if(semop(sem_id, &sbuf, 1) == -1)
    {
        perror("P operation Error");
        return -1;
    }
    return 0;
}
 
// V操作：
//  释放资源并将信号量值+1
//  如果有进程正在挂起等待，则唤醒它们
int sem_v(int sem_id)
{
    struct sembuf sbuf;
    sbuf.sem_num = 0; /*序号*/
    sbuf.sem_op = 1;  /*V操作*/
    sbuf.sem_flg = SEM_UNDO;
 
    if(semop(sem_id, &sbuf, 1) == -1)
    {
        perror("V operation Error");
        return -1;
    }
    return 0;
}
 
 
int main()
{
    key_t key;
    int shmid, semid, msqid;
    char *shm;
    struct msg_form msg;
    int flag = 1; /*while循环条件*/
 
    // 获取key值
    if((key = ftok(".", 'z')) < 0)
    {
        perror("ftok error");
        exit(1);
    }
 
    // 获取共享内存
    if((shmid = shmget(key, 1024, 0)) == -1)
    {
        perror("shmget error");
        exit(1);
    }
 
    // 连接共享内存
    shm = (char*)shmat(shmid, 0, 0);
    if((int)shm == -1)
    {
        perror("Attach Shared Memory Error");
        exit(1);
    }
 
    // 创建消息队列
    if ((msqid = msgget(key, 0)) == -1)
    {
        perror("msgget error");
        exit(1);
    }
 
    // 获取信号量
    if((semid = semget(key, 0, 0)) == -1)
    {
        perror("semget error");
        exit(1);
    }
    
    // 写数据
    printf("***************************************\n");
    printf("*                 IPC                 *\n");
    printf("*    Input r to send data to server.  *\n");
    printf("*    Input q to quit.                 *\n");
    printf("***************************************\n");
    
    while(flag)
    {
        char c;
        printf("Please input command: ");
        scanf("%c", &c);
        switch(c)
        {
            case 'r':
                printf("Data to send: ");
                sem_p(semid);  /*访问资源*/
                scanf("%s", shm);
                sem_v(semid);  /*释放资源*/
                /*清空标准输入缓冲区*/
                while((c=getchar())!='\n' && c!=EOF);
                msg.mtype = 888;  
                msg.mtext = 'r';  /*发送消息通知服务器读数据*/
                msgsnd(msqid, &msg, sizeof(msg.mtext), 0);
                break;
            case 'q':
                msg.mtype = 888;
                msg.mtext = 'q';
                msgsnd(msqid, &msg, sizeof(msg.mtext), 0);
                flag = 0;
                break;
            default:
                printf("Wrong input!\n");
                /*清空标准输入缓冲区*/
                while((c=getchar())!='\n' && c!=EOF);
        }
    }
 
    // 断开连接
    shmdt(shm);
 
    return 0;
}
```

## 三、Linux多线程

### 1.线程创建与等待

#### 1.1 线程创建函数原型与demo

```C
int pthread_create(pthread_t *restrict tidp, const pthread_attr_t *restrict attr, void *(*start_rtn)(void *), void *restrict arg);
```

```C
#include <stdio.h>
#include <pthread.h>
 
//int pthread_create(pthread_t *restrict tidp, const pthread_attr_t *restrict attr,
//       void *(*start_rtn)(void *), void *restrict arg);
 
void *func1(void *arg)
{
        printf("%ld thread is create\n",(unsigned long)pthread_self());    //打印线程的pid
        printf("param is %d\n",*((int *)arg));    
}
 
int main()
{
        int ret;
        pthread_t t1;     /*不使用指针是以免空指针异常*/
        int arg = 100;
 
        //创建 t1 线程
        ret = pthread_create(&t1,NULL,func1,(void *)&arg);    //参数2：线程属性，一般设置为NULL，参数3：线程干活的函数，参数4：往t1线程里面传送数据。
        //ret = pthread_create(&t1,NULL,func1,NULL);   
 
        if(ret == 0){
                printf("create t1 success\n");
        }
 
        return 0;
}
```

#### 1.2线程等待函数原型与demo

```C
int pthread_join(pthread_t thread, void **rval_ptr);

// 返回：若成功返回0，否则返回错误编号

第 2 个参数，可以设置为NULL不关心线程收回的退出状态 

```

```C
#include <stdio.h>
#include <pthread.h>
 
//int pthread_create(pthread_t *restrict tidp, const pthread_attr_t *restrict attr,
//       void *(*start_rtn)(void *), void *restrict arg);
 
//nt pthread_join(pthread_t thread, void **rval_ptr);
 
void *func1(void *arg)
{
        static char *p = "t1 is run out";    //变量必须定义前加 static ，否则二级指针指向它的时候数据会出错
 
        printf("t1:%ld thread is create\n",(unsigned long)pthread_self());
        printf("t1:param is %d\n",*((int *)arg));
 
        pthread_exit((void *)p); //线程退出，并且返回 p 指向的字符串
}
 
int main()
{
        int ret;
        pthread_t t1;
        int arg = 100;
 
        char *pret = NULL;    //不可以直接定义为二级指针
 
        //创建线程
        ret = pthread_create(&t1,NULL,func1,(void *)&arg);
 
        if(ret == 0){
                printf("main:create t1 success\n");
        }
 
        printf("main:%ld \n",(unsigned long)pthread_self());
 
        //线程的等待/阻塞
        pthread_join(t1,(void **)&pret);  //参数2：将指针pret 指向 t1线程的p    
 
        printf("main: t1 quit :%s\n",pret);
 
        return 0;
}
```

### 2.线程同步之互斥量加锁和解锁

> 注： 互斥锁的作用是用来控制线程同步的，但是只能控制一个线程执行完才到下一个线程，不能保证线程的运行顺序。

#### 2.1 创建及销毁互斥锁

```C
#include <pthread.h>
 
int pthread_mutex_init(pthread_mutex_t *restrict mutex, const pthread_mutexattr_t *restrict attr); //初始化互斥量，默认属性attr参数可以设置为NULL。
 
int pthread_mutex_destroy(pthread_mutex_t mutex);  //释放互斥量                             
 
// 返回：若成功返回0，否则返回错误编号
 
要用默认的属性初始化互斥量，只需把attr设置为NULL。
```

#### 2.2 加锁及解锁

``` C
#include <pthread.h>
 
int pthread_mutex_lock(pthread_mutex_t *mutex);    //加锁
 
int pthread_mutex_trylock(pthread_mutex_t *mutex);
 
int pthread_mutex_unlock(pthread_mutex_t *mutex);  //解锁
 
// 返回：若成功返回0，否则返回错误编号
```

如果线程不希望被阻塞，它可以使用pthread_mutex_trylock尝试对互斥量进行加锁。如果调用pthread_mutex_trylock时互斥量处于未锁住状态，那么pthread_mutex_trylock将锁住互斥量，不会出现阻塞并返回0，否则pthread_mutex_trylock就会失败，不能锁住互斥量，而返回EBUSY。

### 3.条件控制实现线程的同步

#### 3.1创建和销毁条件变量

```C
#include <pthread.h>
 
int pthread_cond_init(pthread_cond_t *restrict cond, const pthread_condattr_t *restrict attr);
 
int pthread_cond_destroy(pthread_cond_t cond);
 
// 返回：若成功返回0，否则返回错误编号
```



#### 3.2 等待  

```C
#include <pthread.h>
 
int pthread_cond_wait(pthread_cond_t *restrict cond, pthread_mutex_t *restrict mutex);
 
int pthread_cond_timedwait(pthread_cond_t *restrict cond, pthread_mutex_t *restrict mutex, cond struct timespec *restrict timeout);
 
// 返回：若成功返回0，否则返回错误编号
```

- pthread_cond_wait等待条件变为真。如果在给定的时间内条件不能满足，那么会生成一个代表一个出错码的返回变量。传递给pthread_cond_wait的互斥量对条件进行保护，调用者把锁住的互斥量传给函数。函数把调用线程放到等待条件的线程列表上，然后对互斥量解锁，这两个操作都是原子操作。这样就关闭了条件检查和线程进入休眠状态等待条件改变这两个操作之间的时间通道，这样线程就不会错过条件的任何变化。pthread_cond_wait返回时，互斥量再次被锁住。

- pthread_cond_timedwait函数的工作方式与pthread_cond_wait函数类似，只是多了一个timeout。timeout指定了等待的时间，它是通过timespec结构指定。

#### 3.3 触发

```C
#include <pthread.h>
 
int pthread_cond_signal(pthread_cond_t cond);       //触发
 
int pthread_cond_broadcast(pthread_cond_t cond);    //广播
 
// 返回：若成功返回0，否则返回错误编号

```

- 一定要在改变条件状态以后再给线程发信号
- 这两个函数可以用于通知线程条件已经满足。pthread_cond_signal函数将唤醒等待该条件的某个线程，而pthread_cond_broadcast函数将唤醒等待该条件的所有进程。

#### 3.4 demo

```C
#include <stdio.h>
#include <pthread.h>
 
//int pthread_create(pthread_t *restrict tidp, const pthread_attr_t *restrict attr,
//       void *(*start_rtn)(void *), void *restrict arg);
 
int g_data = 0;
 
pthread_mutex_t mutex;    //锁
pthread_cond_t  cond;     //条件   
 
void *func1(void *arg)
{
        printf("t1: %ld thread is create\n",(unsigned long)pthread_self());
        printf("t1: param is %d\n",*((int *)arg));
 
        while(1){
                pthread_cond_wait(&cond,&mutex);    //条件的等待
                printf("t1 run =========================\n");
 
                printf("t1: %d\n",g_data);
                g_data = 0;
 
                sleep(1);
        }
}
 
void *func2(void *arg)
{
        printf("t2: %ld thread is create\n",(unsigned long)pthread_self());
        printf("t2: param is %d\n",*((int *)arg));
 
        while(1){
                printf("t2: %d\n",g_data);
 
                pthread_mutex_lock(&mutex);            //加锁
                g_data++;
                if(g_data == 3){
                        pthread_cond_signal(&cond);    //条件的信号
                }
                pthread_mutex_unlock(&mutex);          //解锁
                sleep(1);
        }
}
 
int main()
{
        int ret;
        int arg = 100;
        pthread_t t1;
        pthread_t t2;
 
        pthread_mutex_init(&mutex,NULL);    //锁的创建（动态初始化）
        pthread_cond_init(&cond,NULL);      //条件的创建（动态初始化）
 
        ret = pthread_create(&t1,NULL,func1,(void *)&arg);
        if(ret == 0){
//              printf("main:create t1 success\n");
        }
 
        ret = pthread_create(&t2,NULL,func2,(void *)&arg);
        if(ret == 0){
//              printf("main:create t2 success\n");
        }
 
//      printf("main:%ld\n",(unsigned long)pthread_self());
 
        pthread_join(t1,NULL);
        pthread_join(t2,NULL);
 
        pthread_mutex_destroy(&mutex);    //销毁锁
        pthread_cond_destroy(&cond);        //条件的销毁
 
        return 0;
}
```

- 整个代码，主要就是让线程2中的 data++ 到 3 的时候，就触发发出一个信号signal，

- 线程1等待接收到信号signal，之后就会开始运行。

#### 拓展： Linux测试脚本格式

test.c

```bash
int main()
{
        int i;
 
        for(i=0;i<10;i++){
                system("./pthread");
        }
 
        return 0;
}
```

gcc -o pthread pthread.c

gcc -o test test.c

**命令：chmod +x test （\**chmod +x的意思就是给执行权限\**）**

**./test (运行)**

将测试结果保存到文档 ： **./test >>ret.txt & （加个&有点类似于后台运行）**



》==Linux网络编程== 另文档