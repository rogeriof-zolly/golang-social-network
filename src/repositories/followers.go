package repositories

import (
	"database/sql"
)

type Followers struct {
	db *sql.DB
}

func NewFollowersRepository(db *sql.DB) *Followers {
	return &Followers{db}
}

func (repository Followers) FollowUser(loggedUserId uint64, userToFollow uint64) error {
	// 'ignore' avoids the program to throw error
	// if the given user already follows the one it desires
	statement, err := repository.db.Prepare(
		"insert ignore into followers (user_id, follower_id) values (?, ?)",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(loggedUserId, userToFollow); err != nil {
		return err
	}

	return nil
}
