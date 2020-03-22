package user

type (
	Type string
	User struct {
		Id      uint   `json:"id"`
		Type    Type   `json:"type"`
		Email   string `json:"email"`
		Name    string `json:"name"`
		Surname string `json:"surname"`
	}
)

const (
	TypeTeacher Type = "teacher"
	TypeStudent Type = "student"
)
