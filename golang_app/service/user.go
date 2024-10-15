package service

import (
	"context"

	"github.com/asma12a/challenge-s6/ent"
	"github.com/asma12a/challenge-s6/ent/user"
	"github.com/asma12a/challenge-s6/entity"
	"github.com/google/uuid"
)

type User struct {
	db *ent.Client
}

func NewUserService(client *ent.Client) *User {
	return &User{
		db: client,
	}
}

func (repo *User) Create(ctx context.Context, user *entity.User) (*ent.User, error) {
	entUser, err := repo.db.User.Create().
		SetName(user.Name).
		SetEmail(user.Email).
		SetPassword(user.Password).
		SetRole(user.Role).
		Save(ctx)

	if err != nil {
		return nil, entity.ErrCannotBeCreated
	}

	return entUser, nil
}

func (e *User) FindOne(ctx context.Context, id uuid.UUID ) (*ent.User, error) {
	return e.db.User.Query().Where(user.IDEQ(id.String())).Only(ctx)
}

func (repo *User) Update(ctx context.Context, user *entity.User) (*ent.User, error) {

	// Prepare the update query
	entUser, err := repo.db.User.
		UpdateOneID(user.ID).
		SetName(user.Name).
		SetEmail(user.Email).
		SetPassword(user.Password).
		SetRole(user.Role).Save(ctx)

	if err != nil {
		return nil, entity.ErrInvalidEntity
	}
	return entUser, nil
}

func (repo *User) Delete(ctx context.Context, id uuid.UUID) error {
	err := repo.db.User.DeleteOneID(id.String()).Exec(ctx)
	if err != nil {
		return entity.ErrCannotBeDeleted
	}
	return nil
}

func (repo *User) List(ctx context.Context) ([]*ent.User, error) {
	return repo.db.User.Query().All(ctx)
}
