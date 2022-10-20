package pacman

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/humblecandyman/pacman/physics"
	"github.com/humblecandyman/pacman/render"
	"github.com/humblecandyman/pacman/utils"
)

type Food struct {
	size     utils.Vector
	position utils.Vector
	isDead   bool
}

func (food Food) Update() {
}

func (food Food) Draw(target *ebiten.Image) {
	if food.isDead {
		return
	}

	render.Rectangle{
		Size:     food.size,
		Position: food.position,
		Color: color.RGBA{
			R: 0xf5,
			G: 0xd9,
			B: 0xbc,
			A: 0xff,
		},
	}.Draw(target)
}

func (food Food) CanCollisionWith(another string) bool {
	return another == PacmanCollisionMask
}

func (food Food) GetCollisionMask() string {
	return FoodCollisionMask
}

func (food Food) IsDead() bool {
	return food.isDead
}

func (food *Food) OnCollisionWith(_ string) {
	food.isDead = true
}

func (food Food) GetBoundingBox() physics.BoundingBox {
	return physics.BoundingBox{
		Position: food.position,
		Size:     food.size,
	}
}
