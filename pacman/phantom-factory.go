package pacman

import "github.com/humblecandyman/pacman/utils"

type PhantomFactory struct {
	Position utils.Vector
	Size     float64
}

func (factory PhantomFactory) Create() *Phantom {
	return &Phantom{
		Character: Character{
			position: factory.Position,
			size: utils.Vector{
				X: factory.Size,
				Y: factory.Size,
			},
			direction: utils.NoDirection,
		},

		radius: factory.Size,
	}
}
