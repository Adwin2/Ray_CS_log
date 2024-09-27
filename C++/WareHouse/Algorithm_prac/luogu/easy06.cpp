////kmp算法
////寻找匹配子串问题
//// 避免了重复对应问题
////记录需要再次遍历的起点
////关键在于 递推建立next数组
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
//        next.push_back(0);//哨兵
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
//        return j >= p.length() ? i-p.length() : -1;//返回值以及如何判断没有匹配的子串条件  待思考
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
//    ////构造函数 调试
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