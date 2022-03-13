package create

import (
	"github.com/Bananenpro/tmpl/templates"
	"github.com/Bananenpro/tmpl/util"
)

func CreateRust() {
	if !util.IsInstalled("rustc") {
		abort("Rust is not installed.")
	}

	CreateFile("main.rs", templates.RustHelloWorld)
}
