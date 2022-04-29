package main

import (
	"fmt"
	"github.com/MaximZayats/go-typed-di/di"
)

type A struct{ i int }
type B struct{ a A }

func simpleExample() {
	di.AddSingletonByFactory[A](func(c *di.Container) A {
		// Singleton: Will be called only 1 time
		fmt.Println("Init 'A'")
		return A{i: 123123}
	})

	di.AddSingletonByFactory[B](func(c *di.Container) B {
		// Singleton: Will be called only 1 time
		fmt.Println("Init 'B'")
		a := di.GetFromContainer[A](c)
		a.i = 100
		return B{a: a}
	})

	di.AddScopedByFactory[string](func(c *di.Container) string {
		// Default factory: Will be called every 'Get'
		fmt.Println("Init 'string'")
		return "aabbcc"
	})

	di.AddInstance[int](123)

	di.Get[A]()      // A{i: 123123}
	di.Get[B]()      // B{a: A{i: 100}}
	di.Get[B]()      // B{a: A{i: 100}}
	di.Get[B]()      // B{a: A{i: 100}}
	di.Get[int]()    // 123
	di.Get[string]() // "aabbcc"
	di.Get[string]() // "aabbcc"
	di.Get[string]() // "aabbcc"
}
