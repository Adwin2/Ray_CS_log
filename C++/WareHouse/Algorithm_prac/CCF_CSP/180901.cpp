
/////////////////////////////////////////////////////////////////////////////////////
//\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\
//#include<iostream>
//#include<vector>
//using namespace std;
//
//ostream& operator <<(ostream& out,vector<int>& v) {
//	for (vector<int>::iterator it = v.begin(); it != v.end(); it++) {
//		out << *it;
//	}
//	return out;
//}
//
//int main() {
//	int n1;
//	cin >> n1;
//	vector<int>vec;
//	for (int i = 0;i<n1;i++) {
//		int num = 0;
//		cin >> num;
//		vec.push_back(num);
//	}
//	vector<int>v;
//	v.push_back((vec.at(0)+vec.at(1)) / 2);
//	int n = 0;
//	for (vector<int>::iterator it = v.begin() + 1; it != v.end()-1; it++) {
//		v.push_back((vec.at(n) + *it + vec.at(n+2)) / 3);
//		n++;
//	}
//	v.push_back((vec.at(vec.size() - 2) + vec.at(vec.size() - 1)) / 2);
//	cout << v << endl;
//
//	return 0;
//}