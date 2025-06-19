package models

type LoginProject struct {
	ProjectIP string `json:"project_ip"`
	Code      string `json:"code"`
}

type ProjectReq struct {
	Title string `json:"title" gorm:"not null"`
	Code  string `json:"code" gorm:"not null"`
	IP    string `json:"ip" gorm:"not null"`
}
