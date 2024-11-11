package service

import (
	"context"
	"strings"

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

func (repo *User) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	createdUser, err := repo.db.User.Create().
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

	return &entity.User{User: *createdUser}, nil
}

func (u *User) FindOne(ctx context.Context, id ulid.ID) (*entity.User, error) {
	user, err := u.db.User.Query().Where(user.IDEQ(id)).
		Only(ctx)

	if err != nil {
		return nil, entity.ErrEntityNotFound("User")
	}
	return &entity.User{User: *user}, nil
}

func (u *User) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	email = strings.ToLower(email) // Normaliser l'email
	user, err := u.db.User.Query().Where(user.Email(email)).Only(ctx)

	if ent.IsNotFound(err) {
		return nil, entity.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return &entity.User{User: *user}, nil
}

func (repo *User) Update(ctx context.Context, user *entity.User) (*entity.User, error) {

	// Prepare the update query
	entUser, err := repo.db.User.
		UpdateOneID(user.ID).
		SetName(user.Name).
		SetEmail(user.Email).
		SetPassword(user.Password).
		SetRoles(user.Roles).Save(ctx)

	if err != nil {
		return nil, err
	}
	return &entity.User{User: *entUser}, nil
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
