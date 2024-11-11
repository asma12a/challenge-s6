package entity

import "errors"

var ErrNotFound = errors.New("Not found")

func ErrEntityNotFound(entity string) error {
	return errors.New(entity + " not found")
}

var ErrEmailAlreadyRegistred = errors.New("L'adresse email est déjà utilisée.")

var ErrInvalidEntity = errors.New("Invalid entity")

func ErrInvalidRelation(relation string) error {
	return errors.New("Invalid " + relation + ": relation not found")
}

var ErrInvalidInput = errors.New("Invalid input")

var ErrInvalidID = errors.New("Invalid ID")

var ErrCannotBeCreated = errors.New("Cannot be created")

var ErrCannotBeDeleted = errors.New("Cannot be deleted")

var ErrCannotBeUpdated = errors.New("Cannot be updated")

var ErrPasswordGenaration = errors.New("Password cannot be generated")

var ErrPasswordNotStrong = errors.New("Le mot de passe n'est pas assez fort. Veuillez choisir un mot de passe plus fort.")

var ErrInvalidPassword = errors.New("Identifiants incorrects. Veuillez vérifier vos informations et réessayer.")

var ErrCannotParseJSON = errors.New("Cannot parse JSON")
