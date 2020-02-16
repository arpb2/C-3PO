package model

type UserLevel struct {
	UserId  uint   `json:"user_id"`
	LevelId uint   `json:"level_id"`
	Code    string `json:"code"`
}
