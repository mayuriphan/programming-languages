package models

import (
	"rest-api/db"
	"time"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	Date        time.Time `binding:"required"`
	UserID      int
}

var events []Event = []Event{}

func (e Event) Save() error { // *Event an *e can be used too
	query := `
	INSERT INTO events (name, description, location, date, user_id)
	VALUES (?, ?, ?, ?, ?)` // ? is used to prevent SQL injection attacks

	stmt, err := db.DB.Prepare(query) // stored by sql package
	if err != nil {
		return err
	}
	defer stmt.Close()                                                            // to close the statement after execution
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.Date, e.UserID) //updt,insert i.e change the database
	if err != nil {
		return err
	}
	id, err := result.LastInsertId() // to get the last inserted id
	e.ID = id
	return err
}

func GetAllEvents() ([]Event, error) { // uppercase function name is used to export the function to other packages; so can be used fromoutside package
	query := `SELECT id, name, description, location, date, user_id FROM events`
	rows, err := db.DB.Query(query) // fetch
	if err != nil {
		return nil, err
	}
	defer rows.Close() // to close the rows after execution
	var events []Event
	for rows.Next() { // to iterate over the rows returns true or false
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.Date, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	query := `SELECT id, name, description, location, date, user_id FROM events WHERE id = ?`
	row := db.DB.QueryRow(query, id) // fetch exactly one row from the database
	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.Date, &event.UserID)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (event Event) Update() error {
	query := `
	UPDATE events
	SET name = ?, description = ?, location = ?, date = ?, user_id = ?
	WHERE id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.Date, event.UserID, event.ID)
	return err
}

func (event Event) Delete() error {
	query := `DELETE FROM events WHERE id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(event.ID)
	return err
}
