package models

type VarsReq struct {
	Title string `json:"title" gorm:"not null"`
	Value string `json:"value" gorm:"not null"`
}
