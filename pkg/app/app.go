package app

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/micnncim/repoconfig/pkg/github"
	"github.com/micnncim/repoconfig/pkg/http"
	"github.com/micnncim/repoconfig/pkg/logging"
)

type app struct {
	writer    io.Writer
	logLevel  logging.Level
	logFormat logging.Format

	githubAPIBaseURL string
	githubToken      string

	dryRun bool
	debug  bool
}

type repository struct {
	owner, repo string
}

func NewCommand() *cobra.Command {
	app := &app{
		writer:           os.Stderr,
		logLevel:         logging.LevelInfo,
		logFormat:        logging.FormatColorConsole,
		githubAPIBaseURL: github.APIBaseURL,
		githubToken:      os.Getenv("GITHUB_TOKEN"),
	}

	cmd := &cobra.Command{
		Use:   "repoconfig <OWNER> <REPO>",
		Short: "CLI to update repository config",
		RunE:  app.run,
	}

	cmd.Flags().BoolVar(&app.dryRun, "dry-run", false, "Whether user enable dry-run mode")
	cmd.Flags().BoolVar(&app.debug, "debug", false, "Whether user enable debug mode")

	return cmd
}

func (a *app) run(_ *cobra.Command, args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("invalid arguments: %q", args)
	}
	owner, repo := args[0], args[1]

	if a.debug {
		a.logLevel = logging.LevelDebug
	}

	logger, err := logging.NewLogger(a.writer, a.logLevel, a.logFormat)
	if err != nil {
		return err
	}
	logger = logger.Named("app")

	httpClient, err := http.NewClient(
		a.githubAPIBaseURL,
		http.WithLogger(logger),
	)
	if err != nil {
		return err
	}

	githubClient, err := github.NewClient(
		a.githubToken,
		httpClient,
		github.WithDryRun(a.dryRun),
		github.WithLogger(logger),
	)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	repository, err := githubClient.GetRepository(ctx, owner, repo)
	if err != nil {
		return err
	}

	input, err := askUpdateRepositoryInput(repository)
	if err != nil {
		return err
	}

	if err := githubClient.UpdateRepository(ctx, owner, repo, input); err != nil {
		return err
	}

	infof("🚀 %s has been updated\n", repository.GetHTMLURL())

	return nil
}

func infof(format string, a ...interface{}) {
	color.New(color.FgBlue).Printf(format, a...)
}
