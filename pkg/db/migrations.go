package db

import (
	"Gotenv/internal/app/models"
	"Gotenv/internal/app/models/seeds"
	"Gotenv/pkg/logger"
	"errors"
)

func Migrate() error {
	if dbConn == nil {
		logger.Error.Printf("[db.Migrate] Error because database connection is nil")

		return errors.New("database connection is not initialized")
	}

	err := dbConn.AutoMigrate(
		models.User{},
		models.Role{},
		models.Project{},
		models.Vars{},
	)

	if err != nil {
		logger.Error.Printf("[db.Migrate] Error migrating tables: %v", err)

		return err
	}

	err = seeds.SeedRoles(dbConn)
	if err != nil {
		return err
	}

	return nil
}
