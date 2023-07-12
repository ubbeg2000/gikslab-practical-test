package dto

import "errors"

type LoginBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (b LoginBody) Validate() error {
	if b.Username == "" {
		return errors.New("username must not be empty")
	}

	if b.Password == "" {
		return errors.New("password must not be empty")
	}

	return nil
}

type LoginSuccessResponse struct {
	Token   string `json:"token"`
	Profile string `json:"profile"`
}
