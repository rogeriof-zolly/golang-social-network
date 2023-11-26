package repositories

import (
	"database/sql"
	"devbook/src/models"
	"fmt"
	"strings"
)

type Users struct {
	db *sql.DB
}

func NewUsersRepository(db *sql.DB) *Users {
	return &Users{db}
}

func (repository Users) Create(user models.User) (uint64, error) {
	// Prepare is used when values are gonna be passed to Exec after this line
	// This is for safety purposes, as it protects the code from sql injection, for example
	statement, err := repository.db.Prepare("insert into users (name, nickname, email, password) values (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nickname, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastId), nil
}

// Read returns an array of users
// The search parameter receives a query that searches for
// users with that given name or nickname
func (repository Users) Read(search string) ([]models.User, error) {
	query := fmt.Sprintf("%%%s%%", strings.TrimSpace(search)) // This equals to '%nameOrNick%'

	if search == "" {
		return []models.User{}, nil
	}

	rows, err := repository.db.Query("select ID, name, nickname, email, created_at from users where name like ? or nickname like ?", query, query)
	if err != nil {
		return []models.User{}, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Nickname,
			&user.CreatedAt,
		)

		if err != nil {
			return []models.User{}, err
		}

		users = append(users, user)
	}

	return users, nil
}
