#include<iostream>
using namespace std;
class Person {
public:
	//��̬��Ա����
	static void func() {
		m_A = 100;//��̬��Ա���������Է��ʾ�̬��Ա����
		//m_B=200; //��̬��Ա���������Է��ʷǾ�̬��Ա���� ����֪�������ĸ�����
		cout << "static void ����" << endl;
	}
	static int m_A;
};
//��̬��Ա����Ҳ���з���Ȩ�޵�
int Person::m_A = 0;
void test01() {
	//1��ͨ���������
	Person p;
	p.func();
	//2��ͨ����������
	Person::func();
}


int main() {
	test01();
}