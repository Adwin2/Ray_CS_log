#include<iostream>
#include<tr1/unordered_map>
#include <tr1/unordered_set>
using namespace std::tr1;
using namespace std;
#include<vector>
class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        __unordered_map<int,int> map;
        for(int i = 0;  i < nums.size(); i ++){
            auto iter = map.find(target - nums[i]);
            if(iter != map.end()){
                return {iter->second, i};
            }
            map.insert(pair<int,int>(nums[i],i));
        }
        return { };
    }
};

int main(){
    Solution s;
    vector<int> v;
    v.push_back(1);
    v.push_back(2);
    v.push_back(3);
    vector<int> it = s.twoSum(v,3);
    for(int i = 0; i < v.size(); i++) {
        cout<<it[i]<<endl;
    }
    return 0;
}