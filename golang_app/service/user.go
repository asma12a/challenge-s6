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

// @Summary Create a new user
// @Description Create a new user with name, email, and password
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body entity.User true "User to be created"
// @Success 201 {object} entity.User "User created"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 409 {object} map[string]interface{} "Email already registered"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /users [post]
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

// @Summary Get a user by ID
// @Description Get a specific user by their ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} entity.User "User details"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 404 {object} map[string]interface{} "User Not Found"
// @Router /users/{id} [get]
func (u *User) FindOne(ctx context.Context, id ulid.ID) (*entity.User, error) {
	user, err := u.db.User.Query().Where(user.IDEQ(id)).
		Only(ctx)

	if err != nil {
		return nil, entity.ErrEntityNotFound("User")
	}
	return &entity.User{User: *user}, nil
}

// @Summary Get a user by email
// @Description Get a specific user by their email
// @Tags users
// @Accept  json
// @Produce  json
// @Param email path string true "User Email"
// @Success 200 {object} entity.User "User details"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 404 {object} map[string]interface{} "User Not Found"
// @Router /users/email/{email} [get]

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

// @Summary Update a user
// @Description Update a user's details by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Param user body entity.User true "Updated user data"
// @Success 200 {object} entity.User "Updated user"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 404 {object} map[string]interface{} "User Not Found"
// @Router /users/{id} [put]
func (repo *User) Update(ctx context.Context, user *entity.User) (*entity.User, error) {

	// Prepare the update query
	entUser, err := repo.db.User.
		UpdateOneID(user.ID).
		SetName(user.Name).
		SetEmail(user.Email).
		SetPassword(user.Password).
		SetUpdatedAt(user.UpdatedAt).
		SetIsActive(user.IsActive).
		SetRoles(user.Roles).Save(ctx)

	if err != nil {
		return nil, err
	}
	return &entity.User{User: *entUser}, nil
}

// @Summary Delete a user
// @Description Delete a user by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} map[string]interface{} "User deleted"
// @Failure 404 {object} map[string]interface{} "User Not Found"
// @Router /users/{id} [delete]
func (repo *User) Delete(ctx context.Context, id ulid.ID) error {
	err := repo.db.User.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return entity.ErrCannotBeDeleted
	}
	return nil
}

// @Summary List all users
// @Description Get a list of all users
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {array} entity.User "List of users"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /users [get]
func (repo *User) List(ctx context.Context) ([]*ent.User, error) {
	return repo.db.User.Query().Select(user.FieldID, user.FieldName, user.FieldEmail, user.FieldRoles).All(ctx)
}
