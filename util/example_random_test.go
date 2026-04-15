package util_test

import (
	"fmt"

	"github.com/min0625/go-exp/util"
)

func ExamplePickOne() {
	items := []string{"only"}
	fmt.Println(util.PickOne(items))

	// Output:
	// only
}
