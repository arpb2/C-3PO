package service

type Service interface {
	GetUserId(email, password string) (uint, error)
}
