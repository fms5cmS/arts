package backtrackingRela

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func solveNQueens(n int) [][]string {
	result := make([][]string, 0)
	// 保存临时的棋盘，注意，棋盘的边长为皇后的总数量！
	chessboard := make([][]string, n)
	for i := 0; i < n; i++ {
		chessboard[i] = make([]string, n)
		for j := 0; j < n; j++ {
			chessboard[i][j] = "."
		}
	}
	var backtracking func(chessboard [][]string, row int)
	backtracking = func(chessboard [][]string, row int) {
		if row == n {
			temp := make([]string, n)
			for i := 0; i < n; i++ {
				temp[i] = strings.Join(chessboard[i], "")
			}
			result = append(result, temp)
			return
		}
		for col := 0; col < n; col++ {
			if !isValidFor51(row, col, chessboard) {
				continue
			}
			chessboard[row][col] = "Q"
			backtracking(chessboard, row+1)
			chessboard[row][col] = "."
		}
	}
	backtracking(chessboard, 0)
	return result
}

func isValidFor51(row, col int, chessboard [][]string) bool {
	// 棋盘的边长
	length := len(chessboard)
	// 检查同一列是否存在皇后
	for i := 0; i < row; i++ {
		if chessboard[i][col] == "Q" {
			return false
		}
	}
	// 检查当前位置左上对角线上是否存在皇后
	for i, j := row, col; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if chessboard[i][j] == "Q" {
			return false
		}
	}
	// 检查当前位置右上对角线上是否存在皇后
	for i, j := row, col; i >= 0 && j < length; i, j = i-1, j+1 {
		if chessboard[i][j] == "Q" {
			return false
		}
	}
	return true
}

func TestSolveNQueue(t *testing.T) {
	tests := []struct {
		input  int
		output [][]string
	}{
		{input: 4, output: [][]string{{".Q..", "...Q", "Q...", "..Q."}, {"..Q.", "Q...", "...Q", ".Q.."}}},
		{input: 1, output: [][]string{{"Q"}}},
	}
	assert := assert.New(t)
	for _, test := range tests {
		assert.Equal(test.output, solveNQueens(test.input))
	}
}
