//#include<iostream>
//using namespace std;
//
//#include<vector>
//
//class Solution {
//public:
//	int removeElement(vector<int>& nums, int val) {
//		int remain = 0;
//		for (int i = 0; i < nums.size(); i++)
//		{
//			if (nums[i] != val) {
//				nums[remain] = nums[i];
//				remain++;
//			}
//		}
//		return remain;
//	}
//};
//
//int main() {
//	vector<int> nums;
//	for (int i = 0; i < 4; i++) {
//		nums.push_back(i);
//	}
//
//	int val = 3;
//	Solution s;
//	cout<<s.removeElement(nums, val) << endl;
//
//
//	return 0;
//}