package codegen

import (
	"reflect"
	"strings"
)

var az = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

func generateArgNames(amount int) []string {
	return az[0:amount]
}

func generateFullSignature(f any) string {
	return reflect.TypeOf(f).String()
}

func generateRequiredSignature(f any) string {
	// "func(context.Context, float32)"
	t := reflect.TypeOf(f)

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
	for i, _ := range parts {
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

	args := strings.Replace(
		fInStringSig,
		fOutStringSig,
		"",
		1,
	)

	argsSlice := strings.Split(args, ", ")[1:]

	return argsSlice
}
