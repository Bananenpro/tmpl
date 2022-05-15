package create

import (
	"fmt"
	"strings"

	"github.com/Bananenpro/tmpl/external"
	"github.com/Bananenpro/tmpl/input"
)

func CreateCSharp() {
	if !external.IsInstalled("dotnet") {
		abort("Dotnet is not installed.")
	}

	template, cancel := input.Select("Choose a template", []string{"Console", "Web", "Web App", "MVC", "Web API"})
	if cancel {
		stop()
	}
	template = strings.ToLower(strings.ReplaceAll(template, " ", ""))

	out, err := external.Execute("dotnet", "new", template)
	if err != nil {
		abort(fmt.Sprint("Failed to create project: ", out))
	}
}
