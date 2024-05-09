//#include<iostream>
//#include<vector>
//using namespace std;
//
//int main()
//{
//	int row = 0;
//	int column = 0;
//	cin >> row >> column;
//	// 动态分配二维数组
//	int** arr = (int**)malloc(row * sizeof(int*));
//	for (int i = 0; i < row; i++) {
//		arr[i] = (int*)malloc(column * sizeof(int));
//	}
//	for (int m = 0; m < row; m ++) {
//		for (int n = 0; n < column; n ++) {
//			cin >> arr[m][n];
//
//		}
//	}
//
//	for (int m = 1; m <= column; m++) {
//		for (int n = 0; n < row; n++) {
//			cout<<arr[n][column-m]<<" ";
//		}
//		cout << endl;
//	}
//
//
//	for (int i = 0; i < row; i++) {
//		free(arr[i]);
//	}
//	free(arr);
//	return 0;
//}