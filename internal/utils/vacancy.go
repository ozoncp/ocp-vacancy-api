package utils

import (
	"fmt"

	"github.com/ozoncp/ocp-vacancy-api/internal/models"
)

// SplitSliceVacancy converts input Vacancy slice into slice of Vacancy slices.
// Size of chunk is defined by chunkSize. Size of last chunk is less or equal chunkSize.
func SplitSliceVacancy(input []models.Vacancy, chunkSize int) [][]models.Vacancy {
	if chunkSize <= 0 {
		return [][]models.Vacancy{input}
	}
	// calculating number of chunks to allocate memory
	// for resulting slice (len and cap)
	chunksCount := len(input) / chunkSize
	if len(input)%chunkSize > 0 {
		chunksCount = chunksCount + 1
	}
	result := make([][]models.Vacancy, chunksCount)

	for i := 0; i < chunksCount; i++ {
		lowBound := i * chunkSize
		highBound := lowBound + chunkSize
		if highBound > len(input) {
			highBound = len(input)
		}
		result[i] = input[lowBound:highBound]
	}
	return result
}

type IDAlreadyExistsError struct {
	ID uint64
}

func (e *IDAlreadyExistsError) Error() string {
	return fmt.Sprintf("Vacancy with ID %d already exists", e.ID)
}

// GetVacancyMap convert input Vacancy slice into map[uint64]Vacancy,
// where map key is vacancy ID
func GetVacancyMap(input []models.Vacancy) (map[uint64]models.Vacancy, error) {
	result := make(map[uint64]models.Vacancy, len(input))
	for _, v := range input {
		if _, exists := result[v.ID]; exists {
			return nil, &IDAlreadyExistsError{v.ID}
		}
		result[v.ID] = v
	}
	return result, nil
}
