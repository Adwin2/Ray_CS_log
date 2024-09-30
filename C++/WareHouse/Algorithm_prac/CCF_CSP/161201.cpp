#include<iostream>
#include<vector>
#include<algorithm>
using namespace std;

int main() {
	vector<int>v;
	int n = 0;
	cin >> n;
	int t = n;
	for (; n > 0; n--) {
		int num = 0;
		cin >> num;
		v.push_back(num);
	}
	sort(v.begin(), v.end());
	if (t == 1 || t == 2) { cout << "-1" << endl; return 0; }
	if (t % 2 == 0 && (v.at((t / 2) - 1) != v.at(t / 2))) { cout << "-1" << endl; return 0; }
	else {
		if (t % 2 == 1) {
			if()
		}
	}

	return 0;
}
