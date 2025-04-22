package maybe

import "github.com/magicdrive/maybe/result"

type Maybe[T any] struct {
	value T
	valid bool
}

func FromValue[T any](value T, ok bool) Maybe[T] {
	if ok {
		return Some(value)
	}
	return None[T]()
}

func Try[T any](f func() (T, error)) Maybe[T] {
	v, err := f()
	if err != nil {
		return None[T]()
	}
	return Some(v)
}

func Some[T any](v T) Maybe[T] {
	return Maybe[T]{value: v, valid: true}
}

func None[T any]() Maybe[T] {
	var zero T
	return Maybe[T]{value: zero, valid: false}
}

func (m Maybe[T]) IsSome() bool {
	return m.valid
}

func (m Maybe[T]) IsNone() bool {
	return !m.valid
}

func (m Maybe[T]) Unwrap() T {
	if !m.valid {
		panic("called Unwrap on None")
	}
	return m.value
}

func (m Maybe[T]) UnwrapOr(def T) T {
	if !m.valid {
		return def
	}
	return m.value
}

func (m Maybe[T]) OrElse(other Maybe[T]) Maybe[T] {
	if m.valid {
		return m
	}
	return other
}

func (m Maybe[T]) Match(someFn func(T), noneFn func()) {
	if m.valid {
		someFn(m.value)
	} else {
		noneFn()
	}
}

func ToResult[T any, E error](m Maybe[T], err E) result.Result[T, E] {
	if m.valid {
		return result.Ok[T, E](m.value)
	}
	return result.Err[T](err)
}

func Map[T any, U any](m Maybe[T], f func(T) U) Maybe[U] {
	if m.IsNone() {
		return None[U]()
	}
	return Some(f(m.value))
}

func AndThen[T any, U any](m Maybe[T], f func(T) Maybe[U]) Maybe[U] {
	if m.IsNone() {
		return None[U]()
	}
	return f(m.value)
}

func Filter[T any](m Maybe[T], pred func(T) bool) Maybe[T] {
	if m.IsSome() && pred(m.value) {
		return m
	}
	return None[T]()
}

func Fold[T any, R any](m Maybe[T], someFn func(T) R, noneVal R) R {
	if m.IsSome() {
		return someFn(m.value)
	}
	return noneVal
}

func Tap[T any](m Maybe[T], f func(T)) Maybe[T] {
	if m.IsSome() {
		f(m.value)
	}
	return m
}

func Flatten[T any](m Maybe[Maybe[T]]) Maybe[T] {
	if m.IsSome() {
		return m.Unwrap()
	}
	return None[T]()
}

// --- MaybePrimitive ---

func FromValuePrimitive[T Primitive](value T, ok bool) MaybePrimitive[T] {
	if ok {
		return SomePrimitive(value)
	}
	return NonePrimitive[T]()
}

func TryPrimitive[T Primitive](f func() (T, error)) MaybePrimitive[T] {
	v, err := f()
	if err != nil {
		return NonePrimitive[T]()
	}
	return SomePrimitive(v)
}

type Primitive interface {
	~int | ~float64 | ~string | ~bool
}

type MaybePrimitive[T Primitive] struct {
	value *T
}

func SomePrimitive[T Primitive](v T) MaybePrimitive[T] {
	return MaybePrimitive[T]{value: &v}
}

func NonePrimitive[T Primitive]() MaybePrimitive[T] {
	return MaybePrimitive[T]{value: nil}
}

func (m MaybePrimitive[T]) IsSome() bool {
	return m.value != nil
}

func (m MaybePrimitive[T]) IsNone() bool {
	return m.value == nil
}

func (m MaybePrimitive[T]) Unwrap() T {
	if m.value == nil {
		panic("called Unwrap on None")
	}
	return *m.value
}

func (m MaybePrimitive[T]) UnwrapOr(def T) T {
	if m.value == nil {
		return def
	}
	return *m.value
}

func (m MaybePrimitive[T]) OrElse(other MaybePrimitive[T]) MaybePrimitive[T] {
	if m.value != nil {
		return m
	}
	return other
}

func (m MaybePrimitive[T]) Match(someFn func(T), noneFn func()) {
	if m.value != nil {
		someFn(*m.value)
	} else {
		noneFn()
	}
}

func ToResultPrimitive[T Primitive, E error](m MaybePrimitive[T], err E) result.Result[T, E] {
	if m.value != nil {
		return result.Ok[T, E](*m.value)
	}
	return result.Err[T](err)
}

func MapPrimitive[T Primitive, U Primitive](m MaybePrimitive[T], f func(T) U) MaybePrimitive[U] {
	if m.value == nil {
		return NonePrimitive[U]()
	}
	res := f(*m.value)
	return SomePrimitive(res)
}

func AndThenPrimitive[T Primitive, U Primitive](m MaybePrimitive[T], f func(T) MaybePrimitive[U]) MaybePrimitive[U] {
	if m.value == nil {
		return NonePrimitive[U]()
	}
	return f(*m.value)
}

func FilterPrimitive[T Primitive](m MaybePrimitive[T], pred func(T) bool) MaybePrimitive[T] {
	if m.IsSome() && pred(*m.value) {
		return m
	}
	return NonePrimitive[T]()
}

func FoldPrimitive[T Primitive, R any](m MaybePrimitive[T], someFn func(T) R, noneVal R) R {
	if m.IsSome() {
		return someFn(*m.value)
	}
	return noneVal
}

func TapPrimitive[T Primitive](m MaybePrimitive[T], f func(T)) MaybePrimitive[T] {
	if m.IsSome() {
		f(*m.value)
	}
	return m
}

