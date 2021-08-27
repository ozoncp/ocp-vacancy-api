package flusher

import (
	"context"

	"github.com/ozoncp/ocp-vacancy-api/internal/models"
	"github.com/ozoncp/ocp-vacancy-api/internal/repo"
	"github.com/ozoncp/ocp-vacancy-api/internal/utils"
)

// Flusher is an interface to flush Vacancies into Vacancies repo
type Flusher interface {
	Flush(ctx context.Context, vacancies []models.Vacancy) []models.Vacancy
}

// NewFlusher returns Flusher with batch flushing support
func NewFlusher(
	chunkSize int,
	vacancyRepo repo.Repo,
) Flusher {
	return &flusher{
		chunkSize:   chunkSize,
		vacancyRepo: vacancyRepo,
	}
}

type flusher struct {
	chunkSize   int
	vacancyRepo repo.Repo
}

// Flush recieves input slice of Vacancies and flushes them by batches of chunkSize into vacancyRepo.
// If a batch flush failed, entire batch is added to the result slice.
func (f *flusher) Flush(ctx context.Context, vacancies []models.Vacancy) []models.Vacancy {
	var failedVacancies []models.Vacancy
	vacanciesByBatches := utils.SplitSliceVacancy(vacancies, f.chunkSize)
	for i := range vacanciesByBatches {
		err := f.vacancyRepo.AddVacancies(ctx, vacanciesByBatches[i])
		if err != nil {
			// TODO: log the error somewhere
			failedVacancies = append(failedVacancies, vacanciesByBatches[i]...)
		}
	}
	return failedVacancies
}
