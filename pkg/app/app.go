package app

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"

	pkggithub "github.com/micnncim/repoconfig/pkg/github"
	"github.com/micnncim/repoconfig/pkg/http"
	"github.com/micnncim/repoconfig/pkg/logging"
	"github.com/micnncim/repoconfig/pkg/survey"
)

type app struct {
	githubClient pkggithub.Client
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
		pkggithub.APIBaseURL,
		http.WithLogger(logger),
	)
	if err != nil {
		return nil, err
	}

	githubClient, err := pkggithub.NewClient(
		os.Getenv("GITHUB_TOKEN"),
		httpClient,
		pkggithub.WithLogger(logger),
	)
	if err != nil {
		return nil, err
	}

	app := &app{
		githubClient: githubClient,
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

	repository, err := a.githubClient.GetRepository(ctx, owner, repo)
	if err != nil {
		return err
	}

	input, err := askUpdateRepositoryInput(survey.NewSurveyor(), repository)
	if err != nil {
		return err
	}

	if err := a.githubClient.UpdateRepository(ctx, owner, repo, input); err != nil {
		return err
	}

	infof("\n🚀 %s has been updated\n", repository.GetHTMLURL())

	return nil
}
