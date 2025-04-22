package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/magicdrive/maybe"
	"github.com/magicdrive/maybe/result"
)

func divide(x, y int) result.Result[int, error] {
	if y == 0 {
		return result.Err[int](errors.New("division by zero"))
	}
	return result.Ok[int, error](x / y)
}

func DemoMatchIf() {
	fmt.Println("\n=== MatchIf ===")

	m := maybe.Some(30)
	maybe.MatchIf(m, []maybe.MatchCase[int]{
		{Cond: func(x int) bool { return x > 100 }, Then: func(x int) {
			fmt.Println("Too large:", x)
		}},
		{Cond: func(x int) bool { return x > 10 }, Then: func(x int) {
			fmt.Println("Matched OK:", x)
		}},
	}, func() {
		fmt.Println("No match or None")
	})

	mp := maybe.SomePrimitive(5)
	maybe.MatchIfPrimitive(mp, []maybe.MatchPrimitiveCase[int]{
		{Cond: func(x int) bool { return x > 10 }, Then: func(x int) {
			fmt.Println("Primitive large:", x)
		}},
		{Cond: func(x int) bool { return x < 10 }, Then: func(x int) {
			fmt.Println("Primitive small:", x)
		}},
	}, func() {
		fmt.Println("Primitive fallback")
	})

	res := result.Ok[int, error](11)
	result.MatchOkIf(res, []result.MatchOkCase[int, error]{
		{Cond: func(x int) bool { return x > 100 }, Then: func(x int) {
			fmt.Println("Result: huge", x)
		}},
		{Cond: func(x int) bool { return x > 10 }, Then: func(x int) {
			fmt.Println("Result: fine", x)
		}},
	}, func(e error) {
		fmt.Println("Error happened:", e)
	}, func() {
		fmt.Println("No match in Result")
	})
}

func DemoMaybe() {
	fmt.Println("=== Maybe ===")

	// Some / Map / Match
	m := maybe.Some(42)
	str := maybe.Map(m, func(x int) string {
		return fmt.Sprintf("Mapped: %d", x)
	})
	str.Match(
		func(v string) { fmt.Println("Match result:", v) },
		func() { fmt.Println("No value") },
	)

	// FromValue → Filter → Tap → Map → Match
	mv := maybe.FromValue(100, true)
	filtered := maybe.Filter(mv, func(x int) bool { return x > 50 })
	maybe.Tap(filtered, func(x int) { fmt.Println("Tapped:", x) })
	mapped := maybe.Map(filtered, func(x int) string {
		return fmt.Sprintf("String: %d", x)
	})
	mapped.Match(
		func(s string) { fmt.Println("Filtered final:", s) },
		func() { fmt.Println("Filtered out") },
	)

	// Try
	parsed := maybe.Try(func() (int, error) {
		return strconv.Atoi("456")
	})
	fmt.Println("Try parsed:", parsed.UnwrapOr(-1))

	// Flatten
	nested := maybe.Some(maybe.Some(99))
	flattened := maybe.Flatten(nested)
	fmt.Println("Flattened:", flattened.UnwrapOr(-1))

	// Fold
	folded := maybe.Fold(maybe.Some(10),
		func(x int) string { return fmt.Sprintf("folded: %d", x) },
		"none",
	)
	fmt.Println("Fold result:", folded)
}

func DemoMaybePrimitive() {

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
