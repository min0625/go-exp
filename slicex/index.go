package slicex

func FirstOr[S ~[]T, T any](slice S, defaultValue T) T {
	if len(slice) == 0 {
		return defaultValue
	}

	return slice[0]
}

func Last[S ~[]T, T any](slice S) T {
	return slice[len(slice)-1]
}

func LastOr[S ~[]T, T any](slice S, defaultValue T) T {
	if len(slice) == 0 {
		return defaultValue
	}

	return slice[len(slice)-1]
}
