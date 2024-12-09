FULLY COPY from github.com/thuyy/ByteYoungDB 

块存储 文件存储 对象存储 key-value存储
本项目实现的是 内存态  并没有持久化相关内容实现

不考虑并发 
.clang-format 格式化工具 一般有IDE自动生成
编译出的.so动态链接库（sql-parser/编译得）与系统有关 

- parser 语义解析 parser库只支持词法、语法解析 －－ 合理性检验需要自己实现

- metadata 处理存储元信息

- optimizer 根据解析的语法生成计划树 PlanTree

- executor 根据计划树 执行  （火山模型 next_.exec()）依次执行 

- storage 内存态信息 数据结构 的定义

- util 一些工具 ：重命名、类型转换（`_type`转换为string类型方便使用）、debug输出

- trx (transaction) 使用 undo stack 实现commit和Rollback 版本提交和回溯


<h1>操作方法</h1>
```shell
mkdir build
cd build 
cmake ..
make
```
可运行文件生成于./build/bin
- sql-parser-test [SQL语句]
验证正误
- bydb 
进入数据库命令行界面
