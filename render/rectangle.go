package render

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/humblecandyman/pacman/utils"
)

type Rectangle struct {
	Size     utils.Vector
	Position utils.Vector
	Color    color.Color
	Empty    bool
}

func (rect Rectangle) Draw(target *ebiten.Image) {
	halfSize := rect.Size.MultiplyFactor(0.5)
	topLeft := rect.Position.Subtract(halfSize)

	ebitenutil.DrawRect(
		target,
		topLeft.X,
		topLeft.Y,
		rect.Size.X,
		rect.Size.Y,
		rect.Color,
	)
}
