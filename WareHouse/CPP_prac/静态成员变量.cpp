#include <iostream>
using namespace std;

//静态成员变量
class Person {
public:
	//1、所有对象都共享同一份数据
	//2、编译阶段分配内存
	//3、类内声明，类外初始化    --务必注意！
	static int m_A;

	//静态成员变量也是有访问权限的
};

int Person::m_A = 100;//声明是在person作用域下的静态成员 并初始化

//void test01() {
//	Person p;
//	cout << p.m_A << endl;
//	Person p2;
//	p2.m_A = 200;
//	cout << p2.m_A << endl;
//}
void test02() {
	//静态成员变量—-所有对象共享同一份数据，因此有两种访问方式
	//1、通过对象进行访问
	//Person p;
	//cout <<p.m_A <<endl;
	//2、通过类名进行访问
	cout << Person::m_A << endl;

}

int main() {
	//test01();
	test02();
	system("pause");

}