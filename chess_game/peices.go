package main

type Piece interface {
	move(int, int) bool
}

type Pawn struct {
	X int
	Y int
}

func (p *Pawn) move(a, b int) bool {

	// check if can move to (a, b)
	// i need the details of board to check
	return true
}

type Queen struct {
	X int
	Y int
}

func (q *Queen) move(a, b int) bool {

	// check if can move
	return true
}
