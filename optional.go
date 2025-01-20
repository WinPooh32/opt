package opt

type V[T any] struct {
	value T
	set   bool
}

func (o V[T]) Value() T {
	return o.value
}

func (o V[T]) Set() bool {
	return o.set
}

func Wrap[T any](value T) V[T] {
	return V[T]{
		value: value,
		set:   true,
	}
}

func Unwrap[T any](o V[T]) (T, bool) {
	return o.value, o.set
}
