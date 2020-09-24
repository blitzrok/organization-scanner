package main

import (
	"flag"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"organization-scanner/internal/handler"
)

func init() {
	if err := godotenv.Load(); err != nil {
		logrus.Error("No .env file found")
	}
}

func main() {
	dfltOrg := "some-org"
	dfltRepoURL := "git@github.com/some-user/some-repository.git"

	org := flag.String("organization", dfltOrg, "Organization name. Ex: acmecompany")
	repoURL := flag.String("repository-url", dfltRepoURL, "SSH url of the repository. Ex: " + dfltRepoURL)
	flag.Parse()

	if *org != dfltOrg {
		handler.NewGitHubRepositoryScanner().ScanRepositoriesFromOrganization(org)
	} else if *repoURL != dfltRepoURL {
		handler.NewGitHubRepositoryScanner().ScanRepository(repoURL)
	} else {
		println("No matching option. Execute go run . -h to get all available options")
	}
}
