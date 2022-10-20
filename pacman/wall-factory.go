package pacman

import "github.com/humblecandyman/pacman/utils"

type WallFactory struct {
	Position utils.Vector
	Size     utils.Vector
}

func (factory WallFactory) Create() *Wall {
	return &Wall{
		position: factory.Position,
		size:     factory.Size,
	}
}
