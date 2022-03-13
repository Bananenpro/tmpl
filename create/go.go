package create

import (
	"fmt"
	"os"

	"github.com/Bananenpro/tmpl/templates"
	"github.com/Bananenpro/tmpl/util"
)

func CreateGo() {
	if !util.IsInstalled("go") {
		abort("Go is not installed.")
	}

	moduleName := util.Question("Enter a module name:")
	if moduleName == "" {
		fmt.Println("Cancel.")
		os.Exit(0)
	}

	out, err := util.Execute("go", "mod", "init", moduleName)
	if err != nil {
		Clean()
		fmt.Println("Failed to run 'go mod init':", out)
		os.Exit(1)
	}

	err = os.Mkdir("cmd", 0755)
	if err != nil {
		abort(fmt.Sprint("Error while creating directory 'cmd': ", err))
	}

	CreateFile("cmd/main.go", templates.GoHelloWorld)
}
