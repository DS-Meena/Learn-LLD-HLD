package main

import (
	"fmt"
	"math"
)

type Pawn struct {
	X     int
	Y     int
	White bool
}

func (p *Pawn) move(board [9][9]Piece, a, b int) bool {

	// moving straight but block
	if p.Y == b && board[a][b] != nil {
		return false
	}

	// moving backwards
	if p.White && (p.X-a) > 0 {
		return false
	}
	if !p.White && (p.X-a) < 0 {
		return false
	}

	// moving diagonally
	if math.Abs(float64(p.X-a)) == 1 && math.Abs(float64(p.Y-b)) == 1 {

		// empty
		if board[a][b] == nil {
			return false
		}
	}

	board[a][b] = p
	board[p.X][p.Y] = nil
	p.X = a
	p.Y = b

	// add promotion logic
	return true
}

func (p *Pawn) display() {
	fmt.Print(" P ")
}

func (p *Pawn) isWhite() bool {
	return p.White
}
