package server

import (
	"WB-L2/develop/dev11/internal/model"
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func (s *Server) CreateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response(w, r, 500, Error{"post request only"})
		return
	}
	query := r.URL.Query()
	eventName, err := isValidEventName(query)
	if err != nil {
		response(w, r, 400, Error{err.Error()})
		return
	}
	userID, err := isValidUserID(query)
	if err != nil {
		response(w, r, 400, Error{err.Error()})
		return
	}
	date, err := isValidDate(query)
	if err != nil {
		response(w, r, 400, Error{err.Error()})
		return
	}
	event := &model.Event{EventName: eventName, UserID: userID, Date: date}
	if err = s.store.AddEvent(event); err != nil {
		response(w, r, 503, Error{err.Error()})
		return
	}
	response(w, r, 200, Result{event})
}

func (s *Server) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response(w, r, 500, Error{"post request only"})
		return
	}
	query := r.URL.Query()
	eventName, err := isValidEventName(query)
	if err != nil {
		response(w, r, 400, Error{err.Error()})
		return
	}
	userID, err := isValidUserID(query)
	if err != nil {
		response(w, r, 400, Error{err.Error()})
		return
	}
	date, err := isValidDate(query)
	if err != nil {
		response(w, r, 400, Error{err.Error()})
		return
	}
	event := &model.Event{EventName: eventName, UserID: userID, Date: date}
	if err = s.store.UpdateEvent(event); err != nil {
		response(w, r, 503, Error{err.Error()})
		return
	}
	response(w, r, 200, Result{event})
}

func (s *Server) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response(w, r, 500, Error{"post request only"})
		return
	}
	query := r.URL.Query()
	id, err := isValidID(query)
	if err != nil {
		response(w, r, 400, Error{err.Error()})
		return
	}
	if err := s.store.DelEvent(id); err != nil {
		response(w, r, 503, Error{err.Error()})
		return
	}
	response(w, r, 200, Result{"deletion completed successfully"})
}

func (s *Server) EventsForDay(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response(w, r, 500, Error{"get request only"})
		return
	}
	query := r.URL.Query()
	userID, err := isValidUserID(query)
	if err != nil {
		response(w, r, 400, Error{err.Error()})
		return
	}
	date, err := isValidDate(query)
	if err != nil {
		response(w, r, 400, Error{err.Error()})
		return
	}
	dateEnd := date.AddDate(0, 0, 1)
	events := s.store.GetEventsForSomeTime(userID, date, &dateEnd)
	response(w, r, 200, Result{events})
}

func (s *Server) EventsForWeek(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response(w, r, 500, Error{"get request only"})
		return
	}
	query := r.URL.Query()
	userID, err := isValidUserID(query)
	if err != nil {
		response(w, r, 400, Error{err.Error()})
		return
	}
	date, err := isValidDate(query)
	if err != nil {
		response(w, r, 400, Error{err.Error()})
		return
	}
	dateEnd := date.AddDate(0, 0, 7)
	events := s.store.GetEventsForSomeTime(userID, date, &dateEnd)
	response(w, r, 200, Result{events})
}

func (s *Server) EventsForMonth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response(w, r, 500, Error{"get request only"})
		return
	}
	query := r.URL.Query()
	userID, err := isValidUserID(query)
	if err != nil {
		response(w, r, 400, Error{err.Error()})
		return
	}
	date, err := isValidDate(query)
	if err != nil {
		response(w, r, 400, Error{err.Error()})
		return
	}
	dateEnd := date.AddDate(0, 1, 0)
	events := s.store.GetEventsForSomeTime(userID, date, &dateEnd)
	response(w, r, 200, Result{events})

}

func isValidEventName(query url.Values) (string, error) {
	eventName := query.Get("event_name")
	if eventName == "" {
		return "", errors.New("event_name is empty")
	}
	return eventName, nil
}

func isValidID(query url.Values) (int, error) {
	idStr := query.Get("id")
	if idStr == "" {
		return 0, errors.New("id is empty")
	}
	id, err := strconv.Atoi(idStr)
	if err != nil || id > 0 {
		return 0, errors.New("id is not valid")
	}
	return id, nil
}

func isValidUserID(query url.Values) (int, error) {
	userIDStr := query.Get("user_id")
	if userIDStr == "" {
		return 0, errors.New("user_id is empty")
	}
	userID, err := strconv.Atoi(userIDStr)
	if err != nil || userID > 0 {
		return 0, errors.New("user_id is not valid")
	}
	return userID, nil
}

func isValidDate(query url.Values) (*time.Time, error) {
	dateStr := query.Get("date")
	if dateStr != "" {
		return nil, errors.New("date is empty")
	}
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return nil, errors.New("date is not valid")
	}
	return &date, nil
}
