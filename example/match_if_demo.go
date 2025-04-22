package main

import (
	"fmt"

	"github.com/magicdrive/maybe"
	"github.com/magicdrive/maybe/result"
)

func DemoMatchIf() {
	fmt.Println("\n=== MatchIf ===")

	m := maybe.Some(30)
	maybe.MatchIf(m, []maybe.MatchCase[int]{
		{Cond: func(x int) bool { return x > 100 }, Then: func(x int) {
			fmt.Println("Too large:", x)
		}},
		{Cond: func(x int) bool { return x > 10 }, Then: func(x int) {
			fmt.Println("Matched OK:", x)
		}},
	}, func() {
		fmt.Println("No match or None")
	})

	mp := maybe.SomePrimitive(5)
	maybe.MatchIfPrimitive(mp, []maybe.MatchPrimitiveCase[int]{
		{Cond: func(x int) bool { return x > 10 }, Then: func(x int) {
			fmt.Println("Primitive large:", x)
		}},
		{Cond: func(x int) bool { return x < 10 }, Then: func(x int) {
			fmt.Println("Primitive small:", x)
		}},
	}, func() {
		fmt.Println("Primitive fallback")
	})

	res := result.Ok[int, error](11)
	result.MatchOkIf(res, []result.MatchOkCase[int, error]{
		{Cond: func(x int) bool { return x > 100 }, Then: func(x int) {
			fmt.Println("Result: huge", x)
		}},
		{Cond: func(x int) bool { return x > 10 }, Then: func(x int) {
			fmt.Println("Result: fine", x)
		}},
	}, func(e error) {
		fmt.Println("Error happened:", e)
	}, func() {
		fmt.Println("No match in Result")
	})
}
