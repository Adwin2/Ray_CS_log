#include <iostream>
using namespace std;
#include"circle.h"
#include"point.h"
//���Բ����
////����
//class Point {
//public:
//	//����
//	void setX(int x) {
//		m_x = x;
//	}
//	void setY(int y) {
//		m_y = y;
//	}
//	//��ȡ
//	int getX() {
//		return m_x;
//	}
//	int getY() {
//		return m_y;
//	}
//private:
//	int m_x;
//	int m_y;
//};
//
////Բ��
//class circle {
//public:
//	//����
//	void setR(int R) {
//		m_R = R;
//	}
//	void setCenter(Point center) {
//		m_center = center;
//	}
//	//��ȡ
//	int getR() {
//		return m_R;
//	}
//	//�����п�������һ������Ϊ����ĳ�Ա
//	Point getCenter() {
//		return m_center;
//	}
//private:
//	int m_R;
//	Point m_center;
//};
void isInCircle(circle &c,Point &p)
{
	//������������ƽ����
	int distance = ((c.getCenter().getX() - p.getX()) * (c.getCenter().getX() - p.getX())) + ((c.getCenter().getY() - p.getY()) * (c.getCenter().getY() - p.getY()));
	//����뾶��ƽ��
	int rdist = (c.getR()* c.getR());
	//�ж����߹�ϵ
	if (distance > rdist) {
		cout << "����Բ��" << endl;
	}
	else if (distance == rdist) {
		cout << "����Բ��" << endl;
	}
	else {
		cout << "����Բ��" << endl;
	}
}
int main() {
	//����Բ
	circle c;
	c.setR(10);
	Point center;
	center.setX(10);
	center.setY(0);
	c.setCenter(center);
	//������
	Point  p;
	p.setX(10);
	p.setY(9);
	//�жϹ�ϵ
	isInCircle(c, p);
}