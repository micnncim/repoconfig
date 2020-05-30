package github

import (
	"context"
	"fmt"
	"sync"
)

// FakeClient implements Client and returns fake objects.
type FakeClient struct {
	repos map[string]*Repository // key: repo name, value: repo options

	mu sync.RWMutex
}

// Guarantee *FakeClient implements Client.
var _ Client = (*FakeClient)(nil)

func NewFakeClient() *FakeClient {
	return &FakeClient{
		repos: make(map[string]*Repository),
		mu:    sync.RWMutex{},
	}
}

// GetRepository implements Client.
func (c *FakeClient) GetRepository(ctx context.Context, _, repo string) (*Repository, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	r, ok := c.repos[repo]
	if !ok {
		return nil, fmt.Errorf("repo %q not found", repo)
	}

	return r, nil
}

// UpdateRepository implements Client.
func (c *FakeClient) UpdateRepository(ctx context.Context, _, repo string, repository *Repository) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.repos[repo]; !ok {
		return fmt.Errorf("repo %q not found", repo)
	}

	c.repos[repo] = repository
	return nil
}

// SetRepository is an utility method that sets repository for testing.
func (c *FakeClient) SetRepository(repo string, repository *Repository) {
	c.mu.Lock()
	c.repos[repo] = repository
	c.mu.Unlock()
}
