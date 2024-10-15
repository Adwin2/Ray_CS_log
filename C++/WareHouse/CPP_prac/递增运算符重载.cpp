#include<iostream>
using namespace std;

class Point{
	public:
		Point(int a = 0, int b = 0):x(a), y(b){}
		Point& operator++(){
			x ++;
			return *this;
		}
		Point& operator--(){
			x --;
			return *this;
		} //后置运算符返回引用
		
		Point operator++(int) {
			Point old = *this;
			++(*this);
			return old;
		}
		Point operator--(int) {
			Point old = *this;
			--(*this);
			return old;
		}
		//前置重载 为了和后置区分开 没有实际作用

		int getX(){
			return x;
		}
		int getY(){
			return y;
		}
	private:
		int x;
		int y;
};

ostream& operator<<(ostream& out, Point p) {
	out<<p.getX() <<"\t"<<p.getY()<<endl;
	return out;
}

int main(){
	Point p;
	p++;
	cout<<p<<endl;
	p--;
	cout<<p<<endl;
	return 0;
}
