package result_test

import (
	"errors"
	"fmt"
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

func TestResultFold(t *testing.T) {
	r := result.Ok[int, error](10)
	s := result.Fold(r,
		func(v int) string { return fmt.Sprintf("OK: %d", v) },
		func(e error) string { return "ERR: " + e.Error() },
	)
	if s != "OK: 10" {
		t.Errorf("expected 'OK: 10', got '%s'", s)
	}

	r2 := result.Err[int](errors.New("fail"))
	s2 := result.Fold(r2,
		func(v int) string { return "OK" },
		func(e error) string { return "ERR: " + e.Error() },
	)
	if s2 != "ERR: fail" {
		t.Errorf("expected 'ERR: fail', got '%s'", s2)
	}
}

func TestResultTap(t *testing.T) {
	r := result.Ok[int, error](10)
	var logged int
	r2 := result.Tap(r, func(v int) {
		logged = v
	})
	if logged != 10 {
		t.Errorf("expected Tap to log 10, got %d", logged)
	}
	if r2.Unwrap() != 10 {
		t.Errorf("result should still be 10")
	}

	rErr := result.Err[int, error](errors.New("fail"))
	logged = 0
	r3 := result.Tap(rErr, func(v int) {
		logged = 999
	})
	if logged != 0 {
		t.Errorf("expected Tap to not execute on Err")
	}
	if r3.UnwrapErr().Error() != "fail" {
		t.Errorf("expected Err unchanged")
	}
}
