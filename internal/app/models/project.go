package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model

	Title string `json:"title" gorm:"not null"`
	Code  string `json:"-" gorm:"not null"`
	IP    string `json:"-" gorm:"not null"`

	UserID uint   `json:"user_id" gorm:"not null"`
	User   User   `json:"-" gorm:"foreignkey:UserID"`
	Vars   []Vars `json:"vars" gorm:"foreignkey:ProjectID"`
}
