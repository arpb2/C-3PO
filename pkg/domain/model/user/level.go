package user

type Level struct {
	LevelData
	UserId  uint `json:"user_id"`
	LevelId uint `json:"level_id"`
}

type LevelData struct {
	Code      string `json:"code"`
	Workspace string `json:"workspace"`
}
