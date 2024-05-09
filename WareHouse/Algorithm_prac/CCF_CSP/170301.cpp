//#include<iostream>
//#include<vector>
//using namespace std;
//
//int main() {
//	vector<int>v;
//	int weight = 0;
//	int num_cake = 0;
//	cin >> num_cake >> weight;
//	for (; num_cake > 0; num_cake--) {
//		int num = 0;
//		cin >> num;
//		v.push_back(num);
//	}
//	vector<int>::iterator it = v.begin();
//	int sum = 0;
//	int time = 0;
//	while (it != v.end()) {
//		sum += *it;
//		if (sum >= weight) {
//			sum = 0;
//			time += 1;
//		}
//		it++;
//	}
//	if (sum != 0) {
//		time += 1;
//	}
//	cout << time << endl;
//	return 0;
//}