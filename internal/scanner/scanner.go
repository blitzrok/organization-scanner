package scanner

import (
	"encoding/csv"
	"fmt"
	"github.com/sirupsen/logrus"
	gitleaks "github.com/zricethezav/gitleaks/src"
	"organization-scanner/internal/github"
	"os"
)

type ScannerService interface {
	ScanRepositories(repositories []*github.Repository, outputFile *string)
}

type scannerService struct {}

func NewScanService() ScannerService{
	return scannerService{}
}


func (s scannerService) ScanRepositories(repositories []*github.Repository, outputFile *string) {
	var leaks []gitleaks.Leak
	for index, repo := range repositories {
		leaks = append(leaks, scan(*repo.URL)...)
		if index == 8 {
			break
		}
	}

	exportToCSV(leaks)
}

func scan(repoURL string) []gitleaks.Leak {
	logrus.Info("Scanning repo ", repoURL)
	opt := &gitleaks.Options{
		Repo:           repoURL,
		ExcludeForks:   true,
		Entropy:        8.0,
		Log:            logrus.InfoLevel.String(),
		Verbose:        true,
		SampleConfig:   true,
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

func exportToCSV(leaks []gitleaks.Leak) {
	csvFile, err := os.Create("./leaks-report.csv")
	if err != nil {
		logrus.Error("Error creating CSV file", err)
	}
	defer csvFile.Close()
	// TODO write headers
	writer := csv.NewWriter(csvFile)
	for _, leak := range leaks {
		var row []string
		row = append(row, leak.Repo)
		row = append(row, leak.Message)
		row = append(row, leak.Offender)
		row = append(row, leak.Author)
		row = append(row, leak.Type)
		row = append(row, leak.Commit)
		row = append(row, leak.Email)
		row = append(row, leak.File)
		row = append(row, leak.Line)
		row = append(row, leak.Date.String())
		writer.Write(row)
	}

	writer.Flush()
}
