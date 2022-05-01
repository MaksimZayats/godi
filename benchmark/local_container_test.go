package q

import (
	"github.com/MaximZayats/godi/di"
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

func BenchmarkMap(b *testing.B) {
	m := make(map[string]int)

	m["123"] = 123

	for i := 0; i < b.N; i++ {
		_ = m["123"]
	}
}
