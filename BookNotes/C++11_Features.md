# C++11 特性 

> from 《C++ primier》

1. long long 类型

2. 列表初始化

3. nullptr常量

4. constexpr 变量、 constexpr函数、 constexpr构造函数   ；必须返回字面值   隐式指定为内联函数   其构造函数应该是空的 必须初始化所有数据成员

5. typedef  类型别名声明

6. auto 类型指示符

7. decltype 类型指示符

8. 类内初始化（默认初始值）

9. 使用auto 或 decltype 缩写类型   简化声明  简化返回类型定义

10. 范围for语句   `for (declaraion : expression)   \statement`

11. 定义  vector<vector<_Type>>

12. vector对象的列表初始化

13. 容器的cbegin和cend函数    cbegin( )  cend( )   const_iterator类型 不用操作数据时建议

14. 标准库的 begin 和 end 函数  begin( )  end( ) 返回 迭代器  位置同理

15. 除法的舍入规则     趋于0

16. 大括号包围的值列表赋值     

17. 将sizeof用于类成员

18. 标准库 initializer_list 类   成员不可改变  且类型都相同  可以制作error_list

19. 列表初始化返回值    return { , , };

20. 定义尾置返回类型  `auto func ( int i )  ->  int(*)[10];`

21. 使用 = default 生成默认构造函数  编译器生成默认构造函数 类内部内联

22. 类对象成员的类内初始化  将默认值声明为类内初始值  _Type name{...};

23. 委托构造函数   `func(param.s) : func_dest(param.s) { };`

24. 用 string 对象处理文件名    string ifile = “”; ifstram in(ifile);   /旧：const char* 

25. array 和 forward_list 容器    array 大小固定    for_ward容器：单向链表数据结构 （无.size()操作）

26. 容器 （包括关联容器）的列表初始化

27. pair的列表{}初始化 及其 返回类型   ：返回类型->  对返回值进行列表初始化

    ```c++
    pair<string, int> 
    process(vector<string> &v)
    {
    	if(!v.empty())
            return (v.back(),v.back().size()); 
        else 
            return pair<string, int>(); //隐式构造返回值
    }
    ```

    

28. 容器的非成员函数 swap   :非成员版本swap函数在泛型编程中是非常重要的 统一使用是好习惯

29. 容器insert成员的返回类型 ： 指向第一个新加入元素的迭代器

30. 容器emplace成员的返回类型   emplace_front、 emplace_back、  emplace  调用构造函数 可以直接传入参数 少一次拷贝

31. shrink_to_fit   ：并不一定保证可以退回内存空间

32. string 的数值转换函数     ：to_string() 、 stoi()、stol()、stoul()、stoll()、stoull()、stof()、stod()、stold()    重载和返回值有所区别

33. Lambda 表达式      ：` [ capture list ] (parameter list)  ->  return_type {function body} ` 理解为未命名的内联函数    \    auto f = [ ] { return 42; };   f()；//返回42  

34. Lambda 表达式中的尾置返回类型  ： 见33

35. 标准库 bind 函数    `#include <functional>`  ：看作通用函数适配器  `auto newT = bind(T,args_list);` 调用1会调用传入3的2 

36. 无序容器   unordered_map 、 unordered_set  ：性能较好  使用哈希函数和关键字类型的==运算符组织元素

37. 智能指针         shared_ptr\unique_ptr\weak_ptr  `#include <memory>` ：负责自动释放指向的对象  管理底层指针的方式不同 

38. 动态分配对象 和动态分配数组 的列表初始化   

    ```C++
    string *ps = new string(10, '9');
    vector<int> *pv = new vector<int>{0, 1, 2, 3};  //same as int[4]{...}
    ```

39. auto 和 动态分配

    ```C++
    auto p1 = new auto(obj); //该对象用obj进行初始化  只能有单个初始化器
    ```

40. 范围for语句不能应用于动态分配数组 :as it says

41. auto 不能用于分配数组                       :as it says

42. allocator::construct 可使用任意构造函数  ：包含于allocator 类及其算法 扩展了 construct的额外参数

43. 将 = default 用于拷贝控制成员（拷贝构造函数、拷贝赋值运算符） ：显式的要求编译器生成合成**的版本  且将隐式地声明为内联的 （只对类外拷贝运算符操作则取消内联）//可以合成的只有默认..和拷贝..

44. 使用 = delete 阻止拷贝类对象  ： 

    ```C++
    //...（类内）
    Nocopy (const Nocopy&) = delete;//阻止拷贝
    Nocopy &operator=(const Nocopy&) = delete; //阻止赋值
    ```

45. 用 移动类对象 代替 拷贝类对象  ：as it says

46. 右值引用：&&

47. 标准库move函数  ：`#include <utility>  \std::move` 

48. 移动构造函数 和 移动赋值

49. 移动构造函数通常应该是 noexcept  : as it says    ，承诺函数不抛出异常

50. 移动迭代器   ：解引用运算符 生成一个右值引用     `std::make_move_iterator(_iterator);`  返回一个移动迭代器  注意适用性问题

51. 引用限定成员函数      参数列表后放置一个引用限定符 & 

    > 可以是&或&& 分别指出this可以指向一个左值或右值

    ```c++
    //...in a class
    Foo &operator=(const Foo&) &;
    //out of class
    Foo &Foo::operator= (const Foo &rhs) &
    {
        //执行将rhs赋予本对象的操作
        return *this;
    }
    ```

52. function类模版  ：`#include <functional> \ function<T> f;`  

53. explicit 类型转换运算符   ：显式的类型转换运算符explicit  和   显式的构造函数 static_cast<_Type>(p);  

54. 虚函数的override 和 final 指示符  ：

    - override-> 在参数列表(、const、&)后显式的指出 覆盖了继承的虚函数 的成员函数 标记虚函数后 如果并没有覆盖 则编译器报错 ；

    - final->  置于类名称后 使其不能作为基类  或  写在同override的位置 ---不可覆盖

55. 通过定义类为final来阻止继承  : as it says

56. 删除的拷贝控制和继承   ：基类和子类的关系  例：基类析构unaccess 子类不可移动构造

57. 继承的构造函数    using 基类::func;   //继承基类的构造函数

58. 声明模版类型形参为友元  

    ```C++
    template <typename Type> class Bar {
    	friend Type; //将访问权限授予用来实例化Bar的类型
        		//...
    }
    ```

    Foo将成为`Bar<Foo>`的友元

59. 模板类型别名

    ```c++
    template <typename T> using twin = pair<T, T>;
    twin<string> authors;   // authors是一个pair<string, string>
    ```

60. 模板函数的默认模板参数 

    ```C++
    //compare 有一个默认实参less<T> 和一个默认函数实参F()
    template <typename T, typename F = less<T>>
    int compare(const T &v1, const T &v2, F f = F())
    {
    	if(f(v1, v2)) return -1;
        if(f(v2, v1)) return 1;
    	return 0;
    }
    ```

61. 实例化的显式控制    ：模板被使用时才会实例化，大系统中 多个文件中实例化相同模板的额外开销可能非常严重   使用**显式实例化**来避免这种开销

    ```c++
    //实例化声明与定义
    extern template _daclaration;  // 实例化声明  早于首个定义或调用
    template declaration  //实例化定义
    //-- 实现
    extern template class Blob<string>;  //声明
    template int compare(const int&, const int&); //定义
    ```

62. 模板函数与尾置返回类型   ：由于尾置返回出现在参数列表之后， 可以使用函数的参数 常搭配decltype() 使用

63. 引用折叠规则

    - 如果我们间接创建一个引用的引用，则这些引用形成了折叠。引用通常会折叠为左值引用。新标准扩展到右值引用（只在右值引用的右值引用情况下`X&& && ----> X&&`）；只能用于间接创建的引用的引用，如类型别名或模板参数。

64. 使用static_cast强制类型转换将左值转换为右值  是被允许的  :as it says

65. 标准库 forward 函数 : `#include <utility> ` 返回显式实参类型的右值引用 即 forward<T>返回类型是 T&&  必须通过显式模板实参来调用   见66

    - 如果一个函数参数指向模板类型参数的右值引用(T&&)，它对应的实参的const属性和左右值属性将得到保持  (缺点：不能从一个左值实例化 int&&)（函数参数是左值表达式）

    - 当用于一个指向模板参数类型的右值引用函数参数（T&&）时，std::forward<>()会保持实参类型的所有细节 （指const、 左、右值等）

    - ```C++
      // 例：实现函数 (包括参数) 的 转发
      template <typename F, typename T1, typename T2>
      void flip(F f, T1 &&t1, T2 &&t2) {
          f(std::forward<T2>(t2), std::forward<T1>(t1));
      }
      void g(int &&i, int& j) {
      	cout << i << " " << j << endl;
      }
      //调用flip(g,i,42); i以int&类型、42以int&&类型传给g
      ```

66. 可变参数模板    （ 66 - 68 ）

    - 可变参数模板   ：接受可变数目的参数的模板函数或模板类。可变数目的参数称为参数包。存在两种参数包：模板参数包 表示零或多个模板类型参数；函数参数包 表示零或多个函数参数。  用省略号(...)来指出一个模板参数或函数参数表示一个包。

      ```C++
      template <typename T, typename... Args>//Args模板参数包 
      void foo(const T &t, const Args& ... rest)//rest函数参数包 
      {
          std::cout<<sizeof...(Args)<<std::endl;//rest同理
      }
      //编译器通过函数的实参推断模板参数类型 还会推断包中参数的数目 
      //编译器会实例化出对应数量的不同版本
      ```

67. sizeof... 运算符   

    - 但我们需要知道包中有多少元素时，可以使用sizeof...运算符。返回一个常量表达式，而且不会对其实参求值。用法与sizeof()函数基本相同。

68. 可变参数模板 与 转发

    - 标准库容器的emolace_back是一个可变参数成员模板，用实参在内存空间中直接构造一个元素

    - 如果希望能使用string的移动构造函数的话， 还需要保持传递的实参的所有类型信息，步骤：

    - 1.将emplace_back函数参数定义为模板类型参数的右值引用；

      2.emplace_back传递实参给construct时使用std::forward来保持实参的原始类型

      ```C++
      template <class... Args>
      inline
      void StrVec::emplace_back(Args&&... args) {
      	chk_n_alloc(); //如果需要的话重新分配StrVec内存空间,确保有足够空间容纳新元素
          alloc.construct(first_free++, std::forward<Args>(args)...); //如果emplace_back接受了右值实参，construct得到完全相同且继续将此实参传递给string的移动构造函数来创建新元素
      }
      ```

69. 标准库 Tuple 类模板 ：可以有任意数量的成员类型

    ```C++
    tuple<T1, T2, .., Tn> t; //值初始化
    tuple<T1, T2, .., Tn> t(v1, v2, .., vn);//explicit构造函数
    make_tuple(v1,v2,..,vn);
     == / !=
    t1 relop t2   //tuple的关系运算符使用字典序，俩必须有相同数量成员，使用<比较对应成员
    get<i>(t); //返回t的第i个数据成员的引用 左右值类型相对应  ；tuple所有成员都是public的
    //类模板
    tuple_size<tupleType>::value  //public constexpr ststic数据成员 size_t类型 表示给定类型中成员的数量
    tuple_element<i,tupleType>::type  //表示给定tuple类型中指定成员的类型
    ```

    > “快速且随意的数据结构”

70. 新的bitset运算

    - all操作 ， 所有位置位时（ = 1）返回true    拓：any、none

71. ==正则表达式库== ：描述字符序列的方法、强大的计算工具 RE库 `#include <regex>`  Page 646

72. 随机数库   `#include <random>`  组成：1.引擎：类型，生成随机unsigned整数序列；2.分布：类型，使用引擎返回服从特定概率分布的随机数。  两者对象的结合可以叫做 随机数发生器

    - C++程序应该使用default_random_engine类和恰当的分布类对象。

    ```C++
    uniform_int_distribution <unsigned>u(0,9); //此类型决定会生成均匀分布的unsigned值
    default_random_engine e; //生成无符号随机整数
    for (size_t i = 0;i < 10; ++i)
        std::cout << u(e) << std::endl;
    //生成随机的数值序列
    vector<unsigned> good_randVec()
    {
        static default_random_engine e; // !--static
        static uniform_int_distribution <unsigned> u(0,9);
        vector<unsigned> ret;
        for(size_t i = 0; i < 100; ++i)
            ret.push_back(u(e));
        return ret;
    }
    ```

73. 浮点数格式控制

    - hexfloat 强制浮点数使用十六进制格式   defaultfloat 将流恢复到默认状态 根据要打印的值选择计数法       拓：uppercase操作符  输出补白{setw()  left  right  internal  setfill() }

74. noexcept 异常指示符   \    运算符 ： noexcept说明  紧跟在函数的参数列表后面  指定某个函数不会抛出异常。  “做了不抛出说明”  ；一元运算符：返回bool类型的右值常量表达式。

75. 内联命名空间   ：中的名字可以被外层命名空间直接使用。无需添加命名空间前缀。通过外层命名空间名字直接访问。

    - namespace 关键词 前添加 inline 关键字 第二次打开就是隐式内联了
    - 拓：未命名的命名空间的作用范围不会横跨多个不同文件

76. 继承的构造函数和多重继承

    - 格式 `using Base1::Base1;` 如果一个类 从它的多个基类中继承了相同的构造函数，则这个类必须为该构造函数定义他自己的版本  注意出现构造函数必须有`D() = default;`

77. 有作用域的enum   ：enum枚举类型 将整型常量组织在一起 属于 字面值 常量 类型

    - 1.限定作用域的枚举类型

    ```C++
    enum class open_modes {input, output, append}; //关键字class 或 struct
    ```

    - 2.不限定作用域的枚举类型 ： 枚举成员不可重复定义

    ```C++
    enum color {red, yellow, green};   //不限定作用于的枚举类型  省略掉关键字
    //未命名的、不限定作用域的枚举类型  未命名：只能在定义enum时定义它的对象
    enum {floatPrec = 6, doublePrec = 10, double_doublePrec = 10}/*声明列表可放此处(声明 和 定义分开)*/;
    //注意 作用域前缀（使用$的$\显式访问枚举成员）
    ```

78. 说明类型用于保存enum对象

    - 默认情况下 限定作用域的enum成员类型是int ；**不限定作用域**的不存在默认类型，只知道成员的潜在类型足够大，能够容纳枚举值。制定潜在类型后超出范围将引发程序错误。

    ```C++
    enum intValue : unsigned long long {
    	charTyp = 255, shortTyp = 65535, intTyp = 65535, 
        longTyp = 4294967295UL,
        long_longTyp = 18446744073709551615ULL
    };
    ```

79. enum的提前声明：必须指定成员类型（显式（不限作用域）或隐式（限作用域->int））；声明和定义保持一致 不可改变

80. 标准库的mem_fn类模板 `#include <functional>` ：让编译器负责推断成员的类型 

    - 可以从成员指针生成一个可调用对象 编译器由前者的类型推断后者的类型 无需显式指定

    ```C++
    auto f = mem_fn(&string::empty); //f 接受一个 string 或 一个string*
    f(*svec.begin());//传入一个对象 使用 .*调用empty
    f(&svec[0]);     //传入一个指针 使用 ->*调用empty
    ```

81. 类类型的联合成员 ：早期在union中不能含有定义了构造函数或拷贝控制成员的类类型成员，C++11取消了这一限制

    





 