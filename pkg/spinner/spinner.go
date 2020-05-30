package spinner

import (
	"fmt"
	"io"
	"time"

	spin "github.com/briandowns/spinner"
)

type Spinner interface {
	Start(msg string)
	Stop(msg string)
}

type spinner struct {
	s *spin.Spinner
}

// Guarantee *spinner implements Spinner.
var _ Spinner = (*spinner)(nil)

func New(w io.Writer) Spinner {
	return &spinner{
		s: spin.New(
			spin.CharSets[11],
			100*time.Millisecond,
			spin.WithWriter(w),
		),
	}
}

func (s *spinner) Start(format string) {
	s.s.Suffix = fmt.Sprintf(format)
	s.s.Start()
}

func (s *spinner) Stop(format string) {
	s.s.FinalMSG = fmt.Sprintf(format)
	s.s.Stop()

	// reset
	s.s.Suffix = ""
	s.s.FinalMSG = ""
}
