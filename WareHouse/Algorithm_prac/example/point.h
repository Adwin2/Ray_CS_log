#pragma once
#include<iostream>
using namespace std;
//点类
class Point {
public:
	//设置
	void setX(int x);
	void setY(int y);
	//获取
	int getX();
	int getY();
private:
	int m_x;
	int m_y;
};
