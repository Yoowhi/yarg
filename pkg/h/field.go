package h

type Field[T any] struct {
	list []T
	Size Vector
}

func (f *Field[T]) Get(coord Vector) T {
	return f.list[coord.X+coord.Y*f.Size.Y]
}

func (f *Field[T]) Set(coord Vector, value T) bool {
	if coord.X >= f.Size.X || coord.Y >= f.Size.Y {
		return false
	}
	f.list[coord.X+coord.Y*f.Size.Y] = value
	return true
}

func (f *Field[T]) ForEach(fun func(coord Vector, field *Field[T])) Field[T] {
	for x := 0; x < f.Size.X; x++ {
		for y := 0; y < f.Size.Y; y++ {
			coord := Vector{
				X: x,
				Y: y,
			}
			fun(coord, f)
		}
	}
	return *f
}

func (f *Field[T]) Duplicate() Field[T] {
	newField := Field[T]{
		list: make([]T, f.Size.X*f.Size.Y),
		Size: f.Size,
	}
	for i, v := range f.list {
		newField.list[i] = v
	}
	return newField
}

func NewField[T any](size Vector, initVal T) Field[T] {
	field := EmptyField[T](size)
	for i := range field.list {
		field.list[i] = initVal
	}
	return field
}

func EmptyField[T any](size Vector) Field[T] {
	return Field[T]{
		list: make([]T, size.X*size.Y),
		Size: size,
	}
}
