package physics

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/humblecandyman/pacman/render"
)

type World struct {
	bodies []Body
}

func (world World) Update() {
	for leftIndex, leftBody := range world.bodies {
		if leftBody.IsDead() {
			continue
		}

		for rightIndex := leftIndex + 1; rightIndex < len(world.bodies); rightIndex++ {
			rightBody := world.bodies[rightIndex]
			rightcollisionMask := rightBody.GetCollisionMask()

			if leftBody.CanCollisionWith(rightcollisionMask) {
				leftBoundingBox := leftBody.GetBoundingBox()
				rightBoundingBox := rightBody.GetBoundingBox()

				if leftBoundingBox.IsCollidingWith(rightBoundingBox) {
					leftCollisionMask := leftBody.GetCollisionMask()
					leftBody.OnCollisionWith(rightcollisionMask)
					rightBody.OnCollisionWith(leftCollisionMask)

					if leftBody.IsDead() {
						break
					}
				}
			}
		}
	}
}

func (world World) Draw(target *ebiten.Image) {
	for _, body := range world.bodies {
		if body.IsDead() {
			continue
		}

		boundingBox := body.GetBoundingBox()

		render.Rectangle{
			Position: boundingBox.Position,
			Size:     boundingBox.Size,
			Color: color.RGBA{
				G: 0xff,
				A: 0xff,
			},
		}.Draw(target)
	}
}

func (world *World) PushBody(body Body) {
	world.bodies = append(world.bodies, body)
}
