package modells

import (
	"time"

	"example.com/rest-api/database"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name", binding:"required"`
	Description string    `json:"description", binding:"required"`
	Location    string    `json:"location", binding:"required"`
	DateTime    time.Time `json:"date_time", binding:"required"`
	UserId      int       `json:"user_id"`
}

var events = []Event{}

func (e *Event) Save() error {
	eventQuery := `INSERT INTO events (name, description, location, date_time, user_id) VALUES (?, ?, ?, ?, ?)`

	stmt, err := database.Db.Prepare(eventQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)
	if err != nil {
		return err
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	e.ID = lastInsertID
	return err

	//events = append(events, e)
}

func GetEvents() ([]Event, error) {
	fetchQuery := `SELECT * FROM events`
	rows, err := database.Db.Query(fetchQuery)
	if err != nil {
		return nil, err
	}
	var events []Event
	for rows.Next() {
		var event Event
		if err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId); err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	defer rows.Close()
	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	fetchQuery := `SELECT * FROM events WHERE id = ?`
	row := database.Db.QueryRow(fetchQuery, id)
	var event Event
	if err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId); err != nil {
		return nil, err
	}
	return &event, nil
}

// func (event *Event) UpdateEvent() error {
// 	updateQuery := `UPDATE events SET name = ?, description = ?, location = ?, date_time = ? WHERE id = ?`
// 	stmt, err := database.Db.Prepare(updateQuery)
// 	if err != nil {
// 		return err
// 	}
// 	defer stmt.Close()
// 	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)
// 	return err
// }

func UpdateEvent(event *Event) error {
	updateQuery := `UPDATE events SET name = ?, description = ?, location = ?, date_time = ? WHERE id = ?`
	stmt, err := database.Db.Prepare(updateQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)
	return err
}

func DeleteEvent(id int64) error {
	deleteQuery := `DELETE FROM events WHERE id = ?`
	stmt, err := database.Db.Prepare(deleteQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	return err
}
