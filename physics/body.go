package physics

type Body interface {
	GetBoundingBox() BoundingBox

	GetCollisionMask() string
	CanCollideWith(string) bool
	OnCollision(Collision)

	IsDead() bool
}
