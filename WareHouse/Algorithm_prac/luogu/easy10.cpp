////动态规划方面的题目
//
//#include<iostream>
//#include<vector>
//#include<string>
//using namespace std;
//class Solution {                                                      
//public:
//    int minimumTime(vector<int>& nums1, vector<int>& nums2, int x) {
//        int sum = 0;
//        vector<int> v(nums1);
//        for (int i = 0; i < nums1.size(); i++) {
//            v[i] = 0;
//            for (int m = 0; m < v.size(); m++) {
//                if (m == i) continue;
//                v[m] += nums2[m];
//                sum += v[m];
//            }
//            if (sum <= x) {
//                return i + 1;
//            }
//            sum = 0;
//        }
//        return -1;
//    }
//};
//
//int main() {
//    vector<int> v1;
//    v1.push_back(4);
//    v1.push_back(4);
//    v1.push_back(9);
//    v1.push_back(10);
//
//    vector<int>v2;
//    v2.push_back(4);
//    v2.push_back(4);
//    v2.push_back(1);
//    v2.push_back(3);
//
//    int x = 16;
//    Solution s;
//    cout << s.minimumTime(v1,v2,x) << endl;;
//
//    return 0;
//}