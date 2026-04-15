package iterx_test

import (
	"fmt"
	"slices"

	"github.com/min0625/go-exp/iterx"
)

func ExampleMap() {
	seq := slices.Values([]int{1, 2, 3})
	mapped := iterx.Map(seq, func(x int) string {
		return fmt.Sprintf("%d", x*2)
	})

	fmt.Println(slices.Collect(mapped))

	// Output:
	// [2 4 6]
}

func ExampleFilter() {
	seq := slices.Values([]int{1, 2, 3, 4, 5})
	filtered := iterx.Filter(seq, func(x int) bool {
		return x%2 == 0
	})

	fmt.Println(slices.Collect(filtered))

	// Output:
	// [2 4]
}
