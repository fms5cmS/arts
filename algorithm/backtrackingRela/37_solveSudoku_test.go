package backtrackingRela

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// board.length == 9
// board[i].length == 9
// board[i][j] 是一位数字或者 '.'
// 题目数据 保证 输入数独仅有一个解
func solveSudoku(board [][]byte) {
	var backtracking func() bool
	backtracking = func() bool {
		for row := 0; row < len(board); row++ {
			for col := 0; col < len(board); col++ {
				if board[row][col] != '.' {
					continue
				}
				for val := '1'; val <= '9'; val++ {
					// 这里先判断 val 能否放到 board[row][col] 的位置上，可以的话再实际放上去！
					if isValidFor37(board, row, col, byte(val)) {
						board[row][col] = byte(val)
						if backtracking() {
							return true
						}
						board[row][col] = '.'
					}
				}
				return false
			}
		}
		return true
	}
	backtracking()
}

func isValidFor37(board [][]byte, row, col int, val byte) bool {
	// 检查同行
	for j := 0; j < len(board); j++ {
		if board[row][j] == val {
			return false
		}
	}
	// 检查同列
	for i := 0; i < len(board); i++ {
		if board[i][col] == val {
			return false
		}
	}
	// 检查 3x3
	startRow, startCol := (row/3)*3, (col/3)*3 // 注意这里起始位置的计算！
	for i := startRow; i < startRow+3; i++ {
		for j := startCol; j < startCol+3; j++ {
			if board[i][j] == val {
				return false
			}
		}
	}
	return true
}

func TestSolveSudoku(t *testing.T) {
	tests := []struct {
		name   string
		input  [][]byte
		output [][]byte
	}{
		{
			name: "1",
			input: [][]byte{
				{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			},
			output: [][]byte{
				{'5', '3', '4', '6', '7', '8', '9', '1', '2'},
				{'6', '7', '2', '1', '9', '5', '3', '4', '8'},
				{'1', '9', '8', '3', '4', '2', '5', '6', '7'},
				{'8', '5', '9', '7', '6', '1', '4', '2', '3'},
				{'4', '2', '6', '8', '5', '3', '7', '9', '1'},
				{'7', '1', '3', '9', '2', '4', '8', '5', '6'},
				{'9', '6', '1', '5', '3', '7', '2', '8', '4'},
				{'2', '8', '7', '4', '1', '9', '6', '3', '5'},
				{'3', '4', '5', '2', '8', '6', '1', '7', '9'},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			solveSudoku(test.input)
			assert.Equal(t, test.output, test.input)
		})
	}
}
