package ptr

// New returns a pointer to the given value.
func New[T any](v T) *T {
	return &v
}

// Value returns the value pointed to by p.
func Value[T any](p *T) T {
	if p == nil {
		return zero[T]()
	}

	return *p
}

// ValueOr returns the value pointed to by p or defaultValue if p is nil.
func ValueOr[T any](p *T, defaultValue T) T {
	if p == nil {
		return defaultValue
	}

	return *p
}

// Clone creates a shallow copy of the value pointed to by p.
func Clone[T any](p *T) *T {
	if p == nil {
		return nil
	}

	v := *p

	return &v
}

func zero[T any]() T {
	var zero T

	return zero
}
