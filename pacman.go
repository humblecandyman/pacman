package main

import "github.com/hajimehoshi/ebiten/v2"

type Pacman struct {
}

func (pacman *Pacman) Update() error {
	return nil
}

func (pacman *Pacman) Layout(_, _ int) (int, int) {
	return 600, 600
}

func (pacman *Pacman) Draw(screen *ebiten.Image) {
}
