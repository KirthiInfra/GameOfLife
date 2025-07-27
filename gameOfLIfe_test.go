package game

import "testing"

func TestNewBoard(t *testing.T) {
    board := [][]uint8{
		{0, 0},
		{1, 1},
	}

	_, err := NewBoard(board)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestMustNotCreateBoardWithNonBinaryValue(t *testing.T) {
	boardWithNonBinaryValue := [][]uint8{
		{1, 0, 1},
		{0, 2, 0},
		{1, 0, 1},
	}

	_, err := NewBoard(boardWithNonBinaryValue)
	if err == nil {
		t.Errorf("Board must not be created with negative values")
	}
}

func TestGameOfLife(t *testing.T) {
    board := [][]uint8{
        {0, 1, 0},
        {0, 0, 1},
        {1, 1, 1},
		{0, 0, 0},
    }
    expected := [][]uint8{
        {0, 0, 0},
        {1, 0, 1},
        {0, 1, 1},
		{0, 1, 0},
    }
    result := GameOfLife(board)
    for i := range board {
        for j := range board[i] {
            if result.board[i][j] != expected[i][j] {
                t.Errorf("Mismatch at %d,%d: got %d, want %d", i, j, result.board[i][j], expected[i][j])
            }
        }
    }
}
