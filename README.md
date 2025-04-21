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

```go
package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/magicdrive/maybe"
	"github.com/magicdrive/maybe/result"
)

func divide(a, b int) result.Result[int, error] {
	if b == 0 {
		return result.Err[int](errors.New("division by zero"))
	}
	return result.Ok[int](a / b)
}

func main() {
	// Maybe
	m := maybe.Some(42)
	val := maybe.Map(m, func(x int) string {
		return fmt.Sprintf("Value is %d", x)
	})
	fmt.Println(val.UnwrapOr("None"))

	// Maybe.FromValue
	data := map[string]int{"x": 123}
	v := maybe.FromValue(data["x"], true)
	fmt.Println(v.UnwrapOr(0)) // -> 123

	// Maybe.Try
	parsed := maybe.Try(func() (int, error) {
		return strconv.Atoi("456")
	})
	fmt.Println(parsed.UnwrapOr(-1)) // -> 456

	// Result
	r := result.Map(divide(10, 2), func(x int) string {
		return fmt.Sprintf("Success: %d", x)
	})
	fmt.Println(r.UnwrapOr("Error occurred"))

	// Result.From
	f, err := os.Open("file.txt")
	fileResult := result.From(f, err)
	fileResult.Match(
		func(f *os.File) { fmt.Println("Opened:", f.Name()) },
		func(e error) { fmt.Println("Open failed:", e) },
	)

	// Result.Try
	rTry := result.Try(
		func() (*os.File, error) {
			return os.Open("missing.txt")
		},
		func(e error) error {
			return fmt.Errorf("wrapped: %w", e)
		},
	)
	rTry.Match(
		func(f *os.File) { fmt.Println("Try opened:", f.Name()) },
		func(e error) { fmt.Println("Try failed:", e) },
	)
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
