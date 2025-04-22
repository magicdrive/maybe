package maybe_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/magicdrive/maybe"
)

type MyUser struct {
	Name string
}

func (u MyUser) TypeKey() string {
	return "User"
}

type MyAdmin struct {
	Level int
}

func (a MyAdmin) TypeKey() string {
	return "Admin"
}

func TestMatchTypeDynamic(t *testing.T) {
	called := ""

	user := MyUser{Name: "Alice"}
	m := maybe.Some(any(user))

	maybe.MatchTypeDynamic(m, maybe.DynamicTypeHandlers{
		reflect.TypeOf(MyUser{}): func(v any) {
			u := v.(MyUser)
			called = "user:" + u.Name
		},
	}, func() {
		t.Error("should not fallback")
	})

	if called != "user:Alice" {
		t.Errorf("unexpected result: %s", called)
	}
}

func TestMatchTypeKeyed(t *testing.T) {
	called := ""

	admin := MyAdmin{Level: 5}
	var m maybe.Maybe[maybe.Matchable]
	m = maybe.Some[maybe.Matchable](admin)

	maybe.MatchTypeKeyed(m, map[string]func(maybe.Matchable){
		"Admin": func(v maybe.Matchable) {
			a := v.(MyAdmin)
			called = fmt.Sprintf("admin:%d", a.Level)
		},
	}, func() {
		t.Error("should not fallback")
	})

	if called != "admin:5" {
		t.Errorf("unexpected result: %s", called)
	}
}
