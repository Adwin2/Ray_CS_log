#include<iostream>
using namespace std;

#define PI 3.14

class shape{
	public:
		virtual double area() = 0;	//基类虚函数没有定义/纯虚  会导致文件链接的时候出错 type like:'undefined reference to `typeinfo for <基类>'
};

class Square : public shape {
	public:
		virtual double area();
	private:
		int a = 0;
		int b = 0;
};

class Triangle : public shape {
	public:
		virtual double area();
	private:
		int a = 0;
		int h = 0;
};

class Circle : public shape {
	public:
		virtual double area(); 
		
	private:
		double r = 1;
};

double Square::area() {
	return a*b;
}

double Triangle::area() {
	return (a* h)/2;
}

double Circle::area() {
	return PI* r* r;
}

int main() {
	double sum = 0;
	Square s;
	Triangle t;
	Circle c;
	sum += s.area();
	sum += t.area();
	sum += c.area();
	cout<< sum <<endl;

	return 0;
}
