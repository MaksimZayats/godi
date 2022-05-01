package main

import (
	"flag"
	"fmt"
	codegen2 "github.com/MaximZayats/go-typed-di/codegen"
	"github.com/fatih/color"
	"os"
	"strings"
)

const sep = string(os.PathSeparator)

func main() {
	flag.Usage = func(f func()) func() {
		return func() {
			c := color.New(color.FgGreen, color.Bold, color.Underline)
			_, _ = c.Print("Usage:")
			fmt.Println("   godi init path/to/di_storage_folder")

			_, _ = c.Print("Example:")
			fmt.Println(" godi init ./distorage")

			_, _ = c.Print("Note:")
			fmt.Println("    path must be relative!")

			f()
		}
	}(flag.Usage)

	filename := flag.String(
		"filename",
		codegen2.DefaultConfig.StorageFileName,
		"Storage file name",
	)

	packageName := flag.String(
		"package",
		codegen2.DefaultConfig.PackageName,
		"Package name",
	)

	flag.Parse()

	args := flag.Args()
	if len(args) != 2 {
		flag.Usage()
		return
	}

	path := strings.TrimSuffix(args[1], "/")
	path = strings.TrimSuffix(path, "\\")

	action := strings.ToLower(args[0])

	switch action {
	case "init":
		c := color.New(color.FgWhite, color.Bold, color.Underline)
		_, _ = c.Print("Path:")
		fullPath := path + sep + *filename
		fmt.Println("   " + fullPath)

		if _, err := os.Stat(path); os.IsNotExist(err) {
			err = os.Mkdir(path, 777)
			if err != nil {
				panic(err)
			}
		}

		err := codegen2.Generate(codegen2.Config{
			PackageName:         *packageName,
			PathToStorageFolder: path,
			StorageFileName:     *filename,
		})
		if err != nil {
			c = color.New(color.FgRed, color.Bold)
			_, _ = c.Printf("Error when generating file: %s\n", err)
			return
		}

		c = color.New(color.FgGreen, color.Bold, color.Underline)
		_, _ = c.Print("Success:")
		fmt.Println(" Storage file was created")
	default:
		c := color.New(color.FgRed, color.Bold, color.Underline)
		_, _ = c.Print("Error:")
		fmt.Println(" Unrecognized action: " + action)
		return
	}
}
