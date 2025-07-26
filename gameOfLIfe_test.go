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

func TestMustNotCreateBoardWithNonBinaryValue(t *testing.T) {
	boardWithNonBinaryValue := [][]int{
		{1, 0, 1},
		{0, 2, 0},
		{1, 0, 1},
	}
	
	_, err := NewBoard(boardWithNonBinaryValue)
	if err == nil {
		t.Errorf("Board must not be created with negative values")
	}
}