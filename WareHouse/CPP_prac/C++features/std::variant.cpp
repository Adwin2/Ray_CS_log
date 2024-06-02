#include <iostream>
#include <variant>
//4-10  Variant使用实例
enum ErrorCode {
    NoAccess = 0, NotFound = 1, None = 2
};

std::variant<std::string, ErrorCode> ReadFileAsString() {
    return {};
}



int main() {
    std::variant<std::string, int>data;
    data = "Raymond";
    std::cout<<std::get<std::string>(data)<<std::endl;
    if(auto value = std::get_if<std::string>(&data)) 
    {
        std::string& str = *value;
    }
    else
    {
    //process another possibility
    }
    
    data = 2;
    std::cout<<std::get<int>(data);    
    std::cin.get();
}

//BTW std::any pluses the type numbers, but its performance is not better\
(especially for large memory, it will dynamically allocate the memory) \
so just mention it here.