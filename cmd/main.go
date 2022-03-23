package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/Bananenpro/tmpl/create"
	"github.com/Bananenpro/tmpl/input"
)

func printHelp() {
	fmt.Printf("Usage: %s <operation> [...]\n", os.Args[0])

	fmt.Println()

	fmt.Println("Flags:")
	fmt.Println("\t-h, --help\tDisplay help")

	fmt.Println()

	fmt.Println("Operations:")
	fmt.Println("\tnew\t\tCreate a new project in the current directory")
}

func newProject() {
	var template string

	if len(flag.Args()) == 1 {
		template = input.Select("Choose a language", []string{
			"C",
			"C#",
			"Go",
			"Python",
			"Rust",
		}, -1)
	} else {
		template = flag.Arg(1)
	}

	create.Create(strings.ToLower(template))
}

func main() {
	var help bool
	flag.BoolVar(&help, "h", false, "Display help")
	flag.BoolVar(&help, "help", false, "Display help")

	flag.Parse()

	if help {
		printHelp()
		return
	}

	if len(flag.Args()) == 0 {
		fmt.Printf("Usage: %s <operation> [...]\n", os.Args[0])
		fmt.Println("Use -h for help.")
		os.Exit(1)
	}

	switch strings.ToLower(flag.Arg(0)) {
	case "new":
		newProject()
	default:
		fmt.Println("Unknown command:", strings.ToLower(flag.Arg(0)))
		fmt.Printf("Usage: %s <operation> [...]\n", os.Args[0])
		fmt.Println("Use -h for help.")
		os.Exit(1)
	}

	fmt.Println("Done.")
}
