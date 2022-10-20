package main

import "github.com/hajimehoshi/ebiten/v2"

func main() {
	game := PacmanFactory{}.Create()
	ebiten.RunGame(game)
}
