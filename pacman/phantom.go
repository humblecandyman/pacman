package pacman

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/humblecandyman/pacman/physics"
	"github.com/humblecandyman/pacman/render"
)

type Phantom struct {
	Character

	radius float64

	vulnerableTime float64
}

func (phantom *Phantom) Update() {
	phantom.updateVulnerableTime()
}

func (phantom *Phantom) updateVulnerableTime() {
	if !phantom.CanBeEaten() {
		return
	}

	frameTime := 1.0 / 60.0
	phantom.vulnerableTime -= frameTime
}

func (phantom Phantom) Draw(target *ebiten.Image) {
	if phantom.isDead {
		return
	}

	render.Circle{
		Position: phantom.position,
		Radius:   phantom.radius,
		Color:    phantom.getFillColor(),
	}.Draw(target)
}

func (phantom Phantom) getFillColor() color.Color {
	if phantom.CanBeEaten() {
		return color.RGBA{
			B: 0xff,
			A: 0xff,
		}
	}

	return color.RGBA{
		R: 0xff,
		A: 0xff,
	}
}

func (phantom Phantom) GetBoundingBox() physics.BoundingBox {
	return physics.BoundingBox{
		Position: phantom.position,
		Size:     phantom.size.MultiplyFactor(2),
	}
}

func (phantom Phantom) GetCollisionMask() string {
	return PhantomCollisionMask
}

func (phantom Phantom) CanCollideWith(another string) bool {
	return another == PacmanCollisionMask ||
		another == WallCollisionMask
}

func (phantom *Phantom) OnCollision(collision physics.Collision) {
	switch collision.Mask {
	case WallCollisionMask:
		phantom.handleCollisionAgainstWall(collision.Another)
	case PacmanCollisionMask:
		phantom.handleCollisionAgainstPacman()
	}
}

func (phantom Phantom) IsDead() bool {
	return phantom.isDead
}

func (phantom Phantom) CanBeEaten() bool {
	return phantom.vulnerableTime > 0
}

func (phantom *Phantom) MakeItVulnerable() {
	phantom.vulnerableTime = 2
}

func (phantom *Phantom) handleCollisionAgainstPacman() {
	if phantom.CanBeEaten() {
		phantom.isDead = true
	}
}
