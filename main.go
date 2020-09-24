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
	org := flag.String("organization", "some-org", "Organization name")
	flag.Parse()
	handler.NewGitHubRepositoryScanner().ScanRepositoriesFromOrganization(org)
}
