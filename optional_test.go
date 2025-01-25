package opt_test

import (
	"fmt"

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
