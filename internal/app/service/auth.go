package service

import (
	"Gotenv/internal/app/models"
	"Gotenv/internal/repository"
	"Gotenv/pkg/errs"
	"Gotenv/pkg/utils"
	"errors"
)

func SignIn(userDataCheck, password string) (user models.User, accessToken string, refreshToken string, err error) {
	if userDataCheck == "" {
		return user, "", "", errs.ErrInvalidData
	}

	user, err = repository.GetUserByUsernameAndPassword(userDataCheck, password)
	if err != nil {
		if !errors.Is(err, errs.ErrRecordNotFound) {
			return user, "", "", err
		}

		return user, "", "", errs.ErrInvalidCredentials
	}

	accessToken, refreshToken, err = utils.GenerateToken(uint(user.RoleID), user.ID, user.Username)
	if err != nil {
		return user, "", "", err
	}

	return user, accessToken, refreshToken, nil
}
