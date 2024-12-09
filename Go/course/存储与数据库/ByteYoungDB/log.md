2024年 12月 08日 星期日 13:30:22 CST

目前已知 sql/CreateDefinition.h逻辑更新
86rows
导致.so与代码不对应
需要依次更新代码


12月 09日

src/metadata.cpp 行16 `vector --> unordered_set`
编译完成

－－－－－－－－－－
- [ ]模板元编程部分 代码
