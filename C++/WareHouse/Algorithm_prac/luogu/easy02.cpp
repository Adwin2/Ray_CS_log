//
////最长公共前缀  strs = ["flower","flow","flight"]         ->"fl"
//
//
//
//
//#include<iostream>
//#include<vector>
//#include<string>
//using namespace std;
//
//class Solution {
//public:
//    string   longestCommonPrefix(vector<string>& strs) {
//        int min = strs.at(0).length();
//        for (int i = 1; i < strs.size(); i++) {
//            if (strs[i].length() < min) {
//                min = strs[i].length();
//            }
//        }
//        string b = "";
//        string a = strs[0];
//        int num = 0;
//        for (int i = 0; i < min; i++) {
//            for (int m = 0; m < strs.size(); m++) {
//                if (a.at(i) == strs[m].at(i)) num++;
//            }
//            if(num == strs.size()) b += a.at(i);
//            else  return b;
//            num = 0;
//        }
//        return b;
//    }
//};
//
//int main() {
//    Solution s;
//    vector<string> strs;
//    strs.push_back("aaa");
//    strs.push_back("aaaa");
//    strs.push_back("aab");
//    strs.push_back("aac");
//
//    cout<<s.longestCommonPrefix(strs)<<endl;
//
//
//    return 0;
//}