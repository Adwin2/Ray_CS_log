//deque ����  sort �����㷨 ! ! !

#include<iostream>
using namespace std;
#include <deque>
#include <functional>

#include <algorithm>//��׼�㷨ͷ�ļ�

void printDeque(const deque<int>& d)//�����õķ�ʽ��������
{
	for (deque<int>::const_iterator it = d.begin(); it != d.end(); it++) //deque�����ĵ�������const_iterator
	{
		cout << *it << " ";
	}
	cout << endl;
}

//deque ��������
void test01() {
	deque<int>d;
	d.push_back(10);
	d.push_back(20);
	d.push_back(30);
	d.push_front(100);


	printDeque(d);

	/*����  Ĭ�� ��С���� ����
	  ���򡪡�sort(d.rbegin(),d.rend());����
	  ����֧��������ʵĵ�����������������
	  ����ʱ�䳬�޵�����������龺��ʹ��
	*/
	sort(d.begin(), d.end(), std::greater<int>()); //标准库中 greater<_TP>() / lesser<_TP>()
	cout << "�����" << endl;
	printDeque(d);
	
}

int main() {

	test01();

	return 0;

}