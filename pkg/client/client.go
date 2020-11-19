package client

import (
	"context"

	"github.com/google/go-github/v32/github"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
)

const (
	errEmptyToken = "no token provided"
)

// NewClient creates a new client.
func NewClient(token string) (*github.Client, error) {
	if token == "" {
		return nil, errors.New(errEmptyToken)
	}
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	return github.NewClient(tc), nil
}
