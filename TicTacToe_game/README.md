# Tic Tac Toe Game 🎮

In this example, i have implemented a extensible Tic-Tac-Toe game in Go. You can modify the number of players and board size, as you wish.

- **Variable board size** 🛹: The board struct is initialized with a size parameter, allowing for n*n boards
- **Multiple players** 🏂: The Game struct accepts a slice of Player pointers, allowing for m number of players.
- **Winning condition** 🏆: You can modify the IsWinner() method to check for winning conditions.

## Key design aspects: 🔑

- **Encapsulation:** Each struct (Player, Board, Game) encapsulates its data and behavior. 🧱
- **Abstraction:** The Game struct abstracts the game logic from the Board and Player structs. 🎭
- **Extensibility:** The design allows for easy extension to different board sizes and number of players. 🔄