package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/v31/github"
)

// FakeClient implements Client and returns fake objects.
// This doesn't guarantee thread-safety.
type FakeClient struct {
	Repos map[string]*FakeRepository // key: repo name, value: repo options
}

// FakeRepository represents a subset of github.Repository for testing.
type FakeRepository struct {
	Name        string
	Description string
	Private     bool
}

// Guarantee *FakeClient implements Client.
var _ Client = (*FakeClient)(nil)

func (c *FakeClient) GetRepository(ctx context.Context, owner, repo string) (*github.Repository, error) {
	r, ok := c.Repos[repo]
	if !ok {
		return nil, fmt.Errorf("repo %q not found", repo)
	}
	return &github.Repository{
		Owner:       &github.User{Login: github.String(owner)},
		Name:        github.String(repo),
		Description: github.String(r.Description),
		Private:     github.Bool(r.Private),
	}, nil
}

func (c *FakeClient) UpdateRepository(ctx context.Context, owner, repo string, input *UpdateRepositoryInput) error {
	if _, ok := c.Repos[repo]; !ok {
		return fmt.Errorf("repo %q not found", repo)
	}
	c.Repos[repo] = &FakeRepository{
		Description: input.Description,
		Private:     input.Private,
	}
	return nil
}
