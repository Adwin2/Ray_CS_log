#include "circle.h"//���а���point.h�ļ�
//Բ��
	void circle::setR(int R) {
		m_R = R;
	}
	void circle::setCenter(Point center) {
		m_center = center;
	}
	//��ȡ
	int circle::getR() {
		return m_R;
	}
	//�����п�������һ������Ϊ����ĳ�Ա
	Point circle::getCenter() {
		return m_center;
	}
