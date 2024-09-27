#include <iostream>
using namespace std;

class Myinteger {
	friend ostream& operator<<(ostream& cout, Myinteger myint);
public:
	Myinteger() //构造函数
	{
		m_Num = 0;
	}
	//前置递增运算符重载
	Myinteger& operator++() //若返回值，则不会重复存储在相同内存；需要返回引用
	{
		m_Num++;
		return *this;//解引用 返回自身
	}
	//后置……
	
private:
	int m_Num;
};
//重载<<
ostream& operator<<(ostream& cout,Myinteger myint) //**括号内 (ostream& cout（可变） ,Myinteger myint)***重点！！
{
	cout << myint.m_Num << endl;
	return cout;  //注意返回值
}




void test01() {
	Myinteger myint;
	cout << myint;

}