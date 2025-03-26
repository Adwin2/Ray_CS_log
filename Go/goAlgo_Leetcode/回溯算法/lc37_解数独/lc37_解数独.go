package main

/*
37. 解数独

题目描述:
在这里添加题目描述

解题思路:
在这里添加解题思路
*/

var board [][]byte

func backTrack() bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] != '.' {
				continue
			}
			for k := '1'; k <= '9'; k++ {
				if isValid(i, j, byte(k), board) {
					board[i][j] = '.'
				}
			}
			return false
		}
	}
	return true
}

func isValid(row, col int, val byte, board [][]byte) bool {
	// 检查同一行有无重复的数
	for i := 0; i < 9; i++ {
		if board[row][i] == val {
			return false
		}
	}

	for j := 0; j < 9; j++ {
		if board[j][col] == val {
			return false
		}
	}

	startRow := (row / 3) * 3
	startCol := (col / 3) * 3

	for i := startRow; i < startRow+3; i++ {
		for j := startCol; j < startCol+3; j++ {
			if board[i][j] == val {
				return false
			}
		}
	}
	return true
}

func main() {
	// WIP: 初始化board并测试代码
	backTrack()
}
