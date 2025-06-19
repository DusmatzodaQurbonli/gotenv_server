package models

import "gorm.io/gorm"

type Vars struct {
	gorm.Model

	Title     string  `json:"title" gorm:"not null"`
	Value     string  `json:"value" gorm:"not null"`
	ProjectID uint    `json:"project_id" gorm:"not null"`
	Project   Project `json:"-" gorm:"foreignkey:ProjectID"`
}
