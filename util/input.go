package util

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Question(question string) string {
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

	answer := Question(q)
	if answer == "" {
		return defaultValue
	}

	return strings.ToLower(answer) == "y"
}

func Select(msg string, options []string, defaultIndex int) string {
	for i, o := range options {
		fmt.Printf("%d) %s  ", i, o)
	}
	fmt.Println()
	index := -1

	var err error
	for index < 0 || index >= len(options) {
		selection := Question(fmt.Sprintf("%s [%d]: ", msg, defaultIndex))
		index, err = strconv.Atoi(selection)
		if err != nil {
			index = defaultIndex
		}
	}

	return options[index]
}
