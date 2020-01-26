package model

type Code struct {
	UserId uint   `json:"user_id"`
	Id     uint   `json:"id"`
	Code   string `json:"code"`
}
