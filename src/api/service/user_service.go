package service

type UserService interface {

	GetUser(userId string) (user *interface{}, err error)

	CreateUser(data ...interface{}) (user *interface{}, err error)

	UpdateUser(user interface{}) error

	DeleteUser(userId string) error

}

type TeacherService interface {
	UserService

	GetStudents(userId string) (students *[]int, err error)

}