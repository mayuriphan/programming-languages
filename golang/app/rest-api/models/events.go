package models

import "time"

type Event struct {
	ID          int
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	Date        time.Time `binding:"required"`
	UserID      int
}

var events []Event = []Event{}

func (e *Event) Save() { //Event can be used too
	events = append(events, *e) //e
}

func GetAllEvents() []Event { // uppercase function name is used to export the function to other packages; so can be used fromoutside package
	return events
}
