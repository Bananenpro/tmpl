package create

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/Bananenpro/tmpl/external"
	"github.com/Bananenpro/tmpl/input"
	"github.com/Bananenpro/tmpl/templates"
)

func CreateC() {
	compiler := "cc"
	if !external.IsInstalled("cc") {
		if external.IsInstalled("gcc") {
			compiler = "gcc"
		} else {
			if external.IsInstalled("clang") {
				compiler = "clang"
			} else {
				abort("No C compiler installed.")
			}
		}
	}

	availableTemplates := []string{"Empty"}

	if external.IsInstalled("make") {
		availableTemplates = append(availableTemplates, "Makefile")
	}

	if external.IsInstalled("cmake") {
		availableTemplates = append(availableTemplates, "CMake")
	}

	selectedTemplate, cancel := input.Select("Choose a template", availableTemplates)
	if cancel {
		stop()
	}
	selectedTemplate = strings.ToLower(selectedTemplate)

	err := os.Mkdir("src", 0755)
	if err != nil {
		abort(fmt.Sprint("Error while creating directory 'src': ", err))
	}

	CreateFile("src/main.c", templates.CHelloWorld)

	dir, err := os.Getwd()
	if err != nil {
		abort(fmt.Sprint("Failed to get current directory: ", err))
	}
	name := path.Base(dir)

	switch selectedTemplate {
	case "makefile":
		CreateFile("Makefile", strings.ReplaceAll(fmt.Sprintf(templates.CMakefile, compiler, name), "§", "%"))
	case "cmake":
		CreateFile("CMakeLists.txt", fmt.Sprintf(templates.CCMakeLists, name))
	}
}
