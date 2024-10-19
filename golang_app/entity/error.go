package entity

import "errors"

// ErrNotFound
var ErrNotFound = errors.New("Not found")

// Dynamic Not Found
func ErrEntityNotFound(entity string) error {
	return errors.New(entity + " not found")
}

// ErrEmailAlreadyRegistred
var ErrEmailAlreadyRegistred = errors.New("Email already registred")

// ErrInvalidEntity
var ErrInvalidEntity = errors.New("Invalid entity")

// Dynamic Invalid relation
func ErrInvalidRelation(relation string) error {
	return errors.New("Invalid " + relation + ": relation not found")
}

// ErrInvalidInput
var ErrInvalidInput = errors.New("Invalid input")

// ErrCannotBeCreated
var ErrCannotBeCreated = errors.New("Cannot be created")

// ErrCannotBeDeleted
var ErrCannotBeDeleted = errors.New("Cannot be Deleted")

// ErrCannotBeUpdated
var ErrCannotBeUpdated = errors.New("Cannot be updated")

// ErrCannotBeDeleted
var ErrPasswordGenaration = errors.New("Password cannot be generated")

// ErrCannotParseJSON
var ErrCannotParseJSON = errors.New("Cannot parse JSON")
