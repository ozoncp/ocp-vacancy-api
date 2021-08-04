package repo

import "github.com/ozoncp/ocp-vacancy-api/internal/models"

// Repo is a repository interface for Vacancy entity
type Repo interface {
	AddVacancies(vacancies []models.Vacancy) error
	ListVacancies(limit, offset uint64) ([]models.Vacancy, error)
	DescribeVacancy(vacancyId uint64) (*models.Vacancy, error)
}
