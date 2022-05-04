package benchmark

import (
	"github.com/MaximZayats/godi/codegen"
	"github.com/MaximZayats/godi/di"
	"github.com/MaximZayats/godi/injection"
	"testing"
)

func testFunc(a TestType) TestType {
	return a
}

func BenchmarkWithoutInject(b *testing.B) {
	c := di.NewContainer()
	di.AddInstance(TestType{i: 222}, c)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		testFunc(di.MustGetFromContainer[TestType](c))
	}
}

func BenchmarkInject(b *testing.B) {
	c := di.NewContainer()
	di.AddInstance[TestType](TestType{i: 222}, c)

	injection.Configure(codegen.Config{
		GetterFunction: getDecoratorMocked,
	})

	injectedTestFunc := injection.Inject[func() TestType](testFunc, c)

	injection.VerifyInjections()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		injectedTestFunc()
	}
}

func getDecoratorMocked(f any) (any, bool) {
	return func(f func(testType TestType) TestType, c *di.Container) any {
		return func() TestType {
			return f(
				di.MustGetFromContainer[TestType](c),
			)
		}
	}, true
}
