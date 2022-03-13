package create

import (
	"fmt"
	"os"

	"github.com/Bananenpro/tmpl/templates"
	"github.com/Bananenpro/tmpl/util"
)

func CreatePython() {
	if !util.IsInstalled("python") {
		fmt.Println("Python is not installed.")
		os.Exit(1)
	}

	if util.YesNo("Create a virtual environment?", false) {
		out, err := util.Execute("python", "-m", "venv", "venv")
		if err != nil {
			fmt.Println("Failed to create a virtual environment:", out)
		}
	}

	file, err := os.Create("main.py")
	defer file.Close()
	if err != nil {
		Clean()
		fmt.Println("Error while creating file 'main.py':", err)
		os.Exit(1)
	}

	_, err = file.WriteString(templates.PythonHelloWorld)

	if err != nil {
		Clean()
		fmt.Println("Error while writing into file 'main.py':", err)
		os.Exit(1)
	}
}
