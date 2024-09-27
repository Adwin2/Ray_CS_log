#include<iostream>
using namespace std;
class Person {
public:
	//静态成员函数
	static void func() {
		m_A = 100;//静态成员函数仅可以访问静态成员变量
		//m_B=200; //静态成员函数不可以访问非静态成员变量 ，不知道属于哪个对象
		cout << "static void 调用" << endl;
	}
	static int m_A;
};
//静态成员函数也是有访问权限的
int Person::m_A = 0;
void test01() {
	//1、通过对象访问
	Person p;
	p.func();
	//2、通过类名访问
	Person::func();
}


int main() {
	test01();
}