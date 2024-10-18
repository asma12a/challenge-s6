package service

import (
	"context"

	"github.com/asma12a/challenge-s6/ent"
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
	"github.com/asma12a/challenge-s6/ent/user"
	"github.com/asma12a/challenge-s6/entity"
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
		Save(ctx)

	if err != nil {
		if ent.IsConstraintError(err) {
			return nil, entity.ErrEmailAlreadyRegistred
		}
		return nil, entity.ErrCannotBeCreated
	}

	return entUser, nil
}

func (e *User) FindOne(ctx context.Context, id ulid.ID) (*ent.User, error) {
	return e.db.User.Query().Where(user.IDEQ(id)).Only(ctx)
}

func (repo *User) Update(ctx context.Context, user *entity.User) (*ent.User, error) {

	// Prepare the update query
	entUser, err := repo.db.User.
		UpdateOneID(user.ID).
		SetName(user.Name).
		SetEmail(user.Email).
		SetPassword(user.Password).
		SetRoles(user.Roles).Save(ctx)

	if err != nil {
		return nil, entity.ErrCannotBeUpdated
	}
	return entUser, nil
}

func (repo *User) Delete(ctx context.Context, id ulid.ID) error {
	err := repo.db.User.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return entity.ErrCannotBeDeleted
	}
	return nil
}

func (repo *User) List(ctx context.Context) ([]*ent.User, error) {
	return repo.db.User.Query().Select(user.FieldID, user.FieldName, user.FieldEmail, user.FieldRoles).All(ctx)
}
