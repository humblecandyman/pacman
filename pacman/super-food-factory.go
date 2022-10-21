package pacman

import (
	"github.com/humblecandyman/pacman/utils"
)

type SuperFoodFactory struct {
	Position utils.Vector
	OnEaten  utils.Callback
}

func (factory SuperFoodFactory) Create() *SuperFood {
	return &SuperFood{
		Food: Food{
			position: factory.Position,
			size: utils.Vector{
				X: 20,
				Y: 20,
			},
		},
		onEaten: factory.OnEaten,
	}
}
