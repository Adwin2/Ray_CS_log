#include <iostream>
using namespace std;

//��̬��Ա����
class Person {
public:
	//1�����ж��󶼹���ͬһ������
	//2������׶η����ڴ�
	//3�����������������ʼ��    --���ע�⣡
	static int m_A;

	//��̬��Ա����Ҳ���з���Ȩ�޵�
};

int Person::m_A = 100;//��������person�������µľ�̬��Ա ����ʼ��

//void test01() {
//	Person p;
//	cout << p.m_A << endl;
//	Person p2;
//	p2.m_A = 200;
//	cout << p2.m_A << endl;
//}
void test02() {
	//��̬��Ա������-���ж�����ͬһ�����ݣ���������ַ��ʷ�ʽ
	//1��ͨ��������з���
	//Person p;
	//cout <<p.m_A <<endl;
	//2��ͨ���������з���
	cout << Person::m_A << endl;

}

int main() {
	//test01();
	test02();
	system("pause");

}