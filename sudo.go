package main

import "fmt"

// 定義一個數獨矩陣
var matrix = [9][9]int{
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
}

func main() {
	// 解數獨
	solveSudoku(0, 0)

	// 輸出矩陣
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}
}

// solveSudoku 解數獨
func solveSudoku(row, col int) bool {
	// 如果已經遍歷到了最後一行，说明數獨已经解完
	if row == 9 {
		return true
	}

	// 如果當前位置已經有值，則跳過這個位置，直接遍歷下一個位置
	if matrix[row][col] != 0 {
		nextRow, nextCol := nextPos(row, col)
		return solveSudoku(nextRow, nextCol)
	}

	// 否則，枚舉 1-9 的數字，判断能否放入當前位置
	for i := 1; i <= 9; i++ {
		if isValid(row, col, i) {
			matrix[row][col] = i
			nextRow, nextCol := nextPos(row, col)
			// 递归解下一個位置
			if solveSudoku(nextRow, nextCol) {
				return true
			}

			// 如果不能放入，則將當前位置的值清空，繼续枚舉下一個數字
			matrix[row][col] = 0
		}
	}

	// 如果 1-9 的數字都枚舉完了，還是無法放入，那就返回 false
	return false
}

// nextPos 返回下一個位置的行索引和列索引
func nextPos(row, col int) (int, int) {
	if col == 8 {
		return row + 1, 0
	}
	return row, col + 1
}

// isValid 判断能否在指定位置放入指定數字
func isValid(row, col, num int) bool {
	// 判断行
	for i := 0; i < 9; i++ {
		if matrix[row][i] == num {
			return false
		}
	}

	// 判断列
	for i := 0; i < 9; i++ {
		if matrix[i][col] == num {
			return false
		}
	}

	// 判断 3x3 的小矩陣
	rowStart := row / 3 * 3
	colStart := col / 3 * 3
	for i := rowStart; i < rowStart+3; i++ {
		for j := colStart; j < colStart+3; j++ {
			if matrix[i][j] == num {
				return false
			}
		}
	}

	// 如果滿足上述所有條件，則表示可以放入
	return true
}
