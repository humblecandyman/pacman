package pacman

import (
	"github.com/humblecandyman/pacman/controllers"
	"github.com/humblecandyman/pacman/utils"
)

type Character struct {
	position  utils.Vector
	direction utils.Direction
	speed     float64

	movementVector utils.Vector

	controller controllers.IController

	score int
}

func (character *Character) updatePosition() {
	character.position = character.position.Add(character.movementVector)
}

func (character *Character) Do(command string) {
	switch command {
	case "go-up":
		character.setDirection(utils.DirectionUp)
	case "go-right":
		character.setDirection(utils.DirectionRight)
	case "go-down":
		character.setDirection(utils.DirectionDown)
	case "go-left":
		character.setDirection(utils.DirectionLeft)
	}
}

func (character *Character) setDirection(newDirection utils.Direction) {
	if character.direction == newDirection {
		return
	}

	character.direction = newDirection
	character.updateMovementVector()
}

func (character *Character) updateMovementVector() {
	movementVector := utils.Vector{}

	switch character.direction {
	case utils.DirectionUp:
		movementVector.Y -= character.speed
	case utils.DirectionRight:
		movementVector.X += character.speed
	case utils.DirectionDown:
		movementVector.Y += character.speed
	case utils.DirectionLeft:
		movementVector.X -= character.speed
	}

	character.movementVector = movementVector
}

func (character *Character) increaseScore(amout int) {
	character.score += amout
}
