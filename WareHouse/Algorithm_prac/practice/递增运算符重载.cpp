#include <iostream>
using namespace std;

class Myinteger {
	friend ostream& operator<<(ostream& cout, Myinteger myint);
public:
	Myinteger() //���캯��
	{
		m_Num = 0;
	}
	//ǰ�õ������������
	Myinteger& operator++() //������ֵ���򲻻��ظ��洢����ͬ�ڴ棻��Ҫ��������
	{
		m_Num++;
		return *this;//������ ��������
	}
	//���á���
	
private:
	int m_Num;
};
//����<<
ostream& operator<<(ostream& cout,Myinteger myint) //**������ (ostream& cout���ɱ䣩 ,Myinteger myint)***�ص㣡��
{
	cout << myint.m_Num << endl;
	return cout;  //ע�ⷵ��ֵ
}




void test01() {
	Myinteger myint;
	cout << myint;

}