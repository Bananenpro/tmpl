package create

import (
	"fmt"

	"github.com/Bananenpro/tmpl/external"
	"github.com/Bananenpro/tmpl/input"
	"github.com/Bananenpro/tmpl/templates"
)

func CreatePython() {
	if !external.IsInstalled("python") {
		abort("Python is not installed.")
	}

	if input.YesNo("Create a virtual environment?", false) {
		out, err := external.Execute("python", "-m", "venv", "venv")
		if err != nil {
			fmt.Println("Failed to create a virtual environment:", out)
		}
	}

	CreateFile("main.py", templates.PythonHelloWorld)
}
