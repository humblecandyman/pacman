package utils

type Vector struct {
	X float64
	Y float64
}

func (vector Vector) Add(another Vector) Vector {
	return Vector{
		X: vector.X + another.X,
		Y: vector.Y + another.Y,
	}
}

func (vector Vector) Subtract(another Vector) Vector {
	return Vector{
		X: vector.X - another.X,
		Y: vector.Y - another.Y,
	}
}

func (vector Vector) Multiply(another Vector) Vector {
	return Vector{
		X: vector.X * another.X,
		Y: vector.Y * another.Y,
	}
}

func (vector Vector) MultiplyFactor(factor float64) Vector {
	return Vector{
		X: vector.X * factor,
		Y: vector.Y * factor,
	}
}
