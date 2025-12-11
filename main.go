package main

import (
	"fmt"

	"github.com/min0625/go-exp/slicex"
)

func main() {
	v := []int{1, 2, 3}

	// var v []int

	v2 := slicex.Map(v, func(i int) int {
		return i * 2
	})

	fmt.Println(v == nil)
	fmt.Println(v2 == nil)

	fmt.Println(len(v), cap(v))

	fmt.Println(len(v2), cap(v2))

	for _, item := range v2 {
		fmt.Println(item)
	}
}
