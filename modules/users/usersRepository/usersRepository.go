package usersRepository

import (
	"context"
	"fmt"
	"time"

	"github.com/Rayato159/neversuitup-e-commerce-test/modules/users"
	"github.com/jmoiron/sqlx"
)

type IUsersRepository interface {
	InsertUser(req *users.UserCredential) (*users.UserPassport, error)
	FindOneUser(userId string) (*users.User, error)
}

type usersRepository struct {
	db *sqlx.DB
}

func NewUsersRepository(db *sqlx.DB) IUsersRepository {
	return &usersRepository{db}
}

func (r *usersRepository) InsertUser(req *users.UserCredential) (*users.UserPassport, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	query := `
	INSERT INTO "users" (
		"username",
		"password"
	)
	VALUES
	(
		:username,
		:password
	)
	RETURNING
		"id",
		"username";`

	row, err := r.db.NamedQueryContext(ctx, query, req)
	if err != nil {
		return nil, fmt.Errorf("insert user failed: %v", err)
	}
	defer row.Close()

	user := &users.UserPassport{
		User: &users.User{},
	}
	for row.Next() {
		if err := row.Scan(
			&user.User.Id,
			&user.User.Username,
		); err != nil {
			return nil, fmt.Errorf("scan user failed: %v", err)
		}
	}
	return user, nil
}

func (r *usersRepository) FindOneUser(userId string) (*users.User, error) {
	_, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	query := `
	SELECT
		"u"."id",
		"u"."username"
	FROM "users" "u"
	WHERE "u"."id" = $1`

	user := new(users.User)
	if err := r.db.Get(user, query, userId); err != nil {
		return nil, fmt.Errorf("get user failed: %v", err)
	}
	return user, nil
}
