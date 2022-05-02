package codegen

import (
	"reflect"
	"strings"
)

var az = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

func contains(s []string, e string) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

func generateArgNames(amount int) []string {
	return az[0:amount]
}

func generateFullSignature(f any) string {
	return reflect.TypeOf(f).String()
}

func generateRequiredSignature(f any) string {
	// "func(context.Context, float32)"
	t := reflect.TypeOf(f)

	if t.NumIn() == 0 {
		return t.String()
	}

	argNames := generateArgNames(t.NumIn())

	// "func(a context.Context, float32)"
	sig := strings.Replace(
		t.String(),
		"(",
		"("+argNames[0]+" ",
		1,
	)

	// ["func(a context.Context", "float32)"]
	parts := strings.Split(sig, ", ")
	for i := range parts {
		if i == 0 {
			continue
		}

		parts[i] = argNames[i] + " " + parts[i]
	}

	return strings.Join(parts, ", ")
}

func generateArgTypesToGet(fIn any, fOut any) []string {
	fInStringSig := reflect.TypeOf(fIn).String()
	fOutStringSig := reflect.TypeOf(fOut).String()

	fInStringSig = strings.Split(fInStringSig, ")")[0]
	fOutStringSig = strings.Split(fOutStringSig, ")")[0]

	fInStringSig = strings.Replace(fInStringSig, "func(", "", 1)
	fOutStringSig = strings.Replace(fOutStringSig, "func(", "", 1)

	argsIn := strings.Split(fInStringSig, ", ")
	argsOut := strings.Split(fOutStringSig, ", ")

	argsToGet := make([]string, 0)

	for _, argIn := range argsIn {
		if !contains(argsOut, argIn) {
			argsToGet = append(argsToGet, argIn)
		}
	}

	return argsToGet
}
