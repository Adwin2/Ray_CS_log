#include<iostream>
#include<cmath>
using namespace std;

int main() {
	int n_years = 0;
	double rate = 0;
	double sum = 0;
	cin >> n_years >> rate;
	for (int i = 0; i <= n_years; i++) {
		int n_brunch = 0;
		cin >> n_brunch;
		sum += (double)n_brunch * pow((1 + rate), -i);
	}
	cout << sum << endl;
	return 0;
}
