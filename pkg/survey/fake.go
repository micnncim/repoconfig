package survey

import (
	"fmt"
)

// FakeSurveyor implements Surveyor and returns fake objects.
// This doesn't guarantee thread-safety.
type FakeSurveyor struct {
	AskInputMessages       map[string]string
	AskSelectMessages      map[string]string
	AskMultiSelectMessages map[string][]string
}

// Guarantee *FakeSurveyor implements Surveyor.
var _ Surveyor = (*FakeSurveyor)(nil)

func (s *FakeSurveyor) AskInput(message string) (string, error) {
	v, ok := s.AskInputMessages[message]
	if !ok {
		return "", fmt.Errorf("message for %q not found", message)
	}
	return v, nil
}

func (s *FakeSurveyor) AskSelect(message string, options []string) (string, error) {
	v, ok := s.AskSelectMessages[message]
	if !ok {
		return "", fmt.Errorf("option for %q not found", message)
	}
	return v, nil
}

func (s *FakeSurveyor) AskMultiSelect(message string, options []string) ([]string, error) {
	v, ok := s.AskMultiSelectMessages[message]
	if !ok {
		return nil, fmt.Errorf("options for %q not found", message)
	}
	return v, nil
}
