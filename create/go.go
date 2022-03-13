package create

import (
	"fmt"
	"os"

	"github.com/Bananenpro/tmpl/templates"
	"github.com/Bananenpro/tmpl/util"
)

func CreateGo() {
	if !util.IsInstalled("go") {
		fmt.Println("Go is not installed.")
		os.Exit(1)
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
		Clean()
		fmt.Println("Error while creating directory 'cmd':", err)
		os.Exit(1)
	}

	file, err := os.Create("cmd/main.go")
	defer file.Close()
	if err != nil {
		Clean()
		fmt.Println("Error while creating file 'cmd/main.go':", err)
		os.Exit(1)
	}

	_, err = file.WriteString(templates.GoHelloWorld)

	if err != nil {
		Clean()
		fmt.Println("Error while writing into file 'cmd/main.go':", err)
		os.Exit(1)
	}
}
