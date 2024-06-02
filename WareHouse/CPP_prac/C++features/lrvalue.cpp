//引入：左值对应地址 右值临时对象\
&&右值引用相关  可以偷取临时资源进行相关优化

#include <iostream>
//1.const 左值引用可以同时接受两种值
//2.非常量引用(&)必须是左值
//3.(&&)右值引用只接受临时对象
void Print(std::string&& a) {
    std::cout<<a<<std::endl;
}

int main() {
    std::string a = "a";
    std::string b = "b";
    std::string c = a + b;
    Print(a + b);
    //Print(c);
}