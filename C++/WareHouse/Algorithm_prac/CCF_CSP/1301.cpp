//#include<iostream>
//#include<set>
//using namespace std;
//
//int main() {
//	multiset<int>ms;
//	set<int>s;
//	set<int>res;
//	int sum = 0;
//	cin >> sum;
//	for (; sum > 0; sum--) {
//		int num = 0;
//		cin >> num;
//		ms.insert(num);
//		s.insert(num);
//	}
//	int max = 0;
//	for (set<int>::iterator it = s.begin(); it != s.end(); it++) {
//		if (ms.count(*it) > max){
//			max = ms.count(*it);
//		}
//	}
//	for (set<int>::iterator it = s.begin(); it != s.end(); it++) {
//		if (ms.count(*it) == max) {
//			res.insert(*it);
//		}
//	}
//	cout << *res.begin() << endl;
//	return 0;
//}