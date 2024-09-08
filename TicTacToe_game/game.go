package main

import "fmt"

type Game struct {
	Board              *Board
	Players            []*Player
	CurrentPlayerIndex int
}

func NewGame(boardSize int, players []*Player) *Game {
	return &Game{
		Board:   NewBoard(boardSize),
		Players: players,
	}
}

func (g *Game) SwitchPlayer() {
	g.CurrentPlayerIndex = (g.CurrentPlayerIndex + 1) % len(g.Players)
}

func (g *Game) Play() {
	for {
		g.Board.Display()
		currentPlayer := g.Players[g.CurrentPlayerIndex]
		fmt.Printf("%s's turn (%c). Enter position (0-%d): ", currentPlayer.Name, currentPlayer.Symbol, g.Board.Size*g.Board.Size-1)

		var position int
		fmt.Scan(&position)

		if position >= 0 && position < g.Board.Size*g.Board.Size && g.Board.Move(position, currentPlayer.Symbol) {

			// if current player won
			if g.Board.IsWinner(currentPlayer.Symbol) {
				g.Board.Display()
				fmt.Printf("%s wins!\n", currentPlayer.Name)
				break

			} else if g.Board.IsFull() { // check if board is full
				g.Board.Display()
				fmt.Println("It's a tie!")
				break
			}
			g.SwitchPlayer()
		} else {
			fmt.Println("Invalid move. Try again.")
		}
	}
}
