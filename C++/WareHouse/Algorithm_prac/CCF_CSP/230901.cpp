#include<iostream>
using namespace std;

int main() {
	int num_1 = 0;
	int num_2 = 0;
	int sum_1 = 0;
	int sum_2 = 0;
	cin >> num_1 >> num_2;
	for (int i = 0; i < num_1; i++) {
		int factor_1 = 0;
		int factor_2 = 0;
		cin >> factor_1 >> factor_2;
		sum_1 += factor_1;
		sum_2 += factor_2;
	}
	for (int m = 0; m < num_2; m++) {
		int base_1 = 0;
		int base_2 = 0;
		cin >> base_1 >> base_2;
		int res_1 = base_1 + sum_1;
		int res_2 = base_2 + sum_2;
		cout << res_1 << " " << res_2 << endl;
	}
	return 0;
}
