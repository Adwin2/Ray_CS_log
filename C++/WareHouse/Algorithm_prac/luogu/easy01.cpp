////力扣——赎金信  a中字符b都有包含
//#include<iostream>
//#include<string>
//#include<vector>
//using namespace std;
//
//class Solution {
//public:
//    bool canConstruct(string ransomNote, string magazine) {
//        int i = 0;
//        string aa = "";
//        while (magazine.size() != 0) {
//            char a = ransomNote.at(i);
//            int pos = magazine.find(a);
//
//            if (pos != -1) {
//                aa+=a;
//                magazine.erase(pos,1);
//            }
//            if (i == ransomNote.size()-1) {
//                break;
//            }
//            i++;
//        }
//        if (aa.compare(ransomNote) == 0) return true;
//        return false;
//    }
//};
//
//
////cout << magazine.size() << endl;
////cout << i << endl;
////cout << aa << endl;
////cout << ransomNote << endl;
////cout << magazine << endl;
////cout调试法，hhh
//int main() {
//    string a = "aab";
//    string b = "aab";
//    Solution s;
//    if (s.canConstruct(a, b)) cout << "true" << "";
//    else cout << "false" << endl;
//
//    return 0;
//}
//
//
