#include <iostream>
using namespace std;
#include"circle.h"
#include"point.h"
//点和圆案例
////点类
//class Point {
//public:
//	//设置
//	void setX(int x) {
//		m_x = x;
//	}
//	void setY(int y) {
//		m_y = y;
//	}
//	//获取
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
////圆类
//class circle {
//public:
//	//设置
//	void setR(int R) {
//		m_R = R;
//	}
//	void setCenter(Point center) {
//		m_center = center;
//	}
//	//获取
//	int getR() {
//		return m_R;
//	}
//	//在类中可以用另一个类作为本类的成员
//	Point getCenter() {
//		return m_center;
//	}
//private:
//	int m_R;
//	Point m_center;
//};
void isInCircle(circle &c,Point &p)
{
	//计算横纵坐标的平方和
	int distance = ((c.getCenter().getX() - p.getX()) * (c.getCenter().getX() - p.getX())) + ((c.getCenter().getY() - p.getY()) * (c.getCenter().getY() - p.getY()));
	//计算半径的平方
	int rdist = (c.getR()* c.getR());
	//判断两者关系
	if (distance > rdist) {
		cout << "点在圆外" << endl;
	}
	else if (distance == rdist) {
		cout << "点在圆上" << endl;
	}
	else {
		cout << "点在圆内" << endl;
	}
}
int main() {
	//创建圆
	circle c;
	c.setR(10);
	Point center;
	center.setX(10);
	center.setY(0);
	c.setCenter(center);
	//创建点
	Point  p;
	p.setX(10);
	p.setY(9);
	//判断关系
	isInCircle(c, p);
}