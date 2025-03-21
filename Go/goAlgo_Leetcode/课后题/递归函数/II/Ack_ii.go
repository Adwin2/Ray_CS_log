// 递归转迭代
package main

import "fmt"

func ack_ii(m, n int) int {
	// 初始化二维数组 ，var vec [m][n]int 错误 编译时需要确定大小 不可使用变量初始化
	vec := make([][]int, m+2)
	for i := range m + 2 {
		vec[i] = make([]int, n+2)
	}

	// m == 0
	for i := 0; i <= n; i++ {
		vec[0][i] = i + 1
	}

	for i := 1; i <= m; i++ {
		// m != 0 && n == 0
		vec[i][0] = vec[i-1][1]
		// m != 0 && n != 0
		for o := 1; o <= n; o++ {
			vec[i][o] = vec[i-1][vec[i][o-1]]
		}
	}
	return vec[m][n]
}

func main() {
	fmt.Println(ack_ii(2, 1))
}
