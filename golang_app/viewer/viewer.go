package viewer

import (
	"context"
	"errors"

	"github.com/asma12a/challenge-s6/ent/schema/ulid"
)

type User struct {
	ID ulid.ID
}

var userCtxKey = "user-viewer"

func UserFromContext(ctx context.Context) (*User, error) {
	user, ok := ctx.Value(userCtxKey).(*User)
	if !ok {
		return nil, errors.New("user not found in context")
	}
	return user, nil
}

func NewUserContext(ctx context.Context, user *User) context.Context {
	return context.WithValue(ctx, userCtxKey, user)
}
