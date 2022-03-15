package create

import (
	"fmt"

	"github.com/Bananenpro/tmpl/external"
	"github.com/Bananenpro/tmpl/input"
)

func CreateRust() {
	if !external.IsInstalled("cargo") {
		abort("Cargo is not installed.")
	}

	projectType := input.Select("Choose a project type", []string{"binary", "library"}, 0)

	flag := "--bin"

	if projectType == "library" {
		flag = "--lib"
	}

	out, err := external.Execute("cargo", "init", flag, "--vcs", "none")
	if err != nil {
		abort(fmt.Sprint("Failed to create project: ", out))
	}
}
