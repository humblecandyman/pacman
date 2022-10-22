package render

import (
	"image"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

func LoadTexture(textureFileName string) *ebiten.Image {
	textureFile, err := os.Open(textureFileName)

	if err != nil {
		panic(err)
	}

	textureImage, _, _ := image.Decode(textureFile)

	return ebiten.NewImageFromImage(textureImage)
}
