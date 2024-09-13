package main

import "fmt"

type Game struct {
	Board [9][9]Piece
	White bool
}

func initializeGame() *Game {

	board := [9][9]Piece{}

	for i := 1; i <= 8; i++ {
		board[2][i] = &Pawn{X: 2, Y: i, White: true}
		board[7][i] = &Pawn{X: 7, Y: i, White: false}
	}

	// Place other chees pieces on board for both white and black player
	for i := 1; i <= 8; i++ {
		if i == 1 || i == 8 {
			board[1][i] = &Rook{X: 1, Y: i, White: true}
			board[8][i] = &Rook{X: 8, Y: i, White: false}
			continue
		}
		if i == 2 || i == 7 {
			board[1][i] = &Horse{X: 1, Y: i, White: true}
			board[8][i] = &Horse{X: 8, Y: i, White: false}
			continue
		}
		if i == 3 || i == 6 {
			board[1][i] = &Bishop{X: 1, Y: i, White: true}
			board[8][i] = &Bishop{X: 8, Y: i, White: false}
			continue
		}
		if i == 4 {
			board[1][i] = &Queen{X: 1, Y: i, White: true}
			board[8][i] = &Queen{X: 8, Y: i, White: false}
			continue
		}
		board[1][i] = &King{X: 1, Y: i, White: true}
		board[8][i] = &King{X: 8, Y: i, White: false}
	}

	return &Game{Board: board, White: true}
}

// this method moves a piece from x, y to a, b position
// return true if success, otherwise false (will be considered as invalid)
func (g *Game) move(x, y, a, b int) bool {

	// at x, y current player's piece should be present
	if g.Board[x][y] == nil || g.Board[x][y].isWhite() != g.White {
		return false
	}
	// at a, b current player's piece can't be present
	if g.Board[a][b] != nil && g.Board[a][b].isWhite() == g.White {
		return false
	}

	oldAB := g.Board[a][b]
	oldXY := g.Board[x][y]

	// check if the move is valid
	isValid := g.Board[x][y].move(g.Board, a, b)
	if !isValid {
		return false
	}

	// check if the king is in check after moving the piece
	if g.isKingInCheck() {
		// revert the move
		g.Board[x][y] = oldXY
		g.Board[a][b] = oldAB

		fmt.Println("King is in check. Invalid move.")
		return false
	}

	return true
}

func (g *Game) isKingInCheck() bool {

	// return true or false based on whether the king is in check or not
	return false
}

func (g *Game) isGameOver() bool {

	// return true or false based on whether the game is over or not
	return false
}

func (g *Game) displayBoard() {

	for i := 1; i <= 8; i++ {
		for j := 1; j <= 8; j++ {
			if g.Board[i][j] != nil {
				g.Board[i][j].display()
			} else {
				print(" _ ")
			}
		}
		println()
	}
}
