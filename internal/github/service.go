package github

import (
	"context"
	"github.com/google/go-github/v32/github"
)

type Service interface {
	Authenticate() *github.Client
	GetRepositoriesByOrganization(ctx context.Context, organization string, page, elementsPerPage int) ([]*Repository, error)
}

type githubService struct {
}

func NewGitHubService() Service {
	return githubService{}
}


type Repository struct {
	URL *string `json:"URL,omitempty"`
}
