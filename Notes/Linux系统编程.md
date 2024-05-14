[文章来源]: https://blog.csdn.net/qq_47944751/article/details/131568915	"Linux 系统编程"



# Linux系统编程  

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
```
}

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

失败返回-1；

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
        进程就不存在啦8*/
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

​       #include <stdio.h>

       FILE *popen(const char *command, const char *type);
     
       int pclose(FILE *stream);
       
    
#### 8.2 函数说明

​	popen()函数通过创建一个管道，调用fork()产生一个子进程，执行一个shell以运行命令来开启一个进程。这个管道必须由pclose()函数关闭，而不是fclose()函数。	           	pclose()函数关闭标准I/O流，等待命令执行结束，然后返回shell的终止状态。如果shell不能被执行，则pclose()返回的终止状态与shell已执行exit一样。type参数只能是读或者写中的一种，得到的返回值（标准I/O流）也具有和type相应的只读或只写类型。

​	如果type是"r"则文件指针连接到command的标准输出；如果type是"w"则文件指针连接到command的标准输入。

​	command参数是一个指向以NULL结束的shell命令字符串的指针。这行命令将被传到bin/sh并使用-c标志，shell将执行这个命令。

​	popen()的返回值是个标准I/O流，必须由pclose来终止。前面提到这个流是单向的（只能用于读或写）。

​	向这个流写内容相当于写入该命令的标准输入，命令的标准输出和调用popen()的进程相同；与之相反的，从流中读数据相当于读取命令的标准输出，命令的标准输入和调用popen()的进程相同。

#### 8.3 返回值

如果调用fork()或pipe()失败，或者不能分配内存将返回NULL，否则返回标准I/O流。popen()没有为内存分配失败设置errno值。如果调用fork()或pipe()时出现错误，errno被设为相应的错误类型。如果type参数不合法，errno将返回EINVAL。

#### 8.4 比system()函数在应用中的好处：可以获取运行的结果

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