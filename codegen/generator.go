package codegen

import (
	"errors"
	"os"
	"os/exec"
	"text/template"
)

func Generate(
	config Config,
	signatures ...Signature,
) error {
	f, err := os.Create(config.GetPathToFile())
	if err != nil {
		return err
	}
	defer f.Close()

	tmpl, err := template.New("tmpl").Parse(fileTemplate)
	if err != nil {
		return err
	}

	err = tmpl.Execute(f, map[string]any{
		"config":     config,
		"signatures": signatures,
	})
	if err != nil {
		return err
	}

	details, err := exec.Command(
		"go", "run", "golang.org/x/tools/cmd/goimports", "-w", config.GetPathToFile(),
	).CombinedOutput()
	if err != nil {
		return errors.New(err.Error() + " | Details: " + string(details))
	}

	return nil
}
