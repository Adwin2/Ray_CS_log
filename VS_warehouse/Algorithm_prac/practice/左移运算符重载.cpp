#include<iostream>
using namespace std;

//左移运算符重载

class Person{
	//friend void test01();//相当于给封装加缺口
	friend ostream& operator<<(ostream& out, Person& p);
public:
	Person(int a, int b) {
		m_A = a;
		m_B = b;

	}
private:

	//利用成员函数重载 左移运算符 p.operator<<cout 简化为 p<<cout 
	//不会利用成员函数 重载<<运算符，因为无法实现cout在左侧
	/*void operator<<(Person& p) {

	}*/
	int m_A;
	int m_B;
};
//只能利用全局函数重载 左移运算符
//cout 数据类型属于 ostream 全局仅允许一个，所以用引用符&
//左值返回值 用&引用符号
ostream& operator<<(ostream &out,Person &p)//简化 cout << p
{
	//out 通过&引用符指向同一块内存 即cout的别名
	out << "m_A = " << p.m_A << " m_B = "<<p.m_B;
	return  out;
}
void test01() {
	//Person p;
	//p.m_A = 10;
	//p.m_B = 10;
	Person p(10,10);

	//cout << p.m_A << endl;
	//链式编程思想 注意前者 返回值
	cout << p <<endl;
}
int main() {
	test01();
	system("pause");
}