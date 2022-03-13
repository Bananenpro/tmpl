package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
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
		new()
	}

	fmt.Println("Done.")
}
