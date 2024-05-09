//#include<iostream>
//using namespace std;
//#include<bitset>
//#include<vector>
//#include<string>
//class Solution {
//public:
//	string addBinary(string a, string b) {
//		if (a[0] == '0' && b[0] == '0') {
//			string a = "0";
//			return a;
//		}
//		bitset<10000> b1(a);
//		bitset<10000> b2(b);
//		bitset<10000> result = b1.to_ullong() + b2.to_ullong();
//		string str = result.to_string();
//		return str.substr(str.find('1'),str.length() - str.find('1'));
//	}
//};
//
//
//int main() {
//	//string a = " aa";
//	//string b = "bb";
//	string num = "100010010";
//	//char aa = 'c';
//	//a.insert(0,1, aa);
//	//cout << a << endl;
//	bitset < 2> a = 0b11;
//	bitset<2> b = 0b11;
//	cout <<bitset<3> (0b11+0b01)<< endl;
//	const int aa = 9;
//	bitset<aa> num1(num);
//	//bitset<8> num2(b);
//	cout << num1<<endl;
//	//bitset<8> result = num1.to_ulong() + num2.to_ulong();
//	//string str = result.to_string();
//	/*cout << stoi(num) << endl;
//	int aaa = 11111;*/
//	string str = num1.to_string();
//	cout << str << endl;
//	return 0;
//}
