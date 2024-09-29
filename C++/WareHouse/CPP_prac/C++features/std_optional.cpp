#include <iostream>
#include <fstream>
#include <string>
#include <optional>

std::optional<std::string> ReadFileAsString(const std::string& filepath) {
    std::ifstream stream(filepath);
    if(stream) {
        std::string result;
        //read file
        stream.close();
        return result;
    }
    return {};
}

int main() {
    std::optional<std::string> data = ReadFileAsString("data.txt");
    
    std::string value = data.value_or("Not present");
    std::cout<<value<<std::endl;
    if(data) //or data.has_value() they r all bool type ;never mind 
    {
        std::cout<<"Read file successfully"<<std::endl;
    }
    else {
        std::cout<<"Read file failed" <<std::endl;
    }
    std::cin.get();
}