package h

type Field[T any] struct {
	list   []T
	Width  int
	Height int
}

func (f *Field[T]) Get(coord Vector) T {
	return f.list[coord.X+coord.Y*f.Height]
}

func (f *Field[T]) Set(coord Vector, value T) bool {
	if coord.X >= f.Width || coord.Y >= f.Height {
		return false
	}
	f.list[coord.X+coord.Y*f.Height] = value
	return true
}

func (f *Field[T]) ForEach(fun func(coord Vector, field *Field[T])) {
	for x := 0; x < f.Width; x++ {
		for y := 0; y < f.Height; y++ {
			coord := Vector{
				X: x,
				Y: y,
			}
			fun(coord, f)
		}
	}
}

func NewField[T any](width, height int) Field[T] {
	return Field[T]{
		list:   make([]T, width*height),
		Width:  width,
		Height: height,
	}
}
