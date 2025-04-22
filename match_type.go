package maybe

import (
	"reflect"
)

// --- Reflect-based dispatcher ---

type DynamicTypeHandlers map[reflect.Type]func(any)

func MatchTypeDynamic(m Maybe[any], handlers DynamicTypeHandlers, elseFn func()) {
	if m.IsNone() {
		elseFn()
		return
	}
	v := m.Unwrap()
	t := reflect.TypeOf(v)
	if h, ok := handlers[t]; ok {
		h(v)
	} else {
		elseFn()
	}
}

// --- TypeKey-based dispatcher ---

type Matchable interface {
	TypeKey() string
}

func MatchTypeKeyed(m Maybe[Matchable], handlers map[string]func(Matchable), elseFn func()) {
	if m.IsNone() {
		elseFn()
		return
	}
	v := m.Unwrap()
	key := v.TypeKey()
	if h, ok := handlers[key]; ok {
		h(v)
	} else {
		elseFn()
	}
}
