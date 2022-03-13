package create

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/Bananenpro/tmpl/templates"
	"github.com/Bananenpro/tmpl/util"
)

func CreateC() {
	compiler := "cc"
	if !util.IsInstalled("cc") {
		if util.IsInstalled("gcc") {
			compiler = "gcc"
		} else {
			if util.IsInstalled("clang") {
				compiler = "clang"
			} else {
				abort("No C compiler installed.")
			}
		}
	}

	availableTemplates := []string{"empty"}

	if util.IsInstalled("make") {
		availableTemplates = append(availableTemplates, "makefile")
	}

	if util.IsInstalled("cmake") {
		availableTemplates = append(availableTemplates, "cmake")
	}

	selectedTemplate := util.Select("Choose a template", availableTemplates, 0)

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
