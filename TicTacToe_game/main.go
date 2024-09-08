package main

import "fmt"

func main() {
	var boardSize, numPlayers int
	fmt.Print("Enter board size: ")
	fmt.Scan(&boardSize)
	fmt.Print("Enter number of players: ")
	fmt.Scan(&numPlayers)

	players := make([]*Player, numPlayers)
	for i := 0; i < numPlayers; i++ {
		var name string
		var symbol rune

		fmt.Printf("Enter name for Player %d: ", i+1)
		fmt.Scan(&name)
		fmt.Printf("Enter symbol for Player %d: ", i+1)
		fmt.Scanf("%c\n", &symbol)

		players[i] = &Player{
			Name:   name,
			Symbol: symbol,
		}
	}

	game := NewGame(boardSize, players)
	game.Play()
}
