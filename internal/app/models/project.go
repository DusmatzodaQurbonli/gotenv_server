package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model

	Title string `json:"title" gorm:"not null"`
	Code  string `json:"code" gorm:"not null"`
	IP    string `json:"ip" gorm:"not null"`

	UserID uint `json:"user_id" gorm:"not null"`
	User   User `json:"-" gorm:"foreignkey:UserID"`
}
