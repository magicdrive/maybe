# maybe

# 📦 Maybe & Result for Go (Generics Monad Style)

Type-safe optional values and result types for Go — inspired by Rust, implemented with Go generics.

---

## ✨ Features

- ✅ `Some`, `None`, `Unwrap`, `UnwrapOr`, `IsSome`, `IsNone`
- 🔁 Functional helpers: `Map`, `AndThen`, `OrElse`
- 🧩 Pattern matching with `Match()`
- 🔄 `ToResult()` and `ToResultPrimitive()` for conversions
- 🧪 Fully tested with `go test`
- ⚙️ Supports primitive and pointer-safe usage with `MaybePrimitive`
- 🧱 Built for Go 1.18+ (Generics)

---

## 🚀 Example Usage

### 🧪 Overview: Maybe, MaybePrimitive, and Result

See Also [example/main.go](https://github.com/magicdrive/maybe/blob/main/example/main.go)

```go
package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/magicdrive/maybe"
	"github.com/magicdrive/maybe/result"
)

func main() {
	fmt.Println("=== Maybe ===")
	// Some → Map → Match
	m := maybe.Some(10)
	maybe.Map(m, func(x int) string {
		return fmt.Sprintf("Value is %d", x)
	}).Match(
		func(v string) { fmt.Println("Mapped:", v) },
		func() { fmt.Println("No value") },
	)

	// FromValue → Filter → Fold
	mv := maybe.FromValue(99, true)
	filtered := maybe.Filter(mv, func(x int) bool { return x > 50 })
	folded := maybe.Fold(filtered,
		func(x int) string { return fmt.Sprintf("Kept: %d", x) },
		"Filtered out",
	)
	fmt.Println("Folded result:", folded)

	// Try + Flatten
	nested := maybe.Some(maybe.Some(123))
	flat := maybe.Flatten(nested)
	fmt.Println("Flattened:", flat.UnwrapOr(-1))

	// Tap
	maybe.Tap(flat, func(x int) {
		fmt.Println("Tapped value:", x)
	})

	fmt.Println("\n=== MaybePrimitive ===")
	// SomePrimitive → Filter → Map → Fold
	mp := maybe.SomePrimitive(42)
	filteredPrim := maybe.FilterPrimitive(mp, func(x int) bool { return x%2 == 0 })
	mappedPrim := maybe.MapPrimitive(filteredPrim, func(x int) string {
		return fmt.Sprintf("Even: %d", x)
	})
	foldedPrim := maybe.FoldPrimitive(mappedPrim,
		func(s string) string { return "✅ " + s },
		"none",
	)
	fmt.Println("MaybePrimitive result:", foldedPrim)

	// TryPrimitive
	tried := maybe.TryPrimitive(func() (int, error) {
		return strconv.Atoi("456")
	})
	fmt.Println("TryPrimitive parsed:", tried.UnwrapOr(-1))

	fmt.Println("\n=== Result ===")
	// Try + Map + Fold
	res := result.Try(
		func() (int, error) { return divide(10, 2) },
		func(e error) error { return fmt.Errorf("wrapped: %w", e) },
	)
	rmsg := result.Fold(res,
		func(v int) string { return fmt.Sprintf("Divided: %d", v) },
		func(e error) string { return "Error: " + e.Error() },
	)
	fmt.Println("Result fold:", rmsg)
}

func divide(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("division by zero")
	}
	return x / y, nil
}
```

### 🧠 Conditional Match (MatchIf)

```go
package main

import (
	"errors"
	"fmt"

	"github.com/magicdrive/maybe"
	"github.com/magicdrive/maybe/result"
)

func main() {
	// Maybe[T] - MatchIf
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

	// MaybePrimitive[T] - MatchIfPrimitive
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

	// Result[T, E] - MatchOkIf
	r := result.Ok
	result.MatchOkIf(r, []result.MatchOkCase[int, error]{
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
```

---

## 🧪 Run Tests

```bash
make test
```

---

## 🧰 Run Demo

```bash
go run ./example/main.go
```

---

## Author

Copyright (c) 2025 Hiroshi IKEGAMI

---

## License

This project is licensed under the [MIT License](https://github.com/magicdrive/maybe/blob/main/LICENSE)
