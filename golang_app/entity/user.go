package entity

import (
	"github.com/asma12a/challenge-s6/ent"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ent.User
}

func NewUser(email string, name string, password string) (*User, error) {
	user := &User{
		ent.User{
			Email:    email,
			Name:     name,
			Password: password,
		},
	}

	pwd, err := user.GeneratePassword(password)
	if err != nil {
		return nil, err
	}
	user.Password = pwd

	return user, nil
}

func ValidatePassword(user *User, textPassword string) error {

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(textPassword))
	if err != nil {
		return err
	}
	return nil
}

// generatePassword generate password
func (u *User) GeneratePassword(raw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(raw), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
