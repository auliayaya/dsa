package leetcode

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func isValidSudoku(board [][]byte) bool {
	rows := make([]map[byte]bool, 9)
	cols := make([]map[byte]bool, 9)
	boxes := make([]map[byte]bool, 9)

	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
		boxes[i] = make(map[byte]bool)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			dot := board[i][j]
			if dot == '.' {
				continue
			}
			boxIndex := (i/3)*3 + j/3 // Calculate which 3x3 box we're in
			fmt.Println("Dot ", boxIndex)
			fmt.Println("Res ", rows[i][dot], cols[j][dot], boxes[boxIndex][dot])

			// Check if number already exists in row, column, or box
			if rows[i][dot] || cols[j][dot] || boxes[boxIndex][dot] {
				return false
			}

			// Mark dotber as seen
			rows[i][dot] = true
			cols[j][dot] = true
			boxes[boxIndex][dot] = true
		}
	}
	return true
}

func TestIsValidSudoku(t *testing.T) {
	tc := []struct {
		input  [][]byte
		output bool
	}{
		{
			input: [][]byte{
				{'5', '3', '.', '.', ',', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '.', '5', '.', '.', '.'},
				{'.', '9', '.', '.', '.', '.', '.', '.', '.'},
				{'8', '.', '.', '.', '.', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			},
			output: true,
		},
	}
	for _, tt := range tc {
		result := isValidSudoku(tt.input)
		assert.Equal(t, tt.output, result)
	}
}
