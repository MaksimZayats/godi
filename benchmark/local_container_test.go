package q

import (
	"github.com/MaximZayats/go-typed-di/di"
	"testing"
)

type TestType struct{ i int }

// 2.465 ns/op
func BenchmarkGetFromFactorySingleton(b *testing.B) {
	c := di.NewContainer()
	di.AddSingletonByFactory[TestType](func(c *di.Container) TestType {
		return TestType{i: 111}
	}, c)
	for i := 0; i < b.N; i++ {
		di.GetFromContainer[TestType](c)
	}
}

func BenchmarkGetInstance(b *testing.B) {
	c := di.NewContainer()
	di.AddInstance(TestType{i: 222}, c)
	for i := 0; i < b.N; i++ {
		di.GetFromContainer[TestType](c)
	}
}

func BenchmarkGetFromFactory(b *testing.B) {
	c := di.NewContainer()
	di.AddScopedByFactory(func(c *di.Container) TestType {
		return TestType{i: 111}
	}, c)
	for i := 0; i < b.N; i++ {
		di.GetFromContainer[TestType](c)
	}
}
