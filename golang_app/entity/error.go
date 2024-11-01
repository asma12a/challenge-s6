package entity

import "errors"

var ErrNotFound = errors.New("Not found")

func ErrEntityNotFound(entity string) error {
	return errors.New(entity + " not found")
}

var ErrEmailAlreadyRegistred = errors.New("Email already registred")

var ErrInvalidEntity = errors.New("Invalid entity")

func ErrInvalidRelation(relation string) error {
	return errors.New("Invalid " + relation + ": relation not found")
}

var ErrInvalidInput = errors.New("Invalid input")

var ErrCannotBeCreated = errors.New("Cannot be created")

var ErrCannotBeDeleted = errors.New("Cannot be deleted")

var ErrCannotBeUpdated = errors.New("Cannot be updated")

var ErrPasswordGenaration = errors.New("Password cannot be generated")

var ErrInvalidPassword = errors.New("Invalid password")

var ErrCannotParseJSON = errors.New("Cannot parse JSON")
