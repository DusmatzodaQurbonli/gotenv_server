package db

import (
	"Gotenv/pkg/logger"
	"errors"
)

func Migrate() error {
	if dbConn == nil {
		logger.Error.Printf("[db.Migrate] Error because database connection is nil")

		return errors.New("database connection is not initialized")
	}

	err := dbConn.AutoMigrate()
	if err != nil {
		logger.Error.Printf("[db.Migrate] Error migrating tables: %v", err)

		return err
	}

	return nil
}
