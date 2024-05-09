//#include<iostream>
//#include<vector>
//using namespace std;
//
//int main() {
//	vector<int>v;
//	int sum = 0;
//	cin >> sum;
//	for (; sum > 0; sum--) {
//		int num = 0;
//		cin >> num;
//		v.push_back(num);
//	}
//	int a = v.at(0);
//	int num_period = 1;
//	for (vector<int>::iterator it = v.begin()+1; it != v.end(); it++) {
//		if (*it != a) num_period += 1;
//		a = *it;
//	}
//	cout << num_period << endl;
//
//	return 0;
//}