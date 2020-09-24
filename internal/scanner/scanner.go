package scanner

import (
	"fmt"
	"github.com/sirupsen/logrus"
	gitleaks "github.com/zricethezav/gitleaks/src"
	"organization-scanner/internal/github"
)

type ScannerService interface {
	ScanRepositories(repositories []*github.Repository, outputFile *string)
}

type scannerService struct {}

func NewScanService() ScannerService{
	return scannerService{}
}


func (s scannerService) ScanRepositories(repositories []*github.Repository, outputFile *string) {
	for _, s := range repositories {
		scan(*s.URL)
	}
}

func scan(repoURL string) {
	logrus.Info("Scanning repo ", repoURL)
	opt := &gitleaks.Options{
		Repo:           repoURL,
		ExcludeForks:   true,
		Entropy:        8.0,
		Log:            logrus.InfoLevel.String(),
		Verbose:        true,
		Report:         "report.csv",
		SampleConfig:   true,
	}

	res, err := gitleaks.Run(opt)
	if err != nil {
		logrus.Error("Error scanning repo: ", err)
		return
	}
	finishMessage := fmt.Sprintf("Found %s leaks for repository %s", len(res.Leaks), repoURL)
	logrus.Info(finishMessage)
}
