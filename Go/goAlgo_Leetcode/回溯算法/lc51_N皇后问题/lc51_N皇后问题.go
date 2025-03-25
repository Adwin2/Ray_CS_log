package main

import "fmt"

/*
51. N皇后问题

题目描述:
在这里添加题目描述

解题思路:
在这里添加解题思路
*/

var res [][][]byte

// 抽象为树 横向遍历（更新）col：for col ; 纵向遍历（更新）row :更新树层需要 row+1
func backtrack(n, row int, chessboard [][]byte) {
	// 最后一层遍历结束
	if row == n {
		res = append(res, cloneSlice(chessboard))
		return
	}

	for col := 0; col < n; col++ {
		if isValid(row, col, n, chessboard) {
			chessboard[row][col] = 'Q'
			backtrack(n, row+1, chessboard)
			chessboard[row][col] = '.'
		}
	}
}

// 传入当前row, col 并检查之前（向下遍历）是否有重复
func isValid(row, col, n int, chessboard [][]byte) bool {
	// Note: 不用检查行 每行属于一次树层遍历 只会操作一次

	// 检查列
	for i := 0; i < row; i++ {
		// 检查是否放置了'Q'
		if chessboard[i][col] == 'Q' {
			return false
		}
	}
	// 检查45度线上 向左上挪一格
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if chessboard[i][j] == 'Q' {
			return false
		}
	}

	// 检查135度线上 向右上挪一格
	for i, j := row-1, col+1; i >= 0 && j <= n-1; i, j = i-1, j+1 {
		if chessboard[i][j] == 'Q' {
			return false
		}
	}

	return true
}

// utils --
func cloneSlice(slice [][]byte) [][]byte {
	newSlice := make([][]byte, len(slice))
	for i, inner := range slice {
		newSlice[i] = make([]byte, len(inner))
		copy(newSlice[i], inner)
	}
	return newSlice
}

func printSliceOfSliceOfSlices(data [][][]byte) {
	for i, sliceOfSlices := range data {
		fmt.Printf("Slice %d:\n", i+1)
		for j, inner := range sliceOfSlices {
			fmt.Printf("  Inner slice %d: ", j+1)
			for _, b := range inner {
				fmt.Printf("%c ", b)
			}
			fmt.Println()
		}
	}
}

// -- utils

// row = col = n
func combine(n int) {
	chessboard := make([][]byte, n)
	for i := range chessboard {
		chessboard[i] = make([]byte, n)
		for j := range chessboard[i] {
			chessboard[i][j] = '.'
		}
	}

	backtrack(n, 0, chessboard)
}

func main() {
	combine(4)

	printSliceOfSliceOfSlices(res)
}
