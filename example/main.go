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

func main() {
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

	fmt.Println("\n=== MaybePrimitive ===")

	// SomePrimitive / UnwrapOr
	mp := maybe.SomePrimitive(42)
	fmt.Println("Primitive value:", mp.UnwrapOr(0)) // → 42

	// FilterPrimitive
	filteredPrim := maybe.FilterPrimitive(mp, func(x int) bool { return x > 40 })
	fmt.Println("Filtered primitive:", filteredPrim.UnwrapOr(-1)) // → 42

	// TapPrimitive
	maybe.TapPrimitive(filteredPrim, func(x int) {
		fmt.Println("Tapped primitive:", x)
	})

	// MapPrimitive
	mappedPrim := maybe.MapPrimitive(mp, func(x int) string {
		return fmt.Sprintf("val=%d", x)
	})
	fmt.Println("Mapped primitive:", mappedPrim.UnwrapOr("none"))

	// FoldPrimitive
	foldedPrim := maybe.FoldPrimitive(mp,
		func(x int) string { return fmt.Sprintf("prim:%d", x) },
		"none",
	)
	fmt.Println("Fold primitive:", foldedPrim)

	// FromValuePrimitive
	fromVal := maybe.FromValuePrimitive(999, true)
	fmt.Println("FromValuePrimitive:", fromVal.UnwrapOr(-1)) // → 999

	// TryPrimitive
	tryPrim := maybe.TryPrimitive(func() (int, error) {
		return strconv.Atoi("123")
	})
	fmt.Println("TryPrimitive:", tryPrim.UnwrapOr(-1)) // → 123

}
