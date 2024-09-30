#include<iostream>
#include<vector>
#include<algorithm>
using namespace std;

int main() {
	int n = 0;
	vector<int>v;
	cin >> n;
	for (; n > 0; n--) {
		int num = 0;
		cin >> num;
		v.push_back(num);
	}
	sort(v.begin(), v.end());
	int a = v.at(0);
	int num_1 = 0;
	//vector<int>v_1;
	for (vector<int>::iterator it = v.begin() + 1; it != v.end(); it++) {
		//v_1.push_back(abs(*it - a));
		if (*it - a == 1) {
			num_1 += 1;
		}
		a = *it;
	}
	//sort(v_1.begin(), v_1.end());
	//cout << v_1.at(0) << endl;
	cout << num_1 << endl;
	return 0;
}
