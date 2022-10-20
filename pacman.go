package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/humblecandyman/pacman/scenes"
)

type Pacman struct {
	entity scenes.Entity
}

func (pacman *Pacman) Update() error {
	pacman.entity.Update()
	return nil
}

func (pacman *Pacman) Layout(_, _ int) (int, int) {
	return 600, 600
}

func (game *Pacman) Draw(screen *ebiten.Image) {
	game.entity.Draw(screen)
}
