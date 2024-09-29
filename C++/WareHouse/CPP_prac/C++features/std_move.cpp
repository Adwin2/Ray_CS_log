//移动构造函数 和 移动赋值函数

#include <iostream>
#include <string.h>

class String {
public:
    String() = default;
    String(const char* string)  // 构造函数
    {
        printf("Created\n");
        m_size = strlen(string);
        m_data = new char[m_size];
        memcpy(m_data,string,m_size);
    }

    String(const String& other) // 拷贝构造函数
    {
        printf("Copied\n");
        m_size = other.m_size;
        m_data = new char[m_size];
        memcpy(m_data,other.m_data,m_size);
    }
    // 移动构造函数  必须有 noexcept 关键字 这样编译器才会将其视为安全的，否则会回退到复制构造函数 程序可能崩溃
    String(String&& other) noexcept 
    {
        // 浅拷贝 rewire the ptr    复制数据、重新分配内存属于深拷贝
        printf("Moved\n");
        m_size = other.m_size;
        m_data = other.m_data; // 改变指针指向 (move)

        other.m_size = 0;
        other.m_data = nullptr;
    }
    //重载  移动赋值运算符
    String& operator=(String&& other) {
        printf("Moved\n");
    //1.delete the origin data to avoid mem leak \
    2.make sure they are different
        if(this != &other) {
            delete[] m_data;

            m_size = other.m_size;
            m_data = other.m_data; // 改变指针指向 (move)

            other.m_size = 0;
            other.m_data = nullptr;
        }
        return *this;
    }

    void Print() 
    {
        for(uint32_t i = 0;i < m_size; i ++)
            printf("%c",m_data[i]);
        printf("\n");
    }

    ~String() {
        printf("destroyed\n");
        delete m_data;
    }
private:
    char* m_data;
    uint32_t m_size = 0;
};

class Entity{
public:
    Entity(const String& name) 
        :m_name(name)
    {}
    //移动构造函数传入临时变量一定不是const类型的！！！！！
    Entity(String&& name) //接受临时变量 右值引用构造函数
        :m_name(std::move(name))  // 显式转换为临时对象 从而调用移动构造函数 同 m_name((String&&)name)
        {   
        }    
    void PrintName() {
        m_name.Print();
    }

private:
    String m_name;
};

int main() {
    //移动构造函数 例
    Entity entity("raymond");
    entity.PrintName();

    //移动赋值函数 例
    String string = "Hello";
    //创建新变量时 等同于 String dest(std::move(string));&same as $(String&&)string$
    String dest = std::move(string);
    //重载的是 = and .assign() 赋值运算符 因为其默认不接受临时变量
    
    
    String apple = "apple";
    String dest_01 = "";

    std::cout<<"apple : ";
    apple.Print();
    std::cout<<"dest : ";
    dest_01.Print();

    dest_01 = std::move(apple); //  <-- here 相当于调用 .operator=() 函数 前面那个赋值运算符并不是 ——区别
       
    std::cout<<"apple : ";
    apple.Print();
    std::cout<<"dest : ";
    dest_01.Print();

    std::cin.get();
}
