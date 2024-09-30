#include<iostream>
#include<vector>
#include<algorithm>
using namespace std;

//ostream& operator<<(ostream& out, vector<int>& v) {
//	for (vector<int>::iterator it = v.begin(); it != v.end(); it++) {
//		out << *it << " ";
//	}
//	return out;
//}

bool check(vector<int>v) {
	int a = v.at(0);
	for (vector<int>::iterator it = v.begin() + 1; it != v.end(); it++) {
		if (a != *it) return false;
	}
	return true;
}

int _compare(int a, int b) {
	return a > b;
}

int main() {
	int a;
	vector<int>v;
	cin >> a;
	for (; a > 0; a--) {
		int num = 0;
		cin >> num;
		v.push_back(num);
	}
	if (check(v)) {
		cout<<
	}
	sort(v.begin(), v.end(), _compare);
	//cout << v << endl;
	cout << v.at(0) << " ";
	double size = v.size();
	if ((int)size % 2 == 0) {
		printf("%.1f ",(v.at(v.size() / 2 - 1) + v.at(v.size() / 2)) / 2 );
	}
	else {
		cout << v.at((v.size() - 1) / 2) << " ";
	}
	cout << v.at(v.size() - 1) << endl;
	return 0;
}
