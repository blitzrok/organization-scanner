package repository

type Service interface {
	ScanRepositoriesFromOrganization(organization *string) error
}
