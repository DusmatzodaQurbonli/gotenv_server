package models

type VarsReq struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Value string `json:"value"`
}
