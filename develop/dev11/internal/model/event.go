package model

import "time"

type Event struct {
	ID        int
	UserID    int
	EventName string
	Date      *time.Time
}
