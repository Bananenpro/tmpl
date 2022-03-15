package create

import (
	"fmt"

	"github.com/Bananenpro/tmpl/util"
)

func CreateRust() {
	if !util.IsInstalled("cargo") {
		abort("Cargo is not installed.")
	}

	projectType := util.Select("Choose a project type", []string{"binary", "library"}, 0)

	flag := "--bin"

	if projectType == "library" {
		flag = "--lib"
	}

	out, err := util.Execute("cargo", "init", flag, "--vcs", "none")
	if err != nil {
		abort(fmt.Sprint("Failed to create project: ", out))
	}
}
