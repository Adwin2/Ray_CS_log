#pragma once
#include <iostream>
using namespace std;
#include "point.h"
//Բ��
class circle {
public:
	//����
	void setR(int R);
	void setCenter(Point center);
	//��ȡ
	int getR();
	//�����п�������һ������Ϊ����ĳ�Ա
	Point getCenter();
private:
	int m_R;
	Point m_center;
};
