////����tempֵ��ȷ�� �Ƿ���Ҫ��һλ
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
//		//  res.resize(b.size()+1);//���ַ�����������һ����λ
//		// �ĳ�������insert����
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
//				//temp һ�� ��1������Ҫ ���¸�ֵ
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
////���� �� ����ʱ�������Ч�ʽϵ�
////�Ż�˼·������ '0'��'1' ��Ϊ�ַ�֮��Ĳ� ���Դ�����
////�� t ������ a��b�ַ��� ��Ӧ�ַ� ��'0'�Ĳ� 
////  t%2 ��ʾ ��ǰλ Ӧ����д��ֵ  �� t/2 ��ʾ ��Ҫ ����λ
////  һ���зǳ�ţ�Ƶ�˼·����
//// ���ŷ������ײ�
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
////			ans = to_string(carry % 2) + ans; // ��ǰλ��ֵ ����ͷ���ķ�ʽ
////			carry /= 2;                       // carry�����Ƿ��λ��Ϊ0��1���������������һ��ѭ��
////
////		}
////		return ans;
////	}
////};