package create

import (
	"fmt"
	"os"

	"github.com/Bananenpro/tmpl/external"
	"github.com/Bananenpro/tmpl/input"
	"github.com/Bananenpro/tmpl/templates"
)

func CreateGo() {
	if !external.IsInstalled("go") {
		abort("Go is not installed.")
	}

	moduleName := input.Input("Enter a module name:")
	if moduleName == "" {
		fmt.Println("Cancel.")
		os.Exit(0)
	}

	out, err := external.Execute("go", "mod", "init", moduleName)
	if err != nil {
		abort(fmt.Sprint("Failed to run 'go mod init':", out))
	}

	CreateFile("main.go", templates.GoHelloWorld)
}
