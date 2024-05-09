//5������ --------- DFS

//ֻ�� 1��3��9��27��81 �������1~121 �����Է�����������--���Ӽ�������
//����������� �� С��ǰ
#include<iostream>
#include<vector>
#include<string>
using namespace std;
vector<int> weight = { 81, 27, 9, 3, 1 };
int n;

// ���������������
void dfs(int i, int sum, string ans) {
    // �����ǰ��ϵ���������Ŀ�������������ǰ��Ϸ����������ݹ�
    if (sum == n) {
        cout << ans.substr(1); // �����ǰ��Ϸ�����ȥ����ͷ�ļӺŻ���ţ�
        return;
    }

    // �����ǰ��������������Χ��ݹ鵽�ף������ݹ�
    if (i == 5) return;

    // ��������ѡ�񣺽���ǰ����������̡����̻򲻷ţ����ݹ�ص���dfs����
    dfs(i + 1, sum + weight[i], ans + "+" + to_string(weight[i])); // ����ǰ�����������
    dfs(i + 1, sum, ans); // ���ŵ�ǰ����
    dfs(i + 1, sum - weight[i], ans + "-" + to_string(weight[i])); // ����ǰ�����������
}

int main() {
    cin >> n; // ����Ŀ������

    dfs(0, 0, ""); // �������������������Ѱ����Ч���

    return 0;
}