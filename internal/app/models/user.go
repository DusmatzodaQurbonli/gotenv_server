package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Username string `json:"username" gorm:"unique;not null"`
	Password string `json:"-" gorm:"size:255;not null"`

	RoleID int  `json:"role_id" gorm:"not null"`
	Role   Role `json:"-" gorm:"foreignKey:RoleID;not null"`
}

type Role struct {
	ID   int    `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(20);not null;unique"`
}
