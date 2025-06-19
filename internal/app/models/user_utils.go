package models

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	RoleID   uint   `json:"role_id"`
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
