package main

import "fmt"

func main() {
	var x, y, a, b int
	game := initializeGame()

	for {

		// Display the game board
		game.displayBoard()

		// Get the player's move
		fmt.Print("Enter you move (x, y, a, b): ")
		_, err := fmt.Scan(&x, &y, &a, &b)
		if err != nil {
			fmt.Println("Invalid input. Please try again.")
			continue
		}

		// Check if the move is valid
		if !game.move(x, y, a, b) {
			fmt.Println("Invalid move. Please try again.")
		} else {
			// Check if the game is over
			if game.isGameOver() {
				game.displayBoard()
				fmt.Println("Game over!")
				break
			}
		}

		game.White = !game.White
	}
}
