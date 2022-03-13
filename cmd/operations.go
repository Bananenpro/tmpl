package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/Bananenpro/tmpl/create"
)

func new() {
	if len(flag.Args()) == 1 {
		fmt.Printf("Usage: %s new <template>\n", os.Args[0])
		fmt.Println("Available templates:")
		fmt.Println("\tgo")
		fmt.Println("\tpython")
		fmt.Println("\tc#")
		os.Exit(1)
	}

	create.Create(strings.ToLower(flag.Arg(1)))
}
