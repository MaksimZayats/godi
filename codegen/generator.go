package codegen

import (
	"errors"
	"github.com/flosch/pongo2/v5"
	"os"
	"os/exec"
)

func Generate(
	config Config,
	signatures ...Signature,
) error {
	template := pongo2.Must(pongo2.FromString(jinjaTemplate))

	s, err := template.Execute(pongo2.Context{
		"config":      config,
		"signatures":  signatures,
		"packageName": config.PackageName,
	})
	if err != nil {
		return err
	}

	f, err := os.Create(config.GetFullPath())
	if err != nil {
		return err
	}
	defer f.Close()

	f.WriteString(s)

	details, err := exec.Command(
		"goimports", "-w", config.GetFullPath(),
	).CombinedOutput()
	if err != nil {
		return errors.New(err.Error() + " | Details: " + string(details))
	}

	return nil
}
