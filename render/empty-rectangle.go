package render

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/humblecandyman/pacman/utils"
)

type EmptyRectangle struct {
	Position utils.Vector
	Size     utils.Vector
	Color    color.Color
}

func (rect EmptyRectangle) Draw(target *ebiten.Image) {
	corners := rect.getCorners()

	rect.drawLine(target, corners[0], corners[1])
	rect.drawLine(target, corners[1], corners[2])
	rect.drawLine(target, corners[2], corners[3])
	rect.drawLine(target, corners[3], corners[0])
}

func (rect EmptyRectangle) getCorners() []utils.Vector {
	halfSize := rect.Size.MultiplyFactor(0.5)
	corners := []utils.Vector{
		rect.Position.Subtract(halfSize),
		rect.Position.Add(halfSize.Multiply(utils.Vector{
			X: 1,
			Y: -1,
		})),
		rect.Position.Add(halfSize),
		rect.Position.Add(halfSize.Multiply(utils.Vector{
			X: -1,
			Y: 1,
		})),
	}

	return corners
}

func (rect EmptyRectangle) drawLine(target *ebiten.Image, from utils.Vector, to utils.Vector) {
	ebitenutil.DrawLine(
		target,
		from.X,
		from.Y,
		to.X,
		to.Y,
		rect.Color,
	)
}
