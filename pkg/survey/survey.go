package survey

import (
	"github.com/AlecAivazis/survey/v2"
)

type Surveyor interface {
	AskInput(message string) (string, error)
	AskSelect(message string, options []string) (string, error)
	AskMultiSelect(message string, options []string) ([]string, error)
}

type surveyor struct{}

// Guarantee *surveyor implements Surveyor.
var _ Surveyor = (*surveyor)(nil)

func NewSurveyor() *surveyor {
	return &surveyor{}
}

var BoolOptions = []string{"true", "false"}

func (s *surveyor) AskInput(message string) (string, error) {
	var str string

	if err := survey.AskOne(
		&survey.Input{
			Message: message,
		},
		&str,
	); err != nil {
		return "", err
	}

	return str, nil
}

func (s *surveyor) AskSelect(message string, options []string) (string, error) {
	var str string

	if err := survey.Ask(
		[]*survey.Question{
			{
				Prompt: &survey.Select{
					Message:  message,
					Options:  options,
					PageSize: len(options),
				},
			},
		},
		&str,
	); err != nil {
		return "", err
	}

	return str, nil
}

func (s *surveyor) AskMultiSelect(message string, options []string) ([]string, error) {
	var ss []string

	if err := survey.Ask(
		[]*survey.Question{
			{
				Prompt: &survey.MultiSelect{
					Message:  message,
					Options:  options,
					PageSize: len(options),
				},
			},
		},
		&ss,
	); err != nil {
		return nil, err
	}

	return ss, nil
}
