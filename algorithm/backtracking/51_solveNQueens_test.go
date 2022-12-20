package backtracking

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func solveNQueens(n int) [][]string {
	result := make([][]string, 0)
	// 保存棋盘
	chessboard := make([][]string, n)
	for i := range chessboard {
		chessboard[i] = make([]string, n)
		for j := range chessboard[i] {
			chessboard[i][j] = "."
		}
	}
	var backtracking func(row int)
	backtracking = func(row int) {
		if row == n {
			tmp := make([]string, 0, n)
			for _, path := range chessboard {
				tmp = append(tmp, strings.Join(path, ""))
			}
			result = append(result, tmp)
			return
		}
		for col := 0; col < n; col++ {
			if !isValidFor51(chessboard, row, col) {
				continue
			}
			chessboard[row][col] = "Q"
			backtracking(row + 1)
			chessboard[row][col] = "."
		}
	}
	backtracking(0)
	return result
}

func isValidFor51(chessboard [][]string, row, col int) bool {
	length := len(chessboard)
	// 检查同一列
	for i := 0; i < row; i++ {
		if chessboard[i][col] == "Q" {
			return false
		}
	}
	// 检查左上角
	for i, j := row, col; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if chessboard[i][j] == "Q" {
			return false
		}
	}
	// 检查右上角
	for i, j := row, col; i >= 0 && j < length; i, j = i-1, j+1 {
		if chessboard[i][j] == "Q" {
			return false
		}
	}
	return true
}

func TestSolveNQueue(t *testing.T) {
	tests := []struct {
		name   string
		input  int
		output [][]string
	}{
		{name: "1", input: 4, output: [][]string{{".Q..", "...Q", "Q...", "..Q."}, {"..Q.", "Q...", "...Q", ".Q.."}}},
		{name: "2", input: 1, output: [][]string{{"Q"}}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.output, solveNQueens(test.input))
		})
	}
}
