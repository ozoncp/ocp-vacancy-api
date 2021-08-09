package saver

import (
	"sync"
	"time"

	"github.com/ozoncp/ocp-vacancy-api/internal/flusher"
	"github.com/ozoncp/ocp-vacancy-api/internal/models"
)

type Saver interface {
	Save(vacancy models.Vacancy)
	Close()
}

// NewSaver return Saver with periodical flushing support.
// It initializes internal storage with given capacity,
// and starts time.Ticker with given flushInterval.
func NewSaver(capacity uint, flusher flusher.Flusher, flushInterval time.Duration) Saver {

	s := &saver{
		flusher: flusher,
		doneCh:  make(chan bool),
		mu:      sync.RWMutex{},
		storage: make([]models.Vacancy, 0, capacity),
	}

	ticker := time.NewTicker(flushInterval)

	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-s.doneCh: // shutdown signal, need to flush storage
				s.flush()
				return
			case <-ticker.C: // peroidically call flush
				s.flush()
			}
		}
	}()

	return s
}

type saver struct {
	flusher flusher.Flusher
	doneCh  chan bool

	mu      sync.RWMutex
	storage []models.Vacancy
}

// Save adds give vacancy into Saver's internal storage
func (s *saver) Save(vacancy models.Vacancy) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.storage = append(s.storage, vacancy)
}

// Close gracefully stops Saver, closing all its resources.
// This will also cause flush of vacancies of the Saver's internal storage if there are any
func (s *saver) Close() {
	s.doneCh <- true
}

// flush is a helper method to call Saver's Flusher.
func (s *saver) flush() {
	s.mu.Lock()
	defer s.mu.Unlock()

	notFlushedVacanices := s.flusher.Flush(s.storage)
	if len(notFlushedVacanices) > 0 {
		// TODO: Unclear what to de here
	}
	s.storage = s.storage[:0] // clean the slice keeping underlying array to avoid memory allocation
}
