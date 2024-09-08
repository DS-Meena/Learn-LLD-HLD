package main

import "fmt"

type Player struct {
	Name   string
	Symbol rune
}

type Board struct {
	Size  int
	Cells []rune
}

func NewBoard(size int) *Board {
	return &Board{
		Size:  size,
		Cells: make([]rune, size*size), // list of n*n length
	}
}

func (b *Board) Display() {
	for i := 0; i < b.Size; i++ {
		for j := 0; j < b.Size; j++ {
			if b.Cells[i*b.Size+j] == 0 {
				fmt.Printf("_ ")
			} else {
				fmt.Printf("%c ", b.Cells[i*b.Size+j]) // index = i * size + j
			}
		}
		fmt.Println()
	}
}

func (b *Board) IsCellEmpty(position int) bool {
	return b.Cells[position] == 0
}

func (b *Board) Move(position int, symbol rune) bool {
	if b.IsCellEmpty(position) {
		b.Cells[position] = symbol
		return true
	}
	return false
}

func (b *Board) checkLine(start, step int, symbol rune) bool {
	for i := 0; i < b.Size; i++ {
		if b.Cells[start+i*step] != symbol {
			return false
		}
	}

	return true
}

func (b *Board) IsWinner(symbol rune) bool {
	// check horizontal and vertical lines
	for i := 0; i < b.Size; i++ {
		if b.checkLine(i*b.Size, 1, symbol) || b.checkLine(i, b.Size, symbol) {
			return true
		}
	}

	// check diagonal lines
	return b.checkLine(0, b.Size+1, symbol) || b.checkLine(b.Size-1, b.Size-1, symbol)
}

func (b *Board) IsFull() bool {
	for _, cell := range b.Cells {
		if cell == 0 {
			return false
		}
	}

	return true
}
