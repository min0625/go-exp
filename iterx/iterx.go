// Package iterx provides extensions for iterator operations.
package iterx

import (
	"iter"
)

// Map applies the mapper function to each element of the input sequence
// and returns a new sequence containing the results.
func Map[T, V any](input iter.Seq[T], mapper func(T) V) iter.Seq[V] {
	return func(yield func(V) bool) {
		input(func(item T) bool {
			return yield(mapper(item))
		})
	}
}

// Filter returns a new sequence containing only the elements that satisfy the filter function.
func Filter[T any](input iter.Seq[T], filter func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		input(func(item T) bool {
			if filter(item) {
				return yield(item)
			}

			return true
		})
	}
}
