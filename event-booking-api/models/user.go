package models

import (
	"errors"
	"github.com/Lucas-Mol/go-studies/event-booking-api/db"
	"github.com/Lucas-Mol/go-studies/event-booking-api/utils"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `binding:"required" json:"email"`
	Password string `binding:"required" json:"password"`
}

func (u *User) Save() error {
	query := `INSERT INTO tb_users (email, password) VALUES (?, ?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	u.ID = userID
	return nil
}

func (u *User) ValidateCredentials() error {
	query := `SELECT id, password FROM tb_users WHERE email = ?`
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)
	if err != nil {
		return err
	}

	isPasswordValid := utils.CheckPasswordHash(u.Password, retrievedPassword)
	if !isPasswordValid {
		return errors.New("credentials invalid")
	}

	return nil
}
