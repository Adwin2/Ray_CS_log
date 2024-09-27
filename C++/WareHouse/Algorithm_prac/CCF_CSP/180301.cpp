//#include<iostream>
//#include<vector>
//using namespace std;
//
//int main()
//{
//	vector<int>v;
//	int n = 0;
//	while (true) {
//		cin >> n;
//		if (n == 0) break;
//		v.push_back(n);
//	}
//	int sum = 0;
//	int time = 0;
//	int temp = v.at(0);
//	for (vector<int>::iterator it = v.begin(); it != v.end(); it++) {
//		if (*it == 1) {
//			sum += 1;
//			time = 0;
//		}
//		if (*it == 2 && temp == *it) {
//			time += 1;
//			sum += 2 * time;
//			continue;
//		}
//		if (*it == 2) {
//			sum += 2;
//			time += 1;
//		}
//		temp = *it;
//	}
//	cout << sum << endl;
//	return 0;
//}