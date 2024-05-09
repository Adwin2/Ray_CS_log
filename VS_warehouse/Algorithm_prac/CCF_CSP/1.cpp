
//�� �� �� �� ��
//k,q,r,b,n,p  ��д�� Сд�� ����*
#include<iostream>
#include<string>
#include<vector>
//#include<cstdlib>
using namespace std;

int check(string s,vector<string>v) {
	int num = 0;
	for (vector<string>::iterator it = v.begin(); it != v.end(); it++) {
		if (*it == s) num += 1;
	}
	return num;
}

int main() {
	int n = 0;
	cin >> n;
	vector<string>v;
	for (int i = 0; i < n; i++) {
		string str;
		string temp = "";
		for (int m = 0; m < 8; m++) {
			cin >> str;
			temp += str[m];
		}
		v.push_back(temp);
	}
	for (vector<string>::iterator it = v.begin(); it != v.end(); it++) {
		vector<string> v_0(v.begin(), it+1);
		cout << check(*it,v_0) << endl;
	}
	return 0;
}