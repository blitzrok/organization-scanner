package github

import (
	"context"
	"github.com/google/go-github/v32/github"
)

func (g githubService) GetRepositoriesByOrganization(ctx context.Context, organization string, page, elementsPerPage int) ([]*Repository, error) {
	user := g.Authenticate()
	searchParams := &github.RepositoryListByOrgOptions{
		Type: "all",
		ListOptions: github.ListOptions{
			Page:    page,
			PerPage: elementsPerPage,
		},
	}
	result, _, err := user.Repositories.ListByOrg(ctx, organization, searchParams)
	if err != nil {
		return nil, err
	}

	var res []*Repository
	for _, s := range result {
		res = append(res, &Repository{URL: s.SSHURL})
	}
	return res, nil
}
