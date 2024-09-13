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
    3. Check if can move using valid moves. Unique for each piece
    4. Promotion logic for pawn
    5. King should not be in check, after you move the piece.
3. Game finishes.

Few more methods can be implemented, but I think for 1 hour interview this much code should be enough.

[You tube video](https://www.youtube.com/watch?v=kk6QkWxz5jM&list=PL6W8uoQQ2c61X_9e6Net0WdYZidm7zooW&index=22)