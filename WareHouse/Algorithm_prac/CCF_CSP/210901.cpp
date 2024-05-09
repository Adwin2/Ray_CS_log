//#include<iostream>
//#include<numeric>
//#include<vector>
//using namespace std;
//
//int main() {
//	vector<int>vec;
//	int n = 0;
//	cin >> n;
//	for (; n > 0; n--) {
//		int num = 0;
//		cin >> num;
//		vec.push_back(num);
//	}
//	//最大值
//	int sum = accumulate(vec.begin(),vec.end(),0);
//	cout << sum << endl;
//
//	//最小值
//	int a_1 = vec.at(0);
//	for (vector<int>::iterator it = vec.begin()+1; it != vec.end(); it++) {
//		int temp = *it;
//		if (*it == a_1) *it = 0;
//		a_1 = temp;
//	}
//	int sum_1 = accumulate(vec.begin(), vec.end(), 0);
//	cout << sum_1 << endl;
//
//	return 0;
//}