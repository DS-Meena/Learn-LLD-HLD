package main

import "fmt"

type King struct {
	X     int
	Y     int
	White bool
}

func (q *King) move(board [9][9]Piece, a, b int) bool {

	// if (a, b) is not adjacent to (q.X, q.Y)
	if abs(a-q.X) > 1 || abs(b-q.Y) > 1 {
		return false
	}

	board[a][b] = board[q.X][q.Y]
	board[q.X][q.Y] = nil
	q.Y = b
	q.X = a
	return true
}

func (q *King) display() {
	fmt.Print(" K ")
}

func (q *King) isWhite() bool {
	return q.White
}
