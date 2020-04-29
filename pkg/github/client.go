package github

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/go-github/v31/github"
	"go.uber.org/zap"
	"golang.org/x/oauth2"

	"github.com/micnncim/repoconfig/pkg/http"
)

const APIBaseURL = "https://api.github.com"

type Client interface {
	ListRepositories(ctx context.Context, owner string) ([]*github.Repository, error)
	UpdateRepository(ctx context.Context, owner, repo string, opts *UpdateRepositoryOptions) error
}

type client struct {
	githubClient *github.Client
	httpClient   *http.Client

	githubToken string

	dryRun bool

	logger *zap.Logger
}

type Option func(*client)

func WithDryRun(v bool) Option {
	return func(c *client) { c.dryRun = v }
}

func WithLogger(l *zap.Logger) Option {
	return func(c *client) { c.logger = l.Named("github") }
}

func NewClient(token string, httpClient *http.Client, opts ...Option) (Client, error) {
	if token == "" {
		return nil, errors.New("missing github token")
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	c := &client{
		githubClient: github.NewClient(tc),
		httpClient:   httpClient,
		githubToken:  token,
		logger:       zap.NewNop(),
	}
	for _, opt := range opts {
		opt(c)
	}
	return c, nil
}

func (c *client) ListRepositories(ctx context.Context, owner string) ([]*github.Repository, error) {
	logger := c.logger.With(zap.String("owner", owner))

	logger.Debug("started fetching repositories")
	defer logger.Debug("finished fetching repositories")

	opts := github.ListOptions{
		PerPage: 100,
	}

	var allRepos []*github.Repository

	for {
		repos, resp, err := c.githubClient.Repositories.List(ctx, owner, &github.RepositoryListOptions{
			Affiliation: "owner",
			ListOptions: opts,
		})
		if err != nil {
			c.logger.Error("failed to fetch repositories", zap.Error(err))
			return nil, err
		}

		for _, repo := range repos {
			c.logger.Debug("fetched repository",
				zap.String("repo", repo.GetName()),
			)
		}

		allRepos = append(allRepos, repos...)

		if resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}

	return allRepos, nil
}

func (c *client) UpdateRepository(ctx context.Context, owner, repo string, opts *UpdateRepositoryOptions) error {
	logger := c.logger.With(
		zap.String("owner", owner),
		zap.String("repo", repo),
		zap.Any("update_repository_options", opts),
		zap.Bool("dry_run", c.dryRun),
	)

	if !c.dryRun {
		if _, err := c.httpClient.DoRequest(
			ctx,
			"PATCH",
			fmt.Sprintf("repos/%s/%s", owner, repo),
			map[string]string{
				"Content-Type":  "application/json",
				"Authorization": fmt.Sprintf("token %s", c.githubToken),
			},
			opts,
		); err != nil {
			logger.Error("failed to update repository", zap.Error(err))
			return err
		}
	}

	logger.Info("successfully updated repository")
	return nil
}
