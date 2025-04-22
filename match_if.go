package maybe

type MatchCase[T any] struct {
	Cond func(T) bool
	Then func(T)
}

func MatchIf[T any](m Maybe[T], cases []MatchCase[T], elseFn func()) {
	if m.IsNone() {
		elseFn()
		return
	}
	val := m.Unwrap()
	for _, c := range cases {
		if c.Cond(val) {
			c.Then(val)
			return
		}
	}
	elseFn()
}

type MatchPrimitiveCase[T Primitive] struct {
	Cond func(T) bool
	Then func(T)
}

func MatchIfPrimitive[T Primitive](
	m MaybePrimitive[T],
	cases []MatchPrimitiveCase[T],
	elseFn func(),
) {
	if m.IsNone() {
		elseFn()
		return
	}
	val := *m.value
	for _, c := range cases {
		if c.Cond(val) {
			c.Then(val)
			return
		}
	}
	elseFn()
}
