#include<iostream>
#include<vector>
using namespace std;

bool operator>(vector<int>& v_1, vector<int>& v_2) {
	for (int i = 0; i < v_1.size(); i++) {
		if (v_1[i] <= v_2[i]) return false;
	}
	return true;
}

int main() {
	int num = 0;
	int dim = 0;
	cin >> num >> dim;
	vector<vector<int>>v_base;
	for (int i = 0; i < num; i++) {
		vector<int>v;
		for (int m = 0; m < dim; m++) {
			int num = 0;
			cin >> num;
			v.push_back(num);
		}
		v_base.push_back(v);
	}

	for (vector<vector<int>>::iterator it = v_base.begin(); it != v_base.end(); it++) {
		bool istrue = true;
		for (int i = 0; i < v_base.size(); i++) {
			if (v_base.at(i) > *it) {
				cout << i + 1 << endl;
				istrue = false;
				break;
			}
		}
		if (istrue) cout << 0 << endl;
	}

	return 0;
}
