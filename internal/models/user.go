package models

import (
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type User struct {
	GUID      string `json:"GUID"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (user *User) validate() error {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)

	user.Email = strings.TrimSpace(user.Email)
	if user.Email == "" {
		return fmt.Errorf("invalid email")
	}

	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return fmt.Errorf("invalid password")
	}

	return nil
}

// CreateUser encrypts user password and validates a data
func (user *User) CreateUser() error {
	if err := user.validate(); err != nil {
		return err
	}

	pwSlice, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return fmt.Errorf("failed to encrypt the password")
	}
	user.Password = string(pwSlice)

	// TODO guid for user
	user.GUID = uuid.New().String()

	return nil
}
