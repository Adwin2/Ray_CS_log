////引入temp值来确定 是否需要进一位
//
//
//#include<iostream>
//#include<vector>
//#include<string>
//using namespace std;
//
//class Solution {
//public:
//	int trans(char x) {
//		if (x == '1') return 1;
//		if (x == '0') return 0;
//		else return -1;
//	}
//	char trans_1(int x) {
//		if (x == 0) return '0';
//		if (x == 1) return '1';
//		else return '?';
//	}
//	string addBinary(string a, string b) {
//		if (a.length() > b.length()) b.insert(0, a.length() - b.length(), '0');
//		if (a.length() < b.length()) a.insert(0, b.length() - a.length(), '0');
//		int len = a.length();
//		string res(b);
//		res.reserve(b.size() + 1);
//
//		//  res.resize(b.size()+1);//在字符串的最后加了一个空位
//		// 改成在最后加insert函数
//		int temp = 0;
//		for (int i = len-1; i >= 0; i--) {
//
//			int sum = trans(a[i]) + trans(b[i]) + temp;
//			cout << sum << endl;
//			if (sum == 2) {
//				res[i] = '0'; 
//				temp = 1;
//				if (i == 0) res.insert(0, 1, '1');
//			}
//			else if (sum == 3) {
//				//temp 一定 是1，不需要 重新赋值
//				res[i] = '1';
//				if (i == 0) res.insert(0, 1, '1');
//			}
//			else {
//				char mychar = trans_1(sum);
//				res[i] = mychar;
//				cout << mychar << endl;
//				temp = 0;
//			}
//		}
//		return res;
//	}
//};
//
////问题 ： 运行时间过长，效率较低
////优化思路：由于 '0'和'1' 作为字符之间的差 可以代表数
////用 t 来计算 a和b字符串 对应字符 与'0'的差 
////  t%2 表示 当前位 应该填写的值  ， t/2 表示 需要 进的位
////  一大佬非常牛逼的思路！！
//// 最优方案见底部
//
//int main() {
//	string a = "1010";
//	string b = "10111";
//	Solution s;
//	cout << s.addBinary(a, b) << endl;
//
//	return 0;
//}
//
////class Solution {
////public:
////	string addBinary(string a, string b) {
////		int m = a.size() - 1, n = b.size() - 1;
////		string ans;
////		int carry = 0;
////		while (m >= 0 || n >= 0 || carry) {
////			if (m >= 0) carry += a[m--] - '0';
////			if (n >= 0) carry += b[n--] - '0';
////			ans = to_string(carry % 2) + ans; // 当前位的值 加在头部的方式
////			carry /= 2;                       // carry根据是否进位分为0或1两情况，并进入下一次循环
////
////		}
////		return ans;
////	}
////};