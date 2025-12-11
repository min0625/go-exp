package slicex

import (
	"slices"

	"github.com/min0625/go-exp/iterx"
)

func Map[T, V any](input []T, mapper func(T) V) []V {
	in := slices.Values(input)

	out := iterx.Map(in, func(t T) V {
		return mapper(t)
	})

	return slices.AppendSeq(make([]V, 0, cap(input)), out)
}

func Filter[T any](input []T, filter func(T) bool) []T {
	in := slices.Values(input)

	out := iterx.Filter(in, func(t T) bool {
		return filter(t)
	})

	return slices.Collect(out)
}
