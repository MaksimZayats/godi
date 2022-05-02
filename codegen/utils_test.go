package codegen

import (
	"github.com/elliotchance/tf"
	"testing"
)

func TestGenerateArgNames(t *testing.T) {
	generateArgNames := tf.Function(t, generateArgNames)

	generateArgNames(0).Returns([]string{})
	generateArgNames(1).Returns([]string{"a"})
	generateArgNames(2).Returns([]string{"a", "b"})
}

func TestGenerateFullSignature(t *testing.T) {
	generateFullSignature := tf.Function(t, generateFullSignature)

	generateFullSignature(func() {}).
		Returns("func()")
	generateFullSignature(func() int { return 0 }).
		Returns("func() int")
	generateFullSignature(func(a int) {}).
		Returns("func(int)")
	generateFullSignature(func(a int) int { return 0 }).
		Returns("func(int) int")
}

func TestGenerateRequiredSignature(t *testing.T) {
	generateRequiredSignature := tf.Function(t, generateRequiredSignature)

	generateRequiredSignature(func() {}).
		Returns("func()")
	generateRequiredSignature(func() int { return 0 }).
		Returns("func() int")
	generateRequiredSignature(func(a int) {}).
		Returns("func(a int)")
	generateRequiredSignature(func(a int) int { return 0 }).
		Returns("func(a int) int")
	generateRequiredSignature(func(a int, aa string, aaa float64) int { return 0 }).
		Returns("func(a int, b string, c float64) int")
}

func TestGenerateArgTypesToGet(t *testing.T) {
	generateArgTypesToGet := tf.Function(t, generateArgTypesToGet)

	generateArgTypesToGet(func() {}, func() {}).
		Returns([]string{})
	generateArgTypesToGet(func(a int) {}, func() {}).
		Returns([]string{"int"})
	generateArgTypesToGet(func(a int, b string) {}, func() {}).
		Returns([]string{"int", "string"})
	generateArgTypesToGet(func(a int, b string) {}, func(a int) {}).
		Returns([]string{"string"})
}
