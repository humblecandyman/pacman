package physics

import "github.com/humblecandyman/pacman/utils"

type BoundingBox struct {
	Position utils.Vector
	Size     utils.Vector
}

func (box BoundingBox) getMinMaxCorners() (utils.Vector, utils.Vector) {
	halfSize := box.Size.MultiplyFactor(0.5)

	return box.Position.Subtract(halfSize),
		box.Position.Add(halfSize)
}

func (box BoundingBox) IsCollidingWith(another BoundingBox) bool {
	minA, maxA := box.getMinMaxCorners()
	minB, maxB := another.getMinMaxCorners()

	return minB.X < maxA.X &&
		maxB.X > minA.X &&
		minB.Y < maxA.Y &&
		maxB.Y > minA.Y
}

func (box BoundingBox) Top() float64 {
	return box.Position.Y - box.Size.Y*0.5
}

func (box BoundingBox) Right() float64 {
	return box.Position.X + box.Size.X*0.5
}

func (box BoundingBox) Bottom() float64 {
	return box.Position.Y + box.Size.Y*0.5
}

func (box BoundingBox) Left() float64 {
	return box.Position.X - box.Size.X*0.5
}
