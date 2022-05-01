package q

import (
	"github.com/MaximZayats/go-typed-di/di"
	"reflect"
	"testing"
)

type TestType struct{ i int }

func Inject[ReturnType any, InputType any](
	innerFunc InputType,
	container ...di.Container,
) ReturnType {
	innerFuncValue := reflect.ValueOf(innerFunc)

	decorator := reflect.MakeFunc(
		reflect.TypeOf(*new(ReturnType)),
		func(args []reflect.Value) (results []reflect.Value) {
			args = append(args, reflect.ValueOf("Hi from dec"))
			return innerFuncValue.Call(args)
		},
	)

	f, ok := decorator.Interface().(ReturnType)
	if !ok {
		panic("Something went wrong")
	}

	return f
}

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
