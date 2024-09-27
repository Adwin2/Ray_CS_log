//for_each 的实现
//其中 func(value) 使用了函数指针的知识点

#include <iostream>
#include <vector>

void PrintValue(int value)
{
    std::cout << "Value is:" << value << std::endl;
}

void ForEach(const std::vector<int>&values, void(*func)(int))
{
    for(int value : values) {
        func(value);
    }
}

int main() {
    std::vector<int> values = {1, 5, 4, 3, 2};
    ForEach(values,PrintValue);

    std::cin.get();
}