package create

import (
	"fmt"
	"os"

	"github.com/Bananenpro/tmpl/templates"
	"github.com/Bananenpro/tmpl/util"
)

func CreateRust() {
	if !util.IsInstalled("rustc") {
		fmt.Println("Rust is not installed.")
		os.Exit(1)
	}

	file, err := os.Create("main.rs")
	defer file.Close()
	if err != nil {
		Clean()
		fmt.Println("Error while creating file 'main.rs':", err)
		os.Exit(1)
	}

	_, err = file.WriteString(templates.RustHelloWorld)

	if err != nil {
		Clean()
		fmt.Println("Error while writing into file 'main.rs':", err)
		os.Exit(1)
	}
}
