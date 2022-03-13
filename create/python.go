package create

import (
	"fmt"

	"github.com/Bananenpro/tmpl/templates"
	"github.com/Bananenpro/tmpl/util"
)

func CreatePython() {
	if !util.IsInstalled("python") {
		abort("Python is not installed.")
	}

	if util.YesNo("Create a virtual environment?", false) {
		out, err := util.Execute("python", "-m", "venv", "venv")
		if err != nil {
			fmt.Println("Failed to create a virtual environment:", out)
		}
	}

	CreateFile("main.py", templates.PythonHelloWorld)
}
