//重载右移运算符 输入


#include<iostream>
using namespace std;

class Sales_data {
public:
    int bookNum;
    int unit;
    double revenue;
    Sales_data() {
        bookNum = 0;
        unit = 0;
        revenue = 0;
    }
};

istream &operator>>(istream &is,Sales_data &item){
    double price = 0;
    is>>item.bookNum>>item.unit>>price;
    if(is) {
        item.revenue = item.unit * price;
    }
    else {
        item = Sales_data();
    }
    return is;
}

int main(){
    
    while (1)
    {
        Sales_data tmp;
        cin>>tmp;
        cout<<tmp.bookNum<<" "<<tmp.revenue<<" "<<tmp.unit<<endl;
    }
    return 0;
}