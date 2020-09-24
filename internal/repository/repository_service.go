package repository

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"organization-scanner/internal/github"
	"organization-scanner/internal/scanner"
	"time"
)

type repositoryService struct {
	githubService  github.Service
	scannerService scanner.ScannerService
}

func NewRepositoryService(githubService github.Service, scannerService scanner.ScannerService) Service {
	return &repositoryService{githubService: githubService, scannerService: scannerService}
}

func (r repositoryService) ScanRepositoriesFromOrganization(organization *string) error {
	logrus.Info("List repositories for organization ", *organization)
	ctx := context.Background()
	page := 1
	resultsPerPage := 100
	hasNext := true
	var repositories []*github.Repository

	for hasNext {
		result, err := r.githubService.GetRepositoriesByOrganization(ctx, *organization, page, resultsPerPage)
		if err != nil {
			logrus.Error("Error retrieving repository list", err)
			return err
		}

		page++
		hasNext = len(result) == resultsPerPage
		repositories = append(repositories, result...)
	}
	infoMessage := fmt.Sprintf("Found %v repositories for Organization %s. Proceeding to scan.", len(repositories), *organization)
	logrus.Info(infoMessage)
	outputFile := fmt.Sprintf("%s-leaks-report.csv", time.Now().String())
	r.scannerService.ScanRepositories(repositories, &outputFile)
	return nil
}

func (r repositoryService) ScanRepository(repoURL *string) error {
	var repositories []*github.Repository
	repositories = append(repositories, &github.Repository{URL: repoURL})
	outputFile := fmt.Sprintf("%s-leaks-report.csv", time.Now().String())

	r.scannerService.ScanRepositories(repositories, &outputFile)
	return nil
}
