package gameoflife

import "fmt"

type Board struct {
    board [][]int
}

func NewBoard(board [][]int) (*Board, error) {

    for i := range board {
		for j := range board[i] {
			if board[i][j] != 0 && board[i][j] != 1 {
				return nil, fmt.Errorf("invalid value at position [%d,%d]; must be 0 or 1", i, j)
			}
		}
	}

    return &Board{board: board}, nil
}
