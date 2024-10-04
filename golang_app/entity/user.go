package entity

import (
	"github.com/asma12a/challenge-s6/ent"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ent.User
}

func NewUser(email string) (*User, error) {
	user := &User{
		ent.User{
			Email: email,
		},
	}

	pwd, err := user.GeneratePassword("password")
	if err != nil {
		return nil, err
	}
	user.Password = pwd

	return user, nil
}

// ValidatePassword validate user password
func ValidatePassword(user *User, textPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(textPassword))
	if err != nil {
		return err
	}
	return nil
}

// generatePassword generate password
func (u *User) GeneratePassword(raw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(raw), 10)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
