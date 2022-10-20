package pacman

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/humblecandyman/pacman/physics"
	"github.com/humblecandyman/pacman/render"
	"github.com/humblecandyman/pacman/utils"
)

type Teleporter struct {
	position utils.Vector
	size     utils.Vector

	terminal *Teleporter
}

func (teleporter Teleporter) Update() {
}

func (teleporter Teleporter) Draw(target *ebiten.Image) {
	render.EmptyRectangle{
		Position: teleporter.position,
		Size:     teleporter.size,
		Color:    color.White,
	}.Draw(target)
}

func (teleporter Teleporter) GetBoundingBox() physics.BoundingBox {
	return physics.BoundingBox{
		Position: teleporter.position,
		Size:     teleporter.size,
	}
}

func (teleporter Teleporter) GetCollisionMask() string {
	return TeleporterCollisionMask
}

func (teleporter Teleporter) CanCollideWith(another string) bool {
	return another == PacmanCollisionMask
}

func (teleporter Teleporter) OnCollision(physics.Collision) {
}

func (teleporter Teleporter) IsDead() bool {
	return false
}

func (teleporter Teleporter) GetTerminalPosition(direction utils.Direction) utils.Vector {
	switch direction {
	case utils.DirectionUp:
		return utils.Vector{
			X: teleporter.terminal.position.X,
			Y: teleporter.terminal.GetBoundingBox().Top(),
		}
	case utils.DirectionRight:
		return utils.Vector{
			X: teleporter.terminal.GetBoundingBox().Right(),
			Y: teleporter.terminal.position.Y,
		}
	case utils.DirectionDown:
		return utils.Vector{
			X: teleporter.terminal.position.X,
			Y: teleporter.terminal.GetBoundingBox().Bottom(),
		}
	case utils.DirectionLeft:
		return utils.Vector{
			X: teleporter.terminal.GetBoundingBox().Left(),
			Y: teleporter.terminal.position.Y,
		}
	}

	return utils.Vector{}
}
