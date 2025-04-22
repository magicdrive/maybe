package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/magicdrive/maybe/result"
)

func divide(x, y int) result.Result[int, error] {
	if y == 0 {
		return result.Err[int](errors.New("division by zero"))
	}
	return result.Ok[int, error](x / y)
}

func DemoResult() {
	fmt.Println("\n=== Result ===")

	// Tap → Map → UnwrapOr
	r := divide(10, 2)
	r = result.Tap(r, func(x int) { fmt.Println("Divided:", x) })
	rMapped := result.Map(r, func(x int) string {
		return fmt.Sprintf("Result: %d", x)
	})
	fmt.Println("Final result:", rMapped.UnwrapOr("Error"))

	// From
	_, err := strconv.Atoi("xyz")
	from := result.From(0, err)

	// Try
	try := result.Try(func() (int, error) {
		return strconv.Atoi("789")
	}, func(e error) error {
		return fmt.Errorf("wrapped: %w", e)
	})

	// Fold
	foldedFrom := result.Fold(from,
		func(v int) string { return fmt.Sprintf("ok: %d", v) },
		func(e error) string { return "err: " + e.Error() },
	)
	fmt.Println("From folded:", foldedFrom)

	foldedTry := result.Fold(try,
		func(v int) string { return fmt.Sprintf("ok: %d", v) },
		func(e error) string { return "err: " + e.Error() },
	)
	fmt.Println("Try folded:", foldedTry)
}
