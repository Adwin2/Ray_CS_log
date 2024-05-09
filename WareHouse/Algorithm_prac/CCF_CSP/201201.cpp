//#include<iostream>
//using namespace std;
//
//int main() {
//	int n = 0;
//	cin >> n;
//	int** arr = (int**)malloc(n * sizeof(int*));
//	for (int i = 0; i < n; i++) {
//		arr[i] = (int*)malloc(2 * sizeof(int));
//	}
//	int sum = 0;
//	for (int i = 0; i < n; i++) {
//		for (int i_1 = 0; i_1 < 2; i_1++) {
//			int num = 0;
//			cin >> num;
//			arr[i][i_1] = num;
//		}
//		sum += arr[i][0] * arr[i][1];
//	}
//
//	if (sum >= 0)sum = sum;
//	else {
//		sum = 0;
//	}
//	cout << sum << endl;
//	return 0;
//}