package board

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type BoggleBoard struct {
	grid [4][4]rune
}

func NewBoard() *BoggleBoard {

	var board BoggleBoard

	letters := []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			board.grid[i][j] = letters[rand.Intn(len(letters))]
		}
	}

	return &board
}

func (b *BoggleBoard) DrawBoard(screen *ebiten.Image) {
	cellWidth := 100
	cellHeight := 100

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {

			x := j * cellWidth
			y := i * cellHeight

			letter := b.grid[i][j]

			text := string(letter)

			ebitenutil.DebugPrintAt(screen, text, x+cellWidth/3, y+cellHeight/3)
		}
	}
}
