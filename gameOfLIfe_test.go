package gameoflife

import "testing"

func TestNewBoard(t *testing.T) {
    board := [][]int{
		{0, 0},
		{1, 1},
	}

	_, err := NewBoard(board)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}