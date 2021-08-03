package flusher_test

import (
	"errors"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozoncp/ocp-vacancy-api/internal/flusher"
	"github.com/ozoncp/ocp-vacancy-api/internal/mocks"
	"github.com/ozoncp/ocp-vacancy-api/internal/models"
)

var _ = Describe("Flusher", func() {
	var (
		ctrl *gomock.Controller

		mockRepo *mocks.MockRepo

		f flusher.Flusher
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())

		mockRepo = mocks.NewMockRepo(ctrl)

		f = flusher.NewFlusher(3, mockRepo)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("", func() {

		It("Empty slice in Flush must not call repo.AddVacancies", func() {
			mockRepo.EXPECT().
				AddVacancies(gomock.Any()).
				Times(0)

			f.Flush([]models.Vacancy{})
		})

		It("Input slice with len <= flusher.chunkSize in Flush must call repo.AddVacancies just once", func() {
			// case: len < flusher.chunkSize
			mockRepo.EXPECT().
				AddVacancies(gomock.Any()).
				Times(1)

			f.Flush([]models.Vacancy{
				{},
				{},
			})

			// case: len == flusher.chunkSize
			mockRepo.EXPECT().
				AddVacancies(gomock.Any()).
				Times(1)

			f.Flush([]models.Vacancy{
				{},
				{},
				{},
			})
		})

		It("Flush must return some Vacancies as a result, if repo call ends with an error", func() {
			mockRepo.EXPECT().
				AddVacancies(gomock.Any()).
				AnyTimes().
				Return(errors.New("some error"))

			result := f.Flush([]models.Vacancy{
				{ID: 1, Link: "link1", Status: 1},
				{ID: 2, Link: "link2", Status: 2},
				{ID: 3, Link: "link3", Status: 3},
				{ID: 4, Link: "link4", Status: 4},
			})

			Expect(result).ShouldNot(BeEmpty())
		})
	})
})
