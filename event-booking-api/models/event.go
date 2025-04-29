package models

import (
	"github.com/Lucas-Mol/go-studies/event-booking-api/db"
	"time"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `binding:"required" json:"name"`
	Description string    `binding:"required"  json:"description"`
	Location    string    `binding:"required"  json:"location"`
	DateTime    time.Time `binding:"required"  json:"datetime"`
	UserID      int64     `json:"user_id"`
}

func (e *Event) Save() error {
	query := `
	INSERT INTO tb_events(name, description, location, date_time, user_id) 
	VALUES (?, ?, ?, ?, ?);
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	e.ID = id

	return err
}

func (e *Event) Update() error {
	query := `
	UPDATE tb_events
	SET name=?, description=?, location=?, date_time=?
	WHERE id=?;
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.ID)
	return err
}

func (e *Event) Delete() error {
	query := `DELETE FROM tb_events WHERE id=?;`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.ID)
	return err
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM tb_events`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events = []Event{}
	for rows.Next() {
		var event Event
		err := rows.Scan(
			&event.ID,
			&event.Name,
			&event.Description,
			&event.Location,
			&event.DateTime,
			&event.UserID,
		)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func GetEventByID(id int64) (Event, error) {
	query := `SELECT * FROM tb_events WHERE id = ?;`
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(
		&event.ID,
		&event.Name,
		&event.Description,
		&event.Location,
		&event.DateTime,
		&event.UserID,
	)
	if err != nil {
		return Event{}, err
	}

	return event, nil
}

func (e *Event) Register(userId int64) error {
	query := `INSERT INTO tb_registrations(event_id, user_id) VALUES (?, ?);`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)
	return err
}

func (e *Event) CancelRegistration(userId int64) error {
	query := `DELETE FROM tb_registrations WHERE user_id = ? AND event_id = ?;`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)
	return err
}
