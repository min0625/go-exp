// Package slicex provides extensions for slice operations.
package slicex

// FirstOr returns the first element of the slice or the default value if the slice is empty.
func FirstOr[S ~[]T, T any](slice S, defaultValue T) T {
	if len(slice) == 0 {
		return defaultValue
	}

	return slice[0]
}

// Last returns the last element of the slice.
func Last[S ~[]T, T any](slice S) T {
	return slice[len(slice)-1]
}

// LastOr returns the last element of the slice or the default value if the slice is empty.
func LastOr[S ~[]T, T any](slice S, defaultValue T) T {
	if len(slice) == 0 {
		return defaultValue
	}

	return slice[len(slice)-1]
}
