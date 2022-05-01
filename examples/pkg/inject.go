package pkg

import (
	"context"
	"fmt"
	"github.com/MaximZayats/go-typed-di/di"
	"github.com/MaximZayats/go-typed-di/examples/distorage"
	"github.com/MaximZayats/go-typed-di/injection"
	"os"
)

func Handler(c context.Context, a int, b string) int {
	fmt.Println(c, a, b)
	return a
}

func Handler2(c context.Context, a int) int {
	fmt.Println(c, a)
	return a
}

type H = func(context.Context) int

func InjectionExample() {
	di.AddInstance(123)
	di.AddInstance("aabbcc")

	injection.Configure(distorage.Config)

	decoratedHandler := injection.Inject[H](Handler)
	decoratedHandler2 := injection.Inject[H](Handler2)

	ok := injection.VerifyInjections()
	if !ok {
		fmt.Println(
			"The injection functions have been changed.\n" +
				"A restart is required.",
		)
		os.Exit(0)
	}

	decoratedHandler(context.TODO())
	decoratedHandler(context.TODO())
	decoratedHandler2(context.TODO())
	decoratedHandler2(context.TODO())
}
