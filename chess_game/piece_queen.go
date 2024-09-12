package main

import "fmt"

type Queen struct {
	X     int
	Y     int
	White bool
}

func (q *Queen) move(board [9][9]Piece, a, b int) bool {

	// other piece should not be present in path
	// is the way free from (X, Y) to (a, b)

	// on same row
	if q.X == a {
		if isRowClear(board, min(q.Y, b), max(q.Y, b), q.X) {
			board[q.X][b] = board[q.X][q.Y]
			board[q.X][q.Y] = nil
			q.Y = b
			return true
		} else {
			return false
		}
	}

	// on same column
	if q.Y == b {
		if isColumnClear(board, min(q.X, a), max(q.X, a), q.Y) {
			board[a][q.Y] = board[q.X][q.Y]
			board[q.X][q.Y] = nil
			q.X = a
			return true
		} else {
			return false
		}
	}

	// on same diagonal
	if abs(q.X-a) == abs(q.Y-b) {
		if isDiagonalClear(board, q.X, q.Y, a, b) {
			board[a][b] = board[q.X][q.Y]
			board[q.X][q.Y] = nil
			q.X = a
			q.Y = b
			return true
		} else {
			return false
		}
	}

	// on opposite diagonal
	if abs(q.X-a) == abs(q.Y-b) {
		if isDiagonalClear(board, q.X, q.Y, a, b) {
			board[a][b] = board[q.X][q.Y]
			board[q.X][q.Y] = nil
			q.X = a
			q.Y = b
			return true
		} else {
			return false
		}
	}

	// check if can move

	return true
}

func (q *Queen) display() {
	fmt.Print(" Q ")
}

func (q *Queen) isWhite() bool {
	return q.White
}
