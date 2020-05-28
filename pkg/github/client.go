package github

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/google/go-github/v31/github"
	"go.uber.org/zap"
	"golang.org/x/oauth2"

	pkghttp "github.com/micnncim/repoconfig/pkg/http"
)

const APIBaseURL = "https://api.github.com"

type Client interface {
	GetRepository(ctx context.Context, owner, repo string) (*github.Repository, error)
	UpdateRepository(ctx context.Context, owner, repo string, input *UpdateRepositoryInput) error
}

type client struct {
	githubClient *github.Client
	httpClient   *pkghttp.Client

	githubToken string

	dryRun bool

	logger *zap.Logger
}

// Guarantee *client implements Client.
var _ Client = (*client)(nil)

type Option func(*client)

func WithDryRun(v bool) Option {
	return func(c *client) { c.dryRun = v }
}

func WithLogger(l *zap.Logger) Option {
	return func(c *client) { c.logger = l.Named("github") }
}

func NewClient(token string, httpClient *pkghttp.Client, opts ...Option) (Client, error) {
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

func (c *client) GetRepository(ctx context.Context, owner, repo string) (*github.Repository, error) {
	logger := c.logger.With(
		zap.String("owner", owner),
		zap.String("repo", repo),
	)

	repository, _, err := c.githubClient.Repositories.Get(ctx, owner, repo)
	if err != nil {
		logger.Error("failed to get repository", zap.Error(err))
		return nil, err
	}

	logger.Debug("successfully fetched repository", zap.Any("repository", repository))
	return repository, nil
}

func (c *client) UpdateRepository(ctx context.Context, owner, repo string, input *UpdateRepositoryInput) error {
	logger := c.logger.With(
		zap.String("owner", owner),
		zap.String("repo", repo),
		zap.Any("update_repository_input", input),
		zap.Bool("dry_run", c.dryRun),
	)

	if !c.dryRun {
		if _, err := c.httpClient.DoRequest(
			ctx,
			http.MethodPatch,
			fmt.Sprintf("repos/%s/%s", owner, repo),
			map[string]string{
				"Content-Type":  "application/json",
				"Authorization": fmt.Sprintf("token %s", c.githubToken),
			},
			input,
		); err != nil {
			logger.Error("failed to update repository", zap.Error(err))
			return err
		}
	}

	logger.Debug("successfully updated repository")
	return nil
}
