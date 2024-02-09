package repositories

import (
	"database/sql"
	"devbook/src/models"
	"fmt"
)

type Followers struct {
	db *sql.DB
}

func NewFollowersRepository(db *sql.DB) *Followers {
	return &Followers{db}
}

func (repository Followers) FollowUser(
	loggedUserId uint64, userToFollow uint64,
) error {
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

func (respository Followers) UnfollowUser(
	loggedUserID uint64, userToUnfollow uint64,
) error {
	statement, err := respository.db.Prepare(
		"delete from followers where user_id = ? and follower_id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(loggedUserID, userToUnfollow); err != nil {
		return err
	}

	return nil
}

func (repository Followers) GetAllFollowers(userID uint64) (models.Followers, error) {
	followersQuery := fmt.Sprintf(
		`select u.ID, u.name, u.nickname, u.created_at
    from users u 
    inner join followers f on u.id = f.follower_id
    where f.user_id = %d`, userID,
	)

	rows, err := repository.db.Query(followersQuery)
	if err != nil {
		return models.Followers{}, err
	}
	defer rows.Close()

	var followers models.Followers
	var follower models.User

	for rows.Next() {
		follower = models.User{}

		err := rows.Scan(
			&follower.ID,
			&follower.Name,
			&follower.Nickname,
			&follower.CreatedAt,
		)
		if err != nil {
			return models.Followers{}, err
		}

		followers.Followers = append(followers.Followers, follower)
		followers.FollowerCount = followers.FollowerCount + 1
	}

	return followers, nil
}
