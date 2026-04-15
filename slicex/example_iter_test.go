package slicex_test

import (
	"fmt"

	"github.com/min0625/go-exp/slicex"
)

func ExampleMap() {
	input := []int{0, 1, 2, 3, 4}
	output := slicex.Map(input, func(x int) string {
		return fmt.Sprintf("%c", 'A'+x)
	})

	fmt.Println(output) // Output: [A B C D E]
}

func ExampleFilter() {
	input := []int{0, 1, 2, 3, 4}
	output := slicex.Filter(input, func(x int) bool {
		return x%2 == 0
	})

	fmt.Println(output) // Output: [0 2 4]
}
