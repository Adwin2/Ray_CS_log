### Database/SQL与GORM的设计与实践

## 理解database/SQL

## GORM使用简介

## GORM设计原理
每一次调用都生成一个SQL子句（clause）
语句最后 Finishermethod  statement (决定类型)
-->	callback --> 生成 SQL语句
对插件函数的调用
注册删除替换 查询指定顺序 注册时检查条件新（自定义）callback函数
灵活定义，自由扩展 (插件系统)
  - 多租户
  绑定租户ID
  - 多数据库，读写分离
  分为 sources 与 replicas

## connpool
  - 全局模式下所有DB操作都会被预编译（cache不包含参数部分）
  - 会话模式 后续会话操作都会预编译并缓存
  - 全局缓存的语句可被会话使用
查找预编译语句 -- 未找到，将收到的SQL和Vars预编译  -- 使用缓存的预编译SQL执行
MySQLdriver interpolateParams=false

bytedgorm   -- Dialector
  

## GORM最佳实践
数据序列化与SQL表达式
# SQL表达式更新创建 
- 通过gorm.Expr使用表达式
- 使用GORMValuer使用SQL表达式/SubQuery

批量数据操作
- 批量操作 查询 更新
- 批量数据操作加速 关闭默认事务 SkipHooks 使用 Prepared Statement
代码复用、分库分表、sharding
- 代码复用 Scopes方法  
- 分库分表  sharding 抽象逻辑
混沌工程/压测
 
Logger/Trace
- 全局模式  会话模式
Migrator
- 数据库迁移管理 AutoMigrate()  版本管理 Rollback
Gen代码生成/RAW SQL
- db.Raw()  db.Exec()
- Query
安全
- 避免SQL注入（拼接）

