package app

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/micnncim/repoconfig/pkg/github"
	"github.com/micnncim/repoconfig/pkg/http"
	"github.com/micnncim/repoconfig/pkg/logging"
	"github.com/micnncim/repoconfig/pkg/spinner"
	"github.com/micnncim/repoconfig/pkg/survey"
)

type app struct {
	githubClient github.Client
	spinner      *spinner.Spinner
}

type repository struct {
	owner, repo string
}

func NewCommand() (*cobra.Command, error) {
	logger, err := logging.NewLogger(os.Stderr, logging.LevelInfo, logging.FormatColorConsole)
	if err != nil {
		return nil, err
	}
	logger = logger.Named("app")

	httpClient, err := http.NewClient(
		github.APIBaseURL,
		http.WithLogger(logger),
	)
	if err != nil {
		return nil, err
	}

	githubClient, err := github.NewClient(
		os.Getenv("GITHUB_TOKEN"),
		httpClient,
		github.WithLogger(logger),
	)
	if err != nil {
		return nil, err
	}

	app := &app{
		githubClient: githubClient,
		spinner:      spinner.New(),
	}

	cmd := &cobra.Command{
		Use:   "repoconfig <OWNER> <REPO>",
		Short: "CLI to update repository config",
		RunE:  app.run,
	}

	return cmd, nil
}

func (a *app) run(_ *cobra.Command, args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("invalid arguments: %q", args)
	}
	owner, repo := args[0], args[1]

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	repository, err := a.getRepository(ctx, owner, repo)
	if err != nil {
		return err
	}

	input, err := askUpdateRepositoryInput(survey.NewSurveyor(), repository)
	switch err {
	case nil:
	case ErrRepositoryNoChange:
		warnf("\n🤖 %s/%s has not been changed\n", owner, repo)
		return nil
	default:
		return err
	}

	if err := a.githubClient.UpdateRepository(ctx, owner, repo, input); err != nil {
		return err
	}

	infof("\n🚀 https://github.com/%s/%s has been updated\n", owner, repo)

	return nil
}

func (a *app) getRepository(ctx context.Context, owner, repo string) (*github.Repository, error) {
	a.spinner.Start(color.CyanString("🤖 fetching %s/%s...", owner, repo))
	defer a.spinner.Stop()

	return a.githubClient.GetRepository(ctx, owner, repo)
}
