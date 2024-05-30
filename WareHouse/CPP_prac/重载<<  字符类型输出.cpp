//重载左移运算符 输出字符类型

#include<iostream>
#include<string>
using namespace std;

ostream &operator<<(ostream &os, const string &str){
    for(int i = 0;i< str.length();i++){
        os<<str[i];
    }
    return os;
}
int main(){
    string a = "Trial";
    cout<<a<<endl;
    return 0;
}