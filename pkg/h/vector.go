package h

type Vector struct {
	X, Y int
}

func (v1 Vector) Equals(v2 Vector) bool {
	return v1.X == v2.X && v1.Y == v2.Y
}

func (v1 Vector) Substract(v2 Vector) Vector {
	return Vector{X: v1.X - v2.X, Y: v1.Y - v2.Y}
}

func (v1 Vector) Add(v2 Vector) Vector {
	return Vector{X: v1.X + v2.X, Y: v1.Y + v2.Y}
}
