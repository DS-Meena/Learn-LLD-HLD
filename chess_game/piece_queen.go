package main

import "fmt"

type Queen struct {
	X     int
	Y     int
	White bool
}

func (q *Queen) move(board [9][9]Piece, a, b int) bool {

	// other piece should not be present in path
	oldValue := board[a][b]
	board[a][b] = nil

	if q.X == a && !isRowClear(board, min(q.Y, b), max(q.Y, b), q.X) {
		board[a][b] = oldValue
		return false
	}
	if q.Y == b && !isColumnClear(board, min(q.X, a), max(q.X, a), q.Y) {
		board[a][b] = oldValue
		return false
	}
	if q.X-a == q.Y-b && !isDiagonalClear(board, q.X, q.Y, a, b) {
		board[a][b] = oldValue
		return false
	}
	if q.X+a == q.Y+b && !isDiagonalClear2(board, q.X, q.Y, a, b) {
		board[a][b] = oldValue
		return false
	}

	// if (a, b) is outside of attacking range of queen
	if q.X != a && q.Y != b && q.X-a != q.Y-b && q.X+a != q.Y+b {
		board[a][b] = oldValue
		return false
	}

	board[a][b] = board[q.X][q.Y]
	board[q.X][q.Y] = nil
	q.Y = b
	q.X = a
	return true
}

func (q *Queen) display() {
	fmt.Print(" Q ")
}

func (q *Queen) isWhite() bool {
	return q.White
}
