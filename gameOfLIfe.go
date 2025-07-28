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

// GameRules interface - Open for extension with new sets of rules (OCP)
type GameRules interface {
    NextState(current CellState, neighbors []CellState) CellState
}

// ClassicRules implements Conway's Game of Life
type ClassicRules struct{}

func (r ClassicRules) NextState(current CellState, neighbors []CellState) CellState {
    liveCount := 0
    for _, s := range neighbors {
        if s == Alive || s == AliveToDead {
            liveCount++
        }
    }

    switch current {
    case Alive:
        if liveCount < 2 || liveCount > 3 {
            return AliveToDead
        }
        return Alive
    case Dead:
        if liveCount == 3 {
            return DeadToAlive
        }
        return Dead
    default:
        return Dead
    }
}

// Next evolves the board to the next generation using provided GameRules
func (b *Board) Next(rules GameRules) *Board {
    next := b.Clone()
    for i := 0; i < b.Rows(); i++ {
        for j := 0; j < b.Cols(); j++ {
            neighbors := b.Neighbors(i, j)
            newState := rules.NextState(b.At(i, j), neighbors)
            next.Set(i, j, newState)
        }
    }
    // Finalize states by converting transition states to final states
    for i := 0; i < next.Rows(); i++ {
        for j := 0; j < next.Cols(); j++ {
            switch next.At(i, j) {
            case AliveToDead:
                next.Set(i, j, Dead)
            case DeadToAlive:
                next.Set(i, j, Alive)
            }
        }
    }
    return next
}

// GameOfLife function using the above abstractions and default rules
func GameOfLife(board [][]CellState) *Board {
    b, err := NewBoard(board)
    if err != nil {
        panic(err)
    }
    rules := ClassicRules{}
    return b.Next(rules)
}
