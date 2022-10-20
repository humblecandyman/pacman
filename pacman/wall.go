package pacman

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/humblecandyman/pacman/physics"
	"github.com/humblecandyman/pacman/render"
	"github.com/humblecandyman/pacman/utils"
)

type Wall struct {
	position utils.Vector
	size     utils.Vector
}

func (wall Wall) Update() {
}

func (wall Wall) Draw(target *ebiten.Image) {
	render.Rectangle{
		Position: wall.position,
		Size:     wall.size,
		Color: color.RGBA{
			B: 0xff,
			A: 0xff,
		},
	}.Draw(target)
}

func (wall Wall) CanCollideWith(another string) bool {
	return another == PacmanCollisionMask
}

func (wall Wall) GetCollisionMask() string {
	return WallCollisionMask
}

func (wall *Wall) OnCollision(collision physics.Collision) {
}

func (wall Wall) IsDead() bool {
	return false
}

func (wall Wall) GetBoundingBox() physics.BoundingBox {
	return physics.BoundingBox{
		Position: wall.position,
		Size:     wall.size,
	}
}
