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
