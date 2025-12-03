// event.go

package models

import (
	"event-booking-api/db"
	"time"
)

type Event struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"datetime" binding:"required"`
	UserID      int       `json:"user_id"`
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

func GetEventByID(id int64) (*Event, error) {
	query := `SELECT * FROM events WHERE id = ?`
	row := db.DB.QueryRow(query, id)
	var event Event
	err := row.Scan(&event.ID, &event.Title, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func UpdateEvent(updatedEvent Event) error {
	query := `
	UPDATE events 
	SET title = ?, description = ?, location = ?, datetime = ? 
	WHERE id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(updatedEvent.Title, updatedEvent.Description, updatedEvent.Location, updatedEvent.DateTime, updatedEvent.ID)
	return err
}

func DeleteEvent(id int64) error {
	query := `DELETE FROM events WHERE id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	return err
}