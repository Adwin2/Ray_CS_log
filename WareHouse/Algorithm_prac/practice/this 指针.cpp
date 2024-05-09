#include<iostream>
using namespace std;

class Person {
public:
	Person(int age) {
		//this 指针指向 被调用的成员函数所属的对象
		this->age = age;//1、解决名称冲突
	}

	Person& PersonAdd(Person& p) //返回对象本体需要引用   若用Person 返回值 则拷贝一份原数据，形成一个新值
	{
		this->age += p.age;

		//this指向p2的指针，而*this 指向的就是p2这个对象本体
		return *this;//*this指向本体  实现函数自身的反复调用
	}

		
	int age;

};
//1、
void test01() {
	Person p1(18);
	cout <<"年龄是"<<p1.age<<endl;
	
}
//2、返回对象本身用*this
void test02() {
	Person p1(10);

	Person p2(10);

	//链式编程思想
	p2.PersonAdd(p1).PersonAdd(p1).PersonAdd(p1);
	cout << "p2的年龄多大" << p2.age << endl;
}
int main() {
	test01();
	test02();
}