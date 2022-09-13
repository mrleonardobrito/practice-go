package domain

import (
	"math/rand"
	"time"
)

const (
	CELL_BOMB       = "X"
	CELL_REVEALED   = "O"
	CELL_EMPTY      = "-"
	CELL_BOMB_HIDED = "-"
)

type Board [][]string

func NewBoard(size uint, quant_bombs uint) Board {
	board := NewEmptyBoard(size)
	board.FillWithBombs(quant_bombs)

	return board
}

func (board Board) FillWithBombs(quant_bombs uint) {
	rows := len(board)
	cols := len(board[0])
	positions := getRandomPositions(uint(rows*cols), quant_bombs)

	var row, col int

	for _, pos := range positions {
		row = pos / cols
		col = pos - row*rows

		board[row][col] = CELL_BOMB
	}
}

func (board Board) HideBombs() Board {
	size := uint(len(board))
	newBoard := NewEmptyBoard(size)

	for i := range newBoard {
		for j, space := range newBoard[0] {
			if space == CELL_BOMB {
				space = CELL_BOMB_HIDED
			} else {
				space = board[i][j]
			}
		}
	}

	return newBoard
}

func (board Board) IsValidPosition(row uint, col uint) bool {
	return row < uint(len(board)) && col < uint(len(board[0]))
}

func (board Board) Contains(row uint, col uint, element string) bool {
	return board[row][col] == element
}

func (board Board) Set(row uint, col uint, element string) {
	board[row][col] = element
}

func (board Board) HasEmptyCells() bool {
	for row := range board {
		for col := range board[0] {
			if board[row][col] == CELL_EMPTY {
				return true
			}
		}
	}

	return false
}

func NewEmptyBoard(size uint) Board {
	board := make([][]string, size)

	for row := range board {
		board[row] = make([]string, size)
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			board[i][j] = CELL_EMPTY
		}
	}

	return board
}

func getRandomPositions(board_size uint, quant_bombs uint) []int {
	rand.Seed(time.Now().UnixNano())
	p := rand.Perm(int(board_size))

	positions := []int{}

	for _, r := range p[:quant_bombs] {
		positions = append(positions, r)
	}

	return positions
}
