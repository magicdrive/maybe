package result

type MatchOkCase[T any, E error] struct {
	Cond func(T) bool
	Then func(T)
}

func MatchOkIf[T any, E error](
	r Result[T, E],
	cases []MatchOkCase[T, E],
	isErrFn func(E),
	elseFn func(),
) {
	if r.IsErr() {
		isErrFn(r.err)
		return
	}
	val := r.value
	for _, c := range cases {
		if c.Cond(val) {
			c.Then(val)
			return
		}
	}
	elseFn()
}
