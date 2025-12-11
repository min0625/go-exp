package util

import (
	"math/rand/v2"
)

// PickOne returns a random item from the given slice.
func PickOne[T any](items []T) T {
	return items[rand.IntN(len(items))]
}
