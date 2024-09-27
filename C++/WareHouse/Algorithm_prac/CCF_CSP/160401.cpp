//#include<iostream>
//#include<vector>
//using namespace std;
//int main() {
//	vector<int>v;
//	int n = 0;
//	cin >> n;
//	int t = n;
//	for (; n > 0; n--) {
//		int num = 0;
//		cin >> num;
//		v.push_back(num);
//	}
//	if (t == 1 || t == 2){
//		cout << "0" << endl; return 0;
//    }
//	int a = v.at(0);
//	int num_period = 1;
//	vector<int>v_1;
//	for (vector<int>::iterator it = v.begin() + 1; it != v.end(); it++) {
//		if (*it > a) v_1.push_back(1);
//		if (*it < a) v_1.push_back(2);
//		a = *it;
//	}
//	int a_1 = v_1.at(0);
//	int num_period_1 = 1;
//	for (vector<int>::iterator it = v_1.begin()+1; it != v_1.end(); it++) {
//		if (*it != a_1) num_period_1 += 1;
//		a_1 = *it;
//	}
//	cout << num_period_1 - 1 << endl;
//	return 0;
//}