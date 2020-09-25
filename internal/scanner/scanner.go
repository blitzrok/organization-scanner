package scanner

import (
	"fmt"
	"github.com/sirupsen/logrus"
	gitleaks "github.com/zricethezav/gitleaks/src"
	"organization-scanner/internal/exporter"
	"organization-scanner/internal/github"
)

type ScannerService interface {
	ScanRepositories(repositories []*github.Repository, outputFile *string)
}

type scannerService struct{}

func NewScanService() ScannerService {
	return scannerService{}
}

func (s scannerService) ScanRepositories(repositories []*github.Repository, outputFile *string) {
	var leaks []gitleaks.Leak
	for _, repo := range repositories {
		leaks = append(leaks, scan(*repo.URL)...)
	}

	if len(leaks) > 0 {
		exporter.LeaksToCSV(leaks, *outputFile)
	}
}

func scan(repoURL string) []gitleaks.Leak {
	logrus.Info("Scanning repo ", repoURL)
	gitleaksConfigFile := "./scan-config.toml"
	opt := &gitleaks.Options{
		Repo:         repoURL,
		Entropy: 8.0,
		Log:          logrus.InfoLevel.String(),
		Verbose:      true,
		ConfigPath:   gitleaksConfigFile,
	}

	res, err := gitleaks.Run(opt)
	if err != nil {
		logrus.Error("Error scanning repo: ", err)
		return nil
	}

	finishMessage := fmt.Sprintf("Found %v leaks for repository %s", len(res.Leaks), repoURL)
	logrus.Info(finishMessage)
	return res.Leaks
}
