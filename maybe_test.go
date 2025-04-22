package maybe_test

import (
	"errors"
	"testing"

	"github.com/magicdrive/maybe"
)

func TestSomeAndNone(t *testing.T) {
	m := maybe.Some(42)
	if m.IsNone() || !m.IsSome() {
		t.Fatal("expected Some(42)")
	}
	if m.Unwrap() != 42 {
		t.Errorf("expected 42, got %v", m.Unwrap())
	}

	n := maybe.None[int]()
	if n.IsSome() || !n.IsNone() {
		t.Fatal("expected None")
	}
	if n.UnwrapOr(10) != 10 {
		t.Errorf("expected default 10, got %v", n.UnwrapOr(10))
	}
}

func TestMaybeMapAndThen(t *testing.T) {
	m := maybe.Some(2)
	mapped := maybe.Map(m, func(x int) string {
		return "val"
	})
	if !mapped.IsSome() || mapped.Unwrap() != "val" {
		t.Errorf("expected mapped value 'val'")
	}

	chained := maybe.AndThen(m, func(x int) maybe.Maybe[string] {
		return maybe.Some("chain")
	})
	if chained.IsNone() || chained.Unwrap() != "chain" {
		t.Errorf("expected chained value 'chain'")
	}
}

func TestMaybeMatch(t *testing.T) {
	var called string
	maybe.Some(7).Match(
		func(x int) { called = "some" },
		func() { called = "none" },
	)
	if called != "some" {
		t.Errorf("expected some branch to be called")
	}

	maybe.None[int]().Match(
		func(x int) { called = "some" },
		func() { called = "none" },
	)
	if called != "none" {
		t.Errorf("expected none branch to be called")
	}
}

func TestToResult(t *testing.T) {
	r1 := maybe.ToResult(maybe.Some("yes"), errors.New("fail"))
	if r1.IsErr() || r1.Unwrap() != "yes" {
		t.Errorf("expected Result Ok with 'yes'")
	}

	r2 := maybe.ToResult(maybe.None[string](), errors.New("fail"))
	if r2.IsOk() || r2.UnwrapErr().Error() != "fail" {
		t.Errorf("expected Result Err with 'fail'")
	}
}

func TestMaybePrimitive(t *testing.T) {
	p := maybe.SomePrimitive(123)
	if !p.IsSome() || p.Unwrap() != 123 {
		t.Fatal("expected SomePrimitive(123)")
	}

	n := maybe.NonePrimitive[int]()
	if n.IsSome() {
		t.Fatal("expected NonePrimitive")
	}
	if n.UnwrapOr(999) != 999 {
		t.Errorf("expected fallback 999")
	}
}

func TestMapAndThenPrimitive(t *testing.T) {
	p := maybe.SomePrimitive(5)

	pm := maybe.MapPrimitive(p, func(x int) int {
		return x * 2
	})
	if !pm.IsSome() || pm.Unwrap() != 10 {
		t.Errorf("expected mapped primitive 10")
	}

	chained := maybe.AndThenPrimitive(p, func(x int) maybe.MaybePrimitive[int] {
		return maybe.SomePrimitive(x + 10)
	})
	if chained.IsNone() || chained.Unwrap() != 15 {
		t.Errorf("expected AndThenPrimitive result 15")
	}
}

func TestToResultPrimitive(t *testing.T) {
	r1 := maybe.ToResultPrimitive(maybe.SomePrimitive("hello"), errors.New("fail"))
	if r1.IsErr() || r1.Unwrap() != "hello" {
		t.Errorf("expected Result Ok with 'hello'")
	}

	r2 := maybe.ToResultPrimitive(maybe.NonePrimitive[string](), errors.New("fail"))
	if r2.IsOk() || r2.UnwrapErr().Error() != "fail" {
		t.Errorf("expected Result Err with 'fail'")
	}
}

func TestFilter(t *testing.T) {
	m := maybe.Some(10)
	filtered := maybe.Filter(m, func(x int) bool { return x > 5 })
	if filtered.IsNone() || filtered.Unwrap() != 10 {
		t.Errorf("expected value to pass filter")
	}

	filtered2 := maybe.Filter(m, func(x int) bool { return x < 5 })
	if filtered2.IsSome() {
		t.Errorf("expected value to be filtered out")
	}
}

func TestFilterPrimitive(t *testing.T) {
	mp := maybe.SomePrimitive(20)
	filtered := maybe.FilterPrimitive(mp, func(x int) bool { return x == 20 })
	if filtered.IsNone() {
		t.Errorf("expected value to pass primitive filter")
	}

	filtered2 := maybe.FilterPrimitive(mp, func(x int) bool { return false })
	if filtered2.IsSome() {
		t.Errorf("expected value to be filtered out")
	}
}

func TestTap(t *testing.T) {
	m := maybe.Some(123)
	var observed int
	result := maybe.Tap(m, func(x int) {
		observed = x
	})
	if observed != 123 {
		t.Errorf("expected Tap to observe 123, got %d", observed)
	}
	if result.IsNone() || result.Unwrap() != 123 {
		t.Errorf("expected result to remain Some(123)")
	}

	n := maybe.None[int]()
	observed = 0
	result2 := maybe.Tap(n, func(x int) {
		observed = x
	})
	if result2.IsSome() {
		t.Errorf("expected IsSome to result2 false, got true")
	}
	if observed != 0 {
		t.Errorf("expected Tap to not execute on None")
	}
}

func TestTapPrimitive(t *testing.T) {
	m := maybe.SomePrimitive(99)
	var log int
	result := maybe.TapPrimitive(m, func(x int) {
		log = x
	})
	if log != 99 {
		t.Errorf("expected TapPrimitive to observe 99")
	}
	if result.Unwrap() != 99 {
		t.Errorf("expected value to remain 99")
	}
}

func TestFlatten(t *testing.T) {
	// Some(Some(42)) -> Some(42)
	inner := maybe.Some(42)
	outer := maybe.Some(inner)

	result := maybe.Flatten(outer)
	if result.IsNone() || result.Unwrap() != 42 {
		t.Errorf("expected Flatten(Some(Some(42))) to be Some(42), got %+v", result)
	}

	// Some(None) -> None
	outerNone := maybe.Some(maybe.None[int]())
	result2 := maybe.Flatten(outerNone)
	if result2.IsSome() {
		t.Errorf("expected Flatten(Some(None)) to be None")
	}

	// None -> None
	none := maybe.None[maybe.Maybe[int]]()
	result3 := maybe.Flatten(none)
	if result3.IsSome() {
		t.Errorf("expected Flatten(None) to be None")
	}
}
