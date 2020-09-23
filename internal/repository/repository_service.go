package repository

import (
	"context"
	"organization-scanner/internal/github"
)

type repositoryService struct {
	githubService github.Service
}

func NewRepositoryService(githubService github.Service) Service {
	return &repositoryService{githubService: githubService}
}

func (r repositoryService) ListRepositories(organization *string) ([]*github.Repository, error) {
	ctx := context.Background()
	page := 1
	resultsPerPage := 100
	hasNext := true
	var repositories []*github.Repository

	for hasNext {
		result, err := r.githubService.GetRepositoriesByOrganization(ctx, *organization, page, resultsPerPage)
		if err != nil {
			return nil, err
		}

		page++
		hasNext = len(result) == resultsPerPage
		repositories = append(repositories, result...)
	}

	return repositories, nil
}

