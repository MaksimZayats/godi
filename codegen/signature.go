package codegen

import (
	"reflect"
)

type Signature struct {
	FullSignature     string   // "func(context.Context, int)",
	RequiredSignature string   // "func(context.Context)",
	RequiredArgsNames []string // []string{"a"},
	ArgsTypesToGet    []string // []string{"int"},
	WithReturn        bool
}

func NewSignature(inFunc any, outFunc any) Signature {
	outFuncType := reflect.TypeOf(outFunc)
	return Signature{
		FullSignature:     generateFullSignature(inFunc),
		RequiredSignature: generateRequiredSignature(outFunc),
		RequiredArgsNames: generateArgNames(outFuncType.NumIn()),
		ArgsTypesToGet:    generateArgTypesToGet(inFunc, outFunc),
		WithReturn:        outFuncType.NumOut() != 0,
	}
}
