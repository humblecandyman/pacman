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
		Character: Character{
			position:  factory.Position,
			direction: factory.InitialDirection,
			speed:     2,
		},
		radius: factory.Radius,
	}

	controller := controllers.PlayerFactory{
		GoUpKey:      ebiten.KeyW,
		GoRightKey:   ebiten.KeyD,
		GoDownKey:    ebiten.KeyS,
		GoLeftKey:    ebiten.KeyA,
		Controllable: pacman,
	}.Create()

	pacman.controller = controller
	pacman.updateMovementVector()

	return pacman
}
