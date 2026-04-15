package slogx_test

import (
	"fmt"
	"log/slog"

	"github.com/min0625/go-exp/slogx"
)

func ExampleFuncValuer() {
	count := 0
	valuer := slogx.FuncValuer(func() slog.Value {
		count++
		return slog.IntValue(count)
	})

	fmt.Println(valuer.LogValue())
	fmt.Println(valuer.LogValue())

	// Output:
	// 1
	// 2
}
