# maybe

# 📦 Maybe & Result for Go (Generics Monad Style)

Type-safe optional values and result types for Go — inspired by Rust, implemented with Go generics.

---

## ✨ Features

- ✅ `Some`, `None`, `Unwrap`, `UnwrapOr`, `IsSome`, `IsNone`
- 🔁 Functional helpers: `Map`, `AndThen`, `OrElse`, `Filter`, `Flatten`
- 🧩 Pattern matching with `Match()`
- 🧠 `MatchIf()` enables condition-based matching like a functional switch
- 💥 `MatchTypeDynamic()` to dispatch by runtime type via reflect
- ⚡ `MatchTypeKeyed()` for fast dispatch using user-defined TypeKey()
- 🔄 `ToResult()` and `FromValue()`, `Try()` conversions
- 🔍 `Tap()` for side-effect inspection
- ⚙️ Works with both `Maybe[T]`, `MaybePrimitive[T]`, and `Result[T, E]`
- 🧪 Supports primitive and pointer-safe usage with `MaybePrimitive`
- 🧱 Built for Go 1.18+ (Generics)

---

## 🚀 Example Usage

### 🧪 Overview: Maybe, MaybePrimitive, and Result

See Also [example](https://github.com/magicdrive/maybe/blob/main/example)

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

### 🧠 Conditional Match (MatchIf,MatchOkIf)

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
	r := result.Ok[int, error](11)
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

### 🔍TypeMatch (MatchTypeDynamic,MatchTypeKeyed)

```go
package main

import (
	"fmt"
	"reflect"

	"github.com/magicdrive/maybe"
)

type User struct {
	Name string
}

func (u User) TypeKey() string {
	return "User"
}

type Admin struct {
	Level int
}

func (a Admin) TypeKey() string {
	return "Admin"
}

func MatchTypeDemo() {
	// --- MatchTypeDynamic (reflect-based)
	fmt.Println("=== MatchTypeDynamic ===")

	m1 := maybe.Some(any(User{Name: "Taro"}))

	maybe.MatchTypeDynamic(m1, maybe.DynamicTypeHandlers{
		reflect.TypeOf(User{}): func(v any) {
			u := v.(User)
			fmt.Println("User name:", u.Name)
		},
		reflect.TypeOf(Admin{}): func(v any) {
			a := v.(Admin)
			fmt.Println("Admin level:", a.Level)
		},
	}, func() {
		fmt.Println("No match (dynamic)")
	})

	// --- MatchTypeKeyed (TypeKey-based)
	fmt.Println("\n=== MatchTypeKeyed ===")

	var m2 maybe.Maybe[maybe.Matchable]
	m2 = maybe.Some[maybe.Matchable](Admin{Level: 42})

	maybe.MatchTypeKeyed(m2, map[string]func(maybe.Matchable){
		"Admin": func(v maybe.Matchable) {
			a := v.(Admin)
			fmt.Println("Admin match:", a.Level)
		},
		"User": func(v maybe.Matchable) {
			u := v.(User)
			fmt.Println("User match:", u.Name)
		},
	}, func() {
		fmt.Println("No match (keyed)")
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
