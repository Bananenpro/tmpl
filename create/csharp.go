package create

import (
	"fmt"

	"github.com/Bananenpro/tmpl/external"
	"github.com/Bananenpro/tmpl/input"
)

func CreateCSharp() {
	if !external.IsInstalled("dotnet") {
		abort("Dotnet is not installed.")
	}

	template := input.Select("Choose a template", []string{"console", "web", "webapp", "mvc", "webapi"}, 0)

	out, err := external.Execute("dotnet", "new", template)
	if err != nil {
		abort(fmt.Sprint("Failed to create project: ", out))
	}
}
