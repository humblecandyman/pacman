package scenes

import "github.com/hajimehoshi/ebiten/v2"

type Entity interface {
	Update()
	Draw(target *ebiten.Image)
}
