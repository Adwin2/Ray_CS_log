#include<iostream>
#include<vector>

using namespace std;


bool compareMatrices(vector<vector<char>>& matrix1, vector<vector<char>>& matrix2) {
    if (matrix1.size() != matrix2.size() || matrix1[0].size() != matrix2[0].size()) {
        return false; // 矩阵大小不同，直接返回false
    }

    for (int i = 0; i < matrix1.size(); i++) {
        for (int j = 0; j < matrix1[0].size(); j++) {
            if (matrix1[i][j] != matrix2[i][j]) {
                return false; // 存在不相等的元素，返回false
            }
        }
    }

    return true; // 所有元素相等，返回true
}

int main(){
    vector<vector<char>> matrix1;
    vector<vector<char>> matrix2;
    int i_0 = 0;
    cin>>i_0;
 


    if (compareMatrices(matrix1, matrix2)) {
        cout << "两个矩阵相等" << endl;
    } else {
        cout << "两个矩阵不相等" << endl;
    }

    return 0;
}