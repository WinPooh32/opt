package opt_test

import (
	"fmt"
	"math"
	"slices"
	"testing"

	"github.com/WinPooh32/opt"
)

func Example() {
	o1 := opt.Wrap(1)

	if v, ok := opt.Unwrap(o1); ok {
		fmt.Println(v)
	}

	o2 := opt.Wrap(2)

	if o2.Set() {
		fmt.Println(o2.Value())
	}

	fmt.Println(opt.Less(o1, o2))

	fmt.Println(opt.Compare(o1, o2))

	// Output: 1
	// 2
	// true
	// -1
}

func TestLess(t *testing.T) {
	type args struct {
		x opt.T[float64]
		y opt.T[float64]
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"less", args{opt.Wrap(4.0), opt.Wrap(9.0)}, true},
		{"equal", args{opt.Wrap(4.0), opt.Wrap(4.0)}, false},
		{"greater", args{opt.Wrap(9.0), opt.Wrap(4.0)}, false},
		{"x is NaN", args{opt.Wrap(math.NaN()), opt.Wrap(4.0)}, true},
		{"y is NaN", args{opt.Wrap(4.0), opt.Wrap(math.NaN())}, false},
		{"x, y both are NaN", args{opt.Wrap(math.NaN()), opt.Wrap(math.NaN())}, false},
		{"x is not set", args{opt.Empty[float64](), opt.Wrap(1.0)}, true},
		{"y is not set", args{opt.Wrap(1.0), opt.Empty[float64]()}, false},
		{"x,y both are not set", args{opt.Empty[float64](), opt.Empty[float64]()}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := opt.Less(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompare(t *testing.T) {
	type args struct {
		x opt.T[float64]
		y opt.T[float64]
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"less", args{opt.Wrap(4.0), opt.Wrap(9.0)}, -1},
		{"equal", args{opt.Wrap(4.0), opt.Wrap(4.0)}, 0},
		{"greater", args{opt.Wrap(9.0), opt.Wrap(4.0)}, 1},
		{"x is NaN", args{opt.Wrap(math.NaN()), opt.Wrap(4.0)}, -1},
		{"y is NaN", args{opt.Wrap(4.0), opt.Wrap(math.NaN())}, 1},
		{"x, y both are NaN", args{opt.Wrap(math.NaN()), opt.Wrap(math.NaN())}, 0},
		{"x is not set", args{opt.Empty[float64](), opt.Wrap(1.0)}, -1},
		{"y is not set", args{opt.Wrap(1.0), opt.Empty[float64]()}, 1},
		{"x,y both are not set", args{opt.Empty[float64](), opt.Empty[float64]()}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := opt.Compare(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompare_SliceSorting(t *testing.T) {
	type args struct {
		s []opt.T[float64]
	}
	tests := []struct {
		name string
		args args
		want []opt.T[float64]
	}{
		{
			"unordered",
			args{[]opt.T[float64]{
				opt.Wrap(3.0),
				opt.Empty[float64](),
				opt.Wrap(1.0),
				opt.Empty[float64](),
				opt.Wrap(8.0),
			}},
			[]opt.T[float64]{
				opt.Empty[float64](),
				opt.Empty[float64](),
				opt.Wrap(1.0),
				opt.Wrap(3.0),
				opt.Wrap(8.0),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := append([]opt.T[float64](nil), tt.args.s...)

			slices.SortFunc(got, opt.Compare)

			if !equal(got, tt.want) {
				t.Errorf("Sort using Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func equal(s1, s2 []opt.T[float64]) bool {
	a1 := fmt.Sprintf("%v", s1)
	a2 := fmt.Sprintf("%v", s2)
	return a1 == a2
}
