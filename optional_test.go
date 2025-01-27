package opt_test

import (
	"fmt"
	"math"
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := opt.Less(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Less() = %v, want %v", got, tt.want)
			}
		})
	}
}
