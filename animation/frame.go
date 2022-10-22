package animation

import (
	"image"

	"github.com/humblecandyman/pacman/utils"
)

type Frame struct {
	position utils.Vector
	size     utils.Vector
}

func (frame Frame) ToRect() image.Rectangle {
	return image.Rect(
		int(frame.position.X),
		int(frame.position.Y),
		int(frame.position.X+frame.size.X),
		int(frame.position.Y+frame.size.Y),
	)
}
