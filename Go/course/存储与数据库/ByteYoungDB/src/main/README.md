Parser: 解析用户输入的SQL语句，并进行合法性校验
MetaData: 管理用户创建的表、列表元信息
Optimizer: 根据SQL语法树生成计划树
Executor: 根据计划树进行查询执行（此处采用火山模型）
Storage: 内存态数据结构设计
Transection: Commit/Rollback事务支持


main.cpp 创建一个命令行交互界面

使用的SQL解析库hyrise/sql-parser只进行了SQL语句的词法分析和语法分析 ，合法性检验（语义分析）需要自己实现
> 对源库封装成include/和lib/，增加语义分析功能
