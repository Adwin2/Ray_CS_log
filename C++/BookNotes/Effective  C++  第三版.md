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

   1）用声明的依赖性替换定义的依存性

   2）尽量使用对象引用或指针 ——符合3）

   3）尽量使用类型声明式而不是定义式

   4）为声明式和定义式提供不同的头文件——接口和其实现分离：接口类和句柄类

   >增加一层间接访问（开销），当这些开销过于重大以至于类之间的耦合度在相形之下不成为关键时，就以具象类（concrete class）替换句柄类和接口类。

## 六、继承与面向对象设计

1. 确定你的public继承塑模出is-a关系

   >is-a关系：适用于基类的每一件事情也适用于继承类（可以认为每一个派生类都是一个基类对象）

2. 避免遮掩继承而来的名称

   1）using关键字 --将所有版本包含在内

   2）转发函数 --只想要单一版本

   ```C++
   class Base {
   public:
       virtual void mf();
       virtual void mf(double);
   };
   
   class Derived : public Base {
   public:
       virtual void mf() {
           Base::mf();
       }
   };
   ```

3. 区分接口继承和实现继承

   1）public继承下，派生类总是继承基类的接口。

   2）声明一个纯虚函数，为了让派生类只继承该函数接口。

   3）声明简朴的非纯虚函数的目的，是让派生类继承该函数的接口和缺省实现。

   4）声明非虚函数的目的，为了令派生类继承函数接口以及一份强制性实现。

4. 考虑虚函数以外的其他选择

   1）非虚接口设计手法 ( NVI ) 使用一个非虚函数作为wrapper，将虚函数隐藏在封装之下

   ```C++
   class GameCharacter{
   public:
       int HP()const{
           //前置工作
   		int retVal = HPCalculate();
           //后置工作 ：确保得以在一个虚函数被调用之前设定好适当场景，并在调用结束之后清理场景。
           return retVal;
       }
   private:
       //可以是private也可以是protected
       virtual int HPCalculate() const{
           //缺省算法
       }
   };
   ```

   2）函数指针

   3）std::function  :C++11中引入的函数包装器，可以提供闭函数指针更强的灵活度。

   4)古典的Srategy 并非直接利用函数指针（或包装器）调用函数，而是内含一个指针指向来自继承体系的对象

5. 绝不重新定义继承而来的非虚函数*

   - 非虚函数执行的是静态绑定（前期绑定），由对象类型本身（静态类型）决定要调用的函数；
   - 虚函数执行的是动态绑定（后期绑定），决定因素不在对象本身，而在于指向该对象的指针当初的声明类型（即动态类型）

6. 绝不重新定义继承而来的缺省参数值

   ）虚函数是动态绑定而来，意思是调用一个虚函数时，究竟调用哪一份函数实现代码，取决于发出调用的那个对象的动态类型。但与之不同的是，**缺省参数值却是静态绑定**，意思是你可能会在“调用一个定义于派生类的虚函数”的同时，却使用基类为它所指定的缺省参数值。

   ```c++
   class Shape {
   public:
       enum class ShapeColor { Red, Green, Blue };
       virtual void Draw(ShapeColor color = ShapeColor::Red) const = 0;
       ...
   };
   
   class Rectangle : public Shape {
   public:
       virtual void Draw(ShapeColor color = ShapeColor::Green) const;
   };
   
   class Circle : public Shape {
   public:
       virtual void Draw(ShapeColor color) const;
   };
   
   //-----调用时
   
   Shape* pr = new Rectangle;
   Shape* pc = new Circle;
   
   pr->Draw(Shape::ShapeColor::Green);    // 调用 Rectangle::Draw(Shape::Green)
   pr->Draw();                            // 调用 Rectangle::Draw(Shape::Red)
   pc->Draw();                            // 调用 Rectangle::Draw(Shape::Red)
   
   /*-----这就迫使我们在指定虚函数时使用相同的缺省参数值，为了避免不必要的麻烦和错误，可以考虑条款 35 中列出的虚函数的替代设计，例如NVI手法：*/
   class Shape {
   public:
       enum class ShapeColor { Red, Green, Blue };
       void Draw(ShapeColor color = ShapeColor::Red) const { DoDraw(color); }
       ...
   private:
       virtual void DoDraw(ShapeColor color) const = 0;
   };
   
   class Rectangle : public Shape {
   public:
       ...
   private:
       virtual void DoDraw(ShapeColor color) const;
   };
   ```

7. 通过复合塑模出 has-a 或“根据某物实现出”

   )复合（composition），指的是某种类型的对象内含同种类型的对象。通常意味着has-a或“根据某物实现出”的关系。前者发生于应用域，后者发生于实现域。

   ```c++
   //-----has-a
   class Address { ... };
   class PhoneNumber { ... };
   
   class Person {
   public:
       ...
   private:
       std::string name;           // 合成成分物（composed object）
       Address address;            // 同上
       PhoneNumber voiceNumber;    // 同上
       PhoneNumber faxNumber;      // 同上
   };
   //-----根据某物实现出
   // 将 list 应用于 Set
   template<class T>
   class Set {
   public:
       bool member(const T& item) const;
       void insert(const T& item);
       void remove(const T& item);
       std::size_t size() const;
   
   private:
       std::list<T> rep;           // 用来表述 Set 的数据
   };
   ```

8. 明智而审慎地使用private继承

   1）类之间private继承，则编译器不会自动将一个派生类对象转换为一个基类对象。

   2）由private继承来的所有成员，在派生类中都会变成pribate属性，换句话说，private继承只继承实现，不继承接口。

   >private继承的意义是“根据某物实现出”，如果你读过条款 38，就会发现private继承和复合具有相同的意义，事实上也确实如此，绝大部分private继承的使用场合都可以被“public继承+复合”完美解决。
   >
   >> 后者比前者好的原因：
   >>
   >> 1）private继承无法阻止派生类重新定义虚函数，后者可以
   >>
   >> 2）public继承+复合时的类可以仅提供声明，并将具体定义移至实现文件中，从而降低编译依存性。
   >
   >适用于private继承的一个极端情况：空白基类最优化（EBO）
   >
   >```c++
   >class Empty {};
   >class HoldsAnInt {
   >private:
   >    int x;
   >    Empty e;//一个没有非静态成员变量、虚函数的类，看似不需要任何存储空间，但实际上 C++ 规定凡是独立对象都必须有非零大小，因此此处sizeof(HoldsAnInt)必然大于sizeof(int)，通常会多出一字节大小，但有时考虑到内存对齐之类的要求，可能会多出更多的空间。
   >};
   >//-----private继承可以避免产生额外空间：
   >class HoldsAnInt : private Empty {
   >private:
   >    int x;
   >};
   >```

9. 明智而审慎地使用多重继承

   ）可用于结合public继承和private继承，public继承用于提供接口，private继承用于提供实现;谨慎菱形继承

## 七、模板与泛型编程

1. 了解隐式接口和编译期多态

   ）对于模板参数而言，接口是隐式的，依赖于有效表达式：

   ```C++
   template<typename T>
   void DoProcessing(T& w) {
       if (w.size() > 10 && w != someNastyWidget) {
       ...
   ```

   以上代码中，T类型的隐式接口要求：

   1. 提供一个名为`size`的成员函数，该函数的返回值可与`int`（10 的类型）执行`operator>`，或经过隐式转换后可执行`operator>`。

   2. 必须支持一个`operator!=`函数，接受`T`类型和`someNastyWidget`的类型，或其隐式转换后得到的类型。（未考虑重载的情况下）

      >隐式接口和显示接口一样，都在编译期完成检查，不符合要求则代码无法通过编译。

2. 了解typename的双重含义 ：声明式class与typename没什么不同，在内部则拥有更多的含义

   >模板内出现的名称：依于某个模板参数：从属名称；而且呈嵌套状：嵌套从属名称；不依赖：非从属名称。

   ```C++
   template<typename C>
   void Print2nd(const C& container) {
       if (container.size() >= 2) {
           C::const_iterator iter(container.begin());
           ++iter;
           int value = *iter;
           std::cout << value;
       }
   }
   //编译报错
   ```

   原因）`C::const_iterator`是一个指向某类型的**嵌套从属类型名称（nested dependent type name）**，而嵌套从属名称可能会导致解析困难，因为在编译器知道`C`是什么之前，没有任何办法知道`C::const_iterator`是否为一个类型，这就导致出现了歧义状态，而 C++ 默认假设**嵌套从属名称**不是类型名称。

   解决）显式指明嵌套从属类型名称的方法就是将`typename`关键字作为其**前缀词**。

   ```C++
   typename C::const_iterator iter(container.begin());
   //注意typename不可以出现在基类列表内的嵌套从属类型名称之前，也不可以在成员初始化列表中作为基类的修饰符。
   /*名称复杂时记得使用using或typedef来简化：*/
   using value_type = typename std::iterator_traits<IterT>::value_type;
   ```

3. 学习处理模板化基类内的名称

   问题）模板编程中 模板类的继承较特殊：模板类实例化之前拒绝承认模板类中的实现函数

   解决）“C++进入模板基类观察”

    	1）基类函数调用动作之前加上this->；

   ​	 2）使用using声明式 `using 模板类::Func(param)`;

   ​	 3）明确指定来自于模板类：可能会使“虚函数绑定行为”失效。

4. 将与参数无关的代码抽离模板

   	模板编程可能造成代码膨胀，我们需要分析模板中重复使用的部分，将其抽离出模板，减轻模板具现化带来的代码量。

   ​	1）因非类型模板参数而造成的代码膨胀，往往可以消除，	做法是以函数参数或类成员变量替换模板参数。

   ​	2）因类型模板参数而造成的代码膨胀，往往可以降低，做	法是让带有完全相同二进制表述的具现类型共享实现代码

   ```C++
   template<typename T>
   class SquareMatrixBase {
   protected:
       void Invert(std::size_t matrixSize);
       ...
   private:
       std::array<T, n * n> data;
   };
   
   template<typename T, std::size_t n>
   class SquareMatrix : private SquareMatrixBase<T> {  // private 继承实现，见条款 39
       using SquareMatrixBase<T>::Invert;              // 避免掩盖基类函数，见条款 33
   
   public:
       void Invert() { this->Invert(n); }              // 调用模板基类函数，见条款 43
       ...
   };
   ```

   3)`Invert`并不是我们唯一要使用的矩阵操作函数，而且每次都往基类传递矩阵尺寸显得太过繁琐，我们可以考虑将数据放在派生类中，在基类中储存指针和矩阵尺寸。修改代码如下：

   ```C++
   template<typename T>
   class SquareMatrixBase {
   protected:
       SquareMatrixBase(std::size_t n, T* pMem)
           : size(n), pData(pMem) {}
       void SetDataPtr(T* ptr) { pData = ptr; }
       ...
   private:
       std::size_t size;
       T* pData;
   };
   
   template<typename T, std::size_t n>
   class SquareMatrix : private SquareMatrixBase<T> {
   public:
       SquareMatrix() : SquareMatrixBase<T>(n, data.data()) {}
       ...
   private:
       std::array<T, n * n> data;
   };
   ```

   ） 同样地，上面的代码也使用到了牺牲封装性的`protected`，可能会导致资源管理上的混乱和复杂，考虑到这些，也许一点点模板代码的重复并非不可接受。

5. 运用成员函数模板接受所有的兼容类型

   ）C++ 视模板类的不同具现体为完全不同的的类型，但在泛型编程中，我们可能需要一个模板类的不同具现体能够相互类型转换。

   ）考虑设计一个智能指针类，而智能指针需要支持不同类型指针之间的隐式转换（如果可以的话），以及普通指针到智能指针的显式转换。很显然，我们需要的是模板拷贝构造函数：

   ```c++
   c++
   template<typename T>
   class SmartPtr {
   public:
       template<typename U>
       SmartPtr(const SmartPtr<U>& other)
           : heldPtr(other.get()) { ... }
   
       template<typename U>
       explicit SmartPtr(U* p)
           : heldPtr(p) { ... }
   
       T* get() const { return heldPtr; }
       ...
   private:
       T* heldPtr;
   };
   ```

   ）使用`get`获取原始指针，并将在原始指针之间进行类型转换本身提供了一种保障，如果原始指针之间不能隐式转换，那么其对应的智能指针之间的隐式转换会造成编译错误。

   ）模板构造函数并不会阻止编译器暗自生成默认的构造函数，所以如果你想要控制拷贝构造的方方面面，你必须同时声明泛化拷贝构造函数和普通拷贝构造函数，相同规则也适用于赋值运算符：

   ```C++
   c++
   template<typename T>
   class shared_ptr {
   public:
       shared_ptr(shared_ptr const& r);                // 拷贝构造函数
   
       template<typename Y>
       shared_ptr(shared_ptr<Y> const& r);             // 泛化拷贝构造函数
   
       shared_ptr& operator=(shared_ptr const& r);     // 拷贝赋值运算符
   
       template<typename Y>
       shared_ptr& operator=(shared_ptr<Y> const& r);  // 泛化拷贝赋值运算符
   
       ...
   };
   ```

6. 需要类型转换时请为模板定义非成员函数******

7. 请用traits classes表现类型信息******

   >traits classes 可以使我们在编译期就能获取某些类型信息，它被广泛运用于 C++ 标准库中。traits 并不是 C++ 关键字或一个预先定义好的构件：它们是一种技术，也是 C++ 程序员所共同遵守的协议，并要求对用户自定义类型和内置类型表现得一样好。

8. 认识模板元编程

   > 模板元编程（Template metaprogramming，TMP）是编写基于模板的 C++ 程序并执行于编译期的过程。这可以帮助我们在编译期时发现一些原本要在运行期时才能察觉的错误，以及得到较小的可执行文件、较短的运行期、较少的内存需求。当然，副作用就是会使编译时间变长。

   模板元编程已被证明是“图灵完备”的，并且以“函数式语言”的形式发挥作用，因此在模板元编程中没有真正意义上的循环，所有循环效果只能藉由递归实现，而递归在模板元编程中是由 **“递归模板具现化（recursive template instantiation）”** 实现的。

   常用于引入模板元编程的例子是在编译期计算阶乘：

   ```C++
   c++
   template<unsigned n>            // Factorial<n> = n * Factorial<n-1>
   struct Factorial {
       enum { value = n * Factorial<n-1>::value };
   };
   
   template<>
   struct Factorial<0> {           // 处理特殊情况：Factorial<0> = 1
       enum { value = 1 };
   };
   
   std::cout << Factorial<5>::value;
   ```

   模板元编程很酷，但对其进行调试可能是灾难性的，因此在实际应用中并不常见。我们可能会在下面几种情形中见到它的出场：

   1. 确保量度单位正确。
   2. 优化矩阵计算。
   3. 可以生成客户定制之设计模式（custom design pattern）实现品。

## 八、定制new和delete

1. **了解new-handler的行为**

2. **了解new和delete的合理替换时机**

3. **编写new和delete时需固守常规**

4. 为了placement new 也要写placement delete

   placement new 最初的含义指的是“接受一个指针指向对象该被构造之处”的`operator new`版本，它在标准库中的用途广泛，其中之一是负责在 vector 的未使用空间上创建对象，它的声明如下：

   ```C++
   c++
   void* operator new(std::size_t, void* pMemory) noexcept;
   ```

   我们此处要讨论的是广义上的 placement new，即带有附加参数的`operator new`，例如下面这种：

   ```C++
   c++
   void* operator new(std::size_t, std::ostream& logStream);
   
   auto pw = new (std::cerr) Widget;
   ```

   当我们在使用 new 表达式创建对象时，共有两个函数被调用：一个是用以分配内存的`operator new`，一个是对象的构造函数。假设第一个函数调用成功，而第二个函数却抛出异常，那么会由 C++ runtime 调用`operator delete`，归还已经分配好的内存。

   这一切的前提是 C++ runtime 能够找到`operator new`对应的`operator delete`，如果我们使用的是自定义的 placement new，而没有为其准备对应的 placement delete 的话，就无法避免发生内存泄漏。因此，合格的代码应该是这样的：

   ```C++
   c++
   class Widget {
   public:
       static void* operator new(std::size_t size, std::ostream& logStream);   // placement new
   
       static void operator delete(void* pMemory);                             // delete 时调用的正常 operator delete
       static void operator delete(void* pMemory, std::ostream& logStream);    // placement delete
   };
   ```

   另一个要注意的问题是，由于成员函数的名称会掩盖其外部作用域中的相同名称（见条款 33），所以提供 placement new 会导致无法使用正常版本的`operator new`：

   ```C++
   c++
   class Base {
   public:
   	static void* operator new(std::size_t size, std::ostream& logStream);
   	...
   };
   
   auto pb = new Base;             // 无法通过编译！
   auto pb = new (std::cerr) Base; // 正确
   ```

   同样道理，派生类中的`operator new`会掩盖全局版本和继承而得的`operator new`版本：

   ```C++
   c++
   class Derived : public Base {
   public:
   	static void* operator new(std::size_t size);
   	...
   };
   
   auto pd = new (std::clog) Derived;  // 无法通过编译！
   auto pd = new Derived;              // 正确
   ```

   为了避免名称遮掩问题，需要确保以下形式的`operator new`对于定制类型仍然可用，除非你的意图就是阻止客户使用它们：

   ```C++
   c++
   void* operator(std::size_t) throw(std::bad_alloc);           // normal new
   void* operator(std::size_t, void*) noexcept;                 // placement new
   void* operator(std::size_t, const std::nothrow_t&) noexcept; // nothrow new
   ```

   一个最简单的实现方式是，准备一个基类，内含所有正常形式的 new 和 delete：

   ```C++
   c++
   class StadardNewDeleteForms{
   public:
       // normal new/delete
       static void* operator new(std::size_t size){
           return ::operator new(size);
       }
       static void operator delete(void* pMemory) noexcept {
           ::operator delete(pMemory);
       }
   
       // placement new/delete
       static void* operator new(std::size_t size, void* ptr) {
           return ::operator new(size, ptr);
       }
       static void operator delete(void* pMemory, void* ptr) noexcept {
           ::operator delete(pMemory, ptr);
       }
   
       // nothrow new/delete
       static void* operator new(std::size_t size, const std::nothrow_t& nt) {
           return ::operator new(size,nt);
       }
       static void operator delete(void* pMemory,const std::nothrow_t&) noexcept {
           ::operator delete(pMemory);
       }
   };
   ```

   凡是想以自定义形式扩充标准形式的客户，可以利用继承和`using`声明式（见条款 33）取得标准形式：

   ```C++
   c++
   class Widget: public StandardNewDeleteForms{
   public:
       using StandardNewDeleteForms::operator new;
       using StandardNewDeleteForms::operator delete;
   
       static void* operator new(std::size_t size, std::ostream& logStream);
       static void operator detele(std::size_t size, std::ostream& logStream) noexcept;
       ...
   };
   ```

## 九、杂项 

1. 不要轻忽编译器的警告

   1. 严肃对待编译器发出的警告信息。努力在你的编译器的最高（最严苛）警告级别下争取“无任何警告”的荣誉。
   2. 不要过度依赖编译器的警告能力，因为不同的编译器对待事情的态度不同。一旦移植到另一个编译器上，你原本依赖的警告信息可能会消失。

2. 让自己熟悉包括TR1在内的标准程序库

   > 如今 TR1 草案已完全融入 C++ 标准当中，没有再过多了解 TR1 标准库的必要。

3. 让自己熟悉Boost

   > Boost 是若干个程序库的集合，并且当中的许多库已经被 C++ 吸纳为标准库的一部分。不过在现在的 Modern C++ 时代，是否该在项目中使用 Boost 仍然有一定的争议，一些 Boost 组件并无法做到像 C++ 标准库那样高性能，零开销抽象，但毫无疑问的是，Boost 的参考价值是无法忽视的，你可以在 Boost 中找到许多非常值得学习和借鉴的实现。
