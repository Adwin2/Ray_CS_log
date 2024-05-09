//#include<iostream>
//using namespace std;
//
//bool check(int n) {
//	while (n > 0) {
//		if (n % 10 == 7)return true;
//		n /= 10;
//	}
//	return false;
//}
//
//int main() {
//	int n = 0;
//	int n_1 = 0;
//	int n_2 = 0;
//	int n_3 = 0;
//	int n_4 = 0;
//	cin >> n;
//	for (int i = 1; i <= n; i++) {
//		if (check(i) || i % 7 == 0) {
//			cout << i << endl;
//			int m = i % 4;
//			if (m == 1) n_1 += 1;
//			if (m == 2) n_2 += 1;
//			if (m == 3) n_3 += 1;
//			if (m == 0) n_4 += 1;
//			n++;
//		}
//	}
//	cout << n_1 << endl;
//	cout << n_2 << endl;
//	cout << n_3 << endl;
//	cout << n_4 << endl;
//	return 0;
//}