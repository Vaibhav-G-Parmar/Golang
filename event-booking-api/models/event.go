// event.go

package models

import (
	"event-booking-api/db"
	"time"
)

type Event struct {
	ID          int64		
	Title       string		`binding:"required"`
	Description string		`binding:"required"`
	Location    string		`binding:"required"`
	DateTime    time.Time	`binding:"required"`
	UserID      int
}

var events = []Event{}

func AddEvent(newEvent Event) error {

	//adding ? for secure insertion
	query := `
	INSERT INTO events (title, description, location, datetime, user_id) 
	VALUES (?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err :=stmt.Exec(newEvent.Title, newEvent.Description, newEvent.Location, newEvent.DateTime, newEvent.UserID)
	if err != nil {
		return err		
	}
	id, err := result.LastInsertId()
	newEvent.ID = id
	return err
}

func GetAllEvents() []Event {
	return events
}