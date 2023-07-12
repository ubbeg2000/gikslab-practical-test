package dto

import (
	"errors"
	"net/mail"
)

type RegistrationBody struct {
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Username string   `json:"username"`
	Password string   `json:"password"`
	Profile  string   `json:"profile"`
	Skill    []string `json:"skill"`
}

func (b RegistrationBody) Validate() error {
	if b.Name == "" {
		return errors.New("name must not be empty")
	}

	if b.Email == "" {
		return errors.New("email must not be empty")
	}

	if _, err := mail.ParseAddress(b.Email); err != nil {
		return errors.New("invalid email address")
	}

	if b.Username == "" {
		return errors.New("username must not be empty")
	}

	if b.Password == "" {
		return errors.New("password must not be empty")
	}

	if b.Profile == "" {
		return errors.New("password must not be empty")
	}

	if b.Profile != "board" &&
		b.Profile != "expert" &&
		b.Profile != "trainer" &&
		b.Profile != "competitor" {
		return errors.New("profile must be either board, expert, trainer, or competitor")
	}

	return nil
}

type RegistrationResponse BaseResponse
