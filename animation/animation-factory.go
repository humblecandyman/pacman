package animation

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/humblecandyman/pacman/render"
	"github.com/humblecandyman/pacman/utils"
)

type AnimationFactory struct {
	SpriteSheetFileName string
	FrameSize           utils.Vector
	FrameTime           float64
}

func (factory AnimationFactory) Create() *Animation {
	spriteSheet := render.LoadTexture(factory.SpriteSheetFileName)

	return &Animation{
		spriteSheet: spriteSheet,
		frames:      factory.createFrames(spriteSheet),
		frameTime:   factory.FrameTime,
	}
}

func (factory AnimationFactory) createFrames(spriteSheet *ebiten.Image) []Frame {
	frames := []Frame{}

	spriteSheetBounds := spriteSheet.Bounds()
	framePosition := utils.Vector{}

	for framePosition.X < float64(spriteSheetBounds.Max.X) {
		frames = append(frames, Frame{
			position: framePosition,
			size:     factory.FrameSize,
		})

		framePosition.X += factory.FrameSize.X
	}

	return frames
}
