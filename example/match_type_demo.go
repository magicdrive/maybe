package main

import (
	"fmt"
	"reflect"

	"github.com/magicdrive/maybe"
)

type User struct {
	Name string
}

func (u User) TypeKey() string {
	return "User"
}

type Admin struct {
	Level int
}

func (a Admin) TypeKey() string {
	return "Admin"
}

func MatchTypeDemo() {
	// --- MatchTypeDynamic (reflect-based)
	fmt.Println("=== MatchTypeDynamic ===")

	m1 := maybe.Some(any(User{Name: "Taro"}))

	maybe.MatchTypeDynamic(m1, maybe.DynamicTypeHandlers{
		reflect.TypeOf(User{}): func(v any) {
			u := v.(User)
			fmt.Println("User name:", u.Name)
		},
		reflect.TypeOf(Admin{}): func(v any) {
			a := v.(Admin)
			fmt.Println("Admin level:", a.Level)
		},
	}, func() {
		fmt.Println("No match (dynamic)")
	})

	// --- MatchTypeKeyed (TypeKey-based)
	fmt.Println("\n=== MatchTypeKeyed ===")

	var m2 maybe.Maybe[maybe.Matchable]
	m2 = maybe.Some[maybe.Matchable](Admin{Level: 42})

	maybe.MatchTypeKeyed(m2, map[string]func(maybe.Matchable){
		"Admin": func(v maybe.Matchable) {
			a := v.(Admin)
			fmt.Println("Admin match:", a.Level)
		},
		"User": func(v maybe.Matchable) {
			u := v.(User)
			fmt.Println("User match:", u.Name)
		},
	}, func() {
		fmt.Println("No match (keyed)")
	})
}
