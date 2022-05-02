package pkg

import (
	"fmt"
	"github.com/MaximZayats/godi/di"
)

func CustomContainerExample() {
	container := di.NewContainer()

	di.AddInstance[int](123, container)

	di.AddByFactory[string](func(c *di.Container) string {
		return "aabbcc"
	}, container)

	fmt.Println(di.GetFromContainer[int](container))    // 123
	fmt.Println(di.GetFromContainer[string](container)) // "aabbcc"
	fmt.Println(di.Get[string]())                       // "aabbcc"
}
