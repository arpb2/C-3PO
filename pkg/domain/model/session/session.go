package session

type Session struct {
	UserId uint   `json:"user_id"`
	Token  string `json:"token"`
}
