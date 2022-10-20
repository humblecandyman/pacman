package main

import (
	"github.com/humblecandyman/pacman/pacman"
	"github.com/humblecandyman/pacman/utils"
)

type PacmanFactory struct {
}

func (factory PacmanFactory) Create() *Pacman {
	return &Pacman{
		entity: pacman.PacmanFactory{
			Position: utils.Vector{
				X: 300,
				Y: 300,
			},
			Radius:           25,
			InitialDirection: utils.DirectionRight,
		}.Create(),
	}
}
