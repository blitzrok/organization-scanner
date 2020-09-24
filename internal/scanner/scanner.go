package scanner

import (
	"fmt"
	"github.com/sirupsen/logrus"
	gitleaks "github.com/zricethezav/gitleaks/src"
)

func Scan(repoURL string) {
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
