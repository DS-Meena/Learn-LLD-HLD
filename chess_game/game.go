package main

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

	// check if the move is valid
	return g.Board[x][y].move(g.Board, a, b)
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
