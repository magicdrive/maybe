package main

import (
	"fmt"
	"strconv"

	"github.com/magicdrive/maybe"
)

func DemoMaybePrimitive() {

	fmt.Println("\n=== MaybePrimitive ===")

	// SomePrimitive / UnwrapOr
	mp := maybe.SomePrimitive(42)
	fmt.Println("Primitive value:", mp.UnwrapOr(0)) // → 42

	// FilterPrimitive
	filteredPrim := maybe.FilterPrimitive(mp, func(x int) bool { return x > 40 })
	fmt.Println("Filtered primitive:", filteredPrim.UnwrapOr(-1)) // → 42

	// TapPrimitive
	maybe.TapPrimitive(filteredPrim, func(x int) {
		fmt.Println("Tapped primitive:", x)
	})

	// MapPrimitive
	mappedPrim := maybe.MapPrimitive(mp, func(x int) string {
		return fmt.Sprintf("val=%d", x)
	})
	fmt.Println("Mapped primitive:", mappedPrim.UnwrapOr("none"))

	// FoldPrimitive
	foldedPrim := maybe.FoldPrimitive(mp,
		func(x int) string { return fmt.Sprintf("prim:%d", x) },
		"none",
	)
	fmt.Println("Fold primitive:", foldedPrim)

	// FromValuePrimitive
	fromVal := maybe.FromValuePrimitive(999, true)
	fmt.Println("FromValuePrimitive:", fromVal.UnwrapOr(-1)) // → 999

	// TryPrimitive
	tryPrim := maybe.TryPrimitive(func() (int, error) {
		return strconv.Atoi("123")
	})
	fmt.Println("TryPrimitive:", tryPrim.UnwrapOr(-1)) // → 123

}
