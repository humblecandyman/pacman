package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/humblecandyman/pacman/scenes"
	"github.com/humblecandyman/pacman/utils"
)

type Pacman struct {
	screenSize utils.Vector
	entity     scenes.Entity
}

func (pacman *Pacman) Update() error {
	pacman.entity.Update()
	return nil
}

func (pacman *Pacman) Layout(_, _ int) (int, int) {
	return int(pacman.screenSize.X), int(pacman.screenSize.Y)
}

func (pacman *Pacman) Draw(screen *ebiten.Image) {
	ebiten.SetWindowSize(pacman.Layout(0, 0))
	pacman.entity.Draw(screen)
}
