package github

import (
	"context"
	"fmt"
)

// FakeClient implements Client and returns fake objects.
// This doesn't guarantee thread-safety.
type FakeClient struct {
	Repos map[string]*Repository // key: repo name, value: repo options
}

// Guarantee *FakeClient implements Client.
var _ Client = (*FakeClient)(nil)

func (c *FakeClient) GetRepository(ctx context.Context, owner, repo string) (*Repository, error) {
	r, ok := c.Repos[repo]
	if !ok {
		return nil, fmt.Errorf("repo %q not found", repo)
	}
	return &Repository{
		Name:        repo,
		Description: r.Description,
		Private:     r.Private,
	}, nil
}

func (c *FakeClient) UpdateRepository(ctx context.Context, owner, repo string, repository *Repository) error {
	if _, ok := c.Repos[repo]; !ok {
		return fmt.Errorf("repo %q not found", repo)
	}
	c.Repos[repo] = &Repository{
		Name:        repository.Name,
		Description: repository.Description,
		Private:     repository.Private,
	}
	return nil
}
