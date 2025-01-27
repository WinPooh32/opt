package opt

import "cmp"

// T contains optional value of any type.
// Empty value can be used.
type T[U any] struct {
	value U
	set   bool
}

// Value returns contained value.
func (x T[U]) Value() U {
	return x.value
}

// Set returns true when the value is set.
func (x T[U]) Set() bool {
	return x.set
}

// Wrap wraps the value as optional type.
func Wrap[U any](value U) T[U] {
	return T[U]{
		value: value,
		set:   true,
	}
}

// Empty returns empty container of any type U.
func Empty[U any]() T[U] {
	return T[U]{}
}

// Unwrap returns x content as tuple as the value and set indicator.
func Unwrap[U any](x T[U]) (U, bool) {
	return x.value, x.set
}

// Less returns true when x is less than y.
// Unset value is always considered less.
func Less[U cmp.Ordered](x, y T[U]) bool {
	return (!x.set && y.set) || cmp.Less(x.value, y.value)
}

// Compare returns
//
//	-1 if x is less than y,
//	0 if x equals y,
//	+1 if x is greater than y
//
// Unset value is always considered less.
func Compare[U cmp.Ordered](x, y T[U]) int {
	if !x.set && !y.set {
		return 0
	}
	if x.set && !y.set {
		return 1
	}
	if !x.set && y.set {
		return -1
	}
	return cmp.Compare(x.value, y.value)
}
