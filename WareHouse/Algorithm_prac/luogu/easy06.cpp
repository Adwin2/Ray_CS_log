////kmp�㷨
////Ѱ��ƥ���Ӵ�����
//// �������ظ���Ӧ����
////��¼��Ҫ�ٴα��������
////�ؼ����� ���ƽ���next����
////
////
//
//
//
//#include<iostream>
//using namespace std;
//
//#include<string>
//#include<vector>
//
//
//ostream& operator<<(ostream& out, vector<int>& v) {
//
//    for (vector<int>::iterator it = v.begin(); it != v.end(); it++) {
//        out << *it ;
//    }
//    return out;
//}
//class Solution {
//public:
//    vector<int> buildNext(string& p) {
//        vector<int> next;
//        next.reserve(p.length()+1);
//        p = " " + p;
//        next.push_back(0);//�ڱ�
//        next.push_back(0);
//        next.push_back(1);
//        int i = 1;
//        for (int j = 2; j < p.length() - 1; j++) {
//            if (p[i] == p[j]) {
//                i++;
//                next.push_back(i);
//            }
//            else {
//                next.push_back(1);
//                i = 1;
//            }
//        }
//
//        return next;
//    }
//    int strStr(string s, string p, vector<int>nextval) {
//        int j = 1;
//        int i = 1;
//        while (i < s.length() && j < p.length()) {
//            if (s[i] == p[j]) {
//                i++;
//                j++;
//            }
//            else {
//                j = nextval[j];
//            }
//        }
//        return j >= p.length() ? i-p.length() : -1;//����ֵ�Լ�����ж�û��ƥ����Ӵ�����  ��˼��
//    }
//};
//
//int main() {
//    string a = "sadb";
//    string b = "ssadbutsad";
//    string c = "salute";
//    a = ' ' + a;
//
//    Solution s;
//    vector<int> v1 = s.buildNext(b);
//
//    ////���캯�� ����
//    //vector<int> next(++v1.begin(), v1.end());
//    //vector<int> nextval(++v2.begin(), v2.end());
//    //cout << next << " ";
//    //cout << nextval << endl;
//
//    cout << s.strStr(b, a, v1) << endl;
//    
//
//
//    return 0;
//}