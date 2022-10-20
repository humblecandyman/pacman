package pacman

import "github.com/humblecandyman/pacman/utils"

type FoodFactory struct {
	Position utils.Vector
}

func (factory FoodFactory) Create() *Food {
	return &Food{
		position: factory.Position,
		size: utils.Vector{
			X: 10,
			Y: 10,
		},
	}
}
