package entity

import "errors"

var ErrNotFound = errors.New("not found")

func ErrEntityNotFound(entity string) error {
	return errors.New(entity + " not found")
}

var ErrEmailAlreadyRegistred = errors.New("the email address is already in use")

var ErrInvalidEntity = errors.New("invalid entity")

func ErrInvalidRelation(relation string) error {
	return errors.New("invalid " + relation + ": relation not found")
}

var ErrInvalidInput = errors.New("invalid input")

var ErrInvalidID = errors.New("invalid ID")

var ErrCannotBeCreated = errors.New("cannot be created")

var ErrCannotBeDeleted = errors.New("cannot be deleted")

var ErrCannotBeUpdated = errors.New("cannot be updated")

var ErrPasswordNotStrong = errors.New("the password is not strong enough. Please choose a stronger password")

var ErrInvalidPassword = errors.New("incorrect credentials. Please check your information and try again")

var ErrUserNotActive = errors.New("the user is not active. Please verify your account")

var ErrCannotParseJSON = errors.New("cannot parse JSON")

var ErrTeamFull = errors.New("team is full")

var ErrUserAlreadyInATeam = errors.New("user is already in a team")

var ErrUserAlreadyInThisTeam = errors.New("user is already in this team")
