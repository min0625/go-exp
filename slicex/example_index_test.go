package slicex_test

import (
	"fmt"

	"github.com/min0625/go-exp/slicex"
)

func ExampleFirstOr() {
	input := []int{10, 20, 30}
	fmt.Println(slicex.FirstOr(input, -1))

	empty := []int{}
	fmt.Println(slicex.FirstOr(empty, -1))

	// Output:
	// 10
	// -1
}

func ExampleLast() {
	input := []int{10, 20, 30}
	fmt.Println(slicex.Last(input))

	// Output:
	// 30
}

func ExampleLastOr() {
	input := []int{10, 20, 30}
	fmt.Println(slicex.LastOr(input, -1))

	empty := []int{}
	fmt.Println(slicex.LastOr(empty, -1))

	// Output:
	// 30
	// -1
}
