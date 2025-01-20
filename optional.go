package opt

type V[T any] struct {
	value T
	set   bool
}

func (x V[T]) Value() T {
	return x.value
}

func (x V[T]) Set() bool {
	return x.set
}

func Wrap[T any](value T) V[T] {
	return V[T]{
		value: value,
		set:   true,
	}
}

func Unwrap[T any](x V[T]) (T, bool) {
	return x.value, x.set
}
}
