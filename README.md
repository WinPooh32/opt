# opt

![test](https://github.com/WinPooh32/opt/actions/workflows/test.yml/badge.svg)
[![Go Reference](https://pkg.go.dev/badge/github.com/WinPooh32/opt.svg)](https://pkg.go.dev/github.com/WinPooh32/opt)

The `opt` package provides a generic type for optional values.

## Features

- Generic: `T[U]` struct holds any type of value.
- Sortable: implements "cmp"-like helper functions `Less` and `Compare` (only for `U` satisfied [cmp.Ordered](https://pkg.go.dev/cmp#Ordered) constraint).
- Serializable: TODO

## Usage

Here's a simple example demonstrating how to use the `opt` package:

### Example

```go
package main

import (
	"fmt"

	"github.com/WinPooh32/opt"
)

func main() {
	var v opt.T[int]

	v = opt.Wrap(1)

	if v.Set() {
		fmt.Println(v.Value())
	}
}
