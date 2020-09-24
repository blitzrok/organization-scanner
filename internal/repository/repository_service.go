package repository

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"organization-scanner/internal/github"
	"organization-scanner/internal/scanner"
)

type repositoryService struct {
	githubService github.Service
}

func NewRepositoryService(githubService github.Service) Service {
	return &repositoryService{githubService: githubService}
}

func (r repositoryService) ListRepositories(organization *string) ([]*github.Repository, error) {
	logrus.Info("List repositories for organization ", organization)
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
	infoMessage := fmt.Sprintf("Found %s repositories for Organization %s. Proceeding to scan.", len(repositories), organization)
	logrus.Info(infoMessage)
	for _, s := range repositories {
		scanner.Scan(*s.URL)
	}

	return repositories, nil
}

