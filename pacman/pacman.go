package pacman

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/humblecandyman/pacman/controllers"
	"github.com/humblecandyman/pacman/render"
	"github.com/humblecandyman/pacman/utils"
)

type Pacman struct {
	position  utils.Vector
	radius    float64
	direction utils.Direction
	speed     float64

	eyePosition    utils.Vector
	movementVector utils.Vector

	controller controllers.IController
}

func (pacman *Pacman) Update() {
	pacman.controller.Update()
	pacman.updatePosition()
}

func (pacman *Pacman) updatePosition() {
	pacman.position = pacman.position.Add(pacman.movementVector)
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

func (pacman *Pacman) Do(command string) {
	switch command {
	case "go-up":
		pacman.setDirection(utils.DirectionUp)
	case "go-right":
		pacman.setDirection(utils.DirectionRight)
	case "go-down":
		pacman.setDirection(utils.DirectionDown)
	case "go-left":
		pacman.setDirection(utils.DirectionLeft)
	}
}

func (pacman *Pacman) setDirection(newDirection utils.Direction) {
	if pacman.direction == newDirection {
		return
	}

	pacman.direction = newDirection
	pacman.handleDirectionChange()
}

func (pacman *Pacman) handleDirectionChange() {
	pacman.updateMovementVector()
	pacman.updateEyePosition()
}

func (pacman *Pacman) updateMovementVector() {
	movementVector := utils.Vector{}

	switch pacman.direction {
	case utils.DirectionUp:
		movementVector.Y -= pacman.speed
	case utils.DirectionRight:
		movementVector.X += pacman.speed
	case utils.DirectionDown:
		movementVector.Y += pacman.speed
	case utils.DirectionLeft:
		movementVector.X -= pacman.speed
	}

	pacman.movementVector = movementVector
}

func (pacman *Pacman) updateEyePosition() {
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
