// event.go

package models

import "time"

type Event struct {
	ID          int		
	Title       string		`binding:"required"`
	Description string		`binding:"required"`
	Location    string		`binding:"required"`
	DateTime    time.Time	`binding:"required"`
	UserID      int
}

var events = []Event{}

func AddEvent(newEvent Event) {
	events = append(events, newEvent)
}

func GetAllEvents() []Event {
	return events
}