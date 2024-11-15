数据的持久化
- 校验数据的合法性
- 修改内存（用高效的数据结构组织数据）
- 写入存储介质（寿命/性能友好的方式）
缓存 贯穿存储体系 拷贝 尽量减少
需要抽象统一的接入层

## 存储系统 
- RAID技术
高性能性价比可靠性  

关系型 结构化 支持复杂事务（DSL） 非关系型

## 单机存储

Linux 两大数据结构 ：Index node 唯一标识 持久化到磁盘上、 Directory Entry 记录文件名、inode指针、层级关系 内存结构 不会持久化到磁盘  (二者n:1的关系 有关硬链接的实现)

key-value存储
put(k,v) & get(k)
	常见数据结构：LSM-Tree _牺牲读性能，追求写入性能  

## 分布式存储 
	单机存储的基础上实现了分布式协议 设计大量网络交互
1. HDFS 
2. Ceph
数据分布模型：CRUSH算法 (hash + 权重 + 随机抽签) 
一切皆对象 支持对象、块、文件接口
数据写入：主备复制模型

## 单机数据库
单个计算机节点上的数据库系统
> 事务单机内执行，也可能通过网络交互实现分布式事务
- 关系型数据库
MySQL & postgreSQL 
Query Engine 解析query 生成查询计划
Txn Manager 事务并发管理
Lock Manager 锁相关的策略
Storage Engine 组织内存/磁盘数据结构
Replication 主备同步

关键内存数据结构：B－Tree ,B+-Tree ,LRU List等
关键磁盘数据结构：WriteAheadLog(RedoLog), Page

- 非关系型数据库 不使用SQL交互 数据结构不固定 无关系约束时, schema灵活 )在尝试支持SQL与事务
MongoDB & Redis & ElasticSearch

）ElasticSearch
面向文档存储
文档可序列化成JSON，支持嵌套
存在 index index=文档集合
存储和构建索引能力依赖Lucene引擎
实现了大量搜索数据结构和算法
支持RESTFUL API 也支持SQL交互

）MongoDB 灵活
...
存在collection=文档集合
依赖wiredTiger引擎（纯c）
4.0后支持事务（多文档、跨分片多文档）
client/SDK交互，可通过插件转译支持弱SQL
...

）Redis 
数据结构丰富
纯C实现 高性能
基于内存，支持AOF/RDB持久化
常用redis-cli/多语言SDK交互


## RDBMS

SQL引擎 Parser Optimizer Executor

解析器 词法分析 语法分析 语义分析（检查是否合理）
优化器 ）基于规则的优化 RBO ） scan优化 唯一索引 普通索引 全表扫描
）基于代价的优化 CBO cost-- 时间_资源_

Executor 火山模型 --> 向量化Batch利用CPU的simd机制  /  编译执行 LLVM动态编译执行技术

存储引擎 

- InnoDB
- Buffer Pool
- Page
- B+ Tree 二分查找的扩展 从根到叶通过双向链表链接 二分法定位槽－遍历槽

事务引擎

- Atomicity 与 Undo Log
Undo Log: 逻辑日志,记录增量变化。可以进行数据回滚

- Isolation与锁

读 sharelock 写 exclusivelock
MVCC:读写互不堵塞，降低死锁概率，实现一致性读

- Durability 与 Redo Log

对数据修改永久保存
REDOLOG 物理日志 记录页面变化 事务提交前日志写盘


实践案例 
大流量 -- sharding
业务层数据水平拆分，代理层分片路由
线性扩展写入性能和容量

流量激增
-- 扩容
扩容db物理节点数量
影子表压测

-- 代理连接池 

稳定性/可靠性 -- 3Az高可用
proxy 读写分流 /限流 + 监控报警

-- HA管理
监控热切换宕机

## 对象存储 TOS

冷热数据分级存储

拓-----HDFS
弱posix系统语义
目录/文件  append写   无法直接HTTP访问
拓-------

对象存储云端 bucket-object   （object：key data metadata）
http协议访问  适于静态、immutable 结构化和mutable不适合

基本逻辑：申请bucket->业务逻辑开发->上线测试
RESTFUL接口  
Multiupload接口  initupload（id）multiupload  complete
Listprefix接口 分页列举

# 工程实践

接入层 ））接入解析并处理接口请求） 存储引擎层））存储对象内容） 元信息层））存储对象元信息）
容量型-片源 转码 ；qps型-抽帧

可扩展性解法 -- Partition
存储 计算 压力 均匀分布 -分布式 （分布式+ 单机存储）
不同数据映射到不同Partition分区
Partition Logic: Hash/Range
扩容新建partition 新增数据写入映射导向新partition
持久度解法 -- Replication
复制多副本 多机房多region
-- EC存储 冗余编码
-- 温冷转换
高可用 -- 集群拆分 （降低爆炸半径）
-- 镜像灾备
