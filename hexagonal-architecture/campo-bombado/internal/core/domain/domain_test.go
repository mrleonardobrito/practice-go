package domain_test

import (
	"campo-bombado/internal/core/domain"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBoard(t *testing.T) {
	size := uint(4)
	quant_bombs := uint(5)
	board := domain.NewBoard(size, quant_bombs)

	count_bombs := uint(0)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == domain.CELL_BOMB {
				count_bombs++
			}
		}
	}

	assert.Equal(t, uint(len(board)), size)
	assert.Equal(t, uint(len(board[0])), size)
	assert.Equal(t, quant_bombs, count_bombs)

	fmt.Println(board)
}
