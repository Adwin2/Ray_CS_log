//C++标准库的小字符串优化  \
15及以下长度的字符 在栈中分配内存 以上则正常heap allocate\
Visual studio Debug 和 Release 模式不同 \
VScode 相当于是Release模式(按规则来)

//本例码同时包含 查看内存分配的简单方法\
可以使用VS 自带的Valgrind 工具 

#include <iostream>
#include <memory>

class Result {
public:
    uint32_t Allocated = 0;
    uint32_t Freed = 0;

    uint32_t CurrentUsage() {return Allocated - Freed;}
};

static Result res;

void* operator new(size_t size_t) {
    std::cout<< "Allocating : "<<size_t<<" bytes"<<std::endl;
    
    res.Allocated += size_t;
    //std::cout<<res.Allocated<<std::endl;
    
    return malloc(size_t);
}

void operator delete(void* memory,size_t size_t) {
    std::cout<< "Freed " <<size_t <<" bytes"<<std::endl;

    res.Freed += size_t;
    //std::cout<<res.Freed<<std::endl; 
    
    return free(memory);
}

class Object{
    int x,y,z;
};

static void PrintMem(){
    std::cout<<"Memory Usage : "<<res.CurrentUsage()<<std::endl;
}


int main() {
    PrintMem();
    {
        //最大 15个字符 + "\0"即[16]不会堆分配\
        不需要手动释放 内部调用*析构函数*释放内存
        std::string str = "Raymondddddddddk";
    }
    PrintMem();
    {
        PrintMem();
        std::unique_ptr<Object> obj = std::make_unique<Object>(); 
        PrintMem();
    }
    PrintMem();
    std::cin.get();
}