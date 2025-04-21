package maybe_test

import (
	"errors"
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
