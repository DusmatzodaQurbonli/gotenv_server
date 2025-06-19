package service

import (
	"Gotenv/internal/app/models"
	"Gotenv/internal/repository"
	"Gotenv/pkg/errs"
	"Gotenv/pkg/logger"
	"Gotenv/pkg/utils"
	"fmt"
)

func CreateUser(user models.User) (uint, error) {
	usernameExists, err := repository.UserExists(user.Username)
	if err != nil {
		return 0, fmt.Errorf("failed to check existing user: %w", err)
	}

	if user.Password == "" || user.Username == "" {
		return 0, errs.ErrInvalidData
	}

	if usernameExists {
		logger.Error.Printf("[service.CreateUser] user with username %s already exists", user.Username)

		return 0, errs.ErrUsernameUniquenessFailed
	}

	user.Password = utils.GenerateHash(user.Password)

	var userDB models.User

	if userDB, err = repository.CreateUser(user); err != nil {
		return 0, fmt.Errorf("failed to create user: %w", err)
	}

	return userDB.ID, nil
}
