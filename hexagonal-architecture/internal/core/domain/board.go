package domain

const (
	CELL_BOMB       = "X"
	CELL_REVEALED   = "O"
	CELL_EMPTY      = "-"
	CELL_BOMB_HIDED = "-"
)

type BoardSettings struct {
	Size  uint `json:"size"`
	Bombs uint `json:"bombs"`
}

type Board [][]string

func NewBoard(size uint, quant_bombs uint) Board {
	board := NewEmptyBoard(size)
	board.fillWithBombs(quant_bombs)

	return board
}

func (board *Board) fillWithBombs(quant_bombs uint) {
	// positions := getRandomPositions(quant_bombs)

	// fmt.Println(positions)
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
