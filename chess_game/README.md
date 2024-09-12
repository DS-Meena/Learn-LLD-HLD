# Chess Game

## Components

1. Board
2. Pieces black and white
3. 2 Players

## Piece

Every piece will have it's own rules of moving.

## Game 

Game is our main object.
Board will be a matrix of 8x8 size. That will contain 32 pieces initially.

Board

top to down: 0 to 7  (X: represent row) (white is at top and black is at bottom)
left to right: 0 to 7 (Y: represent column)

Game {
    board
}

How do we play:
1. set the board with the pieces.
2. Until check mate or stalemate: 
    1. Each player moves one piece, to any valid position.
    2. If opposite player's piece is at that position, he can capture the piece.