#include<iostream>
using namespace std;

class Person {
public:
	Person(int age) {
		//this ָ��ָ�� �����õĳ�Ա���������Ķ���
		this->age = age;//1��������Ƴ�ͻ
	}

	Person& PersonAdd(Person& p) //���ض�������Ҫ����   ����Person ����ֵ �򿽱�һ��ԭ���ݣ��γ�һ����ֵ
	{
		this->age += p.age;

		//thisָ��p2��ָ�룬��*this ָ��ľ���p2���������
		return *this;//*thisָ����  ʵ�ֺ�������ķ�������
	}

		
	int age;

};
//1��
void test01() {
	Person p1(18);
	cout <<"������"<<p1.age<<endl;
	
}
//2�����ض�������*this
void test02() {
	Person p1(10);

	Person p2(10);

	//��ʽ���˼��
	p2.PersonAdd(p1).PersonAdd(p1).PersonAdd(p1);
	cout << "p2��������" << p2.age << endl;
}
int main() {
	test01();
	test02();
}