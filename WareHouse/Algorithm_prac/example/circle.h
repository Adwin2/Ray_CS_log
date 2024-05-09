#pragma once
#include <iostream>
using namespace std;
#include "point.h"
//圆类
class circle {
public:
	//设置
	void setR(int R);
	void setCenter(Point center);
	//获取
	int getR();
	//在类中可以用另一个类作为本类的成员
	Point getCenter();
private:
	int m_R;
	Point m_center;
};
