package saver_test

import (
	"context"
	"time"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"

	//. "github.com/onsi/gomega"

	"github.com/ozoncp/ocp-vacancy-api/internal/mocks"
	"github.com/ozoncp/ocp-vacancy-api/internal/models"
	"github.com/ozoncp/ocp-vacancy-api/internal/saver"
)

var _ = Describe("Saver", func() {
	var (
		ctrl *gomock.Controller

		mockFlusher *mocks.MockFlusher

		s saver.Saver

		ctx context.Context
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())

		mockFlusher = mocks.NewMockFlusher(ctrl)

		s = saver.NewSaver(10, mockFlusher, 100*time.Millisecond)

		ctx = context.Background()
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("When saving a vacancy", func() {
		It("should call Flush", func() {
			mockFlusher.EXPECT().
				Flush(ctx, gomock.Any()).
				Times(1)

			s.Save(models.Vacancy{})
			s.Close()
			time.Sleep(10 * time.Millisecond) // sleep for 10ms to guarantee s.flush call after s.Close
		})
	})
})
