package ptr_test

import (
	"fmt"

	"github.com/min0625/go-exp/ptr"
)

func ExampleNew() {
	p := ptr.New(42)
	fmt.Println(*p)

	// Output:
	// 42
}

func ExampleValue() {
	n := 99
	fmt.Println(ptr.Value(&n))
	fmt.Println(ptr.Value[int](nil))

	// Output:
	// 99
	// 0
}

func ExampleValueOr() {
	n := 99
	fmt.Println(ptr.ValueOr(&n, -1))
	fmt.Println(ptr.ValueOr[int](nil, -1))

	// Output:
	// 99
	// -1
}

func ExampleClone() {
	n := 42
	p := &n
	c := ptr.Clone(p)

	*p = 100

	fmt.Println(*c)
	fmt.Println(ptr.Clone[int](nil) == nil)

	// Output:
	// 42
	// true
}
