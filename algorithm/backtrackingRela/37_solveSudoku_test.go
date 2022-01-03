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
		for i := 0; i < len(board); i++ { // 遍历行
			for j := 0; j < len(board[i]); j++ { // 遍历列
				// 仅对空位操作
				if board[i][j] != '.' {
					continue
				}
				// 试着在 i,j 坐标放入值 k
				for k := '1'; k <= '9'; k++ {
					if isValidFor37(i, j, byte(k), board) {
						board[i][j] = byte(k) // 放置
						if backtracking() {   // 找到合适的就返回
							return true
						}
						board[i][j] = '.' // 回溯
					}
				}
				return false // 9 个数字都不符合要求
			}
		}
		return true
	}
	backtracking()
}

func isValidFor37(row, col int, val byte, board [][]byte) bool {
	// 检查同行是否有重复元素
	for i := 0; i < 9; i++ {
		if board[row][i] == val {
			return false
		}
	}
	// 检查同列是否有重复元素
	for i := 0; i < 9; i++ {
		if board[i][col] == val {
			return false
		}
	}
	// 判断该位置所在的九宫格中是否已有重复元素
	startRow, startCol := (row/3)*3, (col/3)*3
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
		input  [][]byte
		output [][]byte
	}{
		{
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
	assert := assert.New(t)
	for _, test := range tests {
		solveSudoku(test.input)
		assert.Equal(test.output, test.input)
	}
}
