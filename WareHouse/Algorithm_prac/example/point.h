#pragma once
#include<iostream>
using namespace std;
//����
class Point {
public:
	//����
	void setX(int x);
	void setY(int y);
	//��ȡ
	int getX();
	int getY();
private:
	int m_x;
	int m_y;
};
