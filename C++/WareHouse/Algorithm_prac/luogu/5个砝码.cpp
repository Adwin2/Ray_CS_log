//5个砝码 --------- DFS

//只有 1，3，9，27，81 可以组合1~121 ，可以放在左右两侧--》加减都可以
//输出，，大数 在 小数前
#include<iostream>
#include<vector>
#include<string>
using namespace std;
vector<int> weight = { 81, 27, 9, 3, 1 };
int n;

// 深度优先搜索函数
void dfs(int i, int sum, string ans) {
    // 如果当前组合的重量等于目标重量，输出当前组合方案并结束递归
    if (sum == n) {
        cout << ans.substr(1); // 输出当前组合方案（去除开头的加号或减号）
        return;
    }

    // 如果当前砝码索引超出范围或递归到底，结束递归
    if (i == 5) return;

    // 尝试三种选择：将当前砝码放在左盘、右盘或不放，并递归地调用dfs函数
    dfs(i + 1, sum + weight[i], ans + "+" + to_string(weight[i])); // 将当前砝码放在左盘
    dfs(i + 1, sum, ans); // 不放当前砝码
    dfs(i + 1, sum - weight[i], ans + "-" + to_string(weight[i])); // 将当前砝码放在右盘
}

int main() {
    cin >> n; // 输入目标重量

    dfs(0, 0, ""); // 调用深度优先搜索函数寻找有效组合

    return 0;
}