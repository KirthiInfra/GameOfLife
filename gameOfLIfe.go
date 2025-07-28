package game

import "fmt"

type CellState uint8

const (
    Dead         CellState = 0
    Alive        CellState = 1
    AliveToDead  CellState = 2
    DeadToAlive  CellState = 3
)

type Board struct {
    board [][]CellState
}

func (b *Board) Rows() int {
    return len(b.board)
}

func (b *Board) Cols() int {
    if b.Rows() == 0 {
        return 0
    }
    return len(b.board[0])
}

func (b *Board) At(row, col int) CellState {
    return b.board[row][col]
}

func (b *Board) Set(row, col int, state CellState) {
    b.board[row][col] = state
}

func (b *Board) Clone() *Board {
    rows, cols := b.Rows(), b.Cols()
    clone := make([][]CellState, rows)
    for i := 0; i < rows; i++ {
        clone[i] = make([]CellState, cols)
        copy(clone[i], b.board[i])
    }
    return &Board{board: clone}
}

// Neighbors returns all valid neighbor states of a cell (Tell, don't ask)
func (b *Board) Neighbors(row, col int) []CellState {
    directions := [8][2]int{
        {-1, -1}, {-1, 0}, {-1, 1},
        {0, -1},           {0, 1},
        {1, -1},  {1, 0},  {1, 1},
    }
    neighbors := []CellState{}
    for _, d := range directions {
        nr, nc := row+d[0], col+d[1]
        if nr >= 0 && nr < b.Rows() && nc >= 0 && nc < b.Cols() {
            neighbors = append(neighbors, b.At(nr, nc))
        }
    }
    return neighbors
}

func NewBoard(board [][]CellState) (*Board, error) {

    for i := range board {
		for j := range board[i] {
			if board[i][j] != 0 && board[i][j] != 1 {
				return nil, fmt.Errorf("invalid value at position [%d,%d]; must be 0 or 1", i, j)
			}
		}
	}

    return &Board{board: board}, nil
}

func GameOfLife(board [][]CellState) *Board {
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
