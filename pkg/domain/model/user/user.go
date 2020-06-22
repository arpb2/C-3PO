package user

type (
	Type string
	User struct {
		Id           uint   `json:"id"`
		Type         Type   `json:"type"`
		ClassroomID  uint   `json:"classroom_id,omitempty"`
		CurrentLevel uint   `json:"current_level,omitempty"`
		Email        string `json:"email"`
		Name         string `json:"name"`
		Surname      string `json:"surname"`
	}
)

const (
	TypeTeacher Type = "teacher"
	TypeStudent Type = "student"
)
