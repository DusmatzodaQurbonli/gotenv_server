package db

import (
	"Gotenv/internal/security"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var (
	dbConn     *gorm.DB
	userDBConn *gorm.DB
)

func ConnectToDB() error {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		security.HostName,
		security.Port,
		security.UserName,
		security.Password,
		security.DBName,
		security.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		return err
	}

	dbConn = db

	return nil
}

func GetDBConn() *gorm.DB {
	return dbConn
}

func CloseDBConn() error {
	if sqlDB, err := GetDBConn().DB(); err == nil {
		if err = sqlDB.Close(); err != nil {
			log.Fatalf("Error while closing DB: %s", err)
		}
		fmt.Println("Connection closed successfully")
	} else {
		log.Fatalf("Error while getting *sql.DB from GORM: %s", err)
	}

	return nil
}
