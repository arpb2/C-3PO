package credential

type Service interface {
	GetUserId(email, password string) (uint, error)
}
