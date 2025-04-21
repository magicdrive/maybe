package result_test

import (
	"errors"
	"testing"

	"github.com/magicdrive/maybe/result"
)

func TestOkAndErr(t *testing.T) {
	r1 := result.Ok[int, error](100)
	if !r1.IsOk() || r1.IsErr() {
		t.Fatal("expected Ok")
	}
	if r1.Unwrap() != 100 {
		t.Errorf("expected 100, got %v", r1.Unwrap())
	}

	r2 := result.Err[int](errors.New("fail"))
	if !r2.IsErr() || r2.IsOk() {
		t.Fatal("expected Err")
	}
	if r2.UnwrapOr(50) != 50 {
		t.Errorf("expected fallback 50")
	}
	if r2.UnwrapErr().Error() != "fail" {
		t.Errorf("expected error 'fail'")
	}
}

func TestMap(t *testing.T) {
	r := result.Ok[int, error](3)
	mapped := result.Map(r, func(x int) string {
		return "val"
	})
	if mapped.IsErr() || mapped.Unwrap() != "val" {
		t.Errorf("expected mapped Ok with 'val'")
	}

	err := result.Err[int](errors.New("fail"))
	mappedErr := result.Map(err, func(x int) string {
		return "val"
	})
	if mappedErr.IsOk() {
		t.Errorf("expected mapping over Err to remain Err")
	}
}

func TestAndThen(t *testing.T) {
	r := result.Ok[int, error](5)
	chained := result.AndThen(r, func(x int) result.Result[string, error] {
		return result.Ok[string, error]("chained")
	})
	if chained.IsErr() || chained.Unwrap() != "chained" {
		t.Errorf("expected chained Ok with 'chained'")
	}

	err := result.Err[int](errors.New("fail"))
	chainedErr := result.AndThen(err, func(x int) result.Result[string, error] {
		return result.Ok[string, error]("ignored")
	})
	if chainedErr.IsOk() {
		t.Errorf("expected AndThen over Err to remain Err")
	}
}

func TestOrElse(t *testing.T) {
	r := result.Ok[int, error](42)
	recovered := r.OrElse(func(e error) result.Result[int, error] {
		return result.Ok[int, error](99)
	})
	if recovered.Unwrap() != 42 {
		t.Errorf("expected original value 42, got %v", recovered.Unwrap())
	}

	err := result.Err[int](errors.New("boom"))
	recoveredErr := err.OrElse(func(e error) result.Result[int, error] {
		return result.Ok[int, error](123)
	})
	if recoveredErr.IsErr() || recoveredErr.Unwrap() != 123 {
		t.Errorf("expected recovery to 123")
	}
}

func TestMatch(t *testing.T) {
	var resultText string

	result.Ok[int, error](1).Match(
		func(v int) { resultText = "ok" },
		func(e error) { resultText = "err" },
	)
	if resultText != "ok" {
		t.Errorf("expected ok branch")
	}

	result.Err[int](errors.New("fail")).Match(
		func(v int) { resultText = "ok" },
		func(e error) { resultText = "err" },
	)
	if resultText != "err" {
		t.Errorf("expected err branch")
	}
}
