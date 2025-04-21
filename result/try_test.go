package result_test

import (
	"errors"
	"testing"

	"github.com/magicdrive/maybe/result"
)

type MyErr struct {
	msg string
}

func (e MyErr) Error() string {
	return e.msg
}

func TestFrom(t *testing.T) {
	r := result.From(123, nil)
	if r.IsErr() {
		t.Errorf("expected Ok(123)")
	}

	err := errors.New("fail")
	r2 := result.From(0, err)
	if r2.IsOk() {
		t.Errorf("expected Err(fail)")
	}
}

func TestTry_Success(t *testing.T) {
	res := result.Try(
		func() (int, error) {
			return 42, nil
		},
		func(e error) MyErr {
			return MyErr{msg: e.Error()}
		},
	)

	if res.IsErr() || res.Unwrap() != 42 {
		t.Errorf("expected Ok(42)")
	}
}

func TestTry_Failure(t *testing.T) {
	res := result.Try(
		func() (int, error) {
			return 0, errors.New("something went wrong")
		},
		func(e error) MyErr {
			return MyErr{msg: "wrapped: " + e.Error()}
		},
	)

	if res.IsOk() {
		t.Errorf("expected Err, got Ok(%v)", res.Unwrap())
	}

	if res.UnwrapErr().Error() != "wrapped: something went wrong" {
		t.Errorf("expected wrapped error, got %v", res.UnwrapErr())
	}
}
