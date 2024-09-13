package main

import "fmt"

type Horse struct {
	X     int
	Y     int
	White bool
}

func (q *Horse) move(board [9][9]Piece, a, b int) bool {

	// horse can jump over other pieces
	board[a][b] = board[q.X][q.Y]
	board[q.X][q.Y] = nil
	q.Y = b
	q.X = a
	return true
}

func (q *Horse) display() {
	fmt.Print(" H ")
}

func (q *Horse) isWhite() bool {
	return q.White
}
