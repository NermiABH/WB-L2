package server

import (
	"WB-L2/develop/dev11/internal/store"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Server struct {
	store *store.Store
	mux   *http.ServeMux
}

func NewServer() *Server {
	newServer := &Server{
		store: store.NewStore(),
		mux:   http.NewServeMux(),
	}
	newServer.configureHandlers()
	return newServer
}

func (s *Server) configureHandlers() {
	s.mux.HandleFunc("/create_event", logRequest(s.CreateEvent))
	s.mux.HandleFunc("/update_event", logRequest(s.UpdateEvent))
	s.mux.HandleFunc("/delete_event", logRequest(s.DeleteEvent))
	s.mux.HandleFunc("/events_for_day", logRequest(s.EventsForDay))
	s.mux.HandleFunc("/events_for_week", logRequest(s.EventsForWeek))
	s.mux.HandleFunc("/events_for_month", logRequest(s.EventsForMonth))
}

func (s *Server) Run(addr string) error {
	log.Println("Starting server")
	return http.ListenAndServe(addr, s.mux)
}

func logRequest(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next(w, r)
		log.Printf("method: %s, url: %s, time: %s",
			r.Method, r.URL.EscapedPath(), time.Since(start))
	}
}

type Result struct {
	msg any `json:"result"`
}

type Error struct {
	msg any `json:"error"`
}

func response(w http.ResponseWriter, _ *http.Request, code int, msg any) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(msg)
}
