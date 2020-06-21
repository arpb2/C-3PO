package classroom

import "github.com/arpb2/C-3PO/pkg/domain/model/user"

type Classroom struct {
	Id          uint        `json:"id"`
	Level       uint        `json:"level"`
	Students    []user.User `json:"students"`
	Teacher     user.User   `json:"teacher"`
}
