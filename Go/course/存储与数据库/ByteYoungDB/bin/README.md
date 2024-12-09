[ByteYoungDB](https://bytedance.larkoffice.com/docx/doxcnkFlL3GLJoIgJaG4pKX2JVe)
⬆️  ⬆️  ⬆️

SQL
13
>> show columns db.t;​
14
# Columns in db.t:​
15
a        INT​
39
>> begin;​
40
[BYDB-Info]  Start transaction​
41
​
42
>> update db.t set b = 'apple' where a = 1;​
43
[BYDB-Info]  Update 1 tuple successfully.​
44
​
45
>> select * from db.t where a = 1;​
46
            a           b​
47
-------------------------​
48            
1       apple​
49
-------------------------​
50
1 row​
51
​
52
>> rollback;​
53
[BYDB-Info]  Rollback transaction​
54
​
55
>> select * from db.t where a = 1;​
56
            a           b​
57
-------------------------​
58            
1       first​
59
-------------------------​
60
1 row​
61
​
62
>> begin;​
63
[BYDB-Info]  Start transaction​
64
​
65
>> update db.t set b = 'apple' where a = 1;​
66
[BYDB-Info]  Update 1 tuple successfully.​
67
​
68
>> commit;​
69
[BYDB-Info]  Commit transaction​
70
​
71
>> select * from db.t where a = 1;​
72
            a           b​
73
-------------------------​
74            
1       apple​
75
-------------------------​
76
1 row​
77
​
78
>> exit​
79
# Farewell~~~​
​
￼
​


**扩展演进**​
￼
大家可以在当前项目的基础上继续演进，有如下几个方向：​
1.
实现B+Tree、Hash索引​
2.
实现count()、sum()、min()、max()等简单函数​
3.
实现group by操作​
4.
实现两表join操作​
5.
实现基于磁盘文件的存储引擎，以及数据的持久化


> cloc 工具实现代码统计
