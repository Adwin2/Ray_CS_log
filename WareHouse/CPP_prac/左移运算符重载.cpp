#include<iostream>
using namespace std;

//�������������

class Person{
	//friend void test01();//�൱�ڸ���װ��ȱ��
	friend ostream& operator<<(ostream& out, Person& p);
public:
	Person(int a, int b) {
		m_A = a;
		m_B = b;

	}
private:

	//���ó�Ա�������� ��������� p.operator<<cout ��Ϊ p<<cout 
	//�������ó�Ա���� ����<<���������Ϊ�޷�ʵ��cout�����
	/*void operator<<(Person& p) {

	}*/
	int m_A;
	int m_B;
};
//ֻ������ȫ�ֺ������� ���������
//cout ������������ ostream ȫ�ֽ�����һ�������������÷�&
//��ֵ����ֵ ��&���÷���
ostream& operator<<(ostream &out,Person &p)//�� cout << p
{
	//out ͨ��&���÷�ָ��ͬһ���ڴ� ��cout�ı���
	out << "m_A = " << p.m_A << " m_B = "<<p.m_B;
	return  out;
}
void test01() {
	//Person p;
	//p.m_A = 10;
	//p.m_B = 10;
	Person p(10,10);

	//cout << p.m_A << endl;
	//��ʽ���˼�� ע��ǰ�� ����ֵ
	cout << p <<endl;
}
int main() {
	test01();
	system("pause");
}