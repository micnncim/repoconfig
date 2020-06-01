package spinner

import (
	"io"
	"os"

	"github.com/caarlos0/spin"
)

type Spinner struct {
	s *spin.Spinner

	writer io.Writer
}

type Option func(*Spinner)

func WithWriter(w io.Writer) Option {
	return func(s *Spinner) { s.writer = w }
}

func New(opts ...Option) *Spinner {
	s := &Spinner{
		writer: os.Stdout,
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

func (s *Spinner) Start(msg string) {
	// Reset Spinner
	s.s = spin.New(
		"%s "+msg, // caarlos0/spin.Spinner.Start requires "%s" format as the pointer of indicator.
		spin.WithWriter(s.writer),
	) // reset spin.Spinner
	s.s.Start()
}

func (s *Spinner) Stop() {
	s.s.Stop()
}
