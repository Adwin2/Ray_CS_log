#include<iostream>
#include<vector>
using namespace std;

class number {
public:
	int num = 0;
	int time = 0;
};

void my_push_back(vector<number>& vec, int value) {
	number n;
	n.num = value;
	n.time += 1;
	vec.push_back(n);
}

int main() {
	vector<number>base;
	int a = 0;
	cin >> a;
	for (; a > 0; a--) {
		int num = 0;
		cin >> num;
		my_push_back(base, num);
	}



	return 0;
}
