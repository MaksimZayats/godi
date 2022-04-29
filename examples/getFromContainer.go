package main

import (
	"fmt"
	"github.com/MaximZayats/go-typed-di/di"
)

func customContainerExample() {
	container := di.NewContainer()

	di.AddInstance[int](123, container)

	di.AddScopedByFactory[string](func(c *di.Container) string {
		return "aabbcc"
	}, container)

	fmt.Println(di.GetFromContainer[int](container))    // 123
	fmt.Println(di.GetFromContainer[string](container)) // "aabbcc"
	fmt.Println(di.Get[string]())                       // "aabbcc"
}
