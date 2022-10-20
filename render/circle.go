package render

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/humblecandyman/pacman/utils"
)

type Circle struct {
	Position utils.Vector
	Color    color.Color
	Radius   float64
}

func (circle Circle) Draw(target *ebiten.Image) {
	ebitenutil.DrawCircle(
		target,
		circle.Position.X,
		circle.Position.Y,
		circle.Radius,
		circle.Color,
	)
}
