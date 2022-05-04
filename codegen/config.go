package codegen

import (
	"os/exec"
	"path/filepath"
	"strings"
)

var DefaultConfig = Config{
	PackageName:         "decorators",
	PathToStorageFolder: "./storage/decorators",
	StorageFileName:     "di_injection_decorators_gen.go",
	GetterFunction: func(f any) (any, bool) {
		panic("You must change default config!\nHint: `injection.Configure(decorators.Config)`")
	},
}

type Config struct {
	PackageName         string
	PathToStorageFolder string
	StorageFileName     string
	GetterFunction      func(f any) (any, bool)
}

func (c Config) GetPathToFile() string {
	// Resolving Abs path using git cli
	path, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		return filepath.Join(c.PathToStorageFolder, c.StorageFileName)
	}
	return filepath.Join(strings.TrimSpace(string(path)), c.PathToStorageFolder, c.StorageFileName)
}
