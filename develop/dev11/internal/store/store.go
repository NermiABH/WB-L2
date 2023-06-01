package store

import (
	"WB-L2/develop/dev11/internal/model"
	"errors"
	"sync"
	"time"
)

type Store struct {
	cache map[int]*model.Event
	mutex sync.RWMutex
}

func NewStore() *Store {
	return &Store{cache: make(map[int]*model.Event)}
}

func (s *Store) AddEvent(event *model.Event) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	event.ID = len(s.cache)
	s.cache[event.ID] = event
	return nil
}

func (s *Store) UpdateEvent(event *model.Event) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	_, ok := s.cache[event.ID]
	if !ok {
		return errors.New("event does not exist")
	}
	s.cache[event.ID] = event
	return nil
}

func (s *Store) DelEvent(id int) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	_, ok := s.cache[id]
	if !ok {
		return errors.New("event does not exist")
	}
	return nil
}

func (s *Store) GetEventsForSomeTime(userID int, start, end *time.Time) []*model.Event {
	s.mutex.RLock()
	eventSlice := make([]*model.Event, 0)
	for _, event := range s.cache {
		if event.UserID == userID {
			if event.Date.After(*start) && event.Date.Before(*end) {
				eventSlice = append(eventSlice, event)
			}
		}
	}
	return eventSlice
}
