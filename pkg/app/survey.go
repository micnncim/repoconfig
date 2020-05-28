package app

import "github.com/AlecAivazis/survey/v2"

var surveyBoolOptions = []string{"true", "false"}

func askInput(message string) (string, error) {
	var s string

	if err := survey.AskOne(
		&survey.Input{
			Message: message,
		},
		&s,
	); err != nil {
		return "", err
	}

	return s, nil
}

func askSelect(message string, options []string) (string, error) {
	var s string

	if err := survey.Ask([]*survey.Question{
		{
			Prompt: &survey.Select{
				Message:  message,
				Options:  options,
				PageSize: len(options),
			},
		},
	}, &s); err != nil {
		return "", err
	}

	return s, nil
}

func askMultiSelect(message string, options []string) ([]string, error) {
	var s []string

	if err := survey.Ask([]*survey.Question{
		{
			Prompt: &survey.MultiSelect{
				Message:  message,
				Options:  options,
				PageSize: len(options),
			},
		},
	}, &s); err != nil {
		return nil, err
	}

	return s, nil
}
