package maybe_test

import (
	"testing"

	"github.com/magicdrive/maybe"
)

func TestMatchIf_Multi(t *testing.T) {
	m := maybe.Some(50)
	called := ""

	maybe.MatchIf(m, []maybe.MatchCase[int]{
		{Cond: func(x int) bool { return x > 100 }, Then: func(x int) {
			called = "first"
		}},
		{Cond: func(x int) bool { return x > 10 }, Then: func(x int) {
			called = "second"
		}},
	}, func() {
		called = "else"
	})

	if called != "second" {
		t.Errorf("expected second to be called, got: %s", called)
	}
}

func TestMatchIfPrimitive(t *testing.T) {
	mp := maybe.SomePrimitive(50)
	result := ""

	maybe.MatchIfPrimitive(mp, []maybe.MatchPrimitiveCase[int]{
		{Cond: func(x int) bool { return x > 100 }, Then: func(x int) {
			result = "too big"
		}},
		{Cond: func(x int) bool { return x > 10 }, Then: func(x int) {
			result = "ok"
		}},
	}, func() {
		result = "fallback"
	})

	if result != "ok" {
		t.Errorf("expected ok, got %s", result)
	}
}
