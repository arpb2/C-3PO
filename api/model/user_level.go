package model

type UserLevel struct {
	UserId  uint `json:"user_id"`
	LevelId uint `json:"level_id"`
	*UserLevelData
}

type UserLevelData struct {
	Code      string `json:"code"`
	Workspace string `json:"workspace"`
}
