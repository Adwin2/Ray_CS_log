//#include<iostream>
//#include<set>
//using namespace std;
//
//int main() {
//	set<int>s;
//	int sum = 0;
//	cin >> sum;
//	for (; sum > 0; sum--) {
//		int num = 0;
//		cin >> num;
//		s.insert(num);
//	}
//	int summary = 0;
//	for (set<int>::iterator it = s.begin(); it != s.end(); it++) {
//		int t = 0;
//		t -= *it;
//		if (s.find(t) == s.end()) {
//			continue;
//		}
//		else summary += 1;
//	}
//
//	cout << summary/2 << endl;
//
//	return 0;
//}