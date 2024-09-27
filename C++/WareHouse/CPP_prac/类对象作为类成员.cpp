#include<iostream>
using namespace std;
#include<string>
//类对象作为类成员    称该成员为对象成员

//手机类
class Phone
{
public:

	Phone(string PName)
	{
		m_PName = PName;
	}
	//手机品牌名称
	string m_PName;
};
//人类
class Person
{
public:
	//Phone m_Phone = pName 隐式转换法
	Person(string name, string pName):m_Name(name),m_Phone(pName)
	{

	}
	//姓名
	string m_Name;
	//手机
	Phone m_Phone;
};
//当其它类对象作为本类成员，构造时先构造类对象，再构造自身   析构时顺序与其相反
void test01()
{
	Person p("张三", "iPhone max");

	cout <<p.m_Name<<"拿着" <<p.m_Phone.m_PName << endl;
}

int main() {
	test01();
	
	system("pause");
}