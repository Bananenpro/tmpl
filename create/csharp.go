package create

import (
	"fmt"
	"os"

	"github.com/Bananenpro/tmpl/util"
)

func CreateCSharp() {
	if !util.IsInstalled("dotnet") {
		abort("Dotnet is not installed.")
	}

	template := util.Select("Choose a template", []string{"console", "web", "webapp", "mvc", "webapi"}, 0)

	out, err := util.Execute("dotnet", "new", template)
	if err != nil {
		Clean()
		fmt.Println("Failed to create project:", out)
		os.Exit(1)
	}
}
