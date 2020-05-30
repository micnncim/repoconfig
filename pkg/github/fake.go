package github

import (
	"context"
	"fmt"
	"sync"
)

// FakeClient implements Client and returns fake objects.
type FakeClient struct {
	Repos map[string]*Repository // key: repo name, value: repo options

	mu sync.RWMutex
}

// Guarantee *FakeClient implements Client.
var _ Client = (*FakeClient)(nil)

func NewFakeClient() *FakeClient {
	return &FakeClient{
		Repos: make(map[string]*Repository),
		mu:    sync.RWMutex{},
	}
}

func (c *FakeClient) GetRepository(ctx context.Context, owner, repo string) (*Repository, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

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
	c.mu.Lock()
	defer c.mu.Unlock()

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
