# 《Effective  C++  第三版 》

## 一、让自己习惯C++

1. 视C++为一个语言联邦

   - 传统面向过程：区块、语句、预处理器、内置数据类型、数组、指针
   - 面向对象（C++ with class）类、封装、继承、多态、动态绑定
   - 模板编程Template C++ （TMP）
   - C++ 标准库STL...

2. 尽量以const、enum、inline替换#define

   ```C++
   //C++11后推荐 对于变量
   constexpr auto name = value;
   //enum可以用于替代整型的常量 在模板元编程中应用广泛
   enum{num = 6};
   //使用内联模板函数替代
   template <typename T>
   inline void CallWithMax(const T& a,const T& b) {
       f(a > b? a : b);
   }
   ```

3. 尽可能地使用const

   - 只读常量 标明const

   - 注意const 还是 const_iterator

     ```C++
     const std::vector<int>::iterator it = v.begin();//迭代器不可修改，数据可修改
     std::vector<int>::const_iterator it = v.begin();//迭代器可修改，数据不可修改 
     ```

   - 对于不想要无意义地被当作左值的函数 返回const类型

   - ```C++
     class TextBlock {
     public:
         const char& operator[](std::size_t position) const {
             // 假设这里有非常多的代码
             return text[position];
         }
         char& operator[](std::size_t position) {
             return const_cast<char&>(static_cast<const TextBlock&>(*this)[position]);//为了减少重复代码的转型是适当的
         }
     private:
         std::string text;
     };
     ```

4. 确定对象在使用前已被初始化   定义完对象后一定要尽快赋初值

   ```C++
   //：对于类中的成员变量
   class a{
     private:
       std::size_t TextLength{ 0 };
       bool LengthIsValid{false};
   };/*Since C++11  定义处赋值
   	另一种使用构造函数 成员初始化列表  （亦可括号空着执行默认构造函数）*/
   //类中变量的初始化始终与声明次序一致    （成员变量初始化可选，但引用类型必须初始化）
   ```

- 静态对象的初始化

  ```C++
  //Meyers' Singleton 单例模式的变种 是一种较线程安全的单例模型   --局部静态对象
  FileSystem& tfs(){
      ststic FileSystem fs;
      return fs;
  }
  ```

## 二、构造、析构、赋值运算

1. 了解C++默默编写并调用哪些函数

   - 编译器为C++空类预留的内容

   ```C++
   class Empty {
   public:
       Empty() {  }
       Empty(const Empty&){  }
   	Empty(Empty&& ) {  }
       ~Empty() {  }
       
       Empty& operator= (const Empty&) { }
   	Empty& operator= (Empty&& ) { }
   };
   //注: 类中有const成员 、 基类有private的拷贝赋值运算符都不会自动创建默认的拷贝赋值运算符
   ```

2. 若不想使用编译器自动生成的函数，就应该明确拒绝      //使用   =delete （since C++11）

3. 为多态基类声明虚析构函数

   - problem：当派生类对象经由一个基类指针被删除，而该基类指针带着一个非虚析构函数，其结果是未定义的，可能会无法完全销毁派生类的成员，造成内存泄漏。

   ```C++
   class Base{
     public:
       Base() { };
       virtual ~Base() { };
   };
   //额外存储的虚表指针会使类的体积变大
   ```

   - 把基类作为抽象类

   ```C++
   class Base{
     public:
       virtual ~Base = 0 {}; // 纯虚函数
   }
   ```

4. 别让异常逃离析构函数

   - 为了实现RAII 通常将对象的销毁方法封装在析构函数中。 在析构函数中处理异常的常见做法：

   ```C++
   // 1. 杀死程序
   DBConn::~DBConn() {
   	try{db.close();}
   	catch () {
           //记录运行日志，以便调试
           std::abort();
       }
   }
   //2.重新设计接口，异常处理交给客户端去完成
   class DBConn {
   public:
       //...
       void close() {
           db.close();
           closed = true;
       }
       ~DBConn() {
           if (!closed) {
               try {
                   db.close();
               }
               catch(...) {
                   // 处理异常
               }
           }
       }
   private:
       DBConnection db;
       bool closed;
   };//当一个操作可能会抛出需要客户处理的异常时，暴露在普通函数中而非析构函数中是一个更好的选择
   ```

5. 绝不在构造和析构过程中调用虚函数   （避免间接调用虚函数，难以发现的危险行为）拓：静态成员函数确保不会使用未完成初始化的成员变量

6. 令operator=返回一个指向*this的引用   &类型 返回指针  适用于+= -= *= 为了实现连锁赋值

7. 在operator=处理自我赋值

   - 1、Identity Test ：`if(this == &rhs) return *this;`
   - 2、只关注异常安全性

   ```C++
   Widget& operator=(const Widget& rhs) {
       Resource* pOrigin = pRes; 
       pRes = new Resource(*rhs.pRes);
       delete pOrigin;
       return *this;
   }
   ```

   - 3、Copy and Swap 通过析构函数来实现资源的释放  利用了栈空间会自动释放

   ```C++
   Widget& operator= (const Widget& rhs) {
       Widget tmp(rhs);
       std::swap(*this, tmp);
       return *this;
   }
   //或按值传参时
   Widget& operator= (Widget rhs) {
       std::swap(*this, rhs);
       return *this;
   }
   ```

8. 复制对象时勿忘其每一个成员

## 三、资源管理

1. 以对象管理资源

   - RAII : Resource Acquisition is initialize 资源获取时即是初始化时  析构函数负责资源的释放
   - since C++11 专一所有权：std::unique_ptr  引用计数来管理：std::shared_ptr 

   ```C++
   // Investment* CreateInvestment();
   
   std::unique_ptr<Investment>pUnique_01(CreateInvestment());
   std::unique_ptr<Investment>pUnique_02(std::move(pUnique_01));  //转移资源所有权
   
   std::shared_ptr<Investment>pShared_01(CreateInvestment());
   std::shared_ptr<Investment>pShared_02(pShared_01);        //引用计数 +1
   //智能指针默认可以自动delete所持有的对象  也可以自行管理释放方式 （删除器deleter）
   //void GetRidOfInvestment(Investment*) {}
   
   std::unique_ptr<Investment,decltype(GetRidOfInvestment)*>pUniqueInv(CreateInvestment(), GetRidOfInvestment);
   std::shared_ptr<Investment>pSharedInv(CreateInvestment(),GetRidOfInvestment);
   ```

2. 在资源管理类(RAII对象)中小心拷贝行为

   - 许多时候复制RAII对象不合理，应明确禁止（= delete）
   - 对底层资源 使用引用计数法（std::shared_ptr）每一次复制引用计数+1，每一个离开定义域的对象调用析构函数 引用计数-1，直到引用计数为0时就彻底销毁资源
   - 复制底层资源  ：复制对象又复制底层资源—深拷贝。
   - 转移底层资源的所有权（std::unique_ptr）

3. 在资源管理类中提供对原始资源的访问

   - STL中的智能指针提供了对原始资源的访问

   ```C++
   Investment* pRaw = pSharedInv.get();    // 显式访问原始资源
   Investment raw = *pSharedInv;           // 隐式访问原始资源
   //...
   class Font {
   public:
       FontHandle Get() const { return handle; }       // 显式转换函数
       operator FontHandle() const { return handle; }  // 隐式转换函数
   
   private:
       FontHandle handle;
   };
   //显示转换比较安全，隐式转换对客户比较方便
   ```

4. 成对使用new和delete时要采用相同的形式  注：使用typedef定义数组类型typedef std::string str[4]也需要delete[]，会带来额外的风险

5. 以独立语句将newed对象置入智能指针  _原书方法过时_

   ```C++
   auto pUniqueInv = std::make_unique<Investment>();   // since C++14
   auto pSharedInv = std::make_shared<Investment>();   // since C++11
   ```

## 四、设计与声明

1. 让接口容易被正确使用，不易被误用

   - “容易被正确使用”：接口的一致性、与内置类型的行为兼容

   - “阻止误用”：建立新类型、限制类型上的操作、束缚对象值、消除客户的资源管理责任

     ```C++
     // 三个参数类型相同的函数容易造成误用
     Data::Data(int month, int day, int year) { ... }
     
     // 通过适当定义新的类型加以限制，降低误用的可能性
     Data::Data(const Month& m, const Day& d, const Year& y) { ... }
     ```

2. 设计class犹如设计type

   - 如何创建和销毁、初始化和赋值是否有差别、新type的合法值、继承图系、合理的运算符和函数、哪个函数应该=delete、一般化 模板类

3. 宁以按**常引用**传参替换按值传参

   - 当使用按**值**传参时，程序会调用对象的拷贝构造函数**构建**一个在函数内作用的**局部对象**，这个过程的开销可能会较为昂贵。对于任何用户自定义类型，使用按常引用传参是较为推荐的：

     ```C++
     bool ValidateStudent(const Student& s);
     ```

     没有构建新对象、效率比高很多

   - 常引用传参也可以避免参数切片的问题：原理还是没有创建新对象 继承方面 会调用对应的虚函数
   - 对于内置类型、STL的迭代器、函数对象使用按**值**传递是比较合适的

4. 必须返回对象时，别妄想返回其引用  返回局部对象的引用是严重的错误 返回一个指向局部静态变量的引用也是不推荐的 tho返回对象调用拷贝构造函数会有overhead

5. 将成员变量声明为private

   - 出于对封装行的考虑  隐藏类中的成员变量，并通过对外暴露函数的接口来实现对成员变量的访问

   ```C++
   class AccessLevels {
   public:
       //所谓的Getter和Setter函数
       int GetReadOnly() const { return readOnly; }
       void SetReadWrite(int value) { readWrite = value; }
       int GetReadWrite() const { return readWrite; }
       void SetWriteOnly(int value) { writeOnly = value; }
   
   private:
       int noAccess;
       int readOnly;
       int readWrite;
       int writeOnly;
   };
   ```

6. 宁以非成员、非友元函数替换成员函数

   - 为了遵守一个原则：可以访问数据的代码越少，其封装性越强   置于全局函数 even命名空间中

7. 若所有参数皆需类型转换，请为此采用非成员函数

   - 运算符重载置于类外 作为非成员函数

8. 考虑写一个不抛异常的swap函数

   - 由于`std::swap`函数在 C++11 后改为了用`std::move`实现，因此几乎已经没有性能的缺陷，也不再有像原书中所说的为自定义类型去自己实现的必要。

   - 拓：

     C++ 名称查找法则：编译器会从使用名字的地方开始向上查找，由内向外查找各级作用域（命名空间）直到全局作用域（命名空间），找到同名的声明即停止，若最终没找到则报错。
     函数匹配优先级：普通函数 > 特化函数 > 模板函数

## 五、实现

1. 尽可能延后变量定义式出现的时间    ：默认构造+赋值 效率低于 直接构造

   - 特殊：循环中变量的定义 A.定义于循环外，在循环中赋值；B.定义于循环内
   - 由于做法A会将变量的_作用域扩大_，因此除非知道该变量的赋值成本比“构造+析构”成本低，或者对这段程序的效率要求非常高，否则建议使用做法B

2. 少做转型动作

   ```C++
   //C++式转型
   const_cast<T>(expr);//用于常量型转除  （常量型转除可以减少内存使用，加速程序运行）
   dynamic_cast<T>(expr);//用于安全地向下转型 会执行对继承体系的检查，因此会带来额外的开销。只有拥有虚函数的基类指针能进行dynamic_cast。 也是唯一C替代不了的转型操作
   reinterpret_cast<T>(expr);//用于在任意两个类型间进行低级转型，执行该转型可能会带来风险，也可能不具备移植性。
   static_cast<T>(expr);//用于进行强制隐式转换，也是最常用的转型操作，可以将内置数据类型互相转换，也可以将void*和typed指针，基类指针和派生类指针互相转换。
   ```

   > 当你想知道一个基类指针是否指向一个派生类对象时，你需要用到`dynamic_cast`，如果不满足，则会产生报错。但是对于继承体系的检查可能是非常慢的，所以在注重效率的程序中应当避免使用`dynamic_cast`，改用`static_cast`或别的代替方法。

   ```C++
   class SpecialWindow : public Window {
   public:
       virtual void OnResize() {
           Window::OnResize();//不要试图转型*this 那只会在生成的副本上调用函数
           ...
       }
       ...
   };
   ```

3. 避免返回handles指向对象的内部成分   ：返回成员变量的副本     封装性

   - 避免返回 handles（包括引用、指针、迭代器）指向对象内部。遵循这个条款可增加封装性，使得const成员函数的行为符合常量性，并将发生 “空悬句柄” 的可能性降到最低。

     > 空悬句柄 ：对象销毁后无法通过引用获得返回的数据

4. 为“异常安全”而努力是值得的

   - 当异常被抛出时，带有异常安全性的函数会：不泄漏任何资源。不允许数据败坏。
   - 例：通过以对象管理资源，使用智能指针和调换代码顺序，我们能将其变成一个具有强烈保证的异常安全函数：

   ```C++
   class PrettyMenu {
   public:
       ...
       void ChangeBackground(std::vector<uint8_t>& imgSrc);
       ...
   private:
       Mutex mutex;        // 互斥锁
       Image* bgImage;     // 目前的背景图像
       int imageChanges;   // 背景图像被改变的次数
   };
   /* 错误示范   若抛出异常 则资源泄露
   void PrettyMenu::ChangeBackground(std::vector<uint8_t>& imgSrc) {
       lock(&mutex);
       delete bgImage;
       ++imageChanges;
       bgImage = new Image(imgSrc);
       unlock(&mutex);
   }  */
   
   //通过以对象管理资源，使用智能指针和调换代码顺序--->  (1)
   void PrettyMenu::ChangeBackground(std::vector<uint8_t>& imgSrc) {
       Lock m1(&mutex);
       bgImage.reset(std::make_shared<Image>(imgSrc));
       ++imageChanges;
   }
   //或者Copy and Swap方法----->  (2)
   void PrettyMenu::ChangeBackground(std::vector<uint8_t>& imgSrc) {
       Lock m1(&mutex);
       auto pNew = std::make_shared<PMImpl>(*pImpl);    // 获取副本
       pNew->bgImage.reset(std::make_shared<Image>(imgSrc));
       ++pNew->imageChanges;
       std::swap(pImpl, pNew);
   }
   ```

5. 透彻了解inline的里里外外

   - 将函数声明为内联一共有两种方法，一种是为其显式指定`inline`关键字，另一种是直接将成员函数的定义式写在类中（隐式声明为内联）。

   - C++17中引入： 静态成员变量可以在类中直接定义

     ```C++
     class Person {
     public:
         ...
     private:
         static inline int theAge = 0;  // since C++17
     };
     ```

     

6. 将文件间的编译依存关系降至最低  --31

## 六、继承与面向对象设计

1. 确定你的public继承塑模出is-a关系
2. 避免遮掩继承而来的名称
3. 区分接口继承和实现继承
4. 考虑虚函数以外的其他选择
5. 绝不重新定义继承而来的非虚函数
6. 绝不重新定义继承而来的缺省参数值
7. 通过复合塑模出 has-a 或“根据某物实现出”
8. 明智而审慎地使用private继承
9. 明智而审慎地使用多重继承

## 七、模板与泛型编程

1. 了解隐式接口和编译期多态
2. 了解typename的双重含义
3. 学习处理模板化基类内的名称
4. 将与参数无关的代码抽离模板
5. 运用成员函数模板接受所有的兼容类型
6. 需要类型转换时请为模板定义非成员函数
7. 请用traits classes表现类型信息
8. 认识模板元编程

## 八、定制new和delete

1. 了解new-handler的行为
2. 了解new和delete的合理替换时机
3. 编写new和delete时需固守常规
4. 为了placement new 也要写placement delete

## 九、杂项 

1. 不要轻忽编译器的警告
2. 让自己熟悉包括TR1在内的标准程序库
3. 让自己熟悉Boost