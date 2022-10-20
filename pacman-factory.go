package main

import (
	"github.com/humblecandyman/pacman/scenes"
	"github.com/humblecandyman/pacman/utils"
)

type PacmanFactory struct {
	ScreenSize utils.Vector
}

func (factory PacmanFactory) Create() *Pacman {
	return &Pacman{
		screenSize: factory.ScreenSize,
		entity: scenes.GameSceneFactory{
			ScreenSize: factory.ScreenSize,
		}.Create(),
	}
}
