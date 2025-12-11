package errorx

import "errors"

func As[T any](err error) (T, bool) {
	var t T

	return t, errors.As(err, &t)
}
