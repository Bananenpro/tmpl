package input

import (
	"github.com/AlecAivazis/survey/v2"
)

func Input(question string) (response string, cancel bool) {
	var result string
	err := survey.AskOne(&survey.Input{
		Message: question,
	}, &result, survey.WithValidator(survey.Required))
	return result, err != nil
}

func YesNo(question string, defaultValue bool) (yes bool, cancel bool) {
	err := survey.AskOne(&survey.Confirm{
		Message: question,
		Default: defaultValue,
	}, &yes, survey.WithValidator(survey.Required))
	return yes, err != nil
}

func Select(msg string, options []string) (selected string, cancel bool) {
	var result string
	err := survey.AskOne(&survey.Select{
		Message: msg,
		Options: options,
	}, &result, survey.WithValidator(survey.Required))
	return result, err != nil
}
