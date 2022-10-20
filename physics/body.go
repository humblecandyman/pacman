package physics

type Body interface {
	GetBoundingBox() BoundingBox

	GetCollisionMask() string
	CanCollisionWith(string) bool
	OnCollisionWith(string)

	IsDead() bool
}
