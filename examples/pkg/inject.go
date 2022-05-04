package pkg

import (
	"context"
	"fmt"
	"github.com/MaximZayats/godi/di"
	"github.com/MaximZayats/godi/examples/storage/decorators"
	"github.com/MaximZayats/godi/injection"
)

func Handler(c context.Context, a int, b string) int {
	fmt.Println(c, a, b)
	return a
}

func Handler2(c context.Context, a int) int {
	fmt.Println(c, a)
	return a
}

// H is the type alias for functions after injection
type H = func(context.Context) int

func InjectionExample() {
	di.AddInstance(123)
	di.AddInstance("aabbcc")

	injection.Configure(decorators.Config)

	decoratedHandler := injection.Inject[H](Handler)
	decoratedHandler2 := injection.Inject[H](Handler2)

	injection.MustVerifyInjections()

	decoratedHandler(context.TODO())
	decoratedHandler(context.TODO())
	decoratedHandler2(context.TODO())
	decoratedHandler2(context.TODO())
}
