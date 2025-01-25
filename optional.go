package opt

import "cmp"

type T[U any] struct {
	value U
	set   bool
}

func (x T[U]) Value() U {
	return x.value
}

func (x T[U]) Set() bool {
	return x.set
}

func Wrap[U any](value U) T[U] {
	return T[U]{
		value: value,
		set:   true,
	}
}

func Unwrap[U any](x T[U]) (U, bool) {
	return x.value, x.set
}
}
