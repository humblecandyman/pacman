package physics

type Collision struct {
	Another BoundingBox
	Mask    string
	Data    interface{}
}
