package q

import (
	"context"
	"fmt"
	"github.com/MaximZayats/godi/di"
	"testing"
)

func handler(c context.Context, t TestType) {
	_, _ = c, t
}

func decHandler(c context.Context) {
	handler(c, di.MustGet[TestType]())
}

func BenchmarkDefaultFunction(b *testing.B) {
	di.AddInstance(TestType{123})

	fmt.Println(b.N)

	for i := 0; i < b.N; i++ {
		handler(context.TODO(), di.MustGet[TestType]())
	}
}

func BenchmarkInjectFunction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		decHandler(context.TODO())
	}
}
