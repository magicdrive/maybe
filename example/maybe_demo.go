package main

import (
	"fmt"
	"strconv"

	"github.com/magicdrive/maybe"
)

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
