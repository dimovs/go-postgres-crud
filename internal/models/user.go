package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func CreateUser(db *sql.DB, name, email string) (User, error) {
	var user User
	query := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id, name, email, created_at`

	err := db.QueryRow(query, name, email).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
	return user, err
}

func GetUserByID(db *sql.DB, id int) (User, error) {
	var user User
	query := `SELECT id, name, email, created_at FROM users WHERE id = $1`

	err := db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
	return user, err
}

func GetAllUsers(db *sql.DB) ([]User, error) {
	query := `SELECT id, name, email, created_at from users ORDER BY id`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, rows.Err()
}

func UpdateUser(db *sql.DB, id int, name string, email string) (User, error) {
	var user User
	query := `UPDATE users SET name = $1, email = $2 WHERE id = $3 RETURNING id, name, email, created_at`

	err := db.QueryRow(query, name, email, id).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
	return user, err
}

func DeleteUser(db *sql.DB, id int) error {
	query := `DELETE FROM users WHERE id = $1`

	_, err := db.Exec(query, id)
	return err
}
