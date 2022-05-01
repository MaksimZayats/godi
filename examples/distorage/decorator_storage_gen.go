// Code generated by DI. DO NOT EDIT.

package distorage

import (
	"context"
	"reflect"

	"github.com/MaximZayats/go-typed-di/codegen"
	"github.com/MaximZayats/go-typed-di/di"
)

var Config = codegen.Config{
	PackageName:         "distorage",
	PathToStorageFolder: `.\examples\distorage`,
	StorageFileName:     "decorator_storage_gen.go",
	GetterFunction:      getDecorator,
}

var functions = map[string]any{

	"func(context.Context, int, string) int": func(f func(context.Context, int, string) int, c *di.Container) any {
		return func(a context.Context) int {
			return f(
				a,
				di.MustGetFromContainer[int](c), di.MustGetFromContainer[string](c),
			)
		}
	},

	"func(context.Context, int) int": func(f func(context.Context, int) int, c *di.Container) any {
		return func(a context.Context) int {
			return f(
				a,
				di.MustGetFromContainer[int](c),
			)
		}
	},
}

func getDecorator(f any) (any, bool) {
	v, ok := functions[reflect.TypeOf(f).String()]
	return v, ok
}