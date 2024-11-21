package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/nickwoott/boggle/board"
	"github.com/nickwoott/boggle/game"
	"github.com/nickwoott/boggle/player"
)

func main() {
	player := player.NewPlayer()
	board := board.NewBoard()
	game := game.NewGame(board, player)

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Game Timer")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
