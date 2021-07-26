package utils

import (
	"testing"

	"github.com/ozoncp/ocp-vacancy-api/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestSplitSpliceVacancy_LastChunk(t *testing.T) {
	assert.Equal(
		t,
		[][]models.Vacancy{
			{{ID: 1, Status: 1, Link: "1"}, {ID: 2, Status: 2, Link: "2"}},
			{{ID: 3, Status: 3, Link: "3"}},
		},
		SplitSliceVacancy(
			[]models.Vacancy{
				{ID: 1, Status: 1, Link: "1"},
				{ID: 2, Status: 2, Link: "2"},
				{ID: 3, Status: 3, Link: "3"},
			},
			2),
		"must be equal")

	assert.Equal(
		t,
		[][]models.Vacancy{
			{{ID: 1, Status: 1, Link: "1"}, {ID: 2, Status: 2, Link: "2"}},
			{{ID: 3, Status: 3, Link: "3"}, {ID: 4, Status: 4, Link: "4"}},
		},
		SplitSliceVacancy(
			[]models.Vacancy{
				{ID: 1, Status: 1, Link: "1"},
				{ID: 2, Status: 2, Link: "2"},
				{ID: 3, Status: 3, Link: "3"},
				{ID: 4, Status: 4, Link: "4"},
			},
			2),
		"must be equal")
}

func TestSplitSliceVacancy_NegativeChunkSize(t *testing.T) {
	assert.Equal(
		t,
		[][]models.Vacancy{
			{{ID: 1, Status: 1, Link: "1"}, {ID: 2, Status: 2, Link: "2"}, {ID: 3, Status: 3, Link: "3"}},
		},
		SplitSliceVacancy(
			[]models.Vacancy{
				{ID: 1, Status: 1, Link: "1"},
				{ID: 2, Status: 2, Link: "2"},
				{ID: 3, Status: 3, Link: "3"},
			},
			-1),
		"must return entire slice as first element when chunkSize is -1")
}
func TestSplitSliceVacancy_ZeroChunkSize(t *testing.T) {
	assert.Equal(
		t,
		[][]models.Vacancy{
			{{ID: 1, Status: 1, Link: "1"}, {ID: 2, Status: 2, Link: "2"}, {ID: 3, Status: 3, Link: "3"}},
		},
		SplitSliceVacancy(
			[]models.Vacancy{
				{ID: 1, Status: 1, Link: "1"},
				{ID: 2, Status: 2, Link: "2"},
				{ID: 3, Status: 3, Link: "3"},
			},
			0),
		"must return entire slice as first element when chunkSize is 0")
}

func TestSplitSliceVacancy_BigChunkSize(t *testing.T) {
	assert.Equal(
		t,
		[][]models.Vacancy{
			{{ID: 1, Status: 1, Link: "1"}, {ID: 2, Status: 2, Link: "2"}, {ID: 3, Status: 3, Link: "3"}},
		},
		SplitSliceVacancy(
			[]models.Vacancy{
				{ID: 1, Status: 1, Link: "1"},
				{ID: 2, Status: 2, Link: "2"},
				{ID: 3, Status: 3, Link: "3"},
			},
			10),
		"must return entire slice as first element when chunkSize is bigger than input slice length")
}

func TestGetVacancyMap_Success(t *testing.T) {
	result, err := GetVacancyMap(
		[]models.Vacancy{
			{ID: 1, Status: 1, Link: "1"},
			{ID: 2, Status: 2, Link: "2"},
			{ID: 3, Status: 3, Link: "3"},
		},
	)

	assert.NoError(
		t,
		err,
		"must return no errors",
	)

	assert.Equal(
		t,
		map[uint64]models.Vacancy{
			1: {ID: 1, Status: 1, Link: "1"},
			2: {ID: 2, Status: 2, Link: "2"},
			3: {ID: 3, Status: 3, Link: "3"},
		},
		result,
		"must be equal",
	)
}

func TestGetVacancyMap_DuplicatedVacancy(t *testing.T) {
	_, err := GetVacancyMap(
		[]models.Vacancy{
			{ID: 1, Status: 1, Link: "1"},
			{ID: 1, Status: 1, Link: "1"},
			{ID: 2, Status: 2, Link: "2"},
			{ID: 3, Status: 3, Link: "3"},
		},
	)

	if assert.Errorf(t, err, "must return IDAlreadyExistsError") {
		assert.Equal(t, "Vacancy with ID 1 already exists", err.Error(), "must be equal")
	}
}
