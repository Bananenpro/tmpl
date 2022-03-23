package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Input(question string) string {
	fmt.Printf("%s ", question)
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		return strings.TrimSpace(scanner.Text())
	}

	return ""
}

func YesNo(question string, defaultValue bool) bool {
	q := fmt.Sprintf("%s\t[y/N]", question)
	if defaultValue {
		q = fmt.Sprintf("%s\t[Y/n]", question)
	}

	answer := Input(q)
	if answer == "" {
		return defaultValue
	}

	return strings.ToLower(answer) == "y"
}

func Select(msg string, options []string, defaultIndex int) string {
	prompt := fmt.Sprintf("%s [%d]: ", msg, defaultIndex)
	if defaultIndex < 0 || defaultIndex >= len(options) {
		prompt = fmt.Sprintf("%s: ", msg)
	}

	for i, o := range options {
		fmt.Printf("%d) %s  ", i, o)
	}
	fmt.Println()
	index := -1

	var err error
	for index < 0 || index >= len(options) {
		selection := Input(prompt)
		index, err = strconv.Atoi(selection)
		if err != nil {
			index = defaultIndex
		}
	}

	return options[index]
}
