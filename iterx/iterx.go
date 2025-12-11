package iterx

import (
	"iter"
)

func Map[T, V any](input iter.Seq[T], mapper func(T) V) iter.Seq[V] {
	return func(yield func(V) bool) {
		input(func(item T) bool {
			return yield(mapper(item))
		})
	}
}

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
