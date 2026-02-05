// Package util provides various utility functions.
package util //nolint:revive // var-naming

import (
	"math/rand/v2"
)

// PickOne returns a random element from the given slice.
func PickOne[T any](items []T) T {
	//nolint:gosec // G404 For non-cryptographic purposes, math/rand is sufficient
	return items[rand.IntN(len(items))]
}
