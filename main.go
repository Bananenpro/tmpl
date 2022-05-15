package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Bananenpro/tmpl/create"
	"github.com/Bananenpro/tmpl/input"
	"github.com/ogier/pflag"
)

func newProject() {
	var template string

	if len(pflag.Args()) < 2 {
		var cancel bool
		template, cancel = input.Select("Choose a language", []string{
			"C",
			"C#",
			"Go",
			"Python",
			"Rust",
		})
		if cancel {
			os.Exit(0)
		}
	} else {
		template = pflag.Arg(1)
	}

	create.Create(strings.ToLower(template))
}

func main() {
	pflag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s <operation> [...]\n", os.Args[0])
		fmt.Println("\nOperations:")
		fmt.Println("\tnew\t\tCreate a new project in the current directory (default)")
		pflag.PrintDefaults()
	}
	pflag.Parse()

	operation := strings.ToLower(pflag.Arg(0))

	if len(pflag.Args()) == 0 {
		operation = "new"
	}

	switch operation {
	case "new":
		newProject()
	default:
		fmt.Println("Unknown command:", strings.ToLower(pflag.Arg(0)))
		pflag.Usage()
		os.Exit(1)
	}

	fmt.Println("Done.")
}
