package animation

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Animation struct {
	spriteSheet *ebiten.Image

	frames       []Frame
	currentFrame int

	frameTime        float64
	currentFrameTime float64
	isPaused         bool
}

func (animation *Animation) Update() {
	if animation.isPaused {
		return
	}

	frameTime := 1.0 / 60.0
	animation.currentFrameTime += frameTime

	if animation.currentFrameTime > animation.frameTime {
		animation.increaseCurrentFrame()
		animation.currentFrameTime = 0
	}
}

func (animation *Animation) increaseCurrentFrame() {
	animation.currentFrame++

	if animation.currentFrame > animation.frameCount()-1 {
		animation.currentFrame = 0
	}
}

func (animation Animation) frameCount() int {
	return len(animation.frames)
}

func (animation Animation) GetFrame() *ebiten.Image {
	frame := animation.frames[animation.currentFrame]

	return animation.spriteSheet.SubImage(
		frame.ToRect(),
	).(*ebiten.Image)
}

func (animation *Animation) Pause() {
	animation.isPaused = true
}

func (animation *Animation) Resume() {
	animation.isPaused = false
}
