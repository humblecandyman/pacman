package render

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/humblecandyman/pacman/utils"
)

type Sprite struct {
	Texture *ebiten.Image

	Position utils.Vector
	Size     utils.Vector
	Rotation float64
}

func (sprite Sprite) Draw(target *ebiten.Image) {
	target.DrawImage(
		sprite.Texture,
		sprite.getDrawImageOptions(),
	)

}
func (sprite Sprite) getDrawImageOptions() *ebiten.DrawImageOptions {
	geoM := ebiten.GeoM{}
	geoM.Scale(sprite.getDrawingScale())

	geoM.Translate(sprite.getOrigin())
	geoM.Rotate(sprite.Rotation)

	geoM.Translate(sprite.getDrawingPosition())

	return &ebiten.DrawImageOptions{
		GeoM: geoM,
	}
}

func (sprite Sprite) getDrawingScale() (float64, float64) {
	textureWidth, textureHeight := sprite.Texture.Size()

	return sprite.Size.X / float64(textureWidth),
		sprite.Size.Y / float64(textureHeight)
}

func (sprite Sprite) getDrawingPosition() (float64, float64) {
	return sprite.Position.X,
		sprite.Position.Y
}

func (sprite Sprite) getOrigin() (float64, float64) {
	return -sprite.Size.X * 0.5,
		-sprite.Size.Y * 0.5
}
