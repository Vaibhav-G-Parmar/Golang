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

func AddEvent(newEvent Event) error {
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

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Title, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}