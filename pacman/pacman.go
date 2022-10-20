package pacman

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/humblecandyman/pacman/physics"
	"github.com/humblecandyman/pacman/render"
	"github.com/humblecandyman/pacman/utils"
)

type Pacman struct {
	Character

	radius      float64
	eyePosition utils.Vector
}

func (pacman *Pacman) Update() {
	pacman.controller.Update()
	pacman.updatePosition()
	pacman.updateEyePosition()
}

func (pacman Pacman) Draw(target *ebiten.Image) {
	pacman.drawBody(target)
	pacman.drawEye(target)
}

func (pacman Pacman) drawBody(target *ebiten.Image) {
	render.Circle{
		Position: pacman.position,
		Radius:   pacman.radius,
		Color: color.RGBA{
			R: 0xff,
			G: 0xff,
			A: 0xff,
		},
	}.Draw(target)
}

func (pacman Pacman) drawEye(target *ebiten.Image) {
	render.Circle{
		Position: pacman.position.Add(pacman.eyePosition),
		Radius:   pacman.radius * 0.25,
		Color:    color.Black,
	}.Draw(target)
}

func (pacman *Pacman) updateEyePosition() {
	if pacman.direction == utils.NoDirection {
		return
	}

	eyePositionRelativeToPosition := utils.Vector{
		X: pacman.radius * 0.25,
		Y: pacman.radius * 0.3,
	}

	eyePositions := map[string]utils.Vector{
		utils.DirectionUp: eyePositionRelativeToPosition.Multiply(utils.Vector{
			X: -1,
			Y: -1,
		}),
		utils.DirectionRight: eyePositionRelativeToPosition.Multiply(utils.Vector{
			X: 1,
			Y: -1,
		}),
		utils.DirectionDown: eyePositionRelativeToPosition,
		utils.DirectionLeft: eyePositionRelativeToPosition.Multiply(utils.Vector{
			X: -1,
			Y: -1,
		}),
	}

	pacman.eyePosition = eyePositions[pacman.direction]
}

func (pacman Pacman) CanCollideWith(another string) bool {
	return another != PacmanCollisionMask
}

func (pacman Pacman) GetCollisionMask() string {
	return PacmanCollisionMask
}

func (pacman *Pacman) OnCollision(collision physics.Collision) {
	switch collision.Mask {
	case FoodCollisionMask:
		pacman.increaseScore(100)
	case WallCollisionMask:
		pacman.handleCollisionAgainstWall(collision.Another)
	}
}

func (pacman Pacman) IsDead() bool {
	return false
}
