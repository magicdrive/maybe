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
