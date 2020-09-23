package repository

type Repository struct {
	URL *string `json:"URL,omitempty"`
}

type Service interface {
	ListRepositories() ([]*Repository, error)
}
