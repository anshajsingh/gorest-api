package modells

import (
	"example.com/rest-api/database"
	"example.com/rest-api/utils"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u *User) Save() error {
	userQuery := `INSERT INTO users (email, password) VALUES (?, ?)`
	stmt, err := database.Db.Prepare(userQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hashPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(u.Email, hashPassword)
	if err != nil {
		return err
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	u.ID = lastInsertID
	return nil
}

func GetAllUsers() ([]User, error) {
	fetchQuery := `SELECT * FROM users`
	rows, err := database.Db.Query(fetchQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Email, &user.Password); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (user *User) ValidateUser() (bool, error) {
	fetchQuery := `SELECT * FROM users WHERE email = ?`
	row := database.Db.QueryRow(fetchQuery, user.Email)

	var dbUser User
	if err := row.Scan(&dbUser.ID, &dbUser.Email, &dbUser.Password); err != nil {
		return false, err
	}

	if !utils.CheckPasswordHash(user.Password, dbUser.Password) {
		return false, nil
	}

	// if !utils.CheckPasswordHash(user.Password, dbUser.Password) {
	// 	return false
	// }

	user.ID = dbUser.ID
	return true, nil
}
