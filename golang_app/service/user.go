package service

import (
	"context"

	"github.com/asma12a/challenge-s6/ent"
)

type User struct {
	db *ent.Client
}

func NewUserService(client *ent.Client) *User {
	return &User{
		db: client,
	}
}

func (u *User) ListUsers(ctx context.Context) ([]*ent.User, error) {
	return u.db.User.Query().All(ctx)
}
