package gameoflife

type Board struct {
    board [][]int
}

func NewBoard(m, n int) *Board {
    board := make([][]int, m)
    for i := range board {
        board[i] = make([]int, n)
        for j := range board[i] {
            board[i][j] = 0
        }
    }
    return &Board{board: board}
}
