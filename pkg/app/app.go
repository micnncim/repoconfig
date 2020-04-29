package app

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/micnncim/repoconfig/pkg/github"
	"github.com/micnncim/repoconfig/pkg/http"
	"github.com/micnncim/repoconfig/pkg/log"
)

type app struct {
	writer    io.Writer
	logLevel  log.Level
	logFormat log.Format

	githubAPIBaseURL string
	githubToken      string

	hasIssues           bool
	hasProjects         bool
	hasWiki             bool
	defaultBranch       string
	allowSquashMerge    bool
	allowMergeCommit    bool
	allowRebaseMerge    bool
	deleteBranchOnMerge bool

	dryRun bool
	debug  bool
}

type repository struct {
	owner, repo string
}

func NewCommand() *cobra.Command {
	app := &app{
		writer:           os.Stdout,
		logLevel:         log.LevelInfo,
		logFormat:        log.FormatColorConsole,
		githubAPIBaseURL: github.APIBaseURL,
		githubToken:      os.Getenv("GITHUB_TOKEN"),
	}

	cmd := &cobra.Command{
		Use:   "repoconfig",
		Short: "CLI to update repository configs",
		RunE:  app.run,
	}

	cmd.Flags().BoolVar(&app.hasIssues, "has-issues", true, "Whether a repository has issues")
	cmd.Flags().BoolVar(&app.hasProjects, "has-projects", true, "Whether a repository has projects")
	cmd.Flags().BoolVar(&app.hasWiki, "has-wiki", true, "Whether a repository has wiki")
	cmd.Flags().StringVar(&app.defaultBranch, "default-branch", "master", "The default branch for a repository")
	cmd.Flags().BoolVar(&app.allowSquashMerge, "allow-squash-merge", true, "Whether to allow allow squash-merging pull requests")
	cmd.Flags().BoolVar(&app.allowMergeCommit, "allow-merge-commit", true, "Whether to allow merging pull requests with a merge commit")
	cmd.Flags().BoolVar(&app.allowRebaseMerge, "allow-rebase-merge", true, "Whether to allow rebase-merging pull requests")
	cmd.Flags().BoolVar(&app.deleteBranchOnMerge, "delete-branch-on-merge", false, "Whether to allow automatically deleting head branches when pull requests are merged")

	cmd.Flags().BoolVar(&app.dryRun, "dry-run", false, "Whether user enable dry-run mode")
	cmd.Flags().BoolVar(&app.debug, "debug", false, "Whether user enable debug mode")

	return cmd
}

func (a *app) run(_ *cobra.Command, args []string) error {
	repos, err := parseArgs(args)
	if err != nil {
		return err
	}

	if a.debug {
		a.logLevel = log.LevelDebug
	}

	logger, err := log.NewLogger(a.writer, a.logLevel, a.logFormat)
	if err != nil {
		return err
	}
	logger = logger.Named("app")

	httpClient, err := http.NewClient(a.githubAPIBaseURL, http.WithLogger(logger))
	if err != nil {
		return err
	}

	githubClient, err := github.NewClient(a.githubToken, httpClient, github.WithDryRun(a.dryRun), github.WithLogger(logger))
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	opts := &github.UpdateRepositoryOptions{
		HasIssues:           a.hasIssues,
		HasProjects:         a.hasProjects,
		HasWiki:             a.hasWiki,
		DefaultBranch:       a.defaultBranch,
		AllowSquashMerge:    a.allowSquashMerge,
		AllowMergeCommit:    a.allowMergeCommit,
		AllowRebaseMerge:    a.allowRebaseMerge,
		DeleteBranchOnMerge: a.deleteBranchOnMerge,
	}

	for _, repo := range repos {
		// In the case both owner and repo is specified.
		if repo.repo != "" {
			_ = githubClient.UpdateRepository(ctx, repo.owner, repo.repo, opts)
			// Not return error, just logging
			continue
		}

		// In the case only owner is specified.
		ghRepos, err := githubClient.ListRepositories(ctx, repo.owner)
		if err != nil {
			return err
		}
		for _, ghRepo := range ghRepos {
			// Not return error, just logging
			_ = githubClient.UpdateRepository(ctx, ghRepo.GetOwner().GetLogin(), ghRepo.GetName(), opts)
		}
	}

	return nil
}

func parseArgs(args []string) ([]repository, error) {
	if len(args) == 0 {
		return nil, errors.New("missing arguments")
	}

	repos := make([]repository, 0, len(args))

	for _, arg := range args {
		s := strings.Split(arg, "/")

		switch len(s) {
		case 1:
			repos = append(repos, repository{
				owner: s[0],
			})
		case 2:
			repos = append(repos, repository{
				owner: s[0],
				repo:  s[1],
			})
		default:
			return nil, fmt.Errorf("invalid argument: %q", arg)
		}
	}

	return repos, nil
}
