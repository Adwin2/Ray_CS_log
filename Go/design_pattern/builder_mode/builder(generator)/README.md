# 生成器模式

分步骤创建复杂对象

该demo里的House 拆分为了 Frame Style Door Bed 四个属性 ，
所有建造方法由Build接口实现，
由GetBuild(`TypeName`)获取对应类型的建造函数
manager指定建造(Build)类型
