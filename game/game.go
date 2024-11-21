package game

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/nickwoott/boggle/board"
	"github.com/nickwoott/boggle/player"
)

type Game struct {
	timerDuration  time.Duration
	startTime      time.Time
	board          *board.BoggleBoard
	player         *player.Player
	input          string
	lastDelete     time.Time
	deleteCooldown time.Duration
}

func NewGame(board *board.BoggleBoard, player *player.Player) *Game {
	return &Game{
		// I think some of these fields might need extracting or be used elsewhere?
		timerDuration:  3 * time.Minute,
		startTime:      time.Now(),
		board:          board,
		player:         player,
		input:          "",
		lastDelete:     time.Now(),
		deleteCooldown: 300 * time.Millisecond,
	}
}
func (g *Game) Update(*ebiten.Image) error {

	//remember, these bits here are actually just used to render the input data,
	for _, char := range ebiten.InputChars() {
		g.input += string(char)
	}

	if (ebiten.IsKeyPressed(ebiten.KeyBackspace) || ebiten.IsKeyPressed(ebiten.KeyDelete)) && time.Since(g.lastDelete) > g.deleteCooldown {
		if len(g.input) > 0 {
			g.input = g.input[:len(g.input)-1]
			g.lastDelete = time.Now()
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		g.player.AddWord(g.input)
		g.input = ""
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	//this is timer logic, I would actually extract this here, I do not like it.
	remaining := g.timerDuration - time.Since(g.startTime)
	if remaining < 0 {
		remaining = 0
	}
	minutes := int(remaining.Minutes())
	seconds := int(remaining.Seconds()) % 60
	text := "Time remaining: " + fmt.Sprintf("%02d:%02d", minutes, seconds)

	ebitenutil.DebugPrint(screen, text)

	g.board.DrawBoard(screen)

	inputText := "Word: " + g.input
	ebitenutil.DebugPrintAt(screen, inputText, 10, 440)

	scoreText := fmt.Sprintf("Score: %d", g.player.Score)
	ebitenutil.DebugPrintAt(screen, scoreText, 10, 460)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}
