package errorx_test

import (
	"fmt"

	"github.com/min0625/go-exp/errorx"
)

type myError struct{ code int }

func (e *myError) Error() string { return fmt.Sprintf("code=%d", e.code) }

func ExampleAsType() {
	err := fmt.Errorf("wrapped: %w", &myError{code: 42})

	e, ok := errorx.AsType[*myError](err)
	fmt.Println(ok)
	fmt.Println(e.code)

	// Output:
	// true
	// 42
}

func ExampleAsType_notMatched() {
	err := fmt.Errorf("plain error")

	_, ok := errorx.AsType[*myError](err)
	fmt.Println(ok)

	// Output:
	// false
}
