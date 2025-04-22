package maybe_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/magicdrive/maybe"
)

type MyErr struct {
	msg string
}

func (e MyErr) Error() string {
	return e.msg
}

func TestFromValue(t *testing.T) {
	val := maybe.FromValue(100, true)
	if val.IsNone() || val.Unwrap() != 100 {
		t.Errorf("expected Some(100)")
	}

	none := maybe.FromValue(0, false)
	if none.IsSome() {
		t.Errorf("expected None")
	}
}

func TestTry(t *testing.T) {
	okFn := func() (int, error) {
		return 7, nil
	}
	failFn := func() (int, error) {
		return 0, errors.New("fail")
	}

	res := maybe.Try(okFn)
	if res.IsNone() || res.Unwrap() != 7 {
		t.Errorf("expected Some(7)")
	}

	none := maybe.Try(failFn)
	if none.IsSome() {
		t.Errorf("expected None")
	}
}

func TestFromValuePrimitive(t *testing.T) {
	val := maybe.FromValuePrimitive(10, true)
	if val.IsNone() || val.Unwrap() != 10 {
		t.Errorf("expected SomePrimitive(10)")
	}

	none := maybe.FromValuePrimitive(0, false)
	if none.IsSome() {
		t.Errorf("expected NonePrimitive")
	}
}

func TestTryPrimitive(t *testing.T) {
	okFn := func() (int, error) {
		return 5, nil
	}
	failFn := func() (int, error) {
		return 0, errors.New("fail")
	}

	res := maybe.TryPrimitive(okFn)
	if res.IsNone() || res.Unwrap() != 5 {
		t.Errorf("expected SomePrimitive(5)")
	}

	none := maybe.TryPrimitive(failFn)
	if none.IsSome() {
		t.Errorf("expected NonePrimitive")
	}
}

func TestFold(t *testing.T) {
	m := maybe.Some(5)
	result := maybe.Fold(m, func(x int) string {
		return fmt.Sprintf("val=%d", x)
	}, "none")
	if result != "val=5" {
		t.Errorf("expected val=5, got %s", result)
	}

	none := maybe.None[int]()
	result2 := maybe.Fold(none, func(x int) string {
		return "should not happen"
	}, "none")
	if result2 != "none" {
		t.Errorf("expected 'none', got %s", result2)
	}
}

func TestFoldPrimitive(t *testing.T) {
	m := maybe.SomePrimitive(10)
	result := maybe.FoldPrimitive(m, func(x int) string {
		return fmt.Sprintf("prim=%d", x)
	}, "none")
	if result != "prim=10" {
		t.Errorf("expected prim=10, got %s", result)
	}

	n := maybe.NonePrimitive[int]()
	result2 := maybe.FoldPrimitive(n, func(x int) string {
		return "should not happen"
	}, "none")
	if result2 != "none" {
		t.Errorf("expected 'none', got %s", result2)
	}
}
