package main

import (
	"flag"
	"fmt"
	"github.com/MaximZayats/godi/codegen"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	flag.Usage = func(f func()) func() {
		return func() {
			fmt.Println("ℹ️Usage: godi init relative_path/to/di_storage_folder")
			fmt.Println("ℹ️Example: godi init ./distorage")
			fmt.Println("ℹ️Note: path must be relative!")
			f()
		}
	}(flag.Usage)

	filename := flag.String(
		"filename",
		codegen.DefaultConfig.StorageFileName,
		"Storage file name",
	)

	packageName := flag.String(
		"package",
		codegen.DefaultConfig.PackageName,
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

	if filepath.IsAbs(path) {
		printlnError("Absolute path is unavailable")
		printlnInfo("Use a relative path from the root directory")
		return
	}

	action := strings.ToLower(args[0])

	switch action {
	case "init":
		config := codegen.Config{
			PackageName:         *packageName,
			PathToStorageFolder: path,
			StorageFileName:     *filename,
		}

		printlnInfo("Path to file: " + config.GetPathToFile())

		if _, err := os.Stat(path); os.IsNotExist(err) {
			printlnInfo("Directory " + path + " is not exist...")
			printlnInfo("Creating directory: " + path)
			err = os.MkdirAll(path, 777)
			if err != nil {
				panic(err)
			}
		}

		err := codegen.Generate(config)
		if err != nil {
			printlnError("When generating file: " + err.Error())
			return
		}
		printlnSuccess("Storage file was created")
	default:
		printlnError("Unrecognized action: " + action)
	}
}

func printlnError(message string) {
	fmt.Println("❌Error: " + message)
}

func printlnSuccess(message string) {
	fmt.Println("✅Success: " + message)
}

func printlnInfo(message string) {
	fmt.Println("ℹ️Info: " + message)
}
