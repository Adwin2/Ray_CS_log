////递归解决plusone问题
////效率欠佳，思路可鉴😊
//
//#include<iostream>
//
//#include<vector>
//using namespace std;
//
//ostream& operator<<(ostream& out, vector<int>& v) {
//
//    for (vector<int>::iterator it = v.begin(); it != v.end(); it++) {
//        out << *it ;
//    }
//    return out;
//}
//
//class Solution {
//public:
//    void dealNine(vector<int>& digits, int i) {
//        if (i == 0 && digits[i] == 9) {
//            digits[0] = 1;
//            digits.push_back(0);
//            return;
//        }
//        digits[i] = 0;
//        int j = i - 1;
//        if (digits[j] == 9) {
//            
//            --i;
//            dealNine(digits, i);
//        }
//        else {
//            digits[j] += 1;
//            return;
//        }
//    }
//    vector<int> plusOne(vector<int>& digits) {
//        int i = digits.size() - 1;
//        if (i == 0 && digits[i] == 9) { digits[0] = 1; digits.push_back(0); return digits; }
//        if (digits[i] == 9) dealNine(digits, i);
//        else digits[i] += 1;
//        return digits;
//    }
//};
//
//int main() {
//    vector<int> digits;
//    digits.push_back(9);
//    digits.push_back(9);
//    digits.push_back(9);
//    Solution s;
//    vector<int> result = s.plusOne(digits);
//    cout << result << endl;
//
//
//    return 0;
//}