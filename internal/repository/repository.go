package repository

import "organization-scanner/internal/github"

type Service interface {
	ListRepositories(organization *string) ([]*github.Repository, error)
}
