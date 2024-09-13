package models

import (
	"fmt"
	"strings"
)

type User struct {
	ID        int    `json:"ID"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (user *User) Validate() error {
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
