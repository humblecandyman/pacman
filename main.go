package main

import "github.com/hajimehoshi/ebiten/v2"

func main() {
	game := Pacman{}
	ebiten.RunGame(&game)
}
