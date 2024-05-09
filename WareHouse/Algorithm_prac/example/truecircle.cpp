#include "circle.h"//其中包含point.h文件
//圆类
	void circle::setR(int R) {
		m_R = R;
	}
	void circle::setCenter(Point center) {
		m_center = center;
	}
	//获取
	int circle::getR() {
		return m_R;
	}
	//在类中可以用另一个类作为本类的成员
	Point circle::getCenter() {
		return m_center;
	}
