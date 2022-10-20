package scenes

import (
	"github.com/humblecandyman/pacman/pacman"
	"github.com/humblecandyman/pacman/physics"
	"github.com/humblecandyman/pacman/utils"
)

type GameSceneFactory struct {
	ScreenSize utils.Vector
}

func (factory GameSceneFactory) Create() *Scene {
	world := physics.World{}

	pacmanPlayer := pacman.PacmanFactory{
		Position: factory.ScreenSize.Multiply(utils.Vector{
			X: 0.5,
			Y: 0.5,
		}),
		Radius:           15,
		InitialDirection: utils.DirectionRight,
	}.Create()
	world.PushBody(pacmanPlayer)

	wall := pacman.WallFactory{
		Position: utils.Vector{
			X: factory.ScreenSize.X * 0.5,
			Y: factory.ScreenSize.Y * 0.3,
		},
		Size: utils.Vector{
			X: factory.ScreenSize.X,
			Y: 5,
		},
	}.Create()
	world.PushBody(wall)

	food := pacman.FoodFactory{
		Position: factory.ScreenSize.Multiply(utils.Vector{
			X: 0.8,
			Y: 0.5,
		}),
	}.Create()
	world.PushBody(food)

	return &Scene{
		entities: []Entity{
			wall,
			pacmanPlayer,
			food,
			world,
		},
	}
}
