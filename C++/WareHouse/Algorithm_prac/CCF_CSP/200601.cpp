#include<iostream>
#include<vector>
using namespace std;

class dot {
public:
	int x_dot;
	int y_dot;
	char type;
	bool status;
	dot(int a, int b, char n) {
		x_dot = a;
		y_dot = b;
		type = n;
		status = false;
	}
};

double distance(int t_0, int t_1, int t_2,int x_dot) {
	return (-(double)t_0 - (double)t_1 * x_dot) / (double)t_2;
}

void set(vector<dot>& vec) {
	for (int i = 0; i < vec.size(); i++) {
		vec[i].status = false;
	}
}
int main() {
	vector<dot>type_a;
	vector<dot>type_b;
	int n_dot = 0;
	int m_search = 0;
	cin >> n_dot >> m_search;
	for (int i = 0; i < n_dot; i++) {
		dot d_example(0, 0, 'A');
		cin >> d_example.x_dot >> d_example.y_dot >> d_example.type;
		if (d_example.type == 'A') type_a.push_back(d_example);
		else type_b.push_back(d_example);
	}
	for (int m = 0; m < m_search; m++) {
		int theta_0 = 0;
		int theta_1 = 0;
		int theta_2 = 0;
		int s = 0;
		cin >> theta_0 >> theta_1 >> theta_2;
		set(type_a);
		set(type_b);
		for (vector<dot>::iterator it = type_a.begin(); it != type_a.end(); it++) {
			if (theta_2 == 0) {
				double x = -(double)theta_0 / (double)theta_1;
				if (it->x_dot > x)it->status = true;
			}
			else if (distance(theta_0, theta_1, theta_2, it->x_dot) > it->y_dot) {
				it->status = true;
			}
			if (it->status != type_a[0].status) {
				s += 1;
				break;
			}
		}
		if (s != 0) {
			cout << "No" << endl;
			continue;
		}
		for (vector<dot>::iterator it = type_b.begin(); it != type_b.end(); it++) {
			if (theta_2 == 0) {
				double x = -(double)theta_0 / (double)theta_1;
				if (it->x_dot > x)it->status = true;
			}
			else if (distance(theta_0, theta_1, theta_2, it->x_dot) > it->y_dot) {
				it->status = true;
			}
			if (it->status != type_b[0].status) {
				s += 1;
				break;
			}
		}
		if (s == 0) {
			cout << "Yes" << endl;
		}
		else cout << "No" << endl;
	}
	return 0;
}
