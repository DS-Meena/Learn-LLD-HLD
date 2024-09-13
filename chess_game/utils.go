package main

// check if row ith row is clear from a to b column
func isRowClear(board [9][9]Piece, a, b, ith int) bool {
	for i := a; i <= b; i++ {
		if board[ith][i] != nil {
			return false
		}
	}
	return true
}

// check if column ith column is clear from a to b row
func isColumnClear(board [9][9]Piece, a, b, ith int) bool {
	for i := a; i <= b; i++ {
		if board[i][ith] != nil {
			return false
		}
	}
	return true
}

// check if diagonal from x, y to a, b is clear
func isDiagonalClear(board [9][9]Piece, x, y, a, b int) bool {
	if x < a {
		for i, j := x, y; i <= a && j >= b; i, j = i+1, j-1 {
			if board[i][j] != nil {
				return false
			}
		}
	} else {
		for i, j := x, y; i >= a && j <= b; i, j = i-1, j+1 {
			if board[i][j] != nil {
				return false
			}
		}
	}
	return true
}

// check if diagonal from x, y to a, b is clear
func isDiagonalClear2(board [9][9]Piece, x, y, a, b int) bool {
	if x < a {
		for i, j := x, y; i <= a && j <= b; i, j = i+1, j+1 {
			if board[i][j] != nil {
				return false
			}
		}
	} else {
		for i, j := x, y; i >= a && j >= b; i, j = i-1, j-1 {
			if board[i][j] != nil {
				return false
			}
		}
	}
	return true
}

// check if diagonal from x, y to a, b is clear

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
