package result_test

import (
	"testing"

	"github.com/magicdrive/maybe/result"
)

func TestMatchOkIf(t *testing.T) {
	r := result.Ok[int, error](11)
	called := ""

	result.MatchOkIf(r, []result.MatchOkCase[int, error]{
		{Cond: func(x int) bool { return x > 100 }, Then: func(x int) {
			called = "first"
		}},
		{Cond: func(x int) bool { return x > 10 }, Then: func(x int) {
			called = "second"
		}},
	}, func(e error) {
		called = "error"
	}, func() {
		called = "else"
	})

	if called != "second" {
		t.Errorf("expected 'second', got: %s", called)
	}
}
