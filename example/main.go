package main

import (
	"errors"
	"fmt"
	"os"
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
	fmt.Println("=== Maybe[T] ===")
	m := maybe.Some(10)
	mapped := maybe.Map(m, func(x int) string {
		return fmt.Sprintf("Number: %d", x)
	})
	mapped.Match(
		func(v string) { fmt.Println("Mapped value:", v) },
		func() { fmt.Println("No value") },
	)

	none := maybe.None[int]()
	fallback := none.OrElse(maybe.Some(99))
	fmt.Println("Fallback value:", fallback.Unwrap())

	r := maybe.ToResult(none, errors.New("no value found"))
	r.Match(
		func(v int) { fmt.Println("Got:", v) },
		func(e error) { fmt.Println("Error:", e) },
	)

	// ✅ Maybe.FromValue
	maybeMap := map[string]int{"x": 123}
	v := maybe.FromValue(maybeMap["x"], true)
	fmt.Println("FromValue result:", v.UnwrapOr(0))

	// ✅ Maybe.Try
	maybeParsed := maybe.Try(func() (int, error) {
		return strconv.Atoi("456")
	})
	fmt.Println("Try result:", maybeParsed.UnwrapOr(-1))

	fmt.Println("\n=== MaybePrimitive[T] ===")
	mp := maybe.SomePrimitive(42)
	mpStr := maybe.MapPrimitive(mp, func(x int) string {
		return fmt.Sprintf("Primitive: %d", x)
	})
	if mpStr.IsSome() {
		fmt.Println("Mapped primitive:", mpStr.Unwrap())
	}

	nonePrim := maybe.NonePrimitive[int]()
	fallbackPrim := nonePrim.OrElse(maybe.SomePrimitive(123))
	fmt.Println("Fallback primitive:", fallbackPrim.Unwrap())

	// ✅ MaybePrimitive.FromValuePrimitive
	ok := true
	p := maybe.FromValuePrimitive(777, ok)
	fmt.Println("FromValuePrimitive result:", p.UnwrapOr(0))

	// ✅ MaybePrimitive.TryPrimitive
	envVar := maybe.TryPrimitive(func() (string, error) {
		return os.Hostname()
	})
	fmt.Println("TryPrimitive result:", envVar.UnwrapOr("unknown"))

	fmt.Println("\n=== Result[T, E] ===")
	res := divide(10, 2)
	resStr := result.Map(res, func(x int) string {
		return fmt.Sprintf("Result: %d", x)
	})
	fmt.Println(resStr.UnwrapOr("default"))

	resFail := divide(10, 0)
	resRecovered := resFail.OrElse(func(e error) result.Result[int, error] {
		fmt.Println("Recovered from error:", e)
		return result.Ok[int, error](0)
	})
	fmt.Println("Recovered result:", resRecovered.Unwrap())

	// ✅ Result.From
	f, err := os.Open("nonexistent.txt")
	rFrom := result.From(f, err)
	rFrom.Match(
		func(v *os.File) { fmt.Println("Opened file:", v.Name()) },
		func(e error) { fmt.Println("From failed:", e) },
	)

	// ✅ Result.Try
	rTry := result.Try(
		func() (*os.File, error) {
			return os.Open("nonexistent.txt")
		},
		func(e error) error {
			return fmt.Errorf("wrapped: %w", e)
		},
	)
	rTry.Match(
		func(v *os.File) { fmt.Println("Try opened file:", v.Name()) },
		func(e error) { fmt.Println("Try failed:", e) },
	)
}
