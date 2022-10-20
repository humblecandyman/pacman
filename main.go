package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/humblecandyman/pacman/utils"
)

func main() {
	game := PacmanFactory{
		ScreenSize: utils.Vector{
			X: 600,
			Y: 600,
		},
	}.Create()
	ebiten.RunGame(game)
}
