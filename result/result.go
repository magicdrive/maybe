package result

type Result[T any, E error] struct {
	value T
	err   E
	ok    bool
}

func From[T any](v T, err error) Result[T, error] {
	if err != nil {
		return Err[T](err)
	}
	return Ok[T, error](v)
}

func Try[T any, E error](f func() (T, error), wrap func(error) E) Result[T, E] {
	v, err := f()
	if err != nil {
		return Err[T](wrap(err))
	}
	return Ok[T, E](v)
}

func Ok[T any, E error](v T) Result[T, E] {
	return Result[T, E]{value: v, ok: true}
}

func Err[T any, E error](e E) Result[T, E] {
	return Result[T, E]{err: e, ok: false}
}

func (r Result[T, E]) IsOk() bool {
	return r.ok
}

func (r Result[T, E]) IsErr() bool {
	return !r.ok
}

func (r Result[T, E]) Unwrap() T {
	if !r.ok {
		panic("called Unwrap on Err")
	}
	return r.value
}

func (r Result[T, E]) UnwrapOr(def T) T {
	if r.ok {
		return r.value
	}
	return def
}

func (r Result[T, E]) UnwrapErr() E {
	if r.ok {
		panic("called UnwrapErr on Ok")
	}
	return r.err
}

func (r Result[T, E]) OrElse(f func(E) Result[T, E]) Result[T, E] {
	if r.ok {
		return r
	}
	return f(r.err)
}

func Map[T any, E error, U any](r Result[T, E], f func(T) U) Result[U, E] {
	if r.ok {
		return Ok[U, E](f(r.value))
	}
	return Result[U, E]{err: r.err, ok: false}
}

func AndThen[T any, E error, U any](r Result[T, E], f func(T) Result[U, E]) Result[U, E] {
	if r.ok {
		return f(r.value)
	}
	return Result[U, E]{err: r.err, ok: false}
}

func (r Result[T, E]) Match(okFn func(T), errFn func(E)) {
	if r.ok {
		okFn(r.value)
	} else {
		errFn(r.err)
	}
}

func Fold[T any, E error, R any](r Result[T, E], okFn func(T) R, errFn func(E) R) R {
	if r.IsOk() {
		return okFn(r.value)
	}
	return errFn(r.err)
}

func Tap[T any, E error](r Result[T, E], f func(T)) Result[T, E] {
	if r.IsOk() {
		f(r.value)
	}
	return r
}
