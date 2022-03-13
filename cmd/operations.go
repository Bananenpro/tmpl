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
		fmt.Printf("Type '%s list' to get a list of supported templates.\n", os.Args[0])
	}

	create.Create(strings.ToLower(flag.Arg(1)))
}
