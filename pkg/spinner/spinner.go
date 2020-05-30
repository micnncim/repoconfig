package spinner

import (
	"io"

	"github.com/caarlos0/spin"
)

type Spinner interface {
	Start(msg string)
	Stop()
}

type spinner struct {
	s *spin.Spinner

	w io.Writer
}

// Guarantee *spinner implements Spinner.
var _ Spinner = (*spinner)(nil)

func New(w io.Writer) Spinner {
	return &spinner{
		w: w,
	}
}

func (s *spinner) Start(msg string) {
	// Reset Spinner
	s.s = spin.New(
		"%s "+msg, // caarlos0/spin.Spinner.Start requires "%s" format as the pointer of indicator.
		spin.WithWriter(s.w),
	) // reset spin.Spinner
	s.s.Start()
}

func (s *spinner) Stop() {
	s.s.Stop()
}
