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

func GameOfLife(board [][]int) *Board {
    m, n := len(board), len(board[0])
    directions := [8][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            liveNeighbors := 0
            for _, d := range directions {
                ni, nj := i+d[0], j+d[1]
                if ni >= 0 && ni < m && nj >= 0 && nj < n {
                    if board[ni][nj] == 1 || board[ni][nj] == 2 {
                        liveNeighbors++
                    }
                }
            }
            // Apply the four rules using state-encoding
            if board[i][j] == 1 {
                if liveNeighbors < 2 || liveNeighbors > 3 {
                    board[i][j] = 2        // convert state live to dead
                }
            } else {
                if liveNeighbors == 3 {
                    board[i][j] = 3       // convert state dead to live
                }
            }
        }
    }
    // Finalize the state transition
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if board[i][j] == 2 {
                board[i][j] = 0
            }
            if board[i][j] == 3 {
                board[i][j] = 1
            }
        }
    }
    return &Board{board: board}
}
