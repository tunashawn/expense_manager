package models

import (
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username       string `json:"username"`
	Password       string `json:"password,omitempty" bson:"-"`
	NewPassword    string `json:"new_password,omitempty" bson:"-"`
	HashedPassword []byte `json:"hashed_password,omitempty" bson:"password"`
	FullName       string `json:"full_name"`
	Email          string `json:"email"`
	Active         bool   `json:"active"`
}

// Formats the current user information to save to database, hash and salt the user password
func (u *User) FormatNewUserInformation() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.Wrap(err, "could not hash and salt password")
	}

	u.HashedPassword = hashedPassword
	u.Active = true

	return nil
}
