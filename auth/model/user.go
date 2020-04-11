package model

import (
	"errors"
	"strings"
)

type User struct {
	BaseProperties
	Name     string
	Email    string
	Password string
}

func (user *User) validate() error {
	if len(strings.TrimSpace(user.Name)) == 0 {
		return errors.New("Invalid name")
	}

	if len(strings.TrimSpace(user.Email)) == 0 {
		return errors.New("Invalid email")
	}

	if len(strings.TrimSpace(user.Password)) == 0 {
		return errors.New("Invalid password")
	}

	return nil
}

func (user *User) ValidateNew() error {
	return user.validate()
}

func (user *User) ValidateModify() error {
	if user.Id == 0 {
		return errors.New("Invalid id")
	}

	return user.validate()
}
