package gameoflife

import "testing"

func TestNewBoard(t *testing.T) {
    board := NewBoard(3, 3)
    
    if len(board.board) != 3 || len(board.board[0]) != 3 {
        t.Fatalf("expected 3x3 board, got %dx%d", len(board.board), len(board.board[0]))
    }
}