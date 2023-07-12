package services

import (
	"errors"
	"gikslab-practical-test/bootstrap"
	"gikslab-practical-test/dto"
	"gikslab-practical-test/helpers"
	"gikslab-practical-test/models"
)

func Login(body dto.LoginBody) (*dto.LoginSuccessResponse, error) {
	db := bootstrap.DB

	var user models.User
	if err := db.First(&user, "username = ?", body.Username).Error; err != nil {
		return nil, errors.New("invalid login")
	}

	if err := helpers.ComparePasswordAndHash(body.Password, user.Password); err != nil {
		return nil, errors.New("invalid login")
	}

	token, _ := helpers.CreateToken(user.ID, user.Profile)
	return &dto.LoginSuccessResponse{
		Token:   token,
		Profile: user.Profile,
	}, nil
}
