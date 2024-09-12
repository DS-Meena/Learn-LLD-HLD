package main

type Game struct {
	Board [9][9]Piece
}

func initializeGame() *Game {

	board := [9][9]Piece{}

	for i := 1; i <= 8; i++ {
		board[2][i] = &Pawn{X: 2, Y: i}
		board[7][i] = &Pawn{X: 7, Y: i}
	}

	return &Game{Board: board}
}
