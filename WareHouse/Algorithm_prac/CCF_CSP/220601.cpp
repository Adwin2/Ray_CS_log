//#include<iostream>
//#include<vector>
//#include<algorithm>
//#include<cmath>
//using namespace std;
//
//int main() {
//	int n = 0;
//	cin >> n;
//
//	vector<int>v;
//	int sum = 0;
//	for (int i = 0; i < n; i++) {
//		int num = 0;
//		cin >> num;
//		sum += num;
//		v.push_back(num);
//	}
//	double aver = (double)sum / n;
//	double sum_fc = 0;
//	for_each(v.begin(), v.end(), [&sum_fc,aver] (int num){
//		sum_fc += pow((num - aver), 2);
//		});
//	double D_a = sum_fc / n;
//
//	for_each(v.begin(), v.end(), [aver,D_a](int a) {
//		printf("%.16f\n", (a - aver) / sqrt(D_a));
//		});
//
//	return 0;
//}
