package main

type Piece interface {
	move([9][9]Piece, int, int) bool
	display()
	isWhite() bool
}
