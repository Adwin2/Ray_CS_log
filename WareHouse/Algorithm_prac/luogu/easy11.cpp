//#include<iostream>
//#include<list>
//#include<vector>
//#include<set>
//using namespace std;
//class Solution {
//public:
//	int majorityElement(vector<int>& nums) {
//		set<int> s0(nums.begin(), nums.end());
//		multiset<int> s(nums.begin(), nums.end());
//		int x = nums.size() / 2;
//		int max = 0;
//		for (set<int>::iterator it = s0.begin(); it != s0.end(); it++) {
//			if (s.count(*it) > x) return *it;
//		}
//	}
//};
//int main() {
//	vector<int>v;
//	v.push_back(1);
//	for (int i = 1; i <= 10; i++) {
//		v.push_back(i);
//	}
//	
//	multiset<int>s(v.begin(),v.end());
//	/*for (int i = 0; i < v.size(); i++) {
//		s.insert(v[i]);
//	}*/
//	
//	cout << *s.end()<<" ";
//	cout << v.size()<<" ";
//	cout << s.count(1) << endl;
//
//
//	return 0;
//}