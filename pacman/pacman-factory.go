package pacman

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/humblecandyman/pacman/controllers"
	"github.com/humblecandyman/pacman/utils"
)

type PacmanFactory struct {
	Position         utils.Vector
	Radius           float64
	InitialDirection utils.Direction
}

func (factory PacmanFactory) Create() *Pacman {
	pacman := &Pacman{
		position:  factory.Position,
		radius:    factory.Radius,
		direction: factory.InitialDirection,
		speed:     2,
	}

	controller := controllers.PlayerFactory{
		GoUpKey:      ebiten.KeyW,
		GoRightKey:   ebiten.KeyD,
		GoDownKey:    ebiten.KeyS,
		GoLeftKey:    ebiten.KeyA,
		Controllable: pacman,
	}.Create()

	pacman.controller = controller
	pacman.handleDirectionChange()

	return pacman
}
