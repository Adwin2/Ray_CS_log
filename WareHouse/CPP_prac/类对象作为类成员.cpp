#include<iostream>
using namespace std;
#include<string>
//�������Ϊ���Ա    �Ƹó�ԱΪ�����Ա

//�ֻ���
class Phone
{
public:

	Phone(string PName)
	{
		m_PName = PName;
	}
	//�ֻ�Ʒ������
	string m_PName;
};
//����
class Person
{
public:
	//Phone m_Phone = pName ��ʽת����
	Person(string name, string pName):m_Name(name),m_Phone(pName)
	{

	}
	//����
	string m_Name;
	//�ֻ�
	Phone m_Phone;
};
//�������������Ϊ�����Ա������ʱ�ȹ���������ٹ�������   ����ʱ˳�������෴
void test01()
{
	Person p("����", "iPhone max");

	cout <<p.m_Name<<"����" <<p.m_Phone.m_PName << endl;
}

int main() {
	test01();
	
	system("pause");
}