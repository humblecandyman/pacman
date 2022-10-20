package scenes

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Scene struct {
	entities []Entity
}

func (scene Scene) Update() {
	for _, entity := range scene.entities {
		entity.Update()
	}
}

func (scene Scene) Draw(target *ebiten.Image) {
	for _, entity := range scene.entities {
		entity.Draw(target)
	}

	ebitenutil.DrawLine(
		target,
		300,
		300,
		301,
		300,
		color.White,
	)
}
