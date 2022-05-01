package codegen

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

const (
	sep = string(os.PathSeparator)
)

var baseDir, _ = os.Getwd()
var DefaultConfig = Config{
	PackageName:         "distorage",
	PathToStorageFolder: "./distorage",
	StorageFileName:     "decorator_storage_gen.go",
	GetterFunction: func(f any) (any, bool) {
		panic("You must change default config!\nHint: `injection.Configure(distorage.Config)`")
	},
}

type Config struct {
	PackageName         string
	PathToStorageFolder string
	StorageFileName     string
	GetterFunction      func(f any) (any, bool)
}

func (c Config) GetFullPath() string {
	return baseDir + sep + c.PathToStorageFolder + sep + c.StorageFileName
}

func (c *Config) loadFromToml() {}

func FindConfig(root string) (Config, bool) {
	// var a []string
	fmt.Println("Here")
	filepath.WalkDir(root, func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
		}

		switch filepath.Ext(d.Name()) {
		case ".toml":
			fmt.Println(s)
		case ".ini":
			fmt.Println(s)
		}

		return nil
	})
	return DefaultConfig, false
}
